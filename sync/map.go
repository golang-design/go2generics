// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// Map is a type-safe concurrent-safe map[k]v container.
type Map[K comparable, V any] struct {
	mu     sync.Mutex
	read   atomic.Value // readOnly
	dirty  map[K]*entry
	misses int
}

type readOnly[K comparable, V any] struct {
	m       map[K]*entry
	amended bool // true if the dirty map contains some key not in m.
}

var expunged = unsafe.Pointer(new(any))

type entry struct {
	p unsafe.Pointer // *any
}

func newEntry(i any) *entry {
	return &entry{p: unsafe.Pointer(&i)}
}

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	read, _ := m.read.Load().(readOnly[K, V])
	e, ok := read.m[key]
	if !ok && read.amended {
		m.mu.Lock()
		read, _ = m.read.Load().(readOnly[K, V])
		e, ok = read.m[key]
		if !ok && read.amended {
			e, ok = m.dirty[key]
			m.missLocked()
		}
		m.mu.Unlock()
	}
	if !ok {
		return
	}
	v, loaded := e.load()
	ok = loaded
	if ok {
		value = v.(V)
	}
	return
}

func (e *entry) load() (value any, ok bool) {
	p := atomic.LoadPointer(&e.p)
	if p == nil || p == expunged {
		return nil, false
	}
	return *(*any)(p), true
}

func (m *Map[K, V]) missLocked() {
	m.misses++
	if m.misses < len(m.dirty) {
		return
	}
	m.read.Store(readOnly[K, V]{m: m.dirty})
	m.dirty = nil
	m.misses = 0
}

func (m *Map[K, V]) Delete(key K) {
	m.LoadAndDelete(key)
}

func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	read, _ := m.read.Load().(readOnly[K, V])
	e, ok := read.m[key]
	if !ok && read.amended {
		m.mu.Lock()
		read, _ = m.read.Load().(readOnly[K, V])
		e, ok = read.m[key]
		if !ok && read.amended {
			e, ok = m.dirty[key]
			delete(m.dirty, key)
			m.missLocked()
		}
		m.mu.Unlock()
	}
	if ok {
		v, ok := e.delete()
		if ok {
			value = v.(V)
		}
		loaded = ok
		return
	}
	return
}

func (e *entry) delete() (value any, ok bool) {
	for {
		p := atomic.LoadPointer(&e.p)
		if p == nil || p == expunged {
			return nil, false
		}
		if atomic.CompareAndSwapPointer(&e.p, p, nil) {
			return *(*any)(p), true
		}
	}
}

func (m *Map[K, V]) LoadOrStore(key K, value V) (V, bool) {
	read, _ := m.read.Load().(readOnly[K, V])
	if e, ok := read.m[key]; ok {
		actual, loaded, ok := e.tryLoadOrStore(value)
		if ok {
			return actual.(V), loaded
		}
	}

	m.mu.Lock()
	read, _ = m.read.Load().(readOnly[K, V])
	if e, ok := read.m[key]; ok {
		if e.unexpungeLocked() {
			m.dirty[key] = e
		}
		actual, loaded, _ := e.tryLoadOrStore(value)
		m.mu.Unlock()
		return actual.(V), loaded
	} else if e, ok := m.dirty[key]; ok {
		actual, loaded, _ := e.tryLoadOrStore(value)
		m.missLocked()
		m.mu.Unlock()
		return actual.(V), loaded
	} else {
		if !read.amended {
			m.dirtyLocked()
			m.read.Store(readOnly[K, V]{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry(value)
		actual, loaded := value, false
		m.mu.Unlock()
		return actual, loaded
	}
}

func (e *entry) tryLoadOrStore(i any) (actual any, loaded, ok bool) {
	p := atomic.LoadPointer(&e.p)
	if p == expunged {
		return nil, false, false
	}
	if p != nil {
		return *(*any)(p), true, true
	}

	ic := i
	for {
		if atomic.CompareAndSwapPointer(&e.p, nil, unsafe.Pointer(&ic)) {
			return i, false, true
		}
		p = atomic.LoadPointer(&e.p)
		if p == expunged {
			return nil, false, false
		}
		if p != nil {
			return *(*any)(p), true, true
		}
	}
}

func (e *entry) unexpungeLocked() (wasExpunged bool) {
	return atomic.CompareAndSwapPointer(&e.p, expunged, nil)
}

func (m *Map[K, V]) dirtyLocked() {
	if m.dirty != nil {
		return
	}

	read, _ := m.read.Load().(readOnly[K, V])
	m.dirty = make(map[K]*entry, len(read.m))
	for k, e := range read.m {
		if !e.tryExpungeLocked() {
			m.dirty[k] = e
		}
	}
}

func (e *entry) tryExpungeLocked() (isExpunged bool) {
	p := atomic.LoadPointer(&e.p)
	for p == nil {
		if atomic.CompareAndSwapPointer(&e.p, nil, expunged) {
			return true
		}
		p = atomic.LoadPointer(&e.p)
	}
	return p == expunged
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	// We need to be able to iterate over all of the keys that were already
	// present at the start of the call to Range.
	// If read.amended is false, then read.m satisfies that property without
	// requiring us to hold m.mu for a long time.
	read, _ := m.read.Load().(readOnly[K, V])
	if read.amended {
		// m.dirty contains keys not in read.m. Fortunately, Range is already O(N)
		// (assuming the caller does not break out early), so a call to Range
		// amortizes an entire copy of the map: we can promote the dirty copy
		// immediately!
		m.mu.Lock()
		read, _ = m.read.Load().(readOnly[K, V])
		if read.amended {
			read = readOnly[K, V]{m: m.dirty}
			m.read.Store(read)
			m.dirty = nil
			m.misses = 0
		}
		m.mu.Unlock()
	}

	for k, e := range read.m {
		v, ok := e.load()
		if !ok {
			continue
		}
		if !f(k, v.(V)) {
			break
		}
	}
}

func (m *Map[K, V]) Store(key K, value V) {
	read, _ := m.read.Load().(readOnly[K, V])
	vv := any(value)
	if e, ok := read.m[key]; ok && e.tryStore(&vv) {
		return
	}

	m.mu.Lock()
	read, _ = m.read.Load().(readOnly[K, V])
	if e, ok := read.m[key]; ok {
		if e.unexpungeLocked() {
			m.dirty[key] = e
		}
		vv := any(value)
		e.storeLocked(&vv)
	} else if e, ok := m.dirty[key]; ok {
		vv := any(value)
		e.storeLocked(&vv)
	} else {
		if !read.amended {
			m.dirtyLocked()
			m.read.Store(readOnly[K, V]{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry(value)
	}
	m.mu.Unlock()
}

func (e *entry) tryStore(i *any) bool {
	for {
		p := atomic.LoadPointer(&e.p)
		if p == expunged {
			return false
		}
		if atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i)) {
			return true
		}
	}
}

func (e *entry) storeLocked(i *any) {
	atomic.StorePointer(&e.p, unsafe.Pointer(i))
}
