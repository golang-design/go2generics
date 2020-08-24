// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"fmt"
	"reflect"
	"testing"
)

type Int int

func (i Int) String() string {
	return fmt.Sprintf("%d", i)
}

func TestStringifiable(t *testing.T) {
	ii := []Int{1, 2, 3, 4, 5}
	ss := Stringify(ii)
	if !reflect.DeepEqual(ss, []string{"1", "2", "3", "4", "5"}) {
		t.Errorf("Stringify wrong, got: %v", ss)
	}
}
