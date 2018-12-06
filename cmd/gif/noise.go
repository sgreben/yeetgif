package main

import (
	"image"

	"github.com/sgreben/yeetgif/pkg/gifmath"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandNoise(cmd *cli.Cmd) {
	cmd.Before = InputAndDuplicate
	cmd.Spec = "[OPTIONS] [INTENSITY]"
	var (
		n  = gifcmd.FloatsCSV{Values: []float64{1.0}}
		n1 = gifcmd.FloatsCSV{Values: []float64{0.02}}
		n2 = gifcmd.FloatsCSV{Values: []float64{0.5}}
		n3 = gifcmd.FloatsCSV{Values: []float64{0.1}}
	)
	cmd.VarArg("INTENSITY", &n, "🌀️")
	cmd.VarOpt("u", &n1, "🌀️")
	cmd.VarOpt("s", &n2, "🌀️")
	cmd.VarOpt("l", &n3, "🌀")
	cmd.Action = func() {
		Noise(images, n.PiecewiseLinear(0, 1), n1.PiecewiseLinear(0, 1), n2.PiecewiseLinear(0, 1), n3.PiecewiseLinear(0, 1))
	}
}

func Noise(images []image.Image, noiseF, noise1F, noise2F, noise3F func(float64) float64) {
	n := float64(len(images))
	noise := func(i int) {
		t := float64(i) / n
		noise, noise1, noise2, noise3 := noiseF(t), noise1F(t), noise2F(t), noise3F(t)
		images[i] = imaging.AdjustHSLAFunc(images[i], func(x, y int, h, s, l, a *float64, seed *int) {
			*h = *h + noise*noise1*float64(gifmath.RandomFloat32Signed(seed))
			*s = *s + noise*noise2*float64(gifmath.RandomFloat32Signed(seed))
			*l = *l + noise*noise3*float64(gifmath.RandomFloat32Signed(seed))
		})
	}
	parallel(len(images), noise, "noise")
}
