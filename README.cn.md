# go2generics ![](https://changkun.de/urlstat?mode=github&repo=golang-design/go2generics)

中文 | [English](./README.md)

Go 语言泛型的代码示例（基于类型参数和类型集）

## 支持泛型的 Go 编译器

正在开发阶段的 Go 编译器（称之为 [gotip](https://pkg.go.dev/golang.org/dl/gotip)）支持类型参数和类型集编写的泛型代码。可以通过下面的命令进行安装：

```sh
$ go get golang.org/dl/gotip
$ gotip download
$ gotip version
go version devel go1.18-6e50991 Sat Aug 21 18:23:58 2021 +0000 darwin/arm64
```

## 标准库

下面这些包将出现在 Go 1.18 的发行版中：

### `constraints` 包

该包基于这些提案：

- [golang/go#45458](https://golang.org/issue/45458) proposal: constraints: new package to define standard type parameter constraints
- [golang/go#47319](https://golang.org/issue/47319) proposal: constraints: new package to define standard type parameter constraints (discussion)

可以在[这里](./std/constraints)找到一个可能的实现。

```go
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

### `slices` 包

该包基于这些提案：

- [golang/go#45955](https://golang.org/issue/45955) proposal: slices: new package to provide generic slice functions
- [golang/go#47203](https://golang.org/issue/47203) proposal: slices: new package to provide generic slice functions (discussion)

可以在[这里](./std/slices)找到一个可能的实现。

```go
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

### `maps` 包

该包基于这些提案：

- [golang/go#47649](https://golang.org/issue/47649) proposal: maps: new package to provide generic map functions
- [golang/go#47330](https://golang.org/issue/47330) proposal: maps: new package to provide generic map functions (discussion)

可以在[这里](./std/maps)找到一个可能的实现。

```go
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

### `container/set` 包

该包基于这些提案：

- [golang/go#47331](https://golang.org/issue/47331) proposal: container/set: new package to provide a generic set type (discussion)

可以在[这里](./std/container/set)找到一个可能的实现。

```go
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

### 其他 (仍在讨论中)

- [golang/go#47657](https://golang.org/issue/47657) proposal: sync, sync/atomic: add PoolOf, MapOf, ValueOf
- [golang/go#47632](https://golang.org/issue/47632) proposal: container/heap: add Heap, a heap backed by a slice
- [golang/go#47619](https://golang.org/issue/47619) proposal: generic functions in the sort package

## 更多示例

出了前面的标准库实现示例之外，仓库中还包含了更多其他场景下的的示例。这些示例可以直接
通过安装的 gotip 命令执行：

```
$ gotip run demo/ex1-sort.go
$ gotip run demo/ex2-mapreduce.go
$ gotip run demo/ex3-stack.go
$ gotip run demo/ex4-map.go
$ cd errors      && gotip test
$ cd fmt         && gotip test
$ cd future      && gotip test
$ cd linalg      && gotip test
$ cd list        && gotip test
$ cd math        && gotip test
$ cd metrics     && gotip test
$ cd ring        && gotip test
$ cd stack       && gotip test
$ cd strings     && gotip test
$ cd sync/atomic && gotip test
$ cd tree        && gotip test
```

## 已知的问题

由于当前 Go 的编译器实现上不完整，目前（2021.08.22）已知这些问题：

- 泛型切片表达式尚未实现
- 公开函数的导出和包的导入还需要完善
- 更多类型检查相关的完善

更多满足语言规范的代码（暂时）还不能正常编译执行。例如这些目录下的代码：

```
chans
demo/ex5-loadbalance.go
graph
maps
sched
slices
sync
```

## 进一步阅读

- Changkun Ou. [Go 语言泛型研究总结](./generics.md). 2020.08. 最后更新 2021.08.
- Changkun Ou. [Go 2 泛型: 类型参数](https://changkun.de/s/go2generics/). [Go 夜读 第 80 期](https://talkgo.org). 2020 年 3 月 18 日.

## 许可

BSD-2-Clause

Copyright &copy; 2020-2021 [Changkun Ou](https://changkun.de)
