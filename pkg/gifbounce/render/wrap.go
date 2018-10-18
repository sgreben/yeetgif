package render

import (
	"github.com/sgreben/yeetgif/pkg/box2d"

	"github.com/sgreben/yeetgif/pkg/gifbounce"
)

func suffixUnion(x []box2d.AABB) {
	n := len(x)
	union := x[n-1]
	for i := n - 1; i >= 0; i-- {
		x[i].CombineInPlace(union)
		union = x[i]
	}
}

func prefixUnion(x []box2d.AABB) {
	n := len(x)
	union := x[0]
	for i := 0; i < n; i++ {
		x[i].CombineInPlace(union)
		union = x[i]
	}
}

func coarseOverlap(w *gifbounce.World, worldBounds box2d.AABB, wrapOverlapFramesThreshold, step int) int {
	n := w.NumFrames
	suffixBounds := make([]box2d.AABB, n)
	prefixBounds := make([]box2d.AABB, n)
	for i := 0; i < n; i++ {
		var bounds box2d.AABB
		first := true
		for _, t := range w.Things.Dynamic {
			if !t.Recording.Active[i] {
				continue
			}
			if first {
				bounds = t.Recording.Bounds[i].Clone()
				first = false
				continue
			}
			bounds.CombineInPlace(t.Recording.Bounds[i])
		}
		bounds.IntersectInPlace(worldBounds)
		prefixBounds[i] = bounds
		suffixBounds[i] = bounds
	}
	prefixUnion(prefixBounds)
	suffixUnion(suffixBounds)

	wrapLength := 0
	numOverlapFrames := 0
	for wrapLength < n-step {
		for i := 0; i < step; i++ {
			if box2d.TestOverlapBoundingBoxes(prefixBounds[wrapLength+i], suffixBounds[n-1-wrapLength-i]) {
				numOverlapFrames++
			}
		}
		if numOverlapFrames > wrapOverlapFramesThreshold {
			break
		}
		wrapLength += step
	}
	return wrapLength
}

func fineOverlap(w *gifbounce.World, worldBounds box2d.AABB, wrapOverlapTolerance float64, wrapOverlapFramesThreshold, start, step int) int {
	n := w.NumFrames
	wrapLength := start
	testOverlapQuadratic := func(wrapLength int, tolerance float64) bool {
		for _, t1 := range w.Things.Dynamic {
			if !t1.Recording.Active[wrapLength] {
				continue
			}
			b1 := t1.Recording.Bounds[wrapLength]
			if worldBounds.Contains(b1) {
				continue
			}
			for _, t2 := range w.Things.Dynamic {
				if !t2.Recording.Active[n-1-wrapLength] {
					continue
				}
				b2 := t2.Recording.Bounds[n-1-wrapLength]
				if worldBounds.Contains(b2) {
					continue
				}
				b := b1.Clone()
				b.IntersectInPlace(b2)
				extents := b.GetExtents()
				if extents.X < 0 || extents.Y < 0 {
					continue
				}
				area := (2 * extents.X) * (2 * extents.Y)
				if area > tolerance {
					return true
				}
			}
		}
		return false
	}

	numOverlapFrames := 0
	for wrapLength < n-step {
		for i := 0; i < step; i++ {
			if testOverlapQuadratic(wrapLength+i, wrapOverlapTolerance) {
				numOverlapFrames++
			}
		}
		if numOverlapFrames > wrapOverlapFramesThreshold {
			break
		}
		wrapLength += step
	}

	return wrapLength
}

func WrapV(w *gifbounce.World, worldBounds box2d.AABB, wrapOverlapTolerance float64, wrapOverlapFramesThreshold, step int) {
	n := w.NumFrames
	if n == 0 {
		return
	}

	wrapLength := coarseOverlap(w, worldBounds, wrapOverlapFramesThreshold, step)
	wrapLength = fineOverlap(w, worldBounds, wrapOverlapTolerance, wrapOverlapFramesThreshold, wrapLength, step)
	if wrapLength == 0 {
		return
	}

	w.NumFrames -= wrapLength

	wrapped := make([]*gifbounce.Thing, len(w.Things.Dynamic))
	for i, thing := range w.Things.Dynamic {
		thingCopy := *thing
		thingCopy.Recording = thingCopy.Recording.Slice(n-wrapLength, n)
		thingCopy.Recording.PadRightTo(w.NumFrames)
		thing.Recording = thing.Recording.Slice(0, w.NumFrames)
		wrapped[i] = &thingCopy
	}

	w.Things.Dynamic = append(w.Things.Dynamic, wrapped...)
}
