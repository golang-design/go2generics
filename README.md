# go2generics ![](https://changkun.de/urlstat?mode=github&repo=golang-design/go2generics)

A chunk of demos for Go generics design (based on type parameters and type sets).

## Standard Package

### Package `constraints`

See the official implementation [here](https://github.com/golang/exp/tree/master/constraints).

### Package `slices`

See the official implementation [here](https://github.com/golang/exp/tree/master/slices).

### Package `maps`

See the official implementation [here](https://github.com/golang/exp/tree/master/maps).

### Package `container/set`

This package is from the following discussions:

- [golang/go#47331](https://golang.org/issue/47331) proposal: container/set: new package to provide a generic set type (discussion)

See a possible implementation [here](./std/container/set).

### Others (under discussion)

- [golang/go#47657](https://golang.org/issue/47657) proposal: sync, sync/atomic: add PoolOf, MapOf, ValueOf
- [golang/go#47632](https://golang.org/issue/47632) proposal: container/heap: add Heap, a heap backed by a slice
- [golang/go#47619](https://golang.org/issue/47619) proposal: generic functions in the sort package

## Further Examples

See folders in this repository.
## Known Issues

- https://go.dev/issue/45639
- https://go.dev/issue/51338
- sync/*
- notsupport/*

## References

Here are some documents to get familiar with the spirit of generics:

- Changkun Ou. [A Summary of Go Generics Research](./generics.md) 2020.08. Last Updates: 2021.08.
- Changkun Ou. Go 2 Generics: Type Parameters. https://changkun.de/s/go2generics/.
## Licnese

BSD-2-Clause

Copyright &copy; 2020-2021 [Changkun Ou](https://changkun.de)
