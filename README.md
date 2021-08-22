# go2generics [![PkgGoDev](https://pkg.go.dev/badge/golang.design/x/go2generics)](https://pkg.go.dev/golang.design/x/go2generics) ![](https://changkun.de/urlstat?mode=github&repo=golang-design/go2generics)

A chunk of demos and research regarding Go 2 generics (type parameters and type sets).

## Getting started

There are several documents for you to getting familiar with generics:

- [A Summary of Go Generics Research](./generics.md)
- Changkun Ou. Go 2 Generics: Type Parameters. https://changkun.de/s/go2generics/.

## Code examples

You can use the latest Go tip version to build all code examples in this repository.

```sh
$ git clone https://github.com/golang/go && cd src && ./all.bash

$ ../bin/go test -v ./...
```

## Working Examples

The current compiler implementation is still under development. 
These examples can be run without an error:

```
gotip run demo/ex1-sort.go
gotip run demo/ex2-mapreduce.go
gotip run demo/ex3-stack.go
gotip run demo/ex4-map.go
cd errors && gotip test
cd fmt    && gotip test
cd future && gotip test
cd linalg && gotip test
cd list   && gotip test
cd math   && gotip test
cd metrics&& gotip test
cd ring   && gotip test
cd stack  && gotip test
cd strings&& gotip test
cd sync/atomic && gotip test
cd tree   && gotip test
```

## Upcoming Standard Package

See [changkun/generics](https://github.com/changkun/generics).

## Known Issues

The know issues of the current implementation:

- Type parameter for slices is not supported yet.
- Package import is still buggy

These written packages are not runnable yet (will trigger some internal compiler bug):

```
chans
demo/ex5-loadbalance.go
graph
maps
sched
slices
sync
```

## Licnese

BSD-2-Clause

Copyright &copy; 2020-2021 [Changkun Ou](https://changkun.de)
