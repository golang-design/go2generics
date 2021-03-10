# go2generics [![PkgGoDev](https://pkg.go.dev/badge/golang.design/x/go2generics)](https://pkg.go.dev/golang.design/x/go2generics) ![](https://changkun.de/urlstat?mode=github&repo=golang-design/go2generics)

A chunk of demos and research regarding Go 2 generics (type parameters).

## Getting started

There are several documents for you to getting familiar with generics:

- [A Summary of Go Generics Research](./generics.md)
- Changkun Ou. Go 2 Generics: Type Parameters. https://changkun.de/s/go2generics/.

## Code examples

You can use the latest Go tip version to build all code examples in this repository.

```sh
$ git clone https://github.com/golang/go && cd src && ./all.bash

$ ../bin/go test -gcflags=all=-G=3 -v ./...
```

The `-gcflags=all=-G=3` is the key compiler flag to run Go code with type parameters.

## Licnese

BSD-2-Clause

Copyright &copy; 2020-2021 [Changkun Ou](https://changkun.de)
