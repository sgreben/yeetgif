# piecewiselinear

[![](https://godoc.org/github.com/sgreben/piecewiselinear?status.svg)](http://godoc.org/github.com/sgreben/piecewiselinear) [![](https://goreportcard.com/badge/github.com/sgreben/piecewiselinear/goreportcard)](https://goreportcard.com/report/github.com/sgreben/piecewiselinear) [![cover.run](https://cover.run/go/github.com/sgreben/piecewiselinear.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fsgreben%2Fpiecewiselinear) [![Build Status](https://travis-ci.org/sgreben/piecewiselinear.svg?branch=master)](https://travis-ci.org/sgreben/piecewiselinear)

A tiny library for linear interpolation. `O(log(N))` per evaluation for `N` control points.

```go
import "github.com/sgreben/piecewiselinear"
```

## Get it

```sh
go get -u "github.com/sgreben/piecewiselinear"
```

## Use it

```go
import "github.com/sgreben/piecewiselinear"

func main() {
    f := piecewiselinear.Function{Y:[]float64{0,1,0}} // range: "hat" function
    f.X = piecewiselinear.Span(0, 1, len(f.Y)) // domain: equidistant points along X axis
    fmt.Println(
        f.At(0), // f.At(x) evaluates f at x
        f.At(0.25),
        f.At(0.5),
        f.At(0.75),
        f.At(1.0),
    )
    // Output:
    // 0 0.5 1 0.5 0
}
```
