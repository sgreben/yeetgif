package main

import (
	"image"
	"image/color"

	"github.com/sgreben/yeetgif/pkg/imaging"

	"github.com/fogleman/gg"

	"github.com/sgreben/yeetgif/pkg/piecewiselinear"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
)

func CommandNPC(cmd *cli.Cmd) {
	cmd.Before = ProcessInput
	cmd.Spec = "[OPTIONS]"
	const (
		bgSolid = "solid"
		bgBlur  = "blur"
	)
	var (
		x           = gifcmd.FloatsCSV{Values: []float64{0.5}}
		y           = gifcmd.FloatsCSV{Values: []float64{0.5}}
		scale       = gifcmd.FloatsCSV{Values: []float64{1.0}}
		eyeScale    = gifcmd.FloatsCSV{Values: []float64{1.0}}
		noseScale   = gifcmd.FloatsCSV{Values: []float64{1.0}}
		mouthScaleW = gifcmd.FloatsCSV{Values: []float64{1.0}}
		mouthScaleH = gifcmd.FloatsCSV{Values: []float64{0.0}}
		scaleW      = gifcmd.FloatsCSV{Values: []float64{1.0}}
		scaleH      = gifcmd.FloatsCSV{Values: []float64{1.0}}
		angle       = gifcmd.FloatsCSV{Values: []float64{0}}
		alpha       = gifcmd.FloatsCSV{Values: []float64{1}}
		bgtype      = gifcmd.Enum{Choices: []string{bgSolid, bgBlur}, Value: bgSolid}
	)
	cmd.VarOpt("x", &x, "")
	cmd.VarOpt("y", &y, "")
	cmd.VarOpt("bg", &bgtype, bgtype.Help())
	cmd.VarOpt("s scale", &scale, "")
	cmd.VarOpt("scale-x", &scaleW, "")
	cmd.VarOpt("scale-y", &scaleH, "")
	cmd.VarOpt("eye-scale", &eyeScale, "")
	cmd.VarOpt("nose-scale", &noseScale, "")
	cmd.VarOpt("mouth-scale-x", &mouthScaleW, "")
	cmd.VarOpt("mouth-scale-y", &mouthScaleH, "")
	cmd.VarOpt("r angle", &angle, "")
	cmd.VarOpt("a alpha", &alpha, "")
	cmd.Action = func() {
		npcGrey := color.RGBA{R: 0xA3, G: 0xA3, B: 0xA3, A: 0xFF}
		npcBlack := color.RGBA{R: 0x05, G: 0x05, B: 0x05, A: 0xFF}
		fx := piecewiselinear.Function{Y: x.Values}
		fx.SetDomain(0.0, 1.0)
		fy := piecewiselinear.Function{Y: y.Values}
		fy.SetDomain(0.0, 1.0)
		fscale := piecewiselinear.Function{Y: scale.Values}
		fscale.SetDomain(0.0, 1.0)
		fscaleW := piecewiselinear.Function{Y: scaleW.Values}
		fscaleW.SetDomain(0.0, 1.0)
		fscaleH := piecewiselinear.Function{Y: scaleH.Values}
		fscaleH.SetDomain(0.0, 1.0)
		feyeScale := piecewiselinear.Function{Y: eyeScale.Values}
		feyeScale.SetDomain(0.0, 1.0)
		fnoseScale := piecewiselinear.Function{Y: noseScale.Values}
		fnoseScale.SetDomain(0.0, 1.0)
		fmouthScaleW := piecewiselinear.Function{Y: mouthScaleW.Values}
		fmouthScaleW.SetDomain(0.0, 1.0)
		fmouthScaleH := piecewiselinear.Function{Y: mouthScaleH.Values}
		fmouthScaleH.SetDomain(0.0, 1.0)
		fangle := piecewiselinear.Function{Y: angle.Values}
		fangle.SetDomain(0.0, 1.0)
		falpha := piecewiselinear.Function{Y: alpha.Values}
		falpha.SetDomain(0.0, 1.0)
		NPC(images, npcGrey, npcBlack, fx.At, fy.At, fscale.At, fscaleW.At, fscaleH.At, fangle.At, falpha.At, feyeScale.At, fnoseScale.At, fmouthScaleW.At, fmouthScaleH.At, bgtype.Value == bgBlur)
	}
}

func NPC(images []image.Image, bg, fg color.RGBA, fx, fy, fscale, fscalew, fscaleh, fangle, falpha, feye, fnose, fmouthw, fmouthh func(float64) float64, blur bool) {
	n := float64(len(images))
	eyeDistance := 50
	eyeRadius := 5
	npc := func(i int) {
		t := float64(i) / n
		x := fx(t)
		y := fy(t)
		scale := fscale(t)
		angle := fangle(t)
		alpha := falpha(t)
		eyeDistance := int(float64(eyeDistance) * scale)
		eyeRadius := int(float64(eyeRadius) * scale)
		lineWidth := float64(eyeRadius*4) / 3
		noseHeight := eyeDistance * 7 / 10
		noseMouthDistance := eyeDistance * 8 / 20
		faceOffset := noseMouthDistance / 2

		eyeLeft := image.Point{eyeRadius * 2, eyeRadius * 2}
		eyeRight := eyeLeft.Add(image.Point{X: eyeDistance})

		noseScale := fnose(t)
		noseCorner := image.Point{X: int(noseScale * float64(-eyeDistance) * 2 / 10), Y: noseHeight}
		noseTop := eyeLeft.Add(image.Point{X: int(float64(eyeDistance) * (5 - noseScale*2) / 10), Y: eyeRadius})
		noseLeft := noseTop.Add(noseCorner)
		noseRight := noseLeft.Add(image.Point{X: int(noseScale * float64(eyeDistance) * 7 / 10)})

		mouthScaleW := fmouthw(t)
		mouthLeft := eyeLeft.Add(image.Point{X: (eyeRadius / 2) + int(float64(eyeDistance-eyeRadius/2)*(1-mouthScaleW)), Y: noseHeight + noseMouthDistance})
		mouthRight := eyeRight.Add(image.Point{X: (-eyeRadius / 2) - int(float64(eyeDistance-eyeRadius/2)*(1-mouthScaleW)), Y: noseHeight + noseMouthDistance})
		points := []*image.Point{&eyeLeft, &eyeRight, &noseTop, &noseLeft, &noseRight, &mouthLeft, &mouthRight}
		scaleW := fscalew(t)
		scaleH := fscaleh(t)
		faceW := int(scaleW * 21 * float64(eyeRight.X-eyeLeft.X) / 10)
		faceH := int(scaleH * 22 * float64(mouthLeft.Y-eyeLeft.Y) / 10)
		npcW, npcH := int(float64(faceW)), int(float64(faceH))
		ctx := gg.NewContext(npcW, npcH)
		for j := range points {
			points[j].X = int(float64(points[j].X) * scaleW)
			points[j].Y = int(float64(points[j].Y) * scaleH)
		}
		midpoint := image.Point{
			X: (eyeRight.X + eyeLeft.X) / 2,
			Y: (mouthLeft.Y + eyeLeft.Y) / 2,
		}
		delta := midpoint.Sub(image.Point{X: npcW / 2, Y: npcH/2 + faceOffset})
		for j := range points {
			points[j].X -= delta.X
			points[j].Y -= delta.Y
		}
		midpoint.X -= delta.X
		midpoint.Y -= delta.Y

		if !blur {
			bg.A = uint8(255 * alpha)
		}
		ctx.SetColor(bg)
		ctx.DrawEllipse(float64(midpoint.X), float64(midpoint.Y-faceOffset), float64(faceW)/2, float64(faceH)/2)
		ctx.Fill()
		mask := ctx.AsMask()
		if blur {
			ctx.SetColor(color.Transparent)
			ctx.Clear()
		}

		ctx.SetColor(fg)
		eyeScale := feye(t)
		ctx.DrawPoint(float64(eyeLeft.X), float64(eyeLeft.Y), float64(eyeRadius)*eyeScale)
		ctx.Fill()
		ctx.DrawPoint(float64(eyeRight.X), float64(eyeRight.Y), float64(eyeRadius)*eyeScale)
		ctx.Fill()

		ctx.SetLineWidth(lineWidth)
		ctx.SetColor(fg)
		ctx.DrawLine(float64(noseTop.X), float64(noseTop.Y), float64(noseLeft.X), float64(noseLeft.Y))
		ctx.Stroke()
		ctx.SetColor(fg)
		ctx.DrawLine(float64(noseLeft.X), float64(noseLeft.Y), float64(noseRight.X), float64(noseRight.Y))
		ctx.Stroke()
		ctx.SetColor(fg)

		mouthScaleH := fmouthh(t)
		ctx.DrawRectangle(float64(mouthLeft.X), float64(mouthLeft.Y), float64(mouthRight.X-mouthLeft.X), lineWidth*(mouthScaleH))
		ctx.DrawLine(float64(mouthLeft.X), float64(mouthLeft.Y), float64(mouthRight.X), float64(mouthRight.Y))
		ctx.StrokePreserve()
		ctx.Fill()
		npcImage := imaging.Rotate(ctx.Image(), angle, color.Transparent)
		b := images[i].Bounds()
		w, h := b.Dx(), b.Dy()
		bNPC := npcImage.Bounds()
		pos := image.Point{X: int(x*float64(w)) - bNPC.Dx()/2, Y: int(y*float64(h)) - bNPC.Dy()/2}
		if blur {
			sigma := 10.0
			blurred := imaging.Crop(images[i], bNPC.Add(pos))
			maskRotated := imaging.Rotate(mask, angle, color.Transparent)
			blurred = imaging.Blur(blurred, sigma)
			blurred = imaging.OverlayWithOp(blurred, maskRotated, image.ZP, imaging.OpMinAlpha)
			images[i] = imaging.Overlay(images[i], blurred, pos, 1.0)
		}
		images[i] = imaging.Overlay(images[i], npcImage, pos, 1.0)
	}
	parallel(len(images), npc)
}
