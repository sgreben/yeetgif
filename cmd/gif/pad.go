package main

import (
	"image"
	"image/color"

	"github.com/sgreben/yeetgif/pkg/imaging"

	_ "image/jpeg"
)

func Pad(images []image.Image) {
	width, height := 0, 0
	for i := range images {
		if w := images[i].Bounds().Dx(); w > width {
			width = w
		}
		if h := images[i].Bounds().Dy(); h > height {
			height = h
		}
	}
	pad := func(i int) {
		padded := imaging.New(width, height, color.Transparent)
		images[i] = imaging.PasteCenter(padded, images[i])
	}
	parallel(len(images), pad)
}
