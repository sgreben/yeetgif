package parse

import (
	"encoding/xml"
	"io"
	"strconv"
	"strings"

	"github.com/sgreben/yeetgif/pkg/rendersvg/svg"
)

func Image(r io.Reader) (*svg.Image, error) {
	var img svg.Image
	type xmlShape struct {
		XMLName       xml.Name
		Data          *string    `xml:"d,attr"`
		Cx            *string    `xml:"cx,attr"`
		Cy            *string    `xml:"cy,attr"`
		R             *string    `xml:"r,attr"`
		Rx            *string    `xml:"rx,attr"`
		Ry            *string    `xml:"ry,attr"`
		Color         *string    `xml:"color,attr"`
		Fill          *string    `xml:"fill,attr"`
		FillColor     *string    `xml:"fill-color,attr"`
		FillOpacity   *string    `xml:"fill-opacity,attr"`
		Stroke        *string    `xml:"stroke,attr"`
		StrokeColor   *string    `xml:"stroke-color,attr"`
		StrokeOpacity *string    `xml:"stroke-opacity,attr"`
		StrokeWidth   *string    `xml:"stroke-width,attr"`
		Shapes        []xmlShape `xml:",any"`
	}
	type xmlRoot struct {
		ViewBox *string    `xml:"viewBox,attr"`
		Shapes  []xmlShape `xml:",any"`
	}
	var root xmlRoot
	err := xml.NewDecoder(r).Decode(&root)
	if err != nil {
		return nil, err
	}
	if root.ViewBox != nil {
		viewBox, err := Rectangle(strings.NewReader(*root.ViewBox))
		if err != nil {
			return nil, err
		}
		img.ViewBox = *viewBox
	}
	xmlShapeAttrs := func(x *xmlShape, s *svg.Shape) error {
		if x.Color != nil {
			c, err := Color(*x.Color)
			if err != nil {
				return err
			}
			s.FillColor = c
			s.StrokeColor = c
		}
		if x.Fill != nil {
			fill, err := Color(*x.Fill)
			if err != nil {
				return err
			}
			s.FillColor = fill
		}
		if x.Stroke != nil {
			stroke, err := Color(*x.Stroke)
			if err != nil {
				return err
			}
			s.StrokeColor = stroke
		}
		return nil
	}
	xmlPath := func(x *xmlShape) (*svg.Shape, error) {
		var p svg.Path
		s := svg.Shape{Path: &p}
		if x.Data != nil {
			commands, err := PathCommands(strings.NewReader(*x.Data))
			if err != nil {
				return nil, err
			}
			p.Commands = commands
		}
		err := xmlShapeAttrs(x, &s)
		return &s, err
	}
	xmlCircle := func(x *xmlShape) (*svg.Shape, error) {
		var c svg.Circle
		s := svg.Shape{Circle: &c}
		if x.Cx != nil {
			cx, err := strconv.ParseFloat(*x.Cx, 64)
			if err != nil {
				return nil, err
			}
			c.Center.X = cx
		}
		if x.Cy != nil {
			cy, err := strconv.ParseFloat(*x.Cy, 64)
			if err != nil {
				return nil, err
			}
			c.Center.Y = cy
		}
		if x.R != nil {
			r, err := strconv.ParseFloat(*x.R, 64)
			if err != nil {
				return nil, err
			}
			c.Radius = r
		}
		err := xmlShapeAttrs(x, &s)
		return &s, err
	}
	xmlEllipse := func(x *xmlShape) (*svg.Shape, error) {
		var e svg.Ellipse
		s := svg.Shape{Ellipse: &e}
		if x.Cx != nil {
			cx, err := strconv.ParseFloat(*x.Cx, 64)
			if err != nil {
				return nil, err
			}
			e.Center.X = cx
		}
		if x.Cy != nil {
			cy, err := strconv.ParseFloat(*x.Cy, 64)
			if err != nil {
				return nil, err
			}
			e.Center.Y = cy
		}
		if x.Rx != nil {
			rx, err := strconv.ParseFloat(*x.Rx, 64)
			if err != nil {
				return nil, err
			}
			e.Radius.X = rx
		}
		if x.Ry != nil {
			ry, err := strconv.ParseFloat(*x.Ry, 64)
			if err != nil {
				return nil, err
			}
			e.Radius.Y = ry
		}
		err := xmlShapeAttrs(x, &s)
		return &s, err
	}
	var xmlAny func(*xmlShape) (*svg.Shape, error)
	xmlGroup := func(x *xmlShape) (*svg.Shape, error) {
		var c []svg.Shape
		for _, xs := range x.Shapes {
			s, err := xmlAny(&xs)
			if err != nil {
				return nil, err
			}
			c = append(c, *s)
		}
		s := svg.Shape{Composite: c}
		err := xmlShapeAttrs(x, &s)
		return &s, err
	}
	xmlAny = func(x *xmlShape) (*svg.Shape, error) {
		switch x.XMLName.Local {
		case "g":
			return xmlGroup(x)
		case "path":
			return xmlPath(x)
		case "circle":
			return xmlCircle(x)
		case "ellipse":
			return xmlEllipse(x)
		}
		return nil, nil
	}
	for _, x := range root.Shapes {
		s, err := xmlAny(&x)
		if err != nil {
			return nil, err
		}
		img.Shapes = append(img.Shapes, *s)
	}
	return &img, nil
}
