package piecewiselinear

import "math"

// Span generates `nPoints` equidistant points spanning [min,max]
func Span(min, max float64, nPoints int) []float64 {
	X := make([]float64, nPoints)
	min, max = math.Min(max, min), math.Max(max, min)
	d := max - min
	for i := range X {
		X[i] = min + d*(float64(i)/float64(nPoints-1))
	}
	return X
}
