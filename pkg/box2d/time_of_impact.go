package box2d

import (
	"math"
)

/// Input parameters for b2TimeOfImpact
type TOIInput struct {
	ProxyA DistanceProxy
	ProxyB DistanceProxy
	SweepA Sweep
	SweepB Sweep
	TMax   float64 // defines sweep interval [0, tMax]
}

// Output parameters for b2TimeOfImpact.
type TOIOutputState uint8

const (
	TOIOutputStateUnknown    = 1
	TOIOutputStateFailed     = 2
	TOIOutputStateOverlapped = 3
	TOIOutputStateTouching   = 4
	TOIOutputStateSeparated  = 5
)

type TOIOutput struct {
	State TOIOutputState
	T     float64
}

var _toiCalls, _toiIters, _toiMaxIters int
var _toiRootIters, _toiMaxRootIters int

type SeparationFunctionType uint8

const (
	SeparationFunctionTypePoints SeparationFunctionType = 0
	SeparationFunctionTypeFaceA  SeparationFunctionType = 1
	SeparationFunctionTypeFaceB  SeparationFunctionType = 2
)

type SeparationFunction struct {
	ProxyA         *DistanceProxy
	ProxyB         *DistanceProxy
	SweepA, SweepB Sweep
	Type           SeparationFunctionType
	LocalPoint     Point
	Axis           Point
}

// TODO_ERIN might not need to return the separation
func (sepfunc *SeparationFunction) Initialize(cache *SimplexCache, proxyA *DistanceProxy, sweepA Sweep, proxyB *DistanceProxy, sweepB Sweep, t1 float64) float64 {

	sepfunc.ProxyA = proxyA
	sepfunc.ProxyB = proxyB
	count := cache.Count

	sepfunc.SweepA = sweepA
	sepfunc.SweepB = sweepB

	xfA := Transform{}
	xfB := Transform{}
	sepfunc.SweepA.GetTransform(&xfA, t1)
	sepfunc.SweepB.GetTransform(&xfB, t1)

	if count == 1 {
		sepfunc.Type = SeparationFunctionTypePoints
		localPointA := sepfunc.ProxyA.GetVertex(cache.IndexA[0])
		localPointB := sepfunc.ProxyB.GetVertex(cache.IndexB[0])
		pointA := TransformPointMul(xfA, localPointA)
		pointB := TransformPointMul(xfB, localPointB)
		sepfunc.Axis = PointSub(pointB, pointA)
		s := sepfunc.Axis.Normalize()
		return s
	} else if cache.IndexA[0] == cache.IndexA[1] {
		// Two points on B and one on A.
		sepfunc.Type = SeparationFunctionTypeFaceB
		localPointB1 := proxyB.GetVertex(cache.IndexB[0])
		localPoint := proxyB.GetVertex(cache.IndexB[1])

		sepfunc.Axis = PointCrossVectorScalar(
			PointSub(localPoint, localPointB1),
			1.0,
		)

		sepfunc.Axis.Normalize()
		normal := RotPointMul(xfB.Q, sepfunc.Axis)

		sepfunc.LocalPoint = PointMulScalar(0.5, PointAdd(localPointB1, localPoint))
		pointB := TransformPointMul(xfB, sepfunc.LocalPoint)

		localPointA := proxyA.GetVertex(cache.IndexA[0])
		pointA := TransformPointMul(xfA, localPointA)

		s := PointDot(PointSub(pointA, pointB), normal)
		if s < 0.0 {
			sepfunc.Axis = sepfunc.Axis.OperatorNegate()
			s = -s
		}

		return s
	} else {
		// Two points on A and one or two points on B.
		sepfunc.Type = SeparationFunctionTypeFaceA
		localPointA1 := sepfunc.ProxyA.GetVertex(cache.IndexA[0])
		localPointA2 := sepfunc.ProxyA.GetVertex(cache.IndexA[1])

		sepfunc.Axis = PointCrossVectorScalar(PointSub(localPointA2, localPointA1), 1.0)
		sepfunc.Axis.Normalize()
		normal := RotPointMul(xfA.Q, sepfunc.Axis)

		sepfunc.LocalPoint = PointMulScalar(0.5, PointAdd(localPointA1, localPointA2))
		pointA := TransformPointMul(xfA, sepfunc.LocalPoint)

		localPointB := sepfunc.ProxyB.GetVertex(cache.IndexB[0])
		pointB := TransformPointMul(xfB, localPointB)

		s := PointDot(PointSub(pointB, pointA), normal)
		if s < 0.0 {
			sepfunc.Axis = sepfunc.Axis.OperatorNegate()
			s = -s
		}

		return s
	}
}

//
func (sepfunc *SeparationFunction) FindMinSeparation(indexA *int, indexB *int, t float64) float64 {

	xfA := Transform{}
	xfB := Transform{}

	sepfunc.SweepA.GetTransform(&xfA, t)
	sepfunc.SweepB.GetTransform(&xfB, t)

	switch sepfunc.Type {
	case SeparationFunctionTypePoints:
		{
			axisA := RotPointMulT(xfA.Q, sepfunc.Axis)
			axisB := RotPointMulT(xfB.Q, sepfunc.Axis.OperatorNegate())

			*indexA = sepfunc.ProxyA.GetSupport(axisA)
			*indexB = sepfunc.ProxyB.GetSupport(axisB)

			localPointA := sepfunc.ProxyA.GetVertex(*indexA)
			localPointB := sepfunc.ProxyB.GetVertex(*indexB)

			pointA := TransformPointMul(xfA, localPointA)
			pointB := TransformPointMul(xfB, localPointB)

			separation := PointDot(PointSub(pointB, pointA), sepfunc.Axis)
			return separation
		}

	case SeparationFunctionTypeFaceA:
		{
			normal := RotPointMul(xfA.Q, sepfunc.Axis)
			pointA := TransformPointMul(xfA, sepfunc.LocalPoint)

			axisB := RotPointMulT(xfB.Q, normal.OperatorNegate())

			*indexA = -1
			*indexB = sepfunc.ProxyB.GetSupport(axisB)

			localPointB := sepfunc.ProxyB.GetVertex(*indexB)
			pointB := TransformPointMul(xfB, localPointB)

			separation := PointDot(PointSub(pointB, pointA), normal)
			return separation
		}

	case SeparationFunctionTypeFaceB:
		{
			normal := RotPointMul(xfB.Q, sepfunc.Axis)
			pointB := TransformPointMul(xfB, sepfunc.LocalPoint)

			axisA := RotPointMulT(xfA.Q, normal.OperatorNegate())

			*indexB = -1
			*indexA = sepfunc.ProxyA.GetSupport(axisA)

			localPointA := sepfunc.ProxyA.GetVertex(*indexA)
			pointA := TransformPointMul(xfA, localPointA)

			separation := PointDot(PointSub(pointA, pointB), normal)
			return separation
		}

	default:
		*indexA = -1
		*indexB = -1
		return 0.0
	}
}

//
func (sepfunc *SeparationFunction) Evaluate(indexA int, indexB int, t float64) float64 {

	xfA := Transform{}
	xfB := Transform{}

	sepfunc.SweepA.GetTransform(&xfA, t)
	sepfunc.SweepB.GetTransform(&xfB, t)

	switch sepfunc.Type {
	case SeparationFunctionTypePoints:
		{
			localPointA := sepfunc.ProxyA.GetVertex(indexA)
			localPointB := sepfunc.ProxyB.GetVertex(indexB)

			pointA := TransformPointMul(xfA, localPointA)
			pointB := TransformPointMul(xfB, localPointB)
			separation := PointDot(PointSub(pointB, pointA), sepfunc.Axis)

			return separation
		}

	case SeparationFunctionTypeFaceA:
		{
			normal := RotPointMul(xfA.Q, sepfunc.Axis)
			pointA := TransformPointMul(xfA, sepfunc.LocalPoint)

			localPointB := sepfunc.ProxyB.GetVertex(indexB)
			pointB := TransformPointMul(xfB, localPointB)

			separation := PointDot(PointSub(pointB, pointA), normal)
			return separation
		}

	case SeparationFunctionTypeFaceB:
		{
			normal := RotPointMul(xfB.Q, sepfunc.Axis)
			pointB := TransformPointMul(xfB, sepfunc.LocalPoint)

			localPointA := sepfunc.ProxyA.GetVertex(indexA)
			pointA := TransformPointMul(xfA, localPointA)

			separation := PointDot(PointSub(pointA, pointB), normal)
			return separation
		}

	default:
		return 0.0
	}
}

/// Compute the upper bound on time before two shapes penetrate. Time is represented as
/// a fraction between [0,tMax]. This uses a swept separating axis and may miss some intermediate,
/// non-tunneling collision. If you change the time interval, you should call this function
/// again.
/// Note: use b2Distance to compute the contact point and normal at the time of impact.
// CCD via the local separating axis method. This seeks progression
// by computing the largest time at which separation is maintained.
func TimeOfImpact(output *TOIOutput, input *TOIInput) {

	_toiCalls++

	output.State = TOIOutputStateUnknown
	output.T = input.TMax

	proxyA := &input.ProxyA
	proxyB := &input.ProxyB

	sweepA := input.SweepA
	sweepB := input.SweepB

	// Large rotations can make the root finder fail, so we normalize the
	// sweep angles.
	sweepA.Normalize()
	sweepB.Normalize()

	tMax := input.TMax

	totalRadius := proxyA.Radius + proxyB.Radius
	target := math.Max(_linearSlop, totalRadius-3.0*_linearSlop)
	tolerance := 0.25 * _linearSlop

	t1 := 0.0
	k_maxIterations := 20 // TODO_ERIN b2Settings
	iter := 0

	// Prepare input for distance query.
	cache := SimplexCache{}
	cache.Count = 0
	distanceInput := DistanceInput{}
	distanceInput.ProxyA = input.ProxyA
	distanceInput.ProxyB = input.ProxyB
	distanceInput.UseRadii = false

	// The outer loop progressively attempts to compute new separating axes.
	// This loop terminates when an axis is repeated (no progress is made).
	for {

		xfA := Transform{}
		xfB := Transform{}

		sweepA.GetTransform(&xfA, t1)
		sweepB.GetTransform(&xfB, t1)

		// Get the distance between shapes. We can also use the results
		// to get a separating axis.
		distanceInput.TransformA = xfA
		distanceInput.TransformB = xfB
		distanceOutput := DistanceOutput{}
		Distance(&distanceOutput, &cache, &distanceInput)

		// If the shapes are overlapped, we give up on continuous collision.
		if distanceOutput.Distance <= 0.0 {
			// Failure!
			output.State = TOIOutputStateOverlapped
			output.T = 0.0
			break
		}

		if distanceOutput.Distance < target+tolerance {
			// Victory!
			output.State = TOIOutputStateTouching
			output.T = t1
			break
		}

		// Initialize the separating axis.
		var fcn SeparationFunction
		fcn.Initialize(&cache, proxyA, sweepA, proxyB, sweepB, t1)

		// Compute the TOI on the separating axis. We do this by successively
		// resolving the deepest point. This loop is bounded by the number of vertices.
		done := false
		t2 := tMax
		pushBackIter := 0
		for {
			// Find the deepest point at t2. Store the witness point indices.
			var indexA, indexB int
			s2 := fcn.FindMinSeparation(&indexA, &indexB, t2)

			// Is the final configuration separated?
			if s2 > target+tolerance {
				// Victory!
				output.State = TOIOutputStateSeparated
				output.T = tMax
				done = true
				break
			}

			// Has the separation reached tolerance?
			if s2 > target-tolerance {
				// Advance the sweeps
				t1 = t2
				break
			}

			// Compute the initial separation of the witness points.
			s1 := fcn.Evaluate(indexA, indexB, t1)

			// Check for initial overlap. This might happen if the root finder
			// runs out of iterations.
			if s1 < target-tolerance {
				output.State = TOIOutputStateFailed
				output.T = t1
				done = true
				break
			}

			// Check for touching
			if s1 <= target+tolerance {
				// Victory! t1 should hold the TOI (could be 0.0).
				output.State = TOIOutputStateTouching
				output.T = t1
				done = true
				break
			}

			// Compute 1D root of: f(x) - target = 0
			rootIterCount := 0
			a1 := t1
			a2 := t2

			for {
				// Use a mix of the secant rule and bisection.
				t := 0.0

				if (rootIterCount & 1) != 0x0000 {
					// Secant rule to improve convergence.
					t = a1 + (target-s1)*(a2-a1)/(s2-s1)
				} else {
					// Bisection to guarantee progress.
					t = 0.5 * (a1 + a2)
				}

				rootIterCount++
				_toiRootIters++

				s := fcn.Evaluate(indexA, indexB, t)

				if math.Abs(s-target) < tolerance {
					// t2 holds a tentative value for t1
					t2 = t
					break
				}

				// Ensure we continue to bracket the root.
				if s > target {
					a1 = t
					s1 = s
				} else {
					a2 = t
					s2 = s
				}

				if rootIterCount == 50 {
					break
				}
			}

			_toiMaxRootIters = MaxInt(_toiMaxRootIters, rootIterCount)

			pushBackIter++

			if pushBackIter == _maxPolygonVertices {
				break
			}
		}

		iter++
		_toiIters++

		if done {
			break
		}

		if iter == k_maxIterations {
			// Root finder got stuck. Semi-victory.
			output.State = TOIOutputStateFailed
			output.T = t1
			break
		}
	}

	_toiMaxIters = MaxInt(_toiMaxIters, iter)

}
