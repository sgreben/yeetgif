package box2d

import (
	"image"
	"math"
	"unsafe"
)

type ContactFeatureType uint8

const (
	ContactFeatureTypeVertex ContactFeatureType = 0
	ContactFeatureTypeFace   ContactFeatureType = 1
)

// ContactFeature is features that intersect to form the contact point
type ContactFeature struct {
	IndexA uint8              ///< Feature index on shapeA
	IndexB uint8              ///< Feature index on shapeB
	TypeA  ContactFeatureType ///< The feature type on shapeA
	TypeB  ContactFeatureType ///< The feature type on shapeB
}

type ContactID ContactFeature

/// Key is used to quickly compare contact ids.
func (v ContactID) Key() uint32 {
	return *(*uint32)(unsafe.Pointer(&v.IndexA)) // here we do not care about Endianness; see https://stackoverflow.com/a/7380354
}

func (v ContactID) SetKey(key uint32) {
	*(*uint32)(unsafe.Pointer(&v.IndexA)) = key
}

/// A manifold point is a contact point belonging to a contact
/// manifold. It holds details related to the geometry and dynamics
/// of the contact points.
/// The local point usage depends on the manifold type:
/// -e_circles: the local center of circleB
/// -e_faceA: the local center of cirlceB or the clip point of polygonB
/// -e_faceB: the clip point of polygonA
/// This structure is stored across time steps, so we keep it small.
/// Note: the impulses are used for internal caching and may not
/// provide reliable contact forces, especially for high speed collisions.
type ManifoldPoint struct {
	LocalPoint     Point     ///< usage depends on manifold type
	NormalImpulse  float64   ///< the non-penetration impulse
	TangentImpulse float64   ///< the friction impulse
	Id             ContactID ///< uniquely identifies a contact point between two shapes
}

/// A manifold for two touching convex shapes.
/// Box2D supports multiple types of contact:
/// - clip point versus plane with radius
/// - point versus point with radius (circles)
/// The local point usage depends on the manifold type:
/// -e_circles: the local center of circleA
/// -e_faceA: the center of faceA
/// -e_faceB: the center of faceB
/// Similarly the local normal usage:
/// -e_circles: not used
/// -e_faceA: the normal on polygonA
/// -e_faceB: the normal on polygonB
/// We store contacts in this way so that position correction can
/// account for movement, which is critical for continuous physics.
/// All contact scenarios must be expressed in one of these types.
/// This structure is stored across time steps, so we keep it small.

type ManifoldType uint8

const (
	ManifoldTypeCircles = 0
	ManifoldTypeFaceA   = 1
	ManifoldTypeFaceB   = 2
)

type Manifold struct {
	Points      [_maxManifoldPoints]ManifoldPoint ///< the points of contact
	LocalNormal Point                             ///< not use for Type::e_points
	LocalPoint  Point                             ///< usage depends on manifold type
	Type        ManifoldType                      // Manifold_Type
	PointCount  int                               ///< the number of manifold points
}

// WorldManifold is used to compute the current state of a contact manifold.
type WorldManifold struct {
	Normal      Point                       ///< world vector pointing from A to B
	Points      [_maxManifoldPoints]Point   ///< world contact point (point of intersection)
	Separations [_maxManifoldPoints]float64 ///< a negative value indicates overlap, in meters
}

type PointState uint8

const (
	PointStateNullState    PointState = 0
	PointStateAddState     PointState = 1
	PointStatePersistState PointState = 2
	PointStateRemoveState  PointState = 3
)

/// Used for computing contact manifolds.
type ClipVertex struct {
	V  Point
	Id ContactID
}

/// Ray-cast input data. The ray extends from p1 to p1 + maxFraction * (p2 - p1).
type RayCastInput struct {
	P1, P2      Point
	MaxFraction float64
}

func NewRayCastInput() *RayCastInput {
	return &RayCastInput{}
}

/// Ray-cast output data. The ray hits at p1 + fraction * (p2 - p1), where p1 and p2
/// come from b2RayCastInput.
type RayCastOutput struct {
	Normal   Point
	Fraction float64
}

/// An axis aligned bounding box.
type AABB struct {
	Min Point ///< the lower vertex
	Max Point ///< the upper vertex
}

func (bb AABB) ImageRectangle() image.Rectangle {
	return image.Rectangle{
		Min: bb.Min.ImagePoint(),
		Max: bb.Max.ImagePoint(),
	}
}

/// Get the center of the AABB.
func (bb AABB) GetCenter() Point {
	return PointMulScalar(
		0.5,
		PointAdd(bb.Min, bb.Max),
	)
}

/// Get the extents of the AABB (half-widths).
func (bb AABB) GetExtents() Point {
	return PointMulScalar(
		0.5,
		PointSub(bb.Max, bb.Min),
	)
}

/// Get the perimeter length
func (bb AABB) GetPerimeter() float64 {
	wx := bb.Max.X - bb.Min.X
	wy := bb.Max.Y - bb.Min.Y
	return 2.0 * (wx + wy)
}

func (bb *AABB) IntersectInPlace(s AABB) {
	if bb.Min.X < s.Min.X {
		bb.Min.X = s.Min.X
	}
	if bb.Min.Y < s.Min.Y {
		bb.Min.Y = s.Min.Y
	}
	if bb.Max.X > s.Max.X {
		bb.Max.X = s.Max.X
	}
	if bb.Max.Y > s.Max.Y {
		bb.Max.Y = s.Max.Y
	}
	if bb.Min.X >= bb.Max.X || bb.Min.Y >= bb.Max.Y {
		bb.Min.X = 0
		bb.Min.Y = 0
		bb.Max.X = -1
		bb.Max.Y = -1
	}
}

/// Combine an AABB into this one.
func (bb *AABB) CombineInPlace(aabb AABB) {
	bb.Min = PointMin(bb.Min, aabb.Min)
	bb.Max = PointMax(bb.Max, aabb.Max)
}

/// Combine two AABBs into this one.
func (bb *AABB) CombineTwoInPlace(aabb1, aabb2 AABB) {
	bb.Min = PointMin(aabb1.Min, aabb2.Min)
	bb.Max = PointMax(aabb1.Max, aabb2.Max)
}

/// Does this aabb contain the provided AABB.
func (bb AABB) Contains(aabb AABB) bool {

	return (bb.Min.X <= aabb.Min.X &&
		bb.Min.Y <= aabb.Min.Y &&
		aabb.Max.X <= bb.Max.X &&
		aabb.Max.Y <= bb.Max.Y)
}

func (bb AABB) IsValid() bool {
	d := PointSub(bb.Max, bb.Min)
	valid := d.X >= 0.0 && d.Y >= 0.0
	valid = valid && bb.Min.IsValid() && bb.Max.IsValid()
	return valid
}

func (bb AABB) Clone() AABB {
	clone := AABB{}
	clone.Min = bb.Min.Clone()
	clone.Max = bb.Max.Clone()

	return clone
}

func TestOverlapBoundingBoxes(a, b AABB) bool {

	d1 := PointSub(b.Min, a.Max)
	d2 := PointSub(a.Min, b.Max)

	if d1.X > 0.0 || d1.Y > 0.0 {
		return false
	}

	if d2.X > 0.0 || d2.Y > 0.0 {
		return false
	}

	return true
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// Collision.cpp
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func (wm *WorldManifold) Initialize(manifold *Manifold, xfA Transform, radiusA float64, xfB Transform, radiusB float64) {
	if manifold.PointCount == 0 {
		return
	}

	switch manifold.Type {
	case ManifoldTypeCircles:
		{
			wm.Normal.Set(1.0, 0.0)
			pointA := TransformPointMul(xfA, manifold.LocalPoint)
			pointB := TransformPointMul(xfB, manifold.Points[0].LocalPoint)
			if PointDistanceSquared(pointA, pointB) > _epsilon*_epsilon {
				wm.Normal = PointSub(pointB, pointA)
				wm.Normal.Normalize()
			}

			cA := PointAdd(pointA, PointMulScalar(radiusA, wm.Normal))
			cB := PointSub(pointB, PointMulScalar(radiusB, wm.Normal))

			wm.Points[0] = PointMulScalar(0.5, PointAdd(cA, cB))
			wm.Separations[0] = PointDot(PointSub(cB, cA), wm.Normal)
		}
		break

	case ManifoldTypeFaceA:
		{
			wm.Normal = RotPointMul(xfA.Q, manifold.LocalNormal)
			planePoint := TransformPointMul(xfA, manifold.LocalPoint)

			for i := 0; i < manifold.PointCount; i++ {
				clipPoint := TransformPointMul(xfB, manifold.Points[i].LocalPoint)
				cA := PointAdd(
					clipPoint,
					PointMulScalar(
						radiusA-PointDot(
							PointSub(clipPoint, planePoint),
							wm.Normal,
						),
						wm.Normal,
					),
				)
				cB := PointSub(clipPoint, PointMulScalar(radiusB, wm.Normal))
				wm.Points[i] = PointMulScalar(0.5, PointAdd(cA, cB))
				wm.Separations[i] = PointDot(
					PointSub(cB, cA),
					wm.Normal,
				)
			}
		}
		break

	case ManifoldTypeFaceB:
		{
			wm.Normal = RotPointMul(xfB.Q, manifold.LocalNormal)
			planePoint := TransformPointMul(xfB, manifold.LocalPoint)

			for i := 0; i < manifold.PointCount; i++ {
				clipPoint := TransformPointMul(xfA, manifold.Points[i].LocalPoint)
				cB := PointAdd(clipPoint, PointMulScalar(
					radiusB-PointDot(
						PointSub(clipPoint, planePoint),
						wm.Normal,
					), wm.Normal,
				))
				cA := PointSub(clipPoint, PointMulScalar(radiusA, wm.Normal))
				wm.Points[i] = PointMulScalar(0.5, PointAdd(cA, cB))
				wm.Separations[i] = PointDot(
					PointSub(cA, cB),
					wm.Normal,
				)
			}

			// Ensure normal points from A to B.
			wm.Normal = wm.Normal.OperatorNegate()
		}
		break
	}
}

func GetPointStates(state1 []PointState, state2 []PointState, manifold1 Manifold, manifold2 Manifold) {

	for i := range state1 {
		state1[i] = PointStateNullState
		state2[i] = PointStateNullState
	}

	// Detect persists and removes.
	for i := 0; i < manifold1.PointCount; i++ {
		id := manifold1.Points[i].Id

		state1[i] = PointStateRemoveState

		for j := 0; j < manifold2.PointCount; j++ {
			if manifold2.Points[j].Id.Key() == id.Key() {
				state1[i] = PointStatePersistState
				break
			}
		}
	}

	// Detect persists and adds.
	for i := 0; i < manifold2.PointCount; i++ {
		id := manifold2.Points[i].Id

		state2[i] = PointStateAddState

		for j := 0; j < manifold1.PointCount; j++ {
			if manifold1.Points[j].Id.Key() == id.Key() {
				state2[i] = PointStatePersistState
				break
			}
		}
	}
}

// From Real-time Collision Detection, p179.
func (bb AABB) RayCast(output *RayCastOutput, input RayCastInput) bool {
	tmin := -math.MaxFloat64
	tmax := math.MaxFloat64

	p := input.P1
	d := PointSub(input.P2, input.P1)
	absD := PointAbs(d)

	normal := Point{}

	for i := 0; i < 2; i++ {
		if absD.OperatorIndexGet(i) < _epsilon {
			// Parallel.
			if p.OperatorIndexGet(i) < bb.Min.OperatorIndexGet(i) || bb.Max.OperatorIndexGet(i) < p.OperatorIndexGet(i) {
				return false
			}
		} else {
			inv_d := 1.0 / d.OperatorIndexGet(i)
			t1 := (bb.Min.OperatorIndexGet(i) - p.OperatorIndexGet(i)) * inv_d
			t2 := (bb.Max.OperatorIndexGet(i) - p.OperatorIndexGet(i)) * inv_d

			// Sign of the normal vector.
			s := -1.0

			if t1 > t2 {
				t1, t2 = t2, t1
				s = 1.0
			}

			// Push the min up
			if t1 > tmin {
				normal.SetZero()
				normal.OperatorIndexSet(i, s)
				tmin = t1
			}

			// Pull the max down
			tmax = math.Min(tmax, t2)

			if tmin > tmax {
				return false
			}
		}
	}

	// Does the ray start inside the box?
	// Does the ray intersect beyond the max fraction?
	if tmin < 0.0 || input.MaxFraction < tmin {
		return false
	}

	// Intersection.
	output.Fraction = tmin
	output.Normal = normal
	return true
}

// Sutherland-Hodgman clipping.
func ClipSegmentToLine(vOut []ClipVertex, vIn []ClipVertex, normal Point, offset float64, vertexIndexA int) int {

	// Start with no output points
	numOut := 0

	// Calculate the distance of end points to the line
	distance0 := PointDot(normal, vIn[0].V) - offset
	distance1 := PointDot(normal, vIn[1].V) - offset

	// If the points are behind the plane
	if distance0 <= 0.0 {
		vOut[numOut] = vIn[0]
		numOut++
	}

	if distance1 <= 0.0 {
		vOut[numOut] = vIn[1]
		numOut++
	}

	// If the points are on different sides of the plane
	if distance0*distance1 < 0.0 {
		// Find intersection point of edge and plane
		interp := distance0 / (distance0 - distance1)
		vOut[numOut].V = PointAdd(
			vIn[0].V,
			PointMulScalar(interp, PointSub(vIn[1].V, vIn[0].V)),
		)

		// VertexA is hitting edgeB.
		vOut[numOut].Id.IndexA = uint8(vertexIndexA)
		vOut[numOut].Id.IndexB = vIn[0].Id.IndexB
		vOut[numOut].Id.TypeA = ContactFeatureTypeVertex
		vOut[numOut].Id.TypeB = ContactFeatureTypeFace
		numOut++
	}

	return numOut
}

func TestOverlapShapes(shapeA ShapeInterface, indexA int, shapeB ShapeInterface, indexB int, xfA Transform, xfB Transform) bool {
	input := DistanceInput{}
	input.ProxyA.Set(shapeA, indexA)
	input.ProxyB.Set(shapeB, indexB)
	input.TransformA = xfA
	input.TransformB = xfB
	input.UseRadii = true

	cache := SimplexCache{}
	cache.Count = 0

	output := DistanceOutput{}

	Distance(&output, &cache, &input)

	return output.Distance < 10.0*_epsilon
}
