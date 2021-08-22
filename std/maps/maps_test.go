// Copyright 2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maps

import (
	"reflect"
	"strings"
	"testing"
)

func TestKeys(t *testing.T) {
	tests := []struct {
		m    map[string]string
		want [][]string
	}{
		{
			m:    map[string]string{"a": "b", "c": "d"},
			want: [][]string{[]string{"a", "c"}, []string{"c", "a"}},
		},
	}

	for _, tt := range tests {
		got := Keys(tt.m)
		if !reflect.DeepEqual(got, tt.want[0]) && !reflect.DeepEqual(got, tt.want[1]) {
			t.Fatalf("unexpected Keys, want %v got %v", tt.want, got)
		}
	}
}

func TestValues(t *testing.T) {
	tests := []struct {
		m    map[string]string
		want [][]string
	}{
		{
			m:    map[string]string{"a": "b", "c": "d"},
			want: [][]string{[]string{"b", "d"}, []string{"d", "b"}},
		},
	}

	for _, tt := range tests {
		got := Values(tt.m)
		if !reflect.DeepEqual(got, tt.want[0]) && !reflect.DeepEqual(got, tt.want[1]) {
			t.Fatalf("unexpected Values, want %v got %v", tt.want, got)
		}
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		m1   map[string]string
		m2   map[string]string
		want bool
	}{
		{
			m1:   map[string]string{"a": "b", "c": "d"},
			m2:   map[string]string{"a": "b", "c": "d"},
			want: true,
		},
	}

	for _, tt := range tests {
		got := Equal(tt.m1, tt.m2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("unexpected Equal, want %v got %v", tt.want, got)
		}
	}

	for _, tt := range tests {
		got := EqualFunc(tt.m1, tt.m2, func(v1, v2 string) bool {
			return strings.Compare(v1, v2) == 0
		})
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("unexpected EqualFunc, want %v got %v", tt.want, got)
		}
	}
}

func TestClear(t *testing.T) {
	tests := []struct {
		m    map[string]string
		want map[string]string
	}{
		{
			m:    map[string]string{"a": "b", "c": "d"},
			want: map[string]string{},
		},
	}

	for _, tt := range tests {
		Clear(tt.m)
		if !reflect.DeepEqual(tt.m, tt.want) {
			t.Fatalf("unexpected Clear, want %v got %v", tt.want, tt.m)
		}
	}
}

func TestClone(t *testing.T) {
	tests := []struct {
		m map[string]string
	}{
		{
			m: map[string]string{"a": "b", "c": "d"},
		},
	}

	for _, tt := range tests {
		got := Clone(tt.m)
		if !reflect.DeepEqual(tt.m, got) {
			t.Fatalf("unexpected Clone, want %v got %v", tt.m, got)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		dst  map[string]string
		src  map[string]string
		want map[string]string
	}{
		{
			dst:  map[string]string{"a": "b", "c": "d"},
			src:  map[string]string{"a": "b", "c": "d"},
			want: map[string]string{"a": "b", "c": "d"},
		},
	}

	for _, tt := range tests {
		Add(tt.dst, tt.src)
		if !reflect.DeepEqual(tt.dst, tt.want) {
			t.Fatalf("unexpected Clone, want %v got %v", tt.want, tt.dst)
		}
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		m    map[string]string
		want map[string]string
	}{
		{
			m:    map[string]string{"a": "b", "c": "d"},
			want: map[string]string{"a": "b"},
		},
	}

	for _, tt := range tests {
		Filter(tt.m, func(k, v string) bool {
			if k == "a" {
				return true
			}
			return false
		})
		if !reflect.DeepEqual(tt.m, tt.want) {
			t.Fatalf("unexpected Filter, want %v got %v", tt.want, tt.m)
		}
	}
}
