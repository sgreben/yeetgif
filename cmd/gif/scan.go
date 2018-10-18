package main

import (
	"image"
	"math"
	"math/rand"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandScan(cmd *cli.Cmd) {
	cmd.Before = InputAndDuplicate
	cmd.Spec = "[OPTIONS]"
	var (
		f       = gifcmd.Float{Value: 78.9}
		i       = gifcmd.Float{Value: 0.9}
		l       = gifcmd.FloatsCSV{Values: []float64{1, 1.5, 0, 0, 1}}
		h       = gifcmd.FloatsCSV{Values: []float64{0, 0, -0.2, 0.2, 0}}
		s       = gifcmd.FloatsCSV{Values: []float64{0, 0, 0.3, -0.3, 0}}
		fuzz    = gifcmd.FloatsCSV{Values: []float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0}}
		fuzzMul = gifcmd.Float{Value: 0.1}
		offset  = gifcmd.FloatsCSV{Values: []float64{0}}
	)
	cmd.VarOpt("f frequency", &f, "")
	cmd.VarOpt("l", &l, "")
	cmd.VarOpt("u", &h, "")
	cmd.VarOpt("s", &s, "")
	cmd.VarOpt("z fuzz", &fuzzMul, "")
	cmd.VarOpt("o offset", &offset, "")
	cmd.VarOpt("i intensity", &i, "")
	cmd.Action = func() {
		lF := l.PiecewiseLinear(0, 1)
		hF := h.PiecewiseLinear(0, 1)
		sF := s.PiecewiseLinear(0, 1)
		for i := range fuzz.Values {
			fuzz.Values[i] = fuzzMul.Value * fuzz.Values[i]
		}
		fuzzF := fuzz.PiecewiseLinear(0, 1)
		for i := range fuzz.Values {
			fuzz.Values[i] = fuzzF(float64(i)/float64(len(fuzz.Values))) * rand.Float64()
		}
		fuzzF = fuzz.PiecewiseLinear(0, 1)
		offsetF := offset.PiecewiseLinear(0, 1)
		iInv := 1 - i.Value
		Scan(images, func(x, y, t float64, h, s, l, a *float64) {
			z := offsetF(t) + f.Value*y + fuzzF(t) + fuzzF(1-y)
			z = z - math.Floor(z)
			*h = *h + i.Value*(hF(z))
			*l = iInv**l + i.Value*(*l*lF(z))
			newS := iInv**s + i.Value*(*s+sF(z))
			if newS > 1.0 {
				newS = 1.0
			}
			*s = newS
		})
	}
}

func Scan(images []image.Image, f func(x, y, t float64, h, s, l, a *float64)) {
	n := float64(len(images))
	scan := func(i int) {
		t := float64(i) / n
		b := images[i].Bounds()
		bw := float64(b.Dx())
		bh := float64(b.Dy())
		images[i] = imaging.AdjustHSLAFunc(images[i], func(x, y int, h, s, l, a *float64) {
			xf := float64(x-b.Min.X) / bw
			yf := float64(y-b.Min.Y) / bh
			f(xf, yf, t, h, s, l, a)
		})
	}
	parallel(len(images), scan, "scan")
}
