package main

import (
	"image"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/sgreben/yeetgif/pkg/box2d"
	"github.com/sgreben/yeetgif/pkg/gifbounce/render"
	"github.com/sgreben/yeetgif/pkg/imaging"

	"github.com/sgreben/yeetgif/pkg/gifbounce"

	"github.com/sgreben/yeetgif/pkg/gifmath"

	cli "github.com/jawher/mow.cli"
	"github.com/sgreben/yeetgif/pkg/gifcmd"
)

type BoundingPolygonParams struct {
	NumVertices    int
	AlphaThreshold uint8
}

func CommandRain(cmd *cli.Cmd) {
	cmd.Spec = "[OPTIONS] [INPUT...] [OPTIONS]"
	cmd.Before = Input
	var (
		inputs                     = cmd.StringsArg("INPUT", nil, "")
		boxSize                    = gifcmd.Float{Value: 400}
		gravity                    = gifcmd.Float{Value: 900}
		bounciness                 = gifcmd.FloatsCSV{Values: []float64{0.3}}
		alphaThreshold             = gifcmd.Float{Value: 0.0625}
		wrapOverlapThreshold       = gifcmd.Float{Value: 48.0}
		wrapOverlapFramesThreshold = gifcmd.Float{Value: 24}
		angularVelocity            = gifcmd.Float{Value: 2.0}
		velocity                   = gifcmd.Float{Value: 300.0}
		timeStep                   = gifcmd.FloatsCSV{Values: []float64{1.0}}
		lengthFactor               = gifcmd.Float{Value: 1.0}
		density                    = gifcmd.Float{Value: 3.0 / 4.0}
		x                          = gifcmd.Float{Value: 0.5}
		y                          = gifcmd.Float{Value: 0.5}
		defaultLength              = 90.0
	)
	cmd.VarOpt("d density", &density, "")
	cmd.VarOpt("b bounciness", &bounciness, "")
	cmd.VarOpt("g gravity", &gravity, "")
	cmd.VarOpt("s size", &boxSize, "")
	cmd.VarOpt("a bounds-alpha-threshold", &alphaThreshold, "")
	cmd.VarOpt("v initial-linear-velocity", &velocity, "")
	cmd.VarOpt("l animation-length-factor", &lengthFactor, "")
	cmd.VarOpt("x static-x", &x, "")
	cmd.VarOpt("y static-y", &y, "")
	cmd.VarOpt("wrap-max-overlap", &wrapOverlapThreshold, "")
	cmd.VarOpt("wrap-max-overlap-frames", &wrapOverlapFramesThreshold, "")
	cmd.VarOpt("initial-angular-velocity", &angularVelocity, "")
	numVertices := cmd.IntOpt("bounds-points", 16, "")
	cmd.Action = func() {
		var p gifbounce.Params
		var bp BoundingPolygonParams
		p.Worker = parallel
		lowFrictionF := func(float64) float64 { return 0 }
		highFrictionF := func(float64) float64 { return 0.3 }
		alphaThresholdUint8 := uint8(0xFF * alphaThreshold.Value)
		bp.NumVertices = *numVertices
		bp.AlphaThreshold = alphaThresholdUint8
		p.Gravity = gravity.Value
		p.Things.Walls.Distance = boxSize.Value
		var imagesDynamic [][]image.Image
		bouncinessF := bounciness.PiecewiseLinear(0, 1)
		if len(*inputs) > 0 {
			staticThing := &gifbounce.ThingParams{
				Images:     images,
				Bounciness: bouncinessF,
				Friction:   lowFrictionF,
			}
			b := Bounds(staticThing.Images)
			staticThing.Initial.Position = gifbounce.FromImagePoint(&image.Point{
				X: int(boxSize.Value*x.Value) - b.Dx()/2,
				Y: int(boxSize.Value*y.Value) - b.Dy()/2,
			})
			p.Things.Static = append(p.Things.Static, staticThing)
		}
		if len(*inputs) == 0 {
			p.Things.Dynamic = []*gifbounce.ThingParams{
				{
					Images:     images,
					Bounciness: bouncinessF,
					Friction:   lowFrictionF,
				},
			}
		} else {
			p.Things.Dynamic = make([]*gifbounce.ThingParams, len(*inputs))
			loadDynamicThing := func(i int) {
				f, err := os.Open((*inputs)[i])
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()
				p.Things.Dynamic[i] = &gifbounce.ThingParams{
					Images: Decode(f),
				}
			}
			parallel(len(*inputs), loadDynamicThing)
		}
		for i := range p.Things.Dynamic {
			p.Things.Dynamic[i].Bounciness = bouncinessF
			p.Things.Dynamic[i].Friction = highFrictionF
			p.Things.Dynamic[i].LinearDamping = func(float64) float64 {
				return 0.3
			}
			p.Things.Dynamic[i].AngularDamping = func(float64) float64 {
				return 0.3
			}
		}
		p.NumFrames = int(lengthFactor.Value * defaultLength)
		lcm := len(images)
		for _, v := range imagesDynamic {
			lcm = gifmath.LCM(lcm, len(v))
		}
		p.NumFrames = int(math.Round(float64(p.NumFrames)/float64(lcm))) * lcm
		if p.NumFrames < lcm {
			p.NumFrames = lcm
		}
		timeStepF := timeStep.PiecewiseLinear(0, 1)
		p.Solver.TimeStep = func(t float64) float64 {
			return timeStepF(t) * config.DelayMilliseconds(t) / 1000.0
		}
		p.Solver.VelocityIterations = 8
		p.Solver.PositionIterations = 4
		GenerateBoundingPolygons(&p, &bp)
		RainPlaceThings(&p, &bp, velocity.Value, angularVelocity.Value, density.Value)
		Rain(&p, lcm, wrapOverlapThreshold.Value, int(wrapOverlapFramesThreshold.Value))
	}
}

func RainPlaceThings(p *gifbounce.Params, bp *BoundingPolygonParams, velocity, angularVelocity, density float64) {
	areaStatic := make([]int, len(p.Things.Static))
	parallel(len(p.Things.Static), func(i int) {
		areaStatic[i] = imaging.OpaqueArea(p.Things.Static[i].Images[0], bp.AlphaThreshold)
	})
	areaDynamic := make([]int, len(p.Things.Dynamic))
	boundsDynamic := make([]image.Rectangle, len(p.Things.Dynamic))
	parallel(len(p.Things.Dynamic), func(i int) {
		areaDynamic[i] = imaging.OpaqueArea(p.Things.Dynamic[i].Images[0], bp.AlphaThreshold)
		boundsDynamic[i] = Bounds(p.Things.Dynamic[i].Images)
	})
	w, h := p.Things.Walls.Distance, p.Things.Walls.Distance
	totalArea := w * h
	for _, a := range areaStatic {
		totalArea -= float64(a)
	}
	n := 0
	targetArea := totalArea * (1 - density)
	for {
		area := areaDynamic[n%len(p.Things.Dynamic)]
		totalArea -= float64(area)
		n++
		if totalArea < targetArea {
			break
		}
	}

	dynamicThings := make([]*gifbounce.ThingParams, n)
	for j := range dynamicThings {
		i := j % len(p.Things.Dynamic)
		thingCopy := *p.Things.Dynamic[i]
		dynamicThings[j] = &thingCopy
	}

	batchSize := 3.0
	for j, thing := range dynamicThings {
		i := j % len(p.Things.Dynamic)
		dy := boundsDynamic[i].Dy()
		dx := boundsDynamic[i].Dx()
		thing.Initial.Position = gifbounce.FromImagePoint(&image.Point{
			X: int(rand.Float64()*w) - dx/2,
			Y: int(-(2 + rand.Float64()) * float64(dy)),
		})
		thing.Initial.LinearVelocity = gifbounce.FromImagePoint(&image.Point{
			Y: int(velocity),
		})
		thing.Initial.AngularVelocityDeg = 2 * (rand.Float64() - 0.5) * angularVelocity
		batchIndex := float64(j) / batchSize
		thing.Initial.Time = ((batchIndex + rand.Float64()) * batchSize) / float64(n-1)
	}
	p.Things.Dynamic = dynamicThings
}

func Rain(p *gifbounce.Params, lcm int, wrapOverlapThreshold float64, wrapOverlapFramesThreshold int) {
	world := p.New()
	simulate := func(j int) {
		t := float64(j) / float64(p.NumFrames-1)
		world.Step(t)
	}
	sequential(p.NumFrames, simulate, "simulate")
	renderBounds := image.Rectangle{
		Min: box2d.Point{X: 0, Y: 0}.ImagePoint(),
		Max: box2d.Point{X: p.Things.Walls.Distance, Y: p.Things.Walls.Distance}.ImagePoint(),
	}
	worldBounds := gifbounce.FromImageRect(&renderBounds)
	for world.ContainsDynamicThings(worldBounds) {
		world.NumFrames += lcm
		for i := 0; i < lcm; i++ {
			world.Step(1.0)
		}
	}
	render.WrapV(world, worldBounds, wrapOverlapThreshold, wrapOverlapFramesThreshold, lcm)
	images = render.World(world, renderBounds)
}

func GenerateBoundingPolygons(p *gifbounce.Params, bp *BoundingPolygonParams) {
	generateOutlines := func(t *gifbounce.ThingParams) {
		polygons := make([][]image.Point, len(t.Images))
		generateOutline := func(i int) {
			polygons[i] = imaging.OpaquePolygon(t.Images[i], bp.NumVertices, bp.AlphaThreshold)
		}
		parallel(len(t.Images), generateOutline)
		t.Polygons = gifbounce.FromImagePolygons(polygons)
	}
	parallel(len(p.Things.Dynamic)+len(p.Things.Static), func(i int) {
		if i >= len(p.Things.Dynamic) {
			generateOutlines(p.Things.Static[i-len(p.Things.Dynamic)])
			return
		}
		generateOutlines(p.Things.Dynamic[i])
	}, "outline")
}
