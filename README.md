# go2generics

This repository illustrates a big chunk of demos and researches about Go 2 generics.

## Getting started

There are several documents for you to getting familiar with generics:

- [A Summary of Go Generics Research](./generics.md)
- Changkun Ou. Go 2 Generics? A (P)review. March 2020. https://changkun.de/s/go2generics/.

## Code examples

You can use `go2go` to build all code examples in this repository.

### Using `go2go`

1. Clone `go` repository

2. Download `go2go` code patch

```
git fetch "https://go.googlesource.com/go" refs/changes/17/187317/13 && git checkout FETCH_HEAD
```

3. Place the repository in `GO2PATH`. The `go2generics` repository path should looks like the follows:

```
$GO2PATH/src/github.com/changkun/go2generics
```

## Licnese

BSD-2-Clause &copy; [Changkun Ou](https://changkun.de)
