package main

import (
	"image"
	"image/color"
	"math"

	"github.com/sgreben/yeetgif/pkg/piecewiselinear"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandWobble(cmd *cli.Cmd) {
	cmd.Spec = "[OPTIONS]"
	const (
		wobbleTypeSine   = "sine"
		wobbleTypeSnap   = "snap"
		wobbleTypeSaw    = "saw"
		wobbleTypeSticky = "sticky"
		wobbleTypeCustom = "custom"
	)
	var (
		f  = gifcmd.Float{Value: 1.0}
		a  = gifcmd.Float{Value: 20.0}
		ph = gifcmd.Float{Value: 0.0}
		t  = gifcmd.Enum{
			Choices: []string{
				wobbleTypeSine,
				wobbleTypeSnap,
				wobbleTypeSaw,
				wobbleTypeSticky,
			},
			Value: wobbleTypeSine,
		}
		c = gifcmd.FloatsCSV{}
	)
	cmd.VarOpt("f frequency", &f, "")
	cmd.VarOpt("a amplitude", &a, "")
	cmd.VarOpt("p phase", &ph, "")
	cmd.VarOpt("t type", &t, t.Help())
	cmd.VarOpt("custom", &c, "comma-separated angles (°), e.g. 0,10,0,60,0")
	cmd.Action = func() {
		if len(c.Texts) > 0  {
			t.Value = wobbleTypeCustom
		}
		frequency := f.Value
		amplitude := a.Value
		phase := ph.Value
		sine := func(t float64) float64 {
			return amplitude * math.Sin(2*math.Pi*phase+2*math.Pi*frequency*t)
		}
		sawF := piecewiselinear.Function{
			X: []float64{0.000, 0.250, 0.500, 0.750, 1.000},
			Y: []float64{0.000, 1.000, 0.000, -1.00, 0.000},
		}
		saw := func(t float64) float64 {
			return amplitude * sawF.At(t)
		}
		snap := func(t float64) float64 {
			y := math.Sin(2*math.Pi*phase + 2*math.Pi*frequency*t)
			k := 5
			z := math.Sin(1.0)
			for i := 0; i < k; i++ {
				y = math.Sin(y)
				z = math.Sin(z)
			}
			return amplitude * y / z
		}
		stickyF := piecewiselinear.Function{
			X: []float64{0.000, 0.100, 0.250, 0.400, 0.500, 0.600, 0.750, 0.900, 1.000},
			Y: []float64{0.000, 0.950, 1.000, 0.950, 0.000, -0.95, -1.00, -0.95, 0.000},
		}
		sticky := func(t float64) float64 {
			w := 0.95
			s := math.Sin(2*math.Pi*phase + 2*math.Pi*frequency*t)
			x := w*stickyF.At(t) + (1-w)*s
			return amplitude * x
		}
		customF := piecewiselinear.Function{}
		k := float64(len(c.Values) - 1)
		for i := 0; i < len(c.Values); i++ {
			customF.X = append(customF.X, float64(i)/k)
			customF.Y = append(customF.Y, float64((c.Values)[i]))
		}
		fs := map[string]func(float64) float64{
			wobbleTypeSine:   sine,
			wobbleTypeSaw:    saw,
			wobbleTypeSnap:   snap,
			wobbleTypeSticky: sticky,
			wobbleTypeCustom: func(t float64) float64 {
				return customF.At(t)
			},
		}
		Wobble(images, fs[t.Value])
	}
}

// Wobble `images`
func Wobble(images []image.Image, f func(float64) float64) {
	n := float64(len(images))
	rotate := func(i int) {
		angle := f(float64(i) / n)
		bPre := images[i].Bounds()
		images[i] = imaging.Rotate(images[i], angle, color.Transparent)
		bPost := images[i].Bounds()
		offset := image.Point{
			X: (bPost.Dx() - bPre.Dx()) / 2,
			Y: (bPost.Dy() - bPre.Dy()) / 2,
		}
		bPre.Min = bPre.Min.Add(offset)
		bPre.Max = bPre.Max.Add(offset)
		if !config.Pad {
			images[i] = imaging.Crop(images[i], bPre)
		}
	}
	parallel(len(images), rotate)
}
