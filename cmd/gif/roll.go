package main

import (
	"image"
	"image/color"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandRoll(cmd *cli.Cmd) {
	cmd.Before = InputAndDuplicate
	cmd.Spec = "[OPTIONS]"
	var (
		r = gifcmd.Float{Value: 1.0}
		s = gifcmd.Float{Value: 1.0}
		p = gifcmd.Float{Value: 0.0}
		c = gifcmd.Float{Value: 1.0}
	)
	cmd.VarOpt("r revolutions", &r, "")
	cmd.VarOpt("s scale", &s, "")
	cmd.VarOpt("p phase", &p, "")
	cmd.VarOpt("c crop-scale", &c, "")
	cmd.Action = func() {
		var cropScale *float64
		if c.Text != "" {
			cropScale = &c.Value
		}
		Roll(images, r.Value, s.Value, p.Value, cropScale)
	}
}

// Roll the `images` `rev` times
func Roll(images []image.Image, rev, preScale, phase float64, cropScale *float64) {
	n := len(images)
	rotate := func(i int) {
		angle := 360*rev*float64(i)/float64(n) + phase*360
		bPre := images[i].Bounds()
		if preScale != 1.0 {
			images[i] = imaging.Resize(images[i], int(float64(bPre.Dx())*preScale), int(float64(bPre.Dy())*preScale), imaging.Lanczos)
		}
		images[i] = imaging.Rotate(images[i], angle, color.Transparent)
		bPost := images[i].Bounds()
		offset := image.Point{
			X: (bPost.Dx() - bPre.Dx()) / 2,
			Y: (bPost.Dy() - bPre.Dy()) / 2,
		}
		bPre.Min = bPre.Min.Add(offset)
		bPre.Max = bPre.Max.Add(offset)
		switch {
		case cropScale != nil:
			w := int(float64(bPre.Dx()) * *cropScale)
			h := int(float64(bPre.Dy()) * *cropScale)
			images[i] = imaging.CropCenter(images[i], w, h)
		case !config.Pad:
			images[i] = imaging.Crop(images[i], bPre)
		}
	}
	parallel(len(images), rotate, "roll")
}
