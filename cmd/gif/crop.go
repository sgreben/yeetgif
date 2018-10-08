package main

import (
	"image"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandCrop(cmd *cli.Cmd) {
	cmd.Spec = "[OPTIONS]"
	var (
		t = gifcmd.Float{Value: 0.0}
	)
	cmd.VarOpt("t threshold", &t, "")
	cmd.Action = func() {
		AutoCrop(images, t.Value)
	}
}

func AutoCrop(images []image.Image, threshold float64) {
	width, height := 0, 0
	for i := range images {
		if w := images[i].Bounds().Dx(); w > width {
			width = w
		}
		if h := images[i].Bounds().Dy(); h > height {
			height = h
		}
	}
	sample := imaging.Clone(images[0])
	for i := range images {
		sample = imaging.OverlayWithOp(sample, images[i], image.ZP, imaging.OpMaxAlpha)
	}
	b := imaging.OpaqueBounds(sample, uint8(threshold*255))
	crop := func(i int) {
		images[i] = imaging.Crop(images[i], b)
	}
	parallel(len(images), crop)
}
