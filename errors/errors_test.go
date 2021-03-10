// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import (
	"testing"
)

type e struct{}

func (a e) Error() string { return "error" }

func TestTryCatch(t *testing.T) {
	Try(func() (int, error) {
		return 0, nil
	}).Catch(func(_ int, err error) int {
		t.Fatalf("catch error: %v", err)
		return 0
	}).Final(func(result int) {
		t.Log("errors: everything is good")
	})

	Try(func() (int, error) {
		return 1, e{}
	}).Catch(func(result int, err error) int {
		t.Logf("captured result: %v", result)
		t.Logf("captured error: %v", err)
		return 1
	}).Final(func(result int) {
		if result != 1 {
			t.Fatalf("cannot capture error")
		}
	})

	Try(func() (int, error) {
		return 1, nil
	}).Final(func(r int) {
		if r != 1 {
			t.Fatalf("result from try block is not as expected")
		}
	})
}
