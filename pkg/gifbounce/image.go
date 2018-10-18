package gifbounce

import (
	"image"

	"github.com/sgreben/yeetgif/pkg/box2d"
)

func ToImagePoint(p *box2d.Point) image.Point {
	return image.Point{X: int(p.X), Y: int(-p.Y)}
}

func FromImagePoint(p *image.Point) box2d.Point {
	return box2d.Point{X: float64(p.X), Y: float64(-p.Y)}
}

func FromImageRect(p *image.Rectangle) box2d.AABB {
	min := FromImagePoint(&p.Min)
	max := FromImagePoint(&p.Max)
	min.Y, max.Y = max.Y, min.Y
	return box2d.AABB{Min: min, Max: max}
}

func FromImagePolygons(imagePolygons [][]image.Point) [][]box2d.Point {
	polygons := make([][]box2d.Point, len(imagePolygons))
	for i, points := range imagePolygons {
		polygons[i] = make([]box2d.Point, len(points))
		for j := range points {
			polygons[i][j].SetImagePoint(&points[j])
			polygons[i][j].Y *= -1
		}
	}
	return polygons
}
