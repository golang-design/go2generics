// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chans

import (
	"math/rand"
	"testing"
)

func getInputChan() <-chan int {
	input := make(chan int, 20)
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	go func() {
		for num := range numbers {
			input <- num
		}
		close(input)
	}()
	return input
}

func TestFanin(t *testing.T) {
	chs := make([]<-chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = getInputChan()
	}

	out := Fanin(chs...)
	count := 0
	for range out {
		count++
	}
	if count != 100 {
		t.Fatalf("Fanin failed")
	}
}

func TestLB(t *testing.T) {
	ins := make([]<-chan int, 10)
	for i := 0; i < 10; i++ {
		ins[i] = getInputChan()
	}
	outs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		outs[i] = make(chan int, 10)
	}
	LB(func(m int) int {
		return rand.Intn(m)
	}, ins, outs)
}
