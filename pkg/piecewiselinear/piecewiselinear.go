package piecewiselinear

import (
	"math"
)

type Function struct {
	X []float64
	Y []float64
}

func (f *Function) SetDomain(min, max float64) {
	n := len(f.Y)
	f.X = make([]float64, n)
	d := max - min
	min, max = math.Min(max, min), math.Max(max, min)
	for i := range f.Y {
		f.X[i] = min + d*(float64(i)/float64(n-1))
	}
}

func (f Function) At(x float64) float64 {
	X, Y := f.X, f.Y
	i, j := 0, len(X)
	for i < j {
		h := int(uint(i+j) >> 1)
		if X[h] < x {
			i = h + 1
		} else {
			j = h
		}
	}
	switch i {
	case 0:
		return Y[0]
	case len(X):
		return Y[len(X)-1]
	}
	w := (x - X[i-1]) / (X[i] - X[i-1])
	return (1-w)*Y[i-1] + w*Y[i]
}
