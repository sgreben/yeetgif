package main

import (
	"image"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandCrop(cmd *cli.Cmd) {
	cmd.Before = Input
	cmd.Spec = "[OPTIONS]"
	var (
		t = gifcmd.Float{Value: 0.0}
	)
	cmd.VarOpt("t threshold", &t, "")
	cmd.Action = func() {
		thresholdUint8 := uint8(t.Value * 0xFF)
		AutoCrop(images, thresholdUint8)
	}
}

func AutoCrop(images []image.Image, threshold uint8) {
	var sample image.Image
	sample = images[0]
	for i := range images {
		sample = imaging.OverlayWithOp(sample, images[i], image.ZP, imaging.OpMaxAlpha)
	}
	b := imaging.OpaqueBounds(sample, threshold)
	crop := func(i int) {
		images[i] = imaging.Crop(images[i], b)
	}
	parallel(len(images), crop)
}
