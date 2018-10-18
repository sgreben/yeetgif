package box2d

import (
	"math"
)

// Find the max separation between poly1 and poly2 using edge normals from poly1.
func FindMaxSeparation(edgeIndex *int, poly1 *PolygonShape, xf1 Transform, poly2 *PolygonShape, xf2 Transform) float64 {
	count1 := poly1.Count
	count2 := poly2.Count
	n1s := poly1.Normals
	v1s := poly1.Vertices
	v2s := poly2.Vertices

	xf := TransformMulT(xf2, xf1)

	bestIndex := 0
	maxSeparation := -math.MaxFloat64
	for i := 0; i < count1; i++ {
		// Get poly1 normal in frame2.
		n := RotPointMul(xf.Q, n1s[i])
		v1 := TransformPointMul(xf, v1s[i])

		// Find deepest point for normal i.
		si := math.MaxFloat64
		for j := 0; j < count2; j++ {
			sij := PointDot(n, PointSub(v2s[j], v1))
			if sij < si {
				si = sij
			}
		}

		if si > maxSeparation {
			maxSeparation = si
			bestIndex = i
		}
	}

	*edgeIndex = bestIndex
	return maxSeparation
}

func FindIncidentEdge(c []ClipVertex, poly1 *PolygonShape, xf1 Transform, edge1 int, poly2 *PolygonShape, xf2 Transform) {

	normals1 := poly1.Normals

	count2 := poly2.Count
	vertices2 := poly2.Vertices
	normals2 := poly2.Normals

	// Get the normal of the reference edge in poly2's frame.
	normal1 := RotPointMulT(xf2.Q, RotPointMul(xf1.Q, normals1[edge1]))

	// Find the incident edge on poly2.
	index := 0
	minDot := math.MaxFloat64
	for i := 0; i < count2; i++ {
		dot := PointDot(normal1, normals2[i])
		if dot < minDot {
			minDot = dot
			index = i
		}
	}

	// Build the clip vertices for the incident edge.
	i1 := index
	i2 := 0
	if i1+1 < count2 {
		i2 = i1 + 1
	}

	c[0].V = TransformPointMul(xf2, vertices2[i1])
	c[0].Id.IndexA = uint8(edge1)
	c[0].Id.IndexB = uint8(i1)
	c[0].Id.TypeA = ContactFeatureTypeFace
	c[0].Id.TypeB = ContactFeatureTypeVertex

	c[1].V = TransformPointMul(xf2, vertices2[i2])
	c[1].Id.IndexA = uint8(edge1)
	c[1].Id.IndexB = uint8(i2)
	c[1].Id.TypeA = ContactFeatureTypeFace
	c[1].Id.TypeB = ContactFeatureTypeVertex
}

// Find edge normal of max separation on A - return if separating axis is found
// Find edge normal of max separation on B - return if separation axis is found
// Choose reference edge as min(minA, minB)
// Find incident edge
// Clip

// The normal points from 1 to 2
func CollidePolygons(manifold *Manifold, polyA *PolygonShape, xfA Transform, polyB *PolygonShape, xfB Transform) {

	manifold.PointCount = 0
	totalRadius := polyA.Radius + polyB.Radius

	edgeA := 0
	separationA := FindMaxSeparation(&edgeA, polyA, xfA, polyB, xfB)
	if separationA > totalRadius {
		return
	}

	edgeB := 0
	separationB := FindMaxSeparation(&edgeB, polyB, xfB, polyA, xfA)
	if separationB > totalRadius {
		return
	}

	var poly1 *PolygonShape // reference polygon
	var poly2 *PolygonShape // incident polygon

	xf1 := Transform{}
	xf2 := Transform{}

	edge1 := 0 // reference edge
	var flip uint8
	k_tol := 0.1 * _linearSlop

	if separationB > separationA+k_tol {
		poly1 = polyB
		poly2 = polyA
		xf1 = xfB
		xf2 = xfA
		edge1 = edgeB
		manifold.Type = ManifoldTypeFaceB
		flip = 1
	} else {
		poly1 = polyA
		poly2 = polyB
		xf1 = xfA
		xf2 = xfB
		edge1 = edgeA
		manifold.Type = ManifoldTypeFaceA
		flip = 0
	}

	incidentEdge := make([]ClipVertex, 2)
	FindIncidentEdge(incidentEdge, poly1, xf1, edge1, poly2, xf2)

	count1 := poly1.Count
	vertices1 := poly1.Vertices

	iv1 := edge1
	iv2 := 0
	if edge1+1 < count1 {
		iv2 = edge1 + 1
	}

	v11 := vertices1[iv1]
	v12 := vertices1[iv2]

	localTangent := PointSub(v12, v11)
	localTangent.Normalize()

	localNormal := PointCrossVectorScalar(localTangent, 1.0)
	planePoint := PointMulScalar(0.5, PointAdd(v11, v12))

	tangent := RotPointMul(xf1.Q, localTangent)
	normal := PointCrossVectorScalar(tangent, 1.0)

	v11 = TransformPointMul(xf1, v11)
	v12 = TransformPointMul(xf1, v12)

	// Face offset.
	frontOffset := PointDot(normal, v11)

	// Side offsets, extended by polytope skin thickness.
	sideOffset1 := -PointDot(tangent, v11) + totalRadius
	sideOffset2 := PointDot(tangent, v12) + totalRadius

	// Clip incident edge against extruded edge1 side edges.
	clipPoints1 := make([]ClipVertex, 2)
	clipPoints2 := make([]ClipVertex, 2)
	np := 0

	// Clip to box side 1
	np = ClipSegmentToLine(clipPoints1, incidentEdge, tangent.OperatorNegate(), sideOffset1, iv1)

	if np < 2 {
		return
	}

	// Clip to negative box side 1
	np = ClipSegmentToLine(clipPoints2, clipPoints1, tangent, sideOffset2, iv2)

	if np < 2 {
		return
	}

	// Now clipPoints2 contains the clipped points.
	manifold.LocalNormal = localNormal
	manifold.LocalPoint = planePoint

	pointCount := 0
	for i := 0; i < len(clipPoints2); i++ {
		separation := PointDot(normal, clipPoints2[i].V) - frontOffset

		if separation <= totalRadius {
			cp := &manifold.Points[pointCount]
			cp.LocalPoint = TransformPointMulT(xf2, clipPoints2[i].V)
			cp.Id = clipPoints2[i].Id
			if flip != 0 {
				// Swap features
				cf := cp.Id
				cp.Id.IndexA = cf.IndexB
				cp.Id.IndexB = cf.IndexA
				cp.Id.TypeA = cf.TypeB
				cp.Id.TypeB = cf.TypeA
			}
			pointCount++
		}
	}

	manifold.PointCount = pointCount
}
