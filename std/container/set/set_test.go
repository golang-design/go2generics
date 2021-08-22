// Copyright 2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package set

import (
	"reflect"
	"testing"
)

type tests[T comparable] struct {
	es   []string
	want Set[T]
}

func TestOf(t *testing.T) {
	tests := []tests[string]{
		{
			es: []string{"a", "b", "b"},
			want: Set[string]{
				map[string]struct{}{
					"a": struct{}{},
					"b": struct{}{},
				},
			},
		},
	}

	for _, tt := range tests {
		got := Of(tt.es...)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("unexpected Of, got %v want %v", got, tt.want)
		}
	}
}
