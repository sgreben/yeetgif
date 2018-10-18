package box2d

/// A convex polygon. It is assumed that the interior of the polygon is to
/// the left of each edge.
/// Polygons have a maximum number of vertices equal to b2_maxPolygonVertices.
/// In most cases you should not need many vertices for a convex polygon.

type PolygonShape struct {
	Shape

	Centroid Point
	Vertices []Point
	Normals  []Point
	Count    int
}

func MakePolygonShape() PolygonShape {
	return PolygonShape{
		Shape: Shape{
			Type:   ShapeTypePolygon,
			Radius: _polygonRadius,
		},
		Count:    0,
		Centroid: Point{X: 0, Y: 0},
	}
}

func NewPolygonShape() *PolygonShape {
	res := MakePolygonShape()
	return &res
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// PolygonShape.cpp
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func (poly PolygonShape) Clone() ShapeInterface {

	clone := NewPolygonShape()
	clone.Centroid = poly.Centroid
	clone.Count = poly.Count

	clone.Vertices = make([]Point, len(poly.Vertices))
	for i := range poly.Vertices {
		clone.Vertices[i] = poly.Vertices[i]
	}

	clone.Normals = make([]Point, len(poly.Normals))
	for i := range poly.Normals {
		clone.Normals[i] = poly.Normals[i]
	}

	return clone
}

func (poly *PolygonShape) SetAsBox(hx float64, hy float64) {
	poly.Count = 4
	poly.Vertices = make([]Point, poly.Count)
	poly.Normals = make([]Point, poly.Count)
	poly.Vertices[0].Set(-hx, -hy)
	poly.Vertices[1].Set(hx, -hy)
	poly.Vertices[2].Set(hx, hy)
	poly.Vertices[3].Set(-hx, hy)
	poly.Normals[0].Set(0.0, -1.0)
	poly.Normals[1].Set(1.0, 0.0)
	poly.Normals[2].Set(0.0, 1.0)
	poly.Normals[3].Set(-1.0, 0.0)
	poly.Centroid.SetZero()
}

func (poly *PolygonShape) SetAsBoxFromCenterAndAngle(hx float64, hy float64, center Point, angle float64) {
	poly.Count = 4
	poly.Vertices[0].Set(-hx, -hy)
	poly.Vertices[1].Set(hx, -hy)
	poly.Vertices[2].Set(hx, hy)
	poly.Vertices[3].Set(-hx, hy)
	poly.Normals[0].Set(0.0, -1.0)
	poly.Normals[1].Set(1.0, 0.0)
	poly.Normals[2].Set(0.0, 1.0)
	poly.Normals[3].Set(-1.0, 0.0)
	poly.Centroid = center

	xf := Transform{}
	xf.P = center
	xf.Q.Set(angle)

	// Transform vertices and normals.
	for i := 0; i < poly.Count; i++ {
		poly.Vertices[i] = TransformPointMul(xf, poly.Vertices[i])
		poly.Normals[i] = RotPointMul(xf.Q, poly.Normals[i])
	}
}

func (poly PolygonShape) GetChildCount() int {
	return 1
}

func ComputeCentroid(vs []Point, count int) Point {

	c := Point{X: 0, Y: 0}
	area := 0.0

	// pRef is the reference point for forming triangles.
	// It's location doesn't change the result (except for rounding error).
	pRef := Point{X: 0.0, Y: 0.0}

	inv3 := 1.0 / 3.0

	for i := 0; i < count; i++ {
		// Triangle vertices.
		p1 := pRef
		p2 := vs[i]
		p3 := Point{X: 0, Y: 0}
		if i+1 < count {
			p3 = vs[i+1]
		} else {
			p3 = vs[0]
		}

		e1 := PointSub(p2, p1)
		e2 := PointSub(p3, p1)

		D := PointCross(e1, e2)

		triangleArea := 0.5 * D
		area += triangleArea

		// Area weighted centroid
		c.OperatorPlusInplace(PointMulScalar(triangleArea*inv3, PointAdd(PointAdd(p1, p2), p3)))
	}

	// Centroid
	c.OperatorScalarMulInplace(1.0 / area)
	return c
}

func (poly *PolygonShape) Set(vertices []Point) {
	count := len(vertices)
	if count < 3 {
		poly.SetAsBox(1.0, 1.0)
		return
	}

	n := count

	// Perform welding and copy vertices into local buffer.
	ps := make([]Point, count)
	tempCount := 0

	for i := 0; i < n; i++ {
		v := vertices[i]

		unique := true
		for j := 0; j < tempCount; j++ {
			if PointDistanceSquared(v, ps[j]) < ((0.5 * _linearSlop) * (0.5 * _linearSlop)) {
				unique = false
				break
			}
		}

		if unique {
			ps[tempCount] = v
			tempCount++
		}
	}

	n = tempCount
	if n < 3 {
		// Polygon is degenerate.
		poly.SetAsBox(1.0, 1.0)
		return
	}

	// Create the convex hull using the Gift wrapping algorithm
	// http://en.wikipedia.org/wiki/Gift_wrapping_algorithm

	// Find the right most point on the hull
	i0 := 0
	x0 := ps[0].X
	for i := 1; i < n; i++ {
		x := ps[i].X
		if x > x0 || (x == x0 && ps[i].Y < ps[i0].Y) {
			i0 = i
			x0 = x
		}
	}

	hull := make([]int, count)
	m := 0
	ih := i0

	for {
		hull[m] = ih

		ie := 0
		for j := 1; j < n; j++ {
			if ie == ih {
				ie = j
				continue
			}

			r := PointSub(ps[ie], ps[hull[m]])
			v := PointSub(ps[j], ps[hull[m]])
			c := PointCross(r, v)
			if c < 0.0 {
				ie = j
			}

			// Collinearity check
			if c == 0.0 && v.LengthSquared() > r.LengthSquared() {
				ie = j
			}
		}

		m++
		ih = ie

		if ie == i0 {
			break
		}
	}

	if m < 3 {
		// Polygon is degenerate.
		poly.SetAsBox(1.0, 1.0)
		return
	}

	poly.Count = m

	poly.Vertices = make([]Point, m)
	poly.Normals = make([]Point, m)
	// Copy vertices.
	for i := 0; i < m; i++ {
		poly.Vertices[i] = ps[hull[i]]
	}

	// Compute normals. Ensure the edges have non-zero length.
	for i := 0; i < m; i++ {
		i1 := i
		i2 := 0
		if i+1 < m {
			i2 = i + 1
		}

		edge := PointSub(poly.Vertices[i2], poly.Vertices[i1])
		poly.Normals[i] = PointCrossVectorScalar(edge, 1.0)
		poly.Normals[i].Normalize()
	}

	// Compute the polygon centroid.
	poly.Centroid = ComputeCentroid(poly.Vertices[:], m)
}

func (poly PolygonShape) TestPoint(xf Transform, p Point) bool {
	pLocal := RotPointMulT(xf.Q, PointSub(p, xf.P))

	for i := 0; i < poly.Count; i++ {
		dot := PointDot(poly.Normals[i], PointSub(pLocal, poly.Vertices[i]))
		if dot > 0.0 {
			return false
		}
	}

	return true
}

func (poly PolygonShape) RayCast(output *RayCastOutput, input RayCastInput, xf Transform, childIndex int) bool {

	// Put the ray into the polygon's frame of reference.
	p1 := RotPointMulT(xf.Q, PointSub(input.P1, xf.P))
	p2 := RotPointMulT(xf.Q, PointSub(input.P2, xf.P))
	d := PointSub(p2, p1)

	lower := 0.0
	upper := input.MaxFraction

	index := -1

	for i := 0; i < poly.Count; i++ {
		// p = p1 + a * d
		// dot(normal, p - v) = 0
		// dot(normal, p1 - v) + a * dot(normal, d) = 0
		numerator := PointDot(poly.Normals[i], PointSub(poly.Vertices[i], p1))
		denominator := PointDot(poly.Normals[i], d)

		if denominator == 0.0 {
			if numerator < 0.0 {
				return false
			}
		} else {
			// Note: we want this predicate without division:
			// lower < numerator / denominator, where denominator < 0
			// Since denominator < 0, we have to flip the inequality:
			// lower < numerator / denominator <==> denominator * lower > numerator.
			if denominator < 0.0 && numerator < lower*denominator {
				// Increase lower.
				// The segment enters this half-space.
				lower = numerator / denominator
				index = i
			} else if denominator > 0.0 && numerator < upper*denominator {
				// Decrease upper.
				// The segment exits this half-space.
				upper = numerator / denominator
			}
		}

		// The use of epsilon here causes the assert on lower to trip
		// in some cases. Apparently the use of epsilon was to make edge
		// shapes work, but now those are handled separately.
		//if (upper < lower - b2_epsilon)
		if upper < lower {
			return false
		}
	}

	if index >= 0 {
		output.Fraction = lower
		output.Normal = RotPointMul(xf.Q, poly.Normals[index])
		return true
	}

	return false
}

func (poly PolygonShape) ComputeAABB(aabb *AABB, xf Transform, childIndex int) {

	lower := TransformPointMul(xf, poly.Vertices[0])
	upper := lower

	for i := 1; i < poly.Count; i++ {
		v := TransformPointMul(xf, poly.Vertices[i])
		lower = PointMin(lower, v)
		upper = PointMax(upper, v)
	}

	r := Point{X: poly.Radius, Y: poly.Radius}
	aabb.Min = PointSub(lower, r)
	aabb.Max = PointSub(upper, r)
}

func (poly PolygonShape) ComputeMass(massData *MassData, density float64) {
	// Polygon mass, centroid, and inertia.
	// Let rho be the polygon density in mass per unit area.
	// Then:
	// mass = rho * int(dA)
	// centroid.x = (1/mass) * rho * int(x * dA)
	// centroid.y = (1/mass) * rho * int(y * dA)
	// I = rho * int((x*x + y*y) * dA)
	//
	// We can compute these integrals by summing all the integrals
	// for each triangle of the polygon. To evaluate the integral
	// for a single triangle, we make a change of variables to
	// the (u,v) coordinates of the triangle:
	// x = x0 + e1x * u + e2x * v
	// y = y0 + e1y * u + e2y * v
	// where 0 <= u && 0 <= v && u + v <= 1.
	//
	// We integrate u from [0,1-v] and then v from [0,1].
	// We also need to use the Jacobian of the transformation:
	// D = cross(e1, e2)
	//
	// Simplification: triangle centroid = (1/3) * (p1 + p2 + p3)
	//
	// The rest of the derivation is handled by computer algebra.

	center := Point{X: 0, Y: 0}

	area := 0.0
	I := 0.0

	// s is the reference point for forming triangles.
	// It's location doesn't change the result (except for rounding error).
	s := Point{X: 0.0, Y: 0.0}

	// This code would put the reference point inside the polygon.
	for i := 0; i < poly.Count; i++ {
		s.OperatorPlusInplace(poly.Vertices[i])
	}

	s.OperatorScalarMulInplace(1.0 / float64(poly.Count))

	k_inv3 := 1.0 / 3.0

	for i := 0; i < poly.Count; i++ {
		// Triangle vertices.
		e1 := PointSub(poly.Vertices[i], s)
		e2 := Point{X: 0, Y: 0}

		if i+1 < poly.Count {
			e2 = PointSub(poly.Vertices[i+1], s)
		} else {
			e2 = PointSub(poly.Vertices[0], s)
		}

		D := PointCross(e1, e2)

		triangleArea := 0.5 * D
		area += triangleArea

		// Area weighted centroid
		center.OperatorPlusInplace(PointMulScalar(triangleArea*k_inv3, PointAdd(e1, e2)))

		ex1 := e1.X
		ey1 := e1.Y
		ex2 := e2.X
		ey2 := e2.Y

		intx2 := ex1*ex1 + ex2*ex1 + ex2*ex2
		inty2 := ey1*ey1 + ey2*ey1 + ey2*ey2

		I += (0.25 * k_inv3 * D) * (intx2 + inty2)
	}

	// Total mass
	massData.Mass = density * area

	// Center of mass
	center.OperatorScalarMulInplace(1.0 / area)
	massData.Center = PointAdd(center, s)

	// Inertia tensor relative to the local origin (point s).
	massData.I = density * I

	// Shift to center of mass then to original body origin.
	massData.I += massData.Mass * (PointDot(massData.Center, massData.Center) - PointDot(center, center))
}

func (poly PolygonShape) Validate() bool {

	for i := 0; i < poly.Count; i++ {
		i1 := i
		i2 := 0

		if i < poly.Count-1 {
			i2 = i1 + 1
		}

		p := poly.Vertices[i1]
		e := PointSub(poly.Vertices[i2], p)

		for j := 0; j < poly.Count; j++ {
			if j == i1 || j == i2 {
				continue
			}

			v := PointSub(poly.Vertices[j], p)
			c := PointCross(e, v)
			if c < 0.0 {
				return false
			}
		}
	}

	return true
}
