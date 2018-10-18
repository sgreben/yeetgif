package main

import (
	"image"
	"math"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandResize(cmd *cli.Cmd) {
	cmd.Before = Input
	cmd.Spec = "[OPTIONS]"
	var (
		s = gifcmd.Float{Value: 1.0}
		w = gifcmd.Float{Value: 0}
		h = gifcmd.Float{Value: 0}
	)
	cmd.VarOpt("s scale", &s, "")
	cmd.VarOpt("x width", &w, "width (pixels)")
	cmd.VarOpt("y height", &h, "height (pixels)")
	cmd.Action = func() {
		if s.Text != "" {
			ResizeScale(images, s.Value)
			return
		}
		ResizeTarget(images, w.Value, h.Value)
	}
}

func ResizeScale1(img image.Image, scale float64) image.Image {
	b := img.Bounds()
	width, height := float64(b.Dx()), float64(b.Dy())
	return imaging.Resize(img, int(width*scale), int(height*scale), imaging.Lanczos)
}

// ResizeScale resizes by a factor
func ResizeScale(images []image.Image, scale float64) {
	resize := func(i int) { images[i] = ResizeScale1(images[i], scale) }
	parallel(len(images), resize)
}

// ResizeTarget resizes to fit in the given bounds
func ResizeTarget(images []image.Image, width, height float64) {
	resize := func(i int) {
		b := images[i].Bounds()
		w, h := float64(b.Dx()), float64(b.Dy())
		scale := 1.0
		rw, rh := math.Abs(width/w-1), math.Abs(height/w-1)
		switch {
		case width > 0 && height > 0:
			if rw < rh {
				scale = width / w
			} else {
				scale = height / h
			}
		case width > 0:
			scale = width / w
		case height > 0:
			scale = height / h
		}
		images[i] = imaging.Resize(images[i], int(w*scale), int(h*scale), imaging.Lanczos)
	}
	parallel(len(images), resize)
}
