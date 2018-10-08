package main

import (
	"image"
	"image/color"
	"math/rand"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandCrowd(cmd *cli.Cmd) {
	var (
		n   = cmd.IntOpt("n", 3, "crowd size")
		rpx = gifcmd.Float{Value: 0.5}
		rpy = gifcmd.Float{Value: 0.25}
		rs  = gifcmd.Float{Value: 0.25}
		rr  = gifcmd.Float{Value: 0.1}
		ra  = gifcmd.Float{Value: 0.0}
		ro  = gifcmd.Float{Value: 1.0}
		rf  = cmd.BoolOpt("flip", true, "🌀 flip")
	)
	cmd.VarOpt("x", &rpx, "🌀 x")
	cmd.VarOpt("y", &rpy, "🌀 y")
	cmd.VarOpt("s scale", &rs, "🌀 [0.0,1.0]")
	cmd.VarOpt("r rotate", &rr, "🌀 [0.0,1.0]")
	cmd.VarOpt("a alpha", &ra, "🌀 [0.0,1.0]")
	cmd.VarOpt("o offset", &ro, "🌀 [0.0,1.0]")
	cmd.Action = func() {
		Crowd(images, *n, *rf, rpx.Value, rpy.Value, rs.Value, ra.Value, rr.Value, ro.Value)
		AutoCrop(images, 0.0)
	}
}

func Crowd(images []image.Image, k int, rf bool, rpx, rpy, rs, ra, rr, ro float64) {
	p := make([]image.Point, k)
	r := make([]float64, k)
	s := make([]float64, k)
	a := make([]float64, k)
	o := make([]int, k)
	f := make([]bool, k)
	for j := range s {
		s[j] = 1.0 - (rand.Float64() * rs)
		r[j] = 360 * rr * 2 * (rand.Float64() - 0.5)
		o[j] = rand.Intn(int(ro * float64(len(images)-1)))
		a[j] = 1.0 - (rand.Float64() * ra)
		f[j] = rf && (rand.Float32() < 0.5)
	}
	width, height := 0.0, 0.0
	for i := range images {
		for j := range r {
			tmp := images[i]
			if f[j] {
				tmp = imaging.FlipH(tmp)
			}
			tmp = imaging.Rotate(tmp, r[j], color.Transparent)
			b := tmp.Bounds()
			if w := float64(b.Dx()) * s[j]; w > width {
				width = w
			}
			if h := float64(b.Dy()) * s[j]; h > height {
				height = h
			}
		}
	}
	mid := image.Point{
		X: int(width / 2),
		Y: int(height / 2),
	}
	var b, bOriginal image.Rectangle
	bOriginal.Max.X = int(width)
	bOriginal.Max.Y = int(height)
	b.Max.X = bOriginal.Max.X
	b.Max.Y = bOriginal.Max.Y
	for j := range p {
		p[j].X = int(s[j] * float64(width) * rpx * 2 * (rand.Float64() - 0.5))
		p[j].Y = int(s[j] * float64(height) * rpy * 2 * (rand.Float64() - 0.5))
		b = b.Union(bOriginal.Add(p[j]))
	}
	offset := b.Min
	b = b.Sub(offset)
	originals := images
	images = make([]image.Image, len(originals))
	crowd := func(i int) {
		crowded := imaging.New(b.Dx(), b.Dy(), color.Transparent)
		for j := range p {
			iLayer := (o[j] + i) % len(images)
			layer := originals[iLayer]
			bLayer := layer.Bounds()
			w, h := float64(bLayer.Dx())*s[j], float64(bLayer.Dy())*s[j]
			if f[j] {
				layer = imaging.FlipH(layer)
			}
			layer = imaging.Resize(layer, int(w), int(h), imaging.Lanczos)
			layer = imaging.Rotate(layer, r[j], color.Transparent)
			midLayer := imaging.AnchorPoint(layer, imaging.Center)
			p := p[j].Sub(midLayer).Add(mid).Sub(offset)
			crowded = imaging.Overlay(crowded, layer, p, a[j])
		}
		images[i] = crowded
	}
	overwrite := func(i int) {
		originals[i] = images[i]
	}
	parallel(len(images), crowd)
	parallel(len(images), overwrite)
}
