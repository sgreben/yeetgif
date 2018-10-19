package main

import (
	"image"
	"math"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandHue(cmd *cli.Cmd) {
	cmd.Before = ProcessInput
	cmd.Spec = "[OPTIONS]"
	var (
		f    = gifcmd.Float{Value: 1.0}
		a    = gifcmd.Float{Value: 0.1}
		from = gifcmd.Float{Value: -1.0}
		to   = gifcmd.Float{Value: 1.0}
	)
	cmd.VarOpt("f frequency", &f, "")
	cmd.VarOpt("a amplitude", &a, "")
	cmd.Action = func() {
		HuePulse(images, f.Value, from.Value*a.Value, to.Value*a.Value)
	}
}

func HuePulse(images []image.Image, frequency, from, to float64) {
	n := float64(len(images))
	mid := (from + to) / 2
	dist := to - from
	hue := func(i int) {
		delta := mid + (dist * math.Sin(math.Pi+2*math.Pi*frequency*float64(i)/n) / 2)
		images[i] = imaging.AdjustHueRotate(images[i], delta)
	}
	parallel(len(images), hue)
}
