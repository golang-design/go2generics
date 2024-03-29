// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"testing"
	"testing/quick"
)

type mapOp string

const (
	opLoad          = mapOp("Load")
	opStore         = mapOp("Store")
	opLoadOrStore   = mapOp("LoadOrStore")
	opLoadAndDelete = mapOp("LoadAndDelete")
	opDelete        = mapOp("Delete")
)

var mapOps = [...]mapOp{opLoad, opStore, opLoadOrStore, opLoadAndDelete, opDelete}

// mapCall is a quick.Generator for calls on mapInterface.
type mapCall[K comparable, V any] struct {
	op mapOp
	k  K
	v  V
}

func (c mapCall[K, V]) apply(m mapInterface[K, V]) (any, bool) {
	switch c.op {
	case opLoad:
		return m.Load(c.k)
	case opStore:
		m.Store(c.k, c.v)
		return nil, false
	case opLoadOrStore:
		return m.LoadOrStore(c.k, c.v)
	case opLoadAndDelete:
		return m.LoadAndDelete(c.k)
	case opDelete:
		m.Delete(c.k)
		return nil, false
	default:
		panic("invalid mapOp")
	}
}

type mapResult struct {
	value any
	ok    bool
}

func randValue(r *rand.Rand) string {
	b := make([]byte, r.Intn(4))
	for i := range b {
		b[i] = 'a' + byte(rand.Intn(26))
	}
	return string(b)
}

func (mapCall[K, V]) Generate(r *rand.Rand, size int) reflect.Value {
	c := mapCall[mapOp, string]{op: mapOps[rand.Intn(len(mapOps))], k: mapOp(randValue(r))}
	switch c.op {
	case opStore, opLoadOrStore:
		c.v = string(randValue(r))
	}
	return reflect.ValueOf(c)
}

func applyCalls[K comparable, V any](m mapInterface[K, V], calls []mapCall[K, V]) (results []mapResult, final map[K]V) {
	for _, c := range calls {
		v, ok := c.apply(m)
		results = append(results, mapResult{v, ok})
	}

	final = make(map[K]V)
	m.Range(func(k K, v V) bool {
		final[k] = v
		return true
	})

	return results, final
}

func applyMap[K comparable, V any](calls []mapCall[K, V]) ([]mapResult, map[K]V) {
	return applyCalls[K, V](new(Map[K, V]), calls)
}

func applyRWMutexMap[K comparable, V any](calls []mapCall[K, V]) ([]mapResult, map[K]V) {
	return applyCalls[K, V](new(RWMutexMap[K, V]), calls)
}

func applyDeepCopyMap[K comparable, V any](calls []mapCall[K, V]) ([]mapResult, map[K]V) {
	return applyCalls[K, V](new(DeepCopyMap[K, V]), calls)
}

func TestMapMatchesRWMutex(t *testing.T) {
	if err := quick.CheckEqual(applyMap[mapOp, string], applyRWMutexMap[mapOp, string], nil); err != nil {
		t.Error(err)
	}
}

func TestMapMatchesDeepCopy(t *testing.T) {
	if err := quick.CheckEqual(applyMap[mapOp, string], applyDeepCopyMap[mapOp, string], nil); err != nil {
		t.Error(err)
	}
}

func TestConcurrentRange(t *testing.T) {
	const mapSize = 1 << 10

	m := new(Map[int64, int64])
	for n := int64(1); n <= mapSize; n++ {
		m.Store(n, int64(n))
	}

	done := make(chan struct{})
	var wg sync.WaitGroup
	defer func() {
		close(done)
		wg.Wait()
	}()
	for g := int64(runtime.GOMAXPROCS(0)); g > 0; g-- {
		r := rand.New(rand.NewSource(g))
		wg.Add(1)
		go func(g int64) {
			defer wg.Done()
			for i := int64(0); ; i++ {
				select {
				case <-done:
					return
				default:
				}
				for n := int64(1); n < mapSize; n++ {
					if r.Int63n(mapSize) == 0 {
						m.Store(n, n*i*g)
					} else {
						m.Load(n)
					}
				}
			}
		}(g)
	}

	iters := 1 << 10
	if testing.Short() {
		iters = 16
	}
	for n := iters; n > 0; n-- {
		seen := make(map[int64]bool, mapSize)

		m.Range(func(k, v int64) bool {
			if v%k != 0 {
				t.Fatalf("while Storing multiples of %v, Range saw value %v", k, v)
			}
			if seen[k] {
				t.Fatalf("Range visited key %v twice", k)
			}
			seen[k] = true
			return true
		})

		if len(seen) != mapSize {
			t.Fatalf("Range visited %v elements of %v-element Map", len(seen), mapSize)
		}
	}
}
