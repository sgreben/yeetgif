package piecewiselinear

import "testing"

func TestFunction_At(t *testing.T) {
	tests := []struct {
		name string
		X    []float64
		Y    []float64
		x    float64
		want float64
	}{
		{
			name: "simple",
			X:    []float64{0, 1},
			Y:    []float64{0, 1},
			x:    0.5,
			want: 0.5,
		},
		{
			name: "saw(0.25)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    0.25,
			want: -1,
		},
		{
			name: "saw(0.125)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    0.125,
			want: -0.5,
		},
		{
			name: "saw(1.0)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    1.0,
			want: 0.0,
		},
		{
			name: "saw(0.0)",
			X:    []float64{0, 0.25, 0.5, 0.75, 1.0},
			Y:    []float64{0, -1, 0, 1, 0},
			x:    0.0,
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Function{
				X: tt.X,
				Y: tt.Y,
			}
			if got := f.At(tt.x); got != tt.want {
				t.Errorf("Function.At() = %v, want %v", got, tt.want)
			}
		})
	}
}
