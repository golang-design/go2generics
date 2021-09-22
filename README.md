# go2generics ![](https://changkun.de/urlstat?mode=github&repo=golang-design/go2generics)

English | [中文](./README.cn.md)

A chunk of demos for Go generics design (based on type parameters and type sets).

## The Go Compiler with Generics Support

The latest [Go tip version](https://pkg.go.dev/golang.org/dl/gotip) supports
the type parameters and type sets design. To install `gotip`:

```sh
$ go get golang.org/dl/gotip
$ gotip download
$ gotip version
go version devel go1.18-6e50991 Sat Aug 21 18:23:58 2021 +0000 darwin/arm64
```

## Standard Package

The following packages will appear in the Go 1.18 release:

### Package `constraints`

This package is from the following discussions:

- [golang/go#45458](https://golang.org/issue/45458) proposal: constraints: new package to define standard type parameter constraints
- [golang/go#47319](https://golang.org/issue/47319) proposal: constraints: new package to define standard type parameter constraints (discussion)

See a possible implementation [here](./std/constraints).

```go
// Package constraints defines a set of useful constraints to be used
// with type parameters.
package constraints

type Signed interface
type Unsigned interface
type Integer interface
type Float interface
type Complex interface
type Ordered interface
type Slice[Elem any] interface
type Map[Key comparable, Val any] interface
type Chan[Elem any] interface
```

### Package `slices`

This package is from the following discussions:

- [golang/go#45955](https://golang.org/issue/45955) proposal: slices: new package to provide generic slice functions
- [golang/go#47203](https://golang.org/issue/47203) proposal: slices: new package to provide generic slice functions (discussion)

See a possible implementation [here](./std/slices).

```go
// Package slices defines various functions useful with slices of any type.
// Unless otherwise specified, these functions all apply to the elements
// of a slice at index 0 <= i < len(s).
package slices

import "constraints"

func Equal[T comparable](s1, s2 []T) bool
func EqualFunc[T1, T2 any](s1 []T1, s2 []T2, eq func(T1, T2) bool) bool
func Compare[T constraints.Ordered](s1, s2 []T) int
func CompareFunc[T any](s1, s2 []T, cmp func(T, T) int) int
func Index[T comparable](s []T, v T) int
func IndexFunc[T any](s []T, f func(T) bool) int
func Contains[T comparable](s []T, v T) bool
func Insert[S constraints.Slice[T], T any](s S, i int, v ...T) S
func Delete[S constraints.Slice[T], T any](s S, i, j int) S
func Clone[S constraints.Slice[T], T any](s S) S
func Compact[S constraints.Slice[T], T comparable](s S) S
func CompactFunc[S constraints.Slice[T], T any](s S, cmp func(T, T) bool) S
func Grow[S constraints.Slice[T], T any](s S, n int) S
func Clip[S constraints.Slice[T], T any](s S) S
```

### Package `maps`

This package is from the following discussions:

- [golang/go#47649](https://golang.org/issue/47649) proposal: maps: new package to provide generic map functions
- [golang/go#47330](https://golang.org/issue/47330) proposal: maps: new package to provide generic map functions (discussion)

See a possible implementation [here](./std/maps).

```go
// Package maps defines various functions useful with maps of any type.
package maps

func Keys[K comparable, V any](m map[K]V) []K
func Values[K comparable, V any](m map[K]V) []V
func Equal[K, V comparable](m1, m2 map[K]V) bool
func EqualFunc[K comparable, V1, V2 any](m1 map[K]V1, m2 map[K]V2, cmp func(V1, V2) bool) bool
func Clear[K comparable, V any](m map[K]V)
func Clone[K comparable, V any](m map[K]V) map[K]V
func Add[K comparable, V any](dst, src map[K]V)
func Filter[K comparable, V any](m map[K]V, keep func(K, V) bool)
```

### Package `container/set`

This package is from the following discussions:

- [golang/go#47331](https://golang.org/issue/47331) proposal: container/set: new package to provide a generic set type (discussion)

See a possible implementation [here](./std/container/set).


```go
// Package set defines a Set type that holds a set of elements.
package set

type Set[Elem comparable] struct { ... }
func (s *Set[Elem]) Add(v ...Elem)
func (s *Set[Elem]) AddSet(s2 Set[Elem])
func (s *Set[Elem]) Remove(v ...Elem)
func (s *Set[Elem]) RemoveSet(s2 Set[Elem])
func (s *Set[Elem]) Has(v Elem) bool
func (s *Set[Elem]) HasAny(s2 Set[Elem]) bool
func (s *Set[Elem]) HasAll(s2 Set[Elem]) bool
func (s *Set[Elem]) Values() []Elem
func (s *Set[Elem]) Equal(s2 Set[Elem]) bool
func (s *Set[Elem]) Clear()
func (s *Set[Elem]) Clone() Set[Elem]
func (s *Set[Elem]) Filter(keep func(Elem) bool)
func (s *Set[Elem]) Len() int
func (s *Set[Elem]) Do(f func(Elem) bool)

func Of[Elem comparable](v ...Elem) Set[Elem]
func Union[Elem comparable](s1, s2 Set[Elem]) Set[Elem]
func Intersection[Elem comparable](s1, s2 Set[Elem]) Set[Elem]
func Difference[Elem comparable](s1, s2 Set[Elem]) Set[Elem]
```

### Others (under discussion)

- [golang/go#47657](https://golang.org/issue/47657) proposal: sync, sync/atomic: add PoolOf, MapOf, ValueOf
- [golang/go#47632](https://golang.org/issue/47632) proposal: container/heap: add Heap, a heap backed by a slice
- [golang/go#47619](https://golang.org/issue/47619) proposal: generic functions in the sort package

## Further  Working Examples

The current compiler implementation is still under development.
These further examples can be run without an error (last update: 2021.09.22):

```sh
gotip run demo/ex1-sort.go
gotip run demo/ex2-mapreduce.go
gotip run demo/ex3-stack.go
gotip run demo/ex4-map.go
gotip run demo/ex5-loadbalance.go
cd chans       && gotip test ./... && cd ..
cd errors      && gotip test ./... && cd ..
cd fmt         && gotip test ./... && cd ..
cd future      && gotip test ./... && cd ..
cd graph       && gotip test ./... && cd ..
cd linalg      && gotip test ./... && cd ..
cd list        && gotip test ./... && cd ..
cd math        && gotip test ./... && cd ..
cd metrics     && gotip test ./... && cd ..
cd ring        && gotip test ./... && cd ..
cd stack       && gotip test ./... && cd ..
cd strings     && gotip test ./... && cd ..
cd sync        && gotip test ./... && cd ..
cd tree        && gotip test ./... && cd ..
```
## Known Issues

This package cannot be run yet:

```
cd sched       && gotip test && cd ..
```

## References

Here are some documents to get familiar with the spirit of generics:

- Changkun Ou. [A Summary of Go Generics Research](./generics.md) 2020.08. Last Updates: 2021.08.
- Changkun Ou. Go 2 Generics: Type Parameters. https://changkun.de/s/go2generics/.
## Licnese

BSD-2-Clause

Copyright &copy; 2020-2021 [Changkun Ou](https://changkun.de)
