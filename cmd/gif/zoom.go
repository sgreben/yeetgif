package main

import (
	"image"

	"github.com/sgreben/yeetgif/pkg/floats"
	"github.com/sgreben/yeetgif/pkg/piecewiselinear"

	"github.com/sgreben/yeetgif/pkg/gifcmd"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandZoom(cmd *cli.Cmd) {
	cmd.Before = ProcessInput
	cmd.Spec = "[OPTIONS] [ZOOM_LEVELS_CSV]"
	var (
		from = gifcmd.Float{Value: 1.0}
		to   = gifcmd.Float{Value: 1.5}
		c    = gifcmd.FloatsCSV{}
	)
	cmd.VarArg("ZOOM_LEVELS_CSV", &c, "e.g. 1.0,1.5 to zoom from 100% to 150%")
	cmd.VarOpt("0 from", &from, "(deprecated)")
	cmd.VarOpt("1 to", &to, "(deprecated)")
	cmd.VarArg("c custom", &c, "(deprecated)")
	cmd.Action = func() {
		var f func(float64) float64
		switch {
		case len(c.Texts) > 0:
			customF := piecewiselinear.Function{Y: c.Values}
			customF.X = floats.MakeSpan(0, 1, len(customF.Y))
			f = customF.At
		default:
			f = func(t float64) float64 {
				return from.Value*(1-t) + to.Value*t
			}
		}
		Zoom(images, f)
	}
}

// Zoom `images` once from `from` to `to`
func Zoom(images []image.Image, f func(float64) float64) {
	n := len(images)
	scale := func(i int) {
		t := float64(i) / float64(n)
		scale := f(t)
		bPre := images[i].Bounds()
		width := float64(bPre.Dx()) * scale
		height := float64(bPre.Dy()) * scale
		images[i] = imaging.Resize(images[i], int(width), int(height), imaging.Gaussian)
		bPost := images[i].Bounds()
		offset := image.Point{
			X: (bPost.Dx() - bPre.Dx()) / 2,
			Y: (bPost.Dy() - bPre.Dy()) / 2,
		}
		bPre.Min = bPre.Min.Add(offset)
		bPre.Max = bPre.Max.Add(offset)
		images[i] = imaging.Crop(images[i], bPre)
	}
	parallel(len(images), scale)
}
