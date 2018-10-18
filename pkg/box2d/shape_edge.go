package box2d

/// A line segment (edge) shape. These can be connected in chains or loops
/// to other edge shapes. The connectivity information is used to ensure
/// correct contact normals.
type EdgeShape struct {
	Shape
	/// These are the edge vertices
	Vertex1, Vertex2 Point

	/// Optional adjacent vertices. These are used for smooth collision.
	Vertex0, Vertex3       Point
	HasVertex0, HasVertex3 bool
}

func MakeEdgeShape() EdgeShape {
	return EdgeShape{
		Shape: Shape{
			Type:   ShapeTypeEdge,
			Radius: _polygonRadius,
		},
	}
}

func NewEdgeShape() *EdgeShape {
	res := MakeEdgeShape()
	return &res
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// EdgeShape.cpp
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func (edge *EdgeShape) Set(v1 Point, v2 Point) {
	edge.Vertex1 = v1
	edge.Vertex2 = v2
	edge.HasVertex0 = false
	edge.HasVertex3 = false
}

func (edge EdgeShape) Clone() ShapeInterface {
	clone := NewEdgeShape()
	clone.Vertex0 = edge.Vertex0
	clone.Vertex1 = edge.Vertex1
	clone.Vertex2 = edge.Vertex2
	clone.Vertex3 = edge.Vertex3
	clone.HasVertex0 = edge.HasVertex0
	clone.HasVertex3 = edge.HasVertex3

	return clone
}

func (edge EdgeShape) GetChildCount() int {
	return 1
}

func (edge EdgeShape) TestPoint(xf Transform, p Point) bool {
	return false
}

// p = p1 + t * d
// v = v1 + s * e
// p1 + t * d = v1 + s * e
// s * e - t * d = p1 - v1
func (edge EdgeShape) RayCast(output *RayCastOutput, input RayCastInput, xf Transform, childIndex int) bool {

	// Put the ray into the edge's frame of reference.
	p1 := RotPointMulT(xf.Q, PointSub(input.P1, xf.P))
	p2 := RotPointMulT(xf.Q, PointSub(input.P2, xf.P))
	d := PointSub(p2, p1)

	v1 := edge.Vertex1
	v2 := edge.Vertex2
	e := PointSub(v2, v1)
	normal := Point{X: e.Y, Y: -e.X}
	normal.Normalize()

	// q = p1 + t * d
	// dot(normal, q - v1) = 0
	// dot(normal, p1 - v1) + t * dot(normal, d) = 0
	numerator := PointDot(normal, PointSub(v1, p1))
	denominator := PointDot(normal, d)

	if denominator == 0.0 {
		return false
	}

	t := numerator / denominator
	if t < 0.0 || input.MaxFraction < t {
		return false
	}

	q := PointAdd(p1, PointMulScalar(t, d))

	// q = v1 + s * r
	// s = dot(q - v1, r) / dot(r, r)
	r := PointSub(v2, v1)
	rr := PointDot(r, r)
	if rr == 0.0 {
		return false
	}

	s := PointDot(PointSub(q, v1), r) / rr
	if s < 0.0 || 1.0 < s {
		return false
	}

	output.Fraction = t
	if numerator > 0.0 {
		output.Normal = RotPointMul(xf.Q, normal).OperatorNegate()
	} else {
		output.Normal = RotPointMul(xf.Q, normal)
	}

	return true
}

func (edge EdgeShape) ComputeAABB(aabb *AABB, xf Transform, childIndex int) {

	v1 := TransformPointMul(xf, edge.Vertex1)
	v2 := TransformPointMul(xf, edge.Vertex2)

	lower := PointMin(v1, v2)
	upper := PointMax(v1, v2)

	r := Point{X: edge.Radius, Y: edge.Radius}
	aabb.Min = PointSub(lower, r)
	aabb.Max = PointSub(upper, r)
}

func (edge EdgeShape) ComputeMass(massData *MassData, density float64) {
	massData.Mass = 0.0
	massData.Center = PointMulScalar(0.5, PointAdd(edge.Vertex1, edge.Vertex2))
	massData.I = 0.0
}
