// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chans

import (
	"sync"
	"math/rand"
)

// Fanin implements a generic fan-in for variadic channels.
func Fanin[T any](chans ...<-chan T) <-chan T {
	buf := 0
	for _, ch := range chans {
		if len(ch) > buf { buf = len(ch) }
	}
	out := make(chan T, buf)
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch <-chan T) {
			for v := range ch { out <- v }
			wg.Done()
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Fanout implements a generic fan-out for variadic channels
func Fanout[T any](randomizer func(max int) int, In <-chan T, Outs ...chan T) {
	l := len(Outs)
	for v := range In {
		i := randomizer(l)
		if i < 0 || i > l { i = rand.Intn(l) }
		go func(v T) { Outs[i] <- v }(v)
	}
}

func LB[T any](randomizer func(max int) int, ins []<-chan T, outs []chan T) {
	Fanout(randomizer, Fanin(ins...), outs...)
}
