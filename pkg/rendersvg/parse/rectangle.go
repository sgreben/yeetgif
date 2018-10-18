package parse

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/sgreben/yeetgif/pkg/rendersvg/geom"
)

func Rectangle(r io.Reader) (*geom.Rectangle, error) {
	var floats []float64
	scanner := bufio.NewScanner(r)
	scanner.Split(scanFloat)
	for scanner.Scan() {
		f, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, f)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(floats) != 4 {
		return nil, fmt.Errorf("parse rectangle: expected 4 numbers, got %d: %v", len(floats), floats)
	}
	out := geom.Rectangle{
		Min: geom.Point{
			X: floats[0],
			Y: floats[1],
		},
		Max: geom.Point{
			X: floats[2],
			Y: floats[3],
		},
	}
	return &out, nil
}
