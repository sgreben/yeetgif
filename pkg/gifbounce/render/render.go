package render

import (
	"image"
	"image/color"

	"github.com/sgreben/yeetgif/pkg/gifbounce"
	"github.com/sgreben/yeetgif/pkg/imaging"
)

func World(w *gifbounce.World, bounds image.Rectangle) []image.Image {
	worker := w.Worker
	out := make([]image.Image, w.NumFrames)
	numStaticThings := len(w.Things.Static)
	numDynamicThings := len(w.Things.Dynamic)

	type frame []struct {
		Image image.Image
		Point image.Point
	}
	rendered := make([]frame, w.NumFrames)
	for i := range rendered {
		rendered[i] = make(frame, numStaticThings+numDynamicThings)
	}
	blank := imaging.New(1, 1, color.Transparent)
	renderThing := func(t *gifbounce.Thing, i int) func(int) {
		rec := t.Recording
		return func(j int) {
			out := &rendered[j][i]
			var img image.Image
			if !rec.Active[j] {
				img = blank
			} else {
				img = imaging.RotateAbout(
					t.Images[rec.Frames[j]],
					gifbounce.ToImagePoint(&rec.LocalCenters[j]),
					rec.Angles[j],
					color.Transparent,
				)
				img = imaging.Crop(img, imaging.OpaqueBounds(img, 0))
				out.Point = gifbounce.ToImagePoint(&rec.WorldCenters[j]).
					Sub(imaging.AnchorPoint(img, imaging.Center))
			}
			out.Image = img
		}
	}
	worker(numStaticThings, func(j int) {
		worker(w.NumFrames, renderThing(w.Things.Static[j], j))
	}, "render-static")
	worker(numDynamicThings, func(j int) {
		worker(w.NumFrames, renderThing(w.Things.Dynamic[j], numStaticThings+j))
	}, "render-dynamic")
	worker(w.NumFrames, func(j int) {
		out[j] = imaging.OverlayOnCanvas(bounds.Dx(), bounds.Dy(), color.Transparent, rendered[j])
	}, "compose")
	return out
}
