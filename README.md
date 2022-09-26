# goex

[![Go Reference](https://pkg.go.dev/badge/github.com/rstudio/goex.svg)](https://pkg.go.dev/github.com/rstudio/goex)

Roughly the kind of Go things that one might find under golang.org/x/ ; missing standard library-ish stuff

Go as in [go](https://go.dev); Ex as in ex(perimental|tension).

## caveats and assumptions

This library should be considered a "v0" dependency that cannot be pinned to a version range. Backward compatibility is
*not guaranteed*, but best effort is made to keep the upgrade process pain-free.

The [zap](https://pkg.go.dev/go.uber.org/zap) logging library is used in many places, with no effort to allow for
pluggability of alternate logging libraries, or using an abstract library such as
[logr](https://pkg.go.dev/github.com/go-logr/logr).

The HTTP middlewares assume use with code that can adapt "next" handlers to `http.Handler` and accept an `http.Handler`
in return.
