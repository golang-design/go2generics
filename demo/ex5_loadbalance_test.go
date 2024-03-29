// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package demo_test

import (
	"math/rand"
	"sync"
	"testing"
)

func getInputChan() <-chan int {
	input := make(chan int, 100)
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	go func() {
		for num := range numbers {
			input <- num
		}
		close(input)
	}()

	return input
}

func getSquareChan(input <-chan int) <-chan int {
	output := make(chan int, 100)
	go func() {
		for num := range input {
			output <- num * num
		}
		close(output)
	}()
	return output
}

// generic fan-in -- by changkun
func Fanin[T any](chans ...<-chan T) <-chan T {
	buf := 0
	for _, ch := range chans {
		if len(ch) > buf {
			buf = len(ch)
		}
	}
	out := make(chan T)
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch <-chan T) {
			for v := range ch {
				out <- v
			}
			wg.Done()
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// generic fan-out -- by changkun
func Fanout[T any](in <-chan T, outs ...chan T) {
	for v := range in {
		i := rand.Intn(len(outs))
		outs[i] <- v
	}
	for _, ch := range outs {
		close(ch)
	}
}

func TestLoadBalance(t *testing.T) {
	inA := getInputChan()
	inB := getInputChan()
	outs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		outs[i] = make(chan int, 100)
	}

	// load balancing
	Fanout(Fanin(inA, inB), outs...)

	wg := sync.WaitGroup{}
	wg.Add(len(outs))
	for i := 0; i < 10; i++ {
		go func(i int) {
			for v := range outs[i] {
				println(v)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
