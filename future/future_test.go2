// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package future

import (
	"testing"
	"time"
)

func TestFuture(t *testing.T) {
	s := "done"
	f := Future[string]{}

	go func() {
		time.Sleep(time.Second)
		f.Put(s)
	}()

	if f.Get() != s {
		t.Fatalf("future failed")
	}
}
