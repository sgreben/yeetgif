package box2d

import (
	"math"
)

type EPAxisType uint8

// This structure is used to keep track of the best separating axis.
const (
	EPAxisTypeUnknown EPAxisType = 0
	EPAxisTypeEdgeA   EPAxisType = 1
	EPAxisTypeEdgeB   EPAxisType = 2
)

type EPAxis struct {
	Type       EPAxisType
	Index      int
	Separation float64
}

func MakeEPAxis() EPAxis {
	return EPAxis{}
}

// This holds polygon B expressed in frame A.
type TempPolygon struct {
	Vertices []Point
	Normals  []Point
	Count    int
}

// Reference face used for clipping
type ReferenceFace struct {
	I1, I2 int

	V1, V2 Point

	Normal Point

	SideNormal1 Point
	SideOffset1 float64

	SideNormal2 Point
	SideOffset2 float64
}

func MakeReferenceFace() ReferenceFace {
	return ReferenceFace{}
}

type EPColliderVertexType uint8

const (
	EPColliderVertexTypeIsolated EPColliderVertexType = 0
	EPColliderVertexTypeConcave  EPColliderVertexType = 1
	EPColliderVertexTypeConvex   EPColliderVertexType = 2
)

// This class collides and edge and a polygon, taking into account edge adjacency.
type EPCollider struct {
	PolygonB TempPolygon

	Xf                        Transform
	CentroidB                 Point
	V0, V1, V2, V3            Point
	Normal0, Normal1, Normal2 Point
	Normal                    Point
	Type1, Type2              uint8
	LowerLimit, UpperLimit    Point
	Radius                    float64
	Front                     bool
}

func MakeEPCollider() EPCollider {
	return EPCollider{}
}

// Algorithm:
// 1. Classify v1 and v2
// 2. Classify polygon centroid as front or back
// 3. Flip normal if necessary
// 4. Initialize normal range to [-pi, pi] about face normal
// 5. Adjust normal range according to adjacent edges
// 6. Visit each separating axes, only accept axes within the range
// 7. Return if _any_ axis indicates separation
// 8. Clip
func (collider *EPCollider) Collide(manifold *Manifold, edgeA *EdgeShape, xfA Transform, polygonB *PolygonShape, xfB Transform) {

	collider.Xf = TransformMulT(xfA, xfB)

	collider.CentroidB = TransformPointMul(collider.Xf, polygonB.Centroid)

	collider.V0 = edgeA.Vertex0
	collider.V1 = edgeA.Vertex1
	collider.V2 = edgeA.Vertex2
	collider.V3 = edgeA.Vertex3

	hasVertex0 := edgeA.HasVertex0
	hasVertex3 := edgeA.HasVertex3

	edge1 := PointSub(collider.V2, collider.V1)
	edge1.Normalize()
	collider.Normal1.Set(edge1.Y, -edge1.X)
	offset1 := PointDot(collider.Normal1, PointSub(collider.CentroidB, collider.V1))
	offset0 := 0.0
	offset2 := 0.0
	convex1 := false
	convex2 := false

	// Is there a preceding edge?
	if hasVertex0 {
		edge0 := PointSub(collider.V1, collider.V0)
		edge0.Normalize()
		collider.Normal0.Set(edge0.Y, -edge0.X)
		convex1 = PointCross(edge0, edge1) >= 0.0
		offset0 = PointDot(collider.Normal0, PointSub(collider.CentroidB, collider.V0))
	}

	// Is there a following edge?
	if hasVertex3 {
		edge2 := PointSub(collider.V3, collider.V2)
		edge2.Normalize()
		collider.Normal2.Set(edge2.Y, -edge2.X)
		convex2 = PointCross(edge1, edge2) > 0.0
		offset2 = PointDot(collider.Normal2, PointSub(collider.CentroidB, collider.V2))
	}

	// Determine front or back collision. Determine collision normal limits.
	if hasVertex0 && hasVertex3 {
		if convex1 && convex2 {
			collider.Front = offset0 >= 0.0 || offset1 >= 0.0 || offset2 >= 0.0
			if collider.Front {
				collider.Normal = collider.Normal1
				collider.LowerLimit = collider.Normal0
				collider.UpperLimit = collider.Normal2
			} else {
				collider.Normal = collider.Normal1.OperatorNegate()
				collider.LowerLimit = collider.Normal1.OperatorNegate()
				collider.UpperLimit = collider.Normal1.OperatorNegate()
			}
		} else if convex1 {
			collider.Front = offset0 >= 0.0 || (offset1 >= 0.0 && offset2 >= 0.0)
			if collider.Front {
				collider.Normal = collider.Normal1
				collider.LowerLimit = collider.Normal0
				collider.UpperLimit = collider.Normal1
			} else {
				collider.Normal = collider.Normal1.OperatorNegate()
				collider.LowerLimit = collider.Normal2.OperatorNegate()
				collider.UpperLimit = collider.Normal1.OperatorNegate()
			}
		} else if convex2 {
			collider.Front = offset2 >= 0.0 || (offset0 >= 0.0 && offset1 >= 0.0)
			if collider.Front {
				collider.Normal = collider.Normal1
				collider.LowerLimit = collider.Normal1
				collider.UpperLimit = collider.Normal2
			} else {
				collider.Normal = collider.Normal1.OperatorNegate()
				collider.LowerLimit = collider.Normal1.OperatorNegate()
				collider.UpperLimit = collider.Normal0.OperatorNegate()
			}
		} else {
			collider.Front = offset0 >= 0.0 && offset1 >= 0.0 && offset2 >= 0.0
			if collider.Front {
				collider.Normal = collider.Normal1
				collider.LowerLimit = collider.Normal1
				collider.UpperLimit = collider.Normal1
			} else {
				collider.Normal = collider.Normal1.OperatorNegate()
				collider.LowerLimit = collider.Normal2.OperatorNegate()
				collider.UpperLimit = collider.Normal0.OperatorNegate()
			}
		}
	} else if hasVertex0 {
		if convex1 {
			collider.Front = offset0 >= 0.0 || offset1 >= 0.0
			if collider.Front {
				collider.Normal = collider.Normal1
				collider.LowerLimit = collider.Normal0
				collider.UpperLimit = collider.Normal1.OperatorNegate()
			} else {
				collider.Normal = collider.Normal1.OperatorNegate()
				collider.LowerLimit = collider.Normal1
				collider.UpperLimit = collider.Normal1.OperatorNegate()
			}
		} else {
			collider.Front = offset0 >= 0.0 && offset1 >= 0.0
			if collider.Front {
				collider.Normal = collider.Normal1
				collider.LowerLimit = collider.Normal1
				collider.UpperLimit = collider.Normal1.OperatorNegate()
			} else {
				collider.Normal = collider.Normal1.OperatorNegate()
				collider.LowerLimit = collider.Normal1
				collider.UpperLimit = collider.Normal0.OperatorNegate()
			}
		}
	} else if hasVertex3 {
		if convex2 {
			collider.Front = offset1 >= 0.0 || offset2 >= 0.0
			if collider.Front {
				collider.Normal = collider.Normal1
				collider.LowerLimit = collider.Normal1.OperatorNegate()
				collider.UpperLimit = collider.Normal2
			} else {
				collider.Normal = collider.Normal1.OperatorNegate()
				collider.LowerLimit = collider.Normal1.OperatorNegate()
				collider.UpperLimit = collider.Normal1
			}
		} else {
			collider.Front = offset1 >= 0.0 && offset2 >= 0.0
			if collider.Front {
				collider.Normal = collider.Normal1
				collider.LowerLimit = collider.Normal1.OperatorNegate()
				collider.UpperLimit = collider.Normal1
			} else {
				collider.Normal = collider.Normal1.OperatorNegate()
				collider.LowerLimit = collider.Normal2.OperatorNegate()
				collider.UpperLimit = collider.Normal1
			}
		}
	} else {
		collider.Front = offset1 >= 0.0
		if collider.Front {
			collider.Normal = collider.Normal1
			collider.LowerLimit = collider.Normal1.OperatorNegate()
			collider.UpperLimit = collider.Normal1.OperatorNegate()
		} else {
			collider.Normal = collider.Normal1.OperatorNegate()
			collider.LowerLimit = collider.Normal1
			collider.UpperLimit = collider.Normal1
		}
	}

	// Get polygonB in frameA
	collider.PolygonB.Count = polygonB.Count
	collider.PolygonB.Vertices = make([]Point, len(polygonB.Vertices))
	collider.PolygonB.Normals = make([]Point, len(polygonB.Normals))
	for i := 0; i < polygonB.Count; i++ {
		collider.PolygonB.Vertices[i] = TransformPointMul(collider.Xf, polygonB.Vertices[i])
		collider.PolygonB.Normals[i] = RotPointMul(collider.Xf.Q, polygonB.Normals[i])
	}

	collider.Radius = polygonB.Radius + edgeA.Radius

	manifold.PointCount = 0

	edgeAxis := collider.ComputeEdgeSeparation()

	// If no valid normal can be found than this edge should not collide.
	if edgeAxis.Type == EPAxisTypeUnknown {
		return
	}

	if edgeAxis.Separation > collider.Radius {
		return
	}

	polygonAxis := collider.ComputePolygonSeparation()
	if polygonAxis.Type != EPAxisTypeUnknown && polygonAxis.Separation > collider.Radius {
		return
	}

	// Use hysteresis for jitter reduction.
	k_relativeTol := 0.98
	k_absoluteTol := 0.001

	primaryAxis := MakeEPAxis()
	if polygonAxis.Type == EPAxisTypeUnknown {
		primaryAxis = edgeAxis
	} else if polygonAxis.Separation > k_relativeTol*edgeAxis.Separation+k_absoluteTol {
		primaryAxis = polygonAxis
	} else {
		primaryAxis = edgeAxis
	}

	ie := make([]ClipVertex, 2)
	rf := MakeReferenceFace()
	if primaryAxis.Type == EPAxisTypeEdgeA {
		manifold.Type = ManifoldTypeFaceA

		// Search for the polygon normal that is most anti-parallel to the edge normal.
		bestIndex := 0
		bestValue := PointDot(collider.Normal, collider.PolygonB.Normals[0])
		for i := 1; i < collider.PolygonB.Count; i++ {
			value := PointDot(collider.Normal, collider.PolygonB.Normals[i])
			if value < bestValue {
				bestValue = value
				bestIndex = i
			}
		}

		i1 := bestIndex
		i2 := 0
		if i1+1 < collider.PolygonB.Count {
			i2 = i1 + 1
		}

		ie[0].V = collider.PolygonB.Vertices[i1]
		ie[0].Id.IndexA = 0
		ie[0].Id.IndexB = uint8(i1)
		ie[0].Id.TypeA = ContactFeatureTypeFace
		ie[0].Id.TypeB = ContactFeatureTypeVertex

		ie[1].V = collider.PolygonB.Vertices[i2]
		ie[1].Id.IndexA = 0
		ie[1].Id.IndexB = uint8(i2)
		ie[1].Id.TypeA = ContactFeatureTypeFace
		ie[1].Id.TypeB = ContactFeatureTypeVertex

		if collider.Front {
			rf.I1 = 0
			rf.I2 = 1
			rf.V1 = collider.V1
			rf.V2 = collider.V2
			rf.Normal = collider.Normal1
		} else {
			rf.I1 = 1
			rf.I2 = 0
			rf.V1 = collider.V2
			rf.V2 = collider.V1
			rf.Normal = collider.Normal1.OperatorNegate()
		}
	} else {
		manifold.Type = ManifoldTypeFaceB

		ie[0].V = collider.V1
		ie[0].Id.IndexA = 0
		ie[0].Id.IndexB = uint8(primaryAxis.Index)
		ie[0].Id.TypeA = ContactFeatureTypeVertex
		ie[0].Id.TypeB = ContactFeatureTypeFace

		ie[1].V = collider.V2
		ie[1].Id.IndexA = 0
		ie[1].Id.IndexB = uint8(primaryAxis.Index)
		ie[1].Id.TypeA = ContactFeatureTypeVertex
		ie[1].Id.TypeB = ContactFeatureTypeFace

		rf.I1 = primaryAxis.Index
		if rf.I1+1 < collider.PolygonB.Count {
			rf.I2 = rf.I1 + 1
		} else {
			rf.I2 = 0
		}

		rf.V1 = collider.PolygonB.Vertices[rf.I1]
		rf.V2 = collider.PolygonB.Vertices[rf.I2]
		rf.Normal = collider.PolygonB.Normals[rf.I1]
	}

	rf.SideNormal1.Set(rf.Normal.Y, -rf.Normal.X)
	rf.SideNormal2 = rf.SideNormal1.OperatorNegate()
	rf.SideOffset1 = PointDot(rf.SideNormal1, rf.V1)
	rf.SideOffset2 = PointDot(rf.SideNormal2, rf.V2)

	// Clip incident edge against extruded edge1 side edges.
	clipPoints1 := make([]ClipVertex, 2)
	clipPoints2 := make([]ClipVertex, 2)
	np := 0

	// Clip to box side 1
	np = ClipSegmentToLine(clipPoints1, ie, rf.SideNormal1, rf.SideOffset1, rf.I1)

	if np < _maxManifoldPoints {
		return
	}

	// Clip to negative box side 1
	np = ClipSegmentToLine(clipPoints2, clipPoints1, rf.SideNormal2, rf.SideOffset2, rf.I2)

	if np < _maxManifoldPoints {
		return
	}

	// Now clipPoints2 contains the clipped points.
	if primaryAxis.Type == EPAxisTypeEdgeA {
		manifold.LocalNormal = rf.Normal
		manifold.LocalPoint = rf.V1
	} else {
		manifold.LocalNormal = polygonB.Normals[rf.I1]
		manifold.LocalPoint = polygonB.Vertices[rf.I1]
	}

	pointCount := 0
	for i := 0; i < _maxManifoldPoints; i++ {
		separation := 0.0

		separation = PointDot(rf.Normal, PointSub(clipPoints2[i].V, rf.V1))

		if separation <= collider.Radius {
			cp := &manifold.Points[pointCount]

			if primaryAxis.Type == EPAxisTypeEdgeA {
				cp.LocalPoint = TransformPointMulT(collider.Xf, clipPoints2[i].V)
				cp.Id = clipPoints2[i].Id
			} else {
				cp.LocalPoint = clipPoints2[i].V
				cp.Id.TypeA = clipPoints2[i].Id.TypeB
				cp.Id.TypeB = clipPoints2[i].Id.TypeA
				cp.Id.IndexA = clipPoints2[i].Id.IndexB
				cp.Id.IndexB = clipPoints2[i].Id.IndexA
			}

			pointCount++
		}
	}

	manifold.PointCount = pointCount
}

func (collider *EPCollider) ComputeEdgeSeparation() EPAxis {
	axis := MakeEPAxis()
	axis.Type = EPAxisTypeEdgeA
	if collider.Front {
		axis.Index = 0
	} else {
		axis.Index = 1
	}
	axis.Separation = math.MaxFloat64

	for i := 0; i < collider.PolygonB.Count; i++ {
		s := PointDot(collider.Normal, PointSub(collider.PolygonB.Vertices[i], collider.V1))
		if s < axis.Separation {
			axis.Separation = s
		}
	}

	return axis
}

func (collider *EPCollider) ComputePolygonSeparation() EPAxis {

	axis := MakeEPAxis()
	axis.Type = EPAxisTypeUnknown
	axis.Index = -1
	axis.Separation = -math.MaxFloat64

	perp := Point{X: -collider.Normal.Y, Y: collider.Normal.X}

	for i := 0; i < collider.PolygonB.Count; i++ {
		n := collider.PolygonB.Normals[i].OperatorNegate()

		s1 := PointDot(n, PointSub(collider.PolygonB.Vertices[i], collider.V1))
		s2 := PointDot(n, PointSub(collider.PolygonB.Vertices[i], collider.V2))
		s := math.Min(s1, s2)

		if s > collider.Radius {
			// No collision
			axis.Type = EPAxisTypeEdgeB
			axis.Index = i
			axis.Separation = s
			return axis
		}

		// Adjacency
		if PointDot(n, perp) >= 0.0 {
			if PointDot(PointSub(n, collider.UpperLimit), collider.Normal) < -_angularSlop {
				continue
			}
		} else {
			if PointDot(PointSub(n, collider.LowerLimit), collider.Normal) < -_angularSlop {
				continue
			}
		}

		if s > axis.Separation {
			axis.Type = EPAxisTypeEdgeB
			axis.Index = i
			axis.Separation = s
		}
	}

	return axis
}

func CollideEdgeAndPolygon(manifold *Manifold, edgeA *EdgeShape, xfA Transform, polygonB *PolygonShape, xfB Transform) {
	collider := MakeEPCollider()
	collider.Collide(manifold, edgeA, xfA, polygonB, xfB)
}
