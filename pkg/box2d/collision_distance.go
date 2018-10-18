package box2d

/// A distance proxy is used by the GJK algorithm.
/// It encapsulates any shape.
type DistanceProxy struct {
	Buffer   [2]Point
	Vertices []Point // is a memory blob using pointer arithmetic in original implementation
	Count    int
	Radius   float64
}

/// Used to warm start b2Distance.
/// Set count to zero on first call.
type SimplexCache struct {
	Metric float64 ///< length or area
	Count  int
	IndexA [3]int ///< vertices on shape A
	IndexB [3]int ///< vertices on shape B
}

/// Input for b2Distance.
/// You have to option to use the shape radii
/// in the computation. Even
type DistanceInput struct {
	ProxyA     DistanceProxy
	ProxyB     DistanceProxy
	TransformA Transform
	TransformB Transform
	UseRadii   bool
}

/// Output for b2Distance.
type DistanceOutput struct {
	PointA     Point ///< closest point on shapeA
	PointB     Point ///< closest point on shapeB
	Distance   float64
	Iterations int ///< number of GJK iterations used
}

// //////////////////////////////////////////////////////////////////////////

func (p DistanceProxy) GetVertexCount() int {
	return p.Count
}

func (p DistanceProxy) GetVertex(index int) Point {
	return p.Vertices[index]
}

func (p DistanceProxy) GetSupport(d Point) int {
	bestIndex := 0
	bestValue := PointDot(p.Vertices[0], d)
	for i := 1; i < p.Count; i++ {
		value := PointDot(p.Vertices[i], d)
		if value > bestValue {
			bestIndex = i
			bestValue = value
		}
	}

	return bestIndex
}

func (p DistanceProxy) GetSupportVertex(d Point) Point {
	bestIndex := 0
	bestValue := PointDot(p.Vertices[0], d)

	for i := 1; i < p.Count; i++ {
		value := PointDot(p.Vertices[i], d)
		if value > bestValue {
			bestIndex = i
			bestValue = value
		}
	}

	return p.Vertices[bestIndex]
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// Distance.cpp
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// GJK using Voronoi regions (Christer Ericson) and Barycentric coordinates.
var b2_gjkCalls, b2_gjkIters, b2_gjkMaxIters int

func (p *DistanceProxy) Set(shape ShapeInterface, index int) {
	switch shape.GetType() {
	case ShapeTypePolygon:
		{
			polygon := shape.(*PolygonShape)
			p.Vertices = polygon.Vertices[:]
			p.Count = polygon.Count
			p.Radius = polygon.Radius
		}

	case ShapeTypeEdge:
		{
			edge := shape.(*EdgeShape)
			p.Vertices = []Point{edge.Vertex1, edge.Vertex2}
			p.Count = 2
			p.Radius = edge.Radius
		}

	default:
	}
}

type SimplexVertex struct {
	WA     Point   // support point in proxyA
	WB     Point   // support point in proxyB
	W      Point   // wB - wA
	A      float64 // barycentric coordinate for closest point
	IndexA int     // wA index
	IndexB int     // wB index
}

type Simplex struct {
	//V1, V2, V3 *SimplexVertex
	Vs    [3]SimplexVertex
	Count int
}

func (simplex *Simplex) ReadCache(cache *SimplexCache, proxyA *DistanceProxy, transformA Transform, proxyB *DistanceProxy, transformB Transform) {

	// Copy data from cache.
	simplex.Count = cache.Count
	vertices := &simplex.Vs
	for i := 0; i < simplex.Count; i++ {
		v := &vertices[i]
		v.IndexA = cache.IndexA[i]
		v.IndexB = cache.IndexB[i]
		wALocal := proxyA.GetVertex(v.IndexA)
		wBLocal := proxyB.GetVertex(v.IndexB)
		v.WA = TransformPointMul(transformA, wALocal)
		v.WB = TransformPointMul(transformB, wBLocal)
		v.W = PointSub(v.WB, v.WA)
		v.A = 0.0
	}

	// Compute the new simplex metric, if it is substantially different than
	// old metric then flush the simplex.
	if simplex.Count > 1 {
		metric1 := cache.Metric
		metric2 := simplex.GetMetric()
		if metric2 < 0.5*metric1 || 2.0*metric1 < metric2 || metric2 < _epsilon {
			// Reset the simplex.
			simplex.Count = 0
		}
	}

	// If the cache is empty or invalid ...
	if simplex.Count == 0 {
		v := &vertices[0]
		v.IndexA = 0
		v.IndexB = 0
		wALocal := proxyA.GetVertex(0)
		wBLocal := proxyB.GetVertex(0)
		v.WA = TransformPointMul(transformA, wALocal)
		v.WB = TransformPointMul(transformB, wBLocal)
		v.W = PointSub(v.WB, v.WA)
		v.A = 1.0
		simplex.Count = 1
	}
}

func (simplex Simplex) WriteCache(cache *SimplexCache) {
	cache.Metric = simplex.GetMetric()
	cache.Count = simplex.Count
	vertices := &simplex.Vs
	for i := 0; i < simplex.Count; i++ {
		cache.IndexA[i] = vertices[i].IndexA
		cache.IndexB[i] = vertices[i].IndexB
	}
}

func (simplex Simplex) GetSearchDirection() Point {
	switch simplex.Count {
	case 1:
		return simplex.Vs[0].W.OperatorNegate()

	case 2:
		{
			e12 := PointSub(simplex.Vs[1].W, simplex.Vs[0].W)
			sgn := PointCross(e12, simplex.Vs[0].W.OperatorNegate())
			if sgn > 0.0 {
				// Origin is left of e12.
				return PointCrossScalarVector(1.0, e12)
			} else {
				// Origin is right of e12.
				return PointCrossVectorScalar(e12, 1.0)
			}
		}

	default:
		return Point{}
	}
}

func (simplex Simplex) GetClosestPoint() Point {
	switch simplex.Count {
	case 0:
		return Point{}

	case 1:
		return simplex.Vs[0].W

	case 2:
		return PointAdd(
			PointMulScalar(
				simplex.Vs[0].A,
				simplex.Vs[0].W,
			),
			PointMulScalar(
				simplex.Vs[1].A,
				simplex.Vs[1].W,
			),
		)

	case 3:
		return Point{}

	default:
		return Point{}
	}
}

func (simplex Simplex) GetWitnessPoints(pA *Point, pB *Point) {
	switch simplex.Count {
	case 0:
		break

	case 1:
		*pA = simplex.Vs[0].WA
		*pB = simplex.Vs[0].WB
		break

	case 2:
		*pA = PointAdd(
			PointMulScalar(simplex.Vs[0].A, simplex.Vs[0].WA),
			PointMulScalar(simplex.Vs[1].A, simplex.Vs[1].WA),
		)
		*pB = PointAdd(
			PointMulScalar(simplex.Vs[0].A, simplex.Vs[0].WB),
			PointMulScalar(simplex.Vs[1].A, simplex.Vs[1].WB),
		)
		break

	case 3:
		*pA = PointAdd(
			PointAdd(
				PointMulScalar(simplex.Vs[0].A, simplex.Vs[0].WA),
				PointMulScalar(simplex.Vs[1].A, simplex.Vs[1].WA),
			),
			PointMulScalar(simplex.Vs[2].A, simplex.Vs[2].WA),
		)
		*pB = *pA
		break

	default:
		break
	}
}

func (simplex Simplex) GetMetric() float64 {
	switch simplex.Count {
	case 0:
		return 0.0

	case 1:
		return 0.0

	case 2:
		return PointDistance(simplex.Vs[0].W, simplex.Vs[1].W)

	case 3:
		return PointCross(
			PointSub(simplex.Vs[1].W, simplex.Vs[0].W),
			PointSub(simplex.Vs[2].W, simplex.Vs[0].W),
		)

	default:
		return 0.0
	}
}

////////////////////////////////////////////////////

// Solve a line segment using barycentric coordinates.
func (simplex *Simplex) Solve2() {
	w1 := simplex.Vs[0].W
	w2 := simplex.Vs[1].W
	e12 := PointSub(w2, w1)

	// w1 region
	d12_2 := -PointDot(w1, e12)
	if d12_2 <= 0.0 {
		// a2 <= 0, so we clamp it to 0
		simplex.Vs[0].A = 1.0
		simplex.Count = 1
		return
	}

	// w2 region
	d12_1 := PointDot(w2, e12)
	if d12_1 <= 0.0 {
		// a1 <= 0, so we clamp it to 0
		simplex.Vs[1].A = 1.0
		simplex.Count = 1
		simplex.Vs[0] = simplex.Vs[1]
		return
	}

	// Must be in e12 region.
	inv_d12 := 1.0 / (d12_1 + d12_2)
	simplex.Vs[0].A = d12_1 * inv_d12
	simplex.Vs[1].A = d12_2 * inv_d12
	simplex.Count = 2
}

// // Possible regions:
// // - points[2]
// // - edge points[0]-points[2]
// // - edge points[1]-points[2]
// // - inside the triangle
func (simplex *Simplex) Solve3() {

	w1 := simplex.Vs[0].W
	w2 := simplex.Vs[1].W
	w3 := simplex.Vs[2].W

	// Edge12
	// [1      1     ][a1] = [1]
	// [w1.e12 w2.e12][a2] = [0]
	// a3 = 0
	e12 := PointSub(w2, w1)
	w1e12 := PointDot(w1, e12)
	w2e12 := PointDot(w2, e12)
	d12_1 := w2e12
	d12_2 := -w1e12

	// Edge13
	// [1      1     ][a1] = [1]
	// [w1.e13 w3.e13][a3] = [0]
	// a2 = 0
	e13 := PointSub(w3, w1)
	w1e13 := PointDot(w1, e13)
	w3e13 := PointDot(w3, e13)
	d13_1 := w3e13
	d13_2 := -w1e13

	// Edge23
	// [1      1     ][a2] = [1]
	// [w2.e23 w3.e23][a3] = [0]
	// a1 = 0
	e23 := PointSub(w3, w2)
	w2e23 := PointDot(w2, e23)
	w3e23 := PointDot(w3, e23)
	d23_1 := w3e23
	d23_2 := -w2e23

	// Triangle123
	n123 := PointCross(e12, e13)

	d123_1 := n123 * PointCross(w2, w3)
	d123_2 := n123 * PointCross(w3, w1)
	d123_3 := n123 * PointCross(w1, w2)

	// w1 region
	if d12_2 <= 0.0 && d13_2 <= 0.0 {
		simplex.Vs[0].A = 1.0
		simplex.Count = 1
		return
	}

	// e12
	if d12_1 > 0.0 && d12_2 > 0.0 && d123_3 <= 0.0 {
		inv_d12 := 1.0 / (d12_1 + d12_2)
		simplex.Vs[0].A = d12_1 * inv_d12
		simplex.Vs[1].A = d12_2 * inv_d12
		simplex.Count = 2
		return
	}

	// e13
	if d13_1 > 0.0 && d13_2 > 0.0 && d123_2 <= 0.0 {
		inv_d13 := 1.0 / (d13_1 + d13_2)
		simplex.Vs[0].A = d13_1 * inv_d13
		simplex.Vs[2].A = d13_2 * inv_d13
		simplex.Count = 2
		simplex.Vs[1] = simplex.Vs[2]
		return
	}

	// w2 region
	if d12_1 <= 0.0 && d23_2 <= 0.0 {
		simplex.Vs[1].A = 1.0
		simplex.Count = 1
		simplex.Vs[0] = simplex.Vs[1]
		return
	}

	// w3 region
	if d13_1 <= 0.0 && d23_1 <= 0.0 {
		simplex.Vs[2].A = 1.0
		simplex.Count = 1
		simplex.Vs[0] = simplex.Vs[2]
		return
	}

	// e23
	if d23_1 > 0.0 && d23_2 > 0.0 && d123_1 <= 0.0 {
		inv_d23 := 1.0 / (d23_1 + d23_2)
		simplex.Vs[1].A = d23_1 * inv_d23
		simplex.Vs[2].A = d23_2 * inv_d23
		simplex.Count = 2
		simplex.Vs[0] = simplex.Vs[2]
		return
	}

	// Must be in triangle123
	inv_d123 := 1.0 / (d123_1 + d123_2 + d123_3)
	simplex.Vs[0].A = d123_1 * inv_d123
	simplex.Vs[1].A = d123_2 * inv_d123
	simplex.Vs[2].A = d123_3 * inv_d123
	simplex.Count = 3
}

func Distance(output *DistanceOutput, cache *SimplexCache, input *DistanceInput) {
	b2_gjkCalls++

	proxyA := &input.ProxyA
	proxyB := &input.ProxyB

	transformA := input.TransformA
	transformB := input.TransformB

	// Initialize the simplex.
	simplex := Simplex{}
	simplex.ReadCache(cache, proxyA, transformA, proxyB, transformB)

	// Get simplex vertices as an array.
	vertices := &simplex.Vs
	k_maxIters := 20

	// These store the vertices of the last simplex so that we
	// can check for duplicates and prevent cycling.
	saveA := make([]int, 3)
	saveB := make([]int, 3)
	saveCount := 0

	// Main iteration loop.
	iter := 0
	for iter < k_maxIters {
		// Copy simplex so we can identify duplicates.
		saveCount = simplex.Count
		for i := 0; i < saveCount; i++ {
			saveA[i] = vertices[i].IndexA
			saveB[i] = vertices[i].IndexB
		}

		switch simplex.Count {
		case 1:
			break

		case 2:
			simplex.Solve2()
			break

		case 3:
			simplex.Solve3()
			break

		default:
		}

		// If we have 3 points, then the origin is in the corresponding triangle.
		if simplex.Count == 3 {
			break
		}

		// Get search direction.
		d := simplex.GetSearchDirection()

		// Ensure the search direction is numerically fit.
		if d.LengthSquared() < _epsilon*_epsilon {
			// The origin is probably contained by a line segment
			// or triangle. Thus the shapes are overlapped.

			// We can't return zero here even though there may be overlap.
			// In case the simplex is a point, segment, or triangle it is difficult
			// to determine if the origin is contained in the CSO or very close to it.
			break
		}

		// Compute a tentative new simplex vertex using support points.
		vertex := &vertices[simplex.Count]
		vertex.IndexA = proxyA.GetSupport(
			RotPointMulT(transformA.Q, d.OperatorNegate()),
		)
		vertex.WA = TransformPointMul(transformA, proxyA.GetVertex(vertex.IndexA))
		// b2Point wBLocal;
		vertex.IndexB = proxyB.GetSupport(RotPointMulT(transformB.Q, d))
		vertex.WB = TransformPointMul(transformB, proxyB.GetVertex(vertex.IndexB))
		vertex.W = PointSub(vertex.WB, vertex.WA)

		// Iteration count is equated to the number of support point calls.
		iter++
		b2_gjkIters++

		// Check for duplicate support points. This is the main termination criteria.
		duplicate := false
		for i := 0; i < saveCount; i++ {
			if vertex.IndexA == saveA[i] && vertex.IndexB == saveB[i] {
				duplicate = true
				break
			}
		}

		// If we found a duplicate support point we must exit to avoid cycling.
		if duplicate {
			break
		}

		// New vertex is ok and needed.
		simplex.Count++
	}

	if iter > b2_gjkMaxIters {
		b2_gjkMaxIters = iter
	}

	// Prepare output.
	simplex.GetWitnessPoints(&output.PointA, &output.PointB)
	output.Distance = PointDistance(output.PointA, output.PointB)
	output.Iterations = iter

	// // Cache the simplex.
	simplex.WriteCache(cache)

	// // Apply radii if requested.
	if input.UseRadii {
		rA := proxyA.Radius
		rB := proxyB.Radius

		if output.Distance > rA+rB && output.Distance > _epsilon {
			// Shapes are still no overlapped.
			// Move the witness points to the outer surface.
			output.Distance -= rA + rB
			normal := PointSub(output.PointB, output.PointA)
			normal.Normalize()
			output.PointA.OperatorPlusInplace(
				PointMulScalar(rA, normal),
			)
			output.PointB.OperatorMinusInplace(
				PointMulScalar(rB, normal),
			)
		} else {
			// Shapes are overlapped when radii are considered.
			// Move the witness points to the middle.
			p := PointMulScalar(
				0.5,
				PointAdd(output.PointA, output.PointB),
			)
			output.PointA = p
			output.PointB = p
			output.Distance = 0.0
		}
	}
}
