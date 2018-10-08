package main

import (
	"image"

	"github.com/sgreben/yeetgif/pkg/gifcmd"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandZoom(cmd *cli.Cmd) {
	cmd.Spec = "[OPTIONS]"
	var (
		from = gifcmd.Float{Value: 1.0}
		to   = gifcmd.Float{Value: 1.5}
	)
	cmd.VarOpt("0 from", &from, "")
	cmd.VarOpt("1 to", &to, "")
	cmd.Action = func() {
		Zoom(images, from.Value, to.Value)
	}
}

// Zoom `images` once from `from` to `to`
func Zoom(images []image.Image, from, to float64) {
	n := len(images)
	scale := func(i int) {
		weight := float64(i) / float64(n)
		scale := from*(1-weight) + to*weight
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
