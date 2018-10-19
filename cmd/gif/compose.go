package main

import (
	"image"
	"log"
	"math/big"
	"os"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func CommandCompose(cmd *cli.Cmd) {
	cmd.Before = ProcessInput
	cmd.Spec = "[OPTIONS] INPUT"
	const (
		orderUnder = "under"
		orderOver  = "over"
	)
	const (
		positionCenter   = "center"
		positionLeft     = "left"
		positionRight    = "right"
		positionTop      = "top"
		positionBottom   = "bottom"
		positionAbsolute = "abs"
	)
	var (
		input = cmd.StringArg("INPUT", "", "")
		z     = gifcmd.Enum{
			Choices: []string{orderUnder, orderOver},
			Value:   orderOver,
		}
		p = gifcmd.Enum{
			Choices: []string{
				positionCenter,
				positionLeft,
				positionRight,
				positionTop,
				positionBottom,
				positionAbsolute,
			},
			Value: positionCenter,
		}
		x = cmd.IntOpt("x", 0, "")
		y = cmd.IntOpt("y", 0, "")
		s = gifcmd.Float{Value: 1.0}
	)
	cmd.VarOpt("z z-order", &z, z.Help())
	cmd.VarOpt("p position", &p, p.Help())
	cmd.VarOpt("s scale", &s, "")
	cmd.Action = func() {
		f, err := os.Open(*input)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		layer := Decode(f)
		if s.Value != 1.0 {
			ResizeScale(layer, s.Value)
		}
		offset := image.Point{*x, *y}
		var imageAnchor, layerAnchor imaging.Anchor
		switch p.Value {
		case positionAbsolute:
			imageAnchor = imaging.TopLeft
			layerAnchor = imaging.Center
		case positionCenter:
			imageAnchor = imaging.Center
			layerAnchor = imaging.Center
		case positionLeft:
			imageAnchor = imaging.Left
			layerAnchor = imaging.Right
		case positionRight:
			imageAnchor = imaging.Right
			layerAnchor = imaging.Left
		case positionTop:
			imageAnchor = imaging.Top
			layerAnchor = imaging.Bottom
		case positionBottom:
			imageAnchor = imaging.Bottom
			layerAnchor = imaging.Top
		}
		switch z.Value {
		case orderOver:
			Compose(images, layer, offset, imageAnchor, layerAnchor)
		case orderUnder:
			Compose(layer, images, offset, layerAnchor, imageAnchor)
		}
	}
}

func Compose(a, b []image.Image, p image.Point, anchorA, anchorB imaging.Anchor) {
	if len(a) == 0 || len(b) == 0 {
		return
	}
	compose := func(i int) {
		ai := i % len(a)
		bi := i % len(b)
		under := a[ai]
		over := b[bi]
		overOffset := imaging.AnchorPoint(under, anchorA).Sub(imaging.AnchorPoint(over, anchorB)).Add(p)
		bounds := under.Bounds().Union(over.Bounds().Add(overOffset))
		bg := image.NewNRGBA(bounds.Sub(bounds.Min))
		bg = imaging.Paste(bg, under, bounds.Min.Mul(-1))
		images[i] = imaging.Overlay(bg, over, overOffset.Sub(bounds.Min), 1.0)
	}
	var an, bn, z big.Int
	an.SetInt64(int64(len(a)))
	bn.SetInt64(int64(len(b)))
	n := int(z.Mul(z.Div(&bn, z.GCD(nil, nil, &an, &bn)), &an).Int64())
	images = make([]image.Image, n)
	parallel(n, compose)
}
