package main

import (
	"image"
	"image/color"
	"sync"

	"github.com/sgreben/yeetgif/pkg/imaging"

	_ "image/jpeg"
)

func Pad(images []image.Image) {
	b := Bounds(images)
	width, height := b.Dx(), b.Dy()
	pad := func(i int) {
		if images[i].Bounds() != b {
			padded := imaging.New(width, height, color.Transparent)
			images[i] = imaging.PasteCenter(padded, images[i])
		}
	}
	parallel(len(images), pad)
}

func Bounds(images []image.Image) (out image.Rectangle) {
	var mu sync.Mutex
	bound := func(i int) {
		b := images[i].Bounds()
		b = b.Sub(b.Min)
		mu.Lock()
		out = out.Union(b)
		mu.Unlock()
	}
	parallel(len(images), bound)
	return
}

func OpaqueBounds(images []image.Image, threshold uint8) (out image.Rectangle) {
	var mu sync.Mutex
	bound := func(i int) {
		b := imaging.OpaqueBounds(images[i], threshold)
		mu.Lock()
		out = out.Union(b)
		mu.Unlock()
	}
	parallel(len(images), bound)
	return
}
