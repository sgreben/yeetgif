package box2d

import (
	"math"
)

type VelocityConstraintPoint struct {
	RA             Point
	RB             Point
	NormalImpulse  float64
	TangentImpulse float64
	NormalMass     float64
	TangentMass    float64
	VelocityBias   float64
}

type ContactVelocityConstraint struct {
	Points             [_maxManifoldPoints]VelocityConstraintPoint
	Normal             Point
	NormalMass         Mat22
	K                  Mat22
	IndexA             int
	IndexB             int
	InvMassA, InvMassB float64
	InvIA, InvIB       float64
	Friction           float64
	Restitution        float64
	TangentSpeed       float64
	PointCount         int
	ContactIndex       int
}

type ContactSolverDef struct {
	Step       TimeStep
	Contacts   []ContactInterface // has to be backed by pointers
	Count      int
	Positions  []Position
	Velocities []Velocity
}

func MakeContactSolverDef() ContactSolverDef {
	return ContactSolverDef{
		Contacts:   make([]ContactInterface, 0),
		Positions:  make([]Position, 0),
		Velocities: make([]Velocity, 0),
	}
}

type ContactSolver struct {
	Step                TimeStep
	Positions           []Position
	Velocities          []Velocity
	PositionConstraints []ContactPositionConstraint
	VelocityConstraints []ContactVelocityConstraint
	Contacts            []ContactInterface // has to be backed by pointers
	Count               int
}

var g_blockSolve = true

type ContactPositionConstraint struct {
	LocalPoints                [_maxManifoldPoints]Point
	LocalNormal                Point
	LocalPoint                 Point
	IndexA                     int
	IndexB                     int
	InvMassA, InvMassB         float64
	LocalCenterA, LocalCenterB Point
	InvIA, InvIB               float64
	Type                       ManifoldType
	RadiusA, RadiusB           float64
	PointCount                 int
}

func MakeContactSolver(def *ContactSolverDef) ContactSolver {
	solver := ContactSolver{}

	solver.Step = def.Step
	solver.Count = def.Count
	solver.PositionConstraints = make([]ContactPositionConstraint, solver.Count)
	solver.VelocityConstraints = make([]ContactVelocityConstraint, solver.Count)
	solver.Positions = def.Positions
	solver.Velocities = def.Velocities
	solver.Contacts = def.Contacts

	// Initialize position independent portions of the constraints.
	for i := 0; i < solver.Count; i++ {
		contact := solver.Contacts[i]
		c := contact.Data()

		fixtureA := c.GetFixtureA()
		fixtureB := c.GetFixtureB()
		shapeA := fixtureA.GetShape()
		shapeB := fixtureB.GetShape()
		radiusA := shapeA.GetRadius()
		radiusB := shapeB.GetRadius()
		bodyA := fixtureA.GetBody()
		bodyB := fixtureB.GetBody()
		manifold := c.GetManifold()

		pointCount := manifold.PointCount

		vc := &solver.VelocityConstraints[i]
		vc.Friction = c.GetFriction()
		vc.Restitution = c.GetRestitution()
		vc.TangentSpeed = c.GetTangentSpeed()
		vc.IndexA = bodyA.IslandIndex
		vc.IndexB = bodyB.IslandIndex
		vc.InvMassA = bodyA.InvMass
		vc.InvMassB = bodyB.InvMass
		vc.InvIA = bodyA.InvI
		vc.InvIB = bodyB.InvI
		vc.ContactIndex = i
		vc.PointCount = pointCount
		vc.K.SetZero()
		vc.NormalMass.SetZero()

		pc := &solver.PositionConstraints[i]
		pc.IndexA = bodyA.IslandIndex
		pc.IndexB = bodyB.IslandIndex
		pc.InvMassA = bodyA.InvMass
		pc.InvMassB = bodyB.InvMass
		pc.LocalCenterA = bodyA.Sweep.LocalCenter
		pc.LocalCenterB = bodyB.Sweep.LocalCenter
		pc.InvIA = bodyA.InvI
		pc.InvIB = bodyB.InvI
		pc.LocalNormal = manifold.LocalNormal
		pc.LocalPoint = manifold.LocalPoint
		pc.PointCount = pointCount
		pc.RadiusA = radiusA
		pc.RadiusB = radiusB
		pc.Type = manifold.Type

		for j := 0; j < pointCount; j++ {
			cp := &manifold.Points[j]
			vcp := &vc.Points[j]

			if solver.Step.WarmStarting {
				vcp.NormalImpulse = solver.Step.DtRatio * cp.NormalImpulse
				vcp.TangentImpulse = solver.Step.DtRatio * cp.TangentImpulse
			} else {
				vcp.NormalImpulse = 0.0
				vcp.TangentImpulse = 0.0
			}

			vcp.RA.SetZero()
			vcp.RB.SetZero()
			vcp.NormalMass = 0.0
			vcp.TangentMass = 0.0
			vcp.VelocityBias = 0.0

			pc.LocalPoints[j] = cp.LocalPoint
		}
	}

	return solver
}

func (solver *ContactSolver) Destroy() {
}

// Initialize position dependent portions of the velocity constraints.
func (solver *ContactSolver) InitializeVelocityConstraints() {
	for i := 0; i < solver.Count; i++ {
		vc := &solver.VelocityConstraints[i]
		pc := &solver.PositionConstraints[i]

		radiusA := pc.RadiusA
		radiusB := pc.RadiusB
		manifold := solver.Contacts[vc.ContactIndex].Data().GetManifold()

		indexA := vc.IndexA
		indexB := vc.IndexB

		mA := vc.InvMassA
		mB := vc.InvMassB
		iA := vc.InvIA
		iB := vc.InvIB
		localCenterA := pc.LocalCenterA
		localCenterB := pc.LocalCenterB

		cA := solver.Positions[indexA].C
		aA := solver.Positions[indexA].A
		vA := solver.Velocities[indexA].V
		wA := solver.Velocities[indexA].W

		cB := solver.Positions[indexB].C
		aB := solver.Positions[indexB].A
		vB := solver.Velocities[indexB].V
		wB := solver.Velocities[indexB].W

		xfA := Transform{}
		xfB := Transform{}
		xfA.Q.Set(aA)
		xfB.Q.Set(aB)
		xfA.P = PointSub(cA, RotPointMul(xfA.Q, localCenterA))
		xfB.P = PointSub(cB, RotPointMul(xfB.Q, localCenterB))

		worldManifold := WorldManifold{}
		worldManifold.Initialize(manifold, xfA, radiusA, xfB, radiusB)

		vc.Normal = worldManifold.Normal

		pointCount := vc.PointCount
		for j := 0; j < pointCount; j++ {
			vcp := &vc.Points[j]

			vcp.RA = PointSub(worldManifold.Points[j], cA)
			vcp.RB = PointSub(worldManifold.Points[j], cB)

			rnA := PointCross(vcp.RA, vc.Normal)
			rnB := PointCross(vcp.RB, vc.Normal)

			kNormal := mA + mB + iA*rnA*rnA + iB*rnB*rnB

			if kNormal > 0.0 {
				vcp.NormalMass = 1.0 / kNormal
			} else {
				vcp.NormalMass = 0.0
			}

			tangent := PointCrossVectorScalar(vc.Normal, 1.0)

			rtA := PointCross(vcp.RA, tangent)
			rtB := PointCross(vcp.RB, tangent)

			kTangent := mA + mB + iA*rtA*rtA + iB*rtB*rtB

			if kTangent > 0.0 {
				vcp.TangentMass = 1.0 / kTangent
			} else {
				vcp.TangentMass = 0.0
			}

			// Setup a velocity bias for restitution.
			vcp.VelocityBias = 0.0
			vRel := PointDot(
				vc.Normal,
				PointSub(
					PointSub(
						PointAdd(
							vB,
							PointCrossScalarVector(wB, vcp.RB),
						),
						vA),
					PointCrossScalarVector(wA, vcp.RA),
				),
			)
			if vRel < -_velocityThreshold {
				vcp.VelocityBias = -vc.Restitution * vRel
			}
		}

		// If we have two points, then prepare the block solver.
		if vc.PointCount == 2 && g_blockSolve {
			vcp1 := &vc.Points[0]
			vcp2 := &vc.Points[1]

			rn1A := PointCross(vcp1.RA, vc.Normal)
			rn1B := PointCross(vcp1.RB, vc.Normal)
			rn2A := PointCross(vcp2.RA, vc.Normal)
			rn2B := PointCross(vcp2.RB, vc.Normal)

			k11 := mA + mB + iA*rn1A*rn1A + iB*rn1B*rn1B
			k22 := mA + mB + iA*rn2A*rn2A + iB*rn2B*rn2B
			k12 := mA + mB + iA*rn1A*rn2A + iB*rn1B*rn2B

			// Ensure a reasonable condition number.
			k_maxConditionNumber := 1000.0
			if k11*k11 < k_maxConditionNumber*(k11*k22-k12*k12) {
				// K is safe to invert.
				vc.K.Ex.Set(k11, k12)
				vc.K.Ey.Set(k12, k22)
				vc.NormalMass = vc.K.GetInverse()
			} else {
				// The constraints are redundant, just use one.
				// TODO_ERIN use deepest?
				vc.PointCount = 1
			}
		}
	}
}

func (solver *ContactSolver) WarmStart() {
	// Warm start.
	for i := 0; i < solver.Count; i++ {
		vc := &solver.VelocityConstraints[i]

		indexA := vc.IndexA
		indexB := vc.IndexB
		mA := vc.InvMassA
		iA := vc.InvIA
		mB := vc.InvMassB
		iB := vc.InvIB
		pointCount := vc.PointCount

		vA := solver.Velocities[indexA].V
		wA := solver.Velocities[indexA].W
		vB := solver.Velocities[indexB].V
		wB := solver.Velocities[indexB].W

		normal := vc.Normal
		tangent := PointCrossVectorScalar(normal, 1.0)

		for j := 0; j < pointCount; j++ {
			vcp := &vc.Points[j]
			P := PointAdd(PointMulScalar(vcp.NormalImpulse, normal), PointMulScalar(vcp.TangentImpulse, tangent))
			wA -= iA * PointCross(vcp.RA, P)
			vA.OperatorMinusInplace(PointMulScalar(mA, P))
			wB += iB * PointCross(vcp.RB, P)
			vB.OperatorPlusInplace(PointMulScalar(mB, P))
		}

		solver.Velocities[indexA].V = vA
		solver.Velocities[indexA].W = wA
		solver.Velocities[indexB].V = vB
		solver.Velocities[indexB].W = wB
	}
}

func (solver *ContactSolver) SolveVelocityConstraints() {
	for i := 0; i < solver.Count; i++ {
		vc := &solver.VelocityConstraints[i]

		indexA := vc.IndexA
		indexB := vc.IndexB
		mA := vc.InvMassA
		iA := vc.InvIA
		mB := vc.InvMassB
		iB := vc.InvIB
		pointCount := vc.PointCount

		vA := solver.Velocities[indexA].V
		wA := solver.Velocities[indexA].W
		vB := solver.Velocities[indexB].V
		wB := solver.Velocities[indexB].W

		normal := vc.Normal
		tangent := PointCrossVectorScalar(normal, 1.0)
		friction := vc.Friction

		// Solve tangent constraints first because non-penetration is more important
		// than friction.
		for j := 0; j < pointCount; j++ {
			vcp := &vc.Points[j]

			// Relative velocity at contact
			dv := PointAdd(
				vB,
				PointSub(
					PointSub(
						PointCrossScalarVector(wB, vcp.RB),
						vA,
					),
					PointCrossScalarVector(wA, vcp.RA),
				),
			)

			// Compute tangent force
			vt := PointDot(dv, tangent) - vc.TangentSpeed
			lambda := vcp.TangentMass * (-vt)

			// b2Clamp the accumulated force
			maxFriction := friction * vcp.NormalImpulse
			newImpulse := FloatClamp(vcp.TangentImpulse+lambda, -maxFriction, maxFriction)
			lambda = newImpulse - vcp.TangentImpulse
			vcp.TangentImpulse = newImpulse

			// Apply contact impulse
			P := PointMulScalar(lambda, tangent)

			vA.OperatorMinusInplace(PointMulScalar(mA, P))
			wA -= iA * PointCross(vcp.RA, P)

			vB.OperatorPlusInplace(PointMulScalar(mB, P))
			wB += iB * PointCross(vcp.RB, P)
		}

		// Solve normal constraints
		if pointCount == 1 || g_blockSolve == false {
			for j := 0; j < pointCount; j++ {
				vcp := &vc.Points[j]

				// Relative velocity at contact
				dv := PointAdd(
					vB,
					PointSub(
						PointSub(
							PointCrossScalarVector(wB, vcp.RB),
							vA,
						),
						PointCrossScalarVector(wA, vcp.RA),
					),
				)

				// Compute normal impulse
				vn := PointDot(dv, normal)
				lambda := -vcp.NormalMass * (vn - vcp.VelocityBias)

				// b2Clamp the accumulated impulse
				newImpulse := math.Max(vcp.NormalImpulse+lambda, 0.0)
				lambda = newImpulse - vcp.NormalImpulse
				vcp.NormalImpulse = newImpulse

				// Apply contact impulse
				P := PointMulScalar(lambda, normal)
				vA.OperatorMinusInplace(PointMulScalar(mA, P))
				wA -= iA * PointCross(vcp.RA, P)

				vB.OperatorPlusInplace(PointMulScalar(mB, P))
				wB += iB * PointCross(vcp.RB, P)
			}
		} else {
			// Block solver developed in collaboration with Dirk Gregorius (back in 01/07 on Box2D_Lite).
			// Build the mini LCP for this contact patch
			//
			// vn = A * x + b, vn >= 0, x >= 0 and vn_i * x_i = 0 with i = 1..2
			//
			// A = J * W * JT and J = ( -n, -r1 x n, n, r2 x n )
			// b = vn0 - velocityBias
			//
			// The system is solved using the "Total enumeration method" (s. Murty). The complementary constraint vn_i * x_i
			// implies that we must have in any solution either vn_i = 0 or x_i = 0. So for the 2D contact problem the cases
			// vn1 = 0 and vn2 = 0, x1 = 0 and x2 = 0, x1 = 0 and vn2 = 0, x2 = 0 and vn1 = 0 need to be tested. The first valid
			// solution that satisfies the problem is chosen.
			//
			// In order to account of the accumulated impulse 'a' (because of the iterative nature of the solver which only requires
			// that the accumulated impulse is clamped and not the incremental impulse) we change the impulse variable (x_i).
			//
			// Substitute:
			//
			// x = a + d
			//
			// a := old total impulse
			// x := new total impulse
			// d := incremental impulse
			//
			// For the current iteration we extend the formula for the incremental impulse
			// to compute the new total impulse:
			//
			// vn = A * d + b
			//    = A * (x - a) + b
			//    = A * x + b - A * a
			//    = A * x + b'
			// b' = b - A * a;

			cp1 := &vc.Points[0]
			cp2 := &vc.Points[1]

			a := Point{X: cp1.NormalImpulse, Y: cp2.NormalImpulse}

			// Relative velocity at contact
			dv1 := PointAdd(vB, PointSub(PointSub(PointCrossScalarVector(wB, cp1.RB), vA), PointCrossScalarVector(wA, cp1.RA)))
			dv2 := PointAdd(vB, PointSub(PointSub(PointCrossScalarVector(wB, cp2.RB), vA), PointCrossScalarVector(wA, cp2.RA)))

			// Compute normal velocity
			vn1 := PointDot(dv1, normal)
			vn2 := PointDot(dv2, normal)

			b := Point{}
			b.X = vn1 - cp1.VelocityBias
			b.Y = vn2 - cp2.VelocityBias

			// Compute b'
			b.OperatorMinusInplace(PointMat22Mul(vc.K, a))

			const k_errorTol = 0.001
			// _NOT_USED(k_errorTol);

			for {
				//
				// Case 1: vn = 0
				//
				// 0 = A * x + b'
				//
				// Solve for x:
				//
				// x = - inv(A) * b'
				//
				x := PointMat22Mul(vc.NormalMass, b).OperatorNegate()

				if x.X >= 0.0 && x.Y >= 0.0 {
					// Get the incremental impulse
					d := PointSub(x, a)

					// Apply incremental impulse
					P1 := PointMulScalar(d.X, normal)
					P2 := PointMulScalar(d.Y, normal)
					vA.OperatorMinusInplace(PointMulScalar(mA, PointAdd(P1, P2)))
					wA -= iA * (PointCross(cp1.RA, P1) + PointCross(cp2.RA, P2))

					vB.OperatorPlusInplace(PointMulScalar(mB, PointAdd(P1, P2)))
					wB += iB * (PointCross(cp1.RB, P1) + PointCross(cp2.RB, P2))

					// Accumulate
					cp1.NormalImpulse = x.X
					cp2.NormalImpulse = x.Y

					break
				}

				//
				// Case 2: vn1 = 0 and x2 = 0
				//
				//   0 = a11 * x1 + a12 * 0 + b1'
				// vn2 = a21 * x1 + a22 * 0 + b2'
				//
				x.X = -cp1.NormalMass * b.X
				x.Y = 0.0
				vn1 = 0.0
				vn2 = vc.K.Ex.Y*x.X + b.Y
				if x.X >= 0.0 && vn2 >= 0.0 {
					// Get the incremental impulse
					d := PointSub(x, a)

					// Apply incremental impulse
					P1 := PointMulScalar(d.X, normal)
					P2 := PointMulScalar(d.Y, normal)
					vA.OperatorMinusInplace(PointMulScalar(mA, PointAdd(P1, P2)))
					wA -= iA * (PointCross(cp1.RA, P1) + PointCross(cp2.RA, P2))

					vB.OperatorPlusInplace(PointMulScalar(mB, PointAdd(P1, P2)))
					wB += iB * (PointCross(cp1.RB, P1) + PointCross(cp2.RB, P2))

					// Accumulate
					cp1.NormalImpulse = x.X
					cp2.NormalImpulse = x.Y

					break
				}

				//
				// Case 3: vn2 = 0 and x1 = 0
				//
				// vn1 = a11 * 0 + a12 * x2 + b1'
				//   0 = a21 * 0 + a22 * x2 + b2'
				//
				x.X = 0.0
				x.Y = -cp2.NormalMass * b.Y
				vn1 = vc.K.Ey.X*x.Y + b.X
				vn2 = 0.0

				if x.Y >= 0.0 && vn1 >= 0.0 {
					// Resubstitute for the incremental impulse
					d := PointSub(x, a)

					// Apply incremental impulse
					P1 := PointMulScalar(d.X, normal)
					P2 := PointMulScalar(d.Y, normal)
					vA.OperatorMinusInplace(PointMulScalar(mA, PointAdd(P1, P2)))
					wA -= iA * (PointCross(cp1.RA, P1) + PointCross(cp2.RA, P2))

					vB.OperatorPlusInplace(PointMulScalar(mB, PointAdd(P1, P2)))
					wB += iB * (PointCross(cp1.RB, P1) + PointCross(cp2.RB, P2))

					// Accumulate
					cp1.NormalImpulse = x.X
					cp2.NormalImpulse = x.Y

					break
				}

				//
				// Case 4: x1 = 0 and x2 = 0
				//
				// vn1 = b1
				// vn2 = b2;
				x.X = 0.0
				x.Y = 0.0
				vn1 = b.X
				vn2 = b.Y

				if vn1 >= 0.0 && vn2 >= 0.0 {
					// Resubstitute for the incremental impulse
					d := PointSub(x, a)

					// Apply incremental impulse
					P1 := PointMulScalar(d.X, normal)
					P2 := PointMulScalar(d.Y, normal)
					vA.OperatorMinusInplace(PointMulScalar(mA, PointAdd(P1, P2)))
					wA -= iA * (PointCross(cp1.RA, P1) + PointCross(cp2.RA, P2))

					vB.OperatorPlusInplace(PointMulScalar(mB, PointAdd(P1, P2)))
					wB += iB * (PointCross(cp1.RB, P1) + PointCross(cp2.RB, P2))

					// Accumulate
					cp1.NormalImpulse = x.X
					cp2.NormalImpulse = x.Y

					break
				}

				// No solution, give up. This is hit sometimes, but it doesn't seem to matter.
				break
			}
		}

		solver.Velocities[indexA].V = vA
		solver.Velocities[indexA].W = wA
		solver.Velocities[indexB].V = vB
		solver.Velocities[indexB].W = wB
	}
}

func (solver *ContactSolver) StoreImpulses() {
	for i := 0; i < solver.Count; i++ {
		vc := &solver.VelocityConstraints[i]
		manifold := solver.Contacts[vc.ContactIndex].Data().GetManifold()

		for j := 0; j < vc.PointCount; j++ {
			manifold.Points[j].NormalImpulse = vc.Points[j].NormalImpulse
			manifold.Points[j].TangentImpulse = vc.Points[j].TangentImpulse
		}
	}
}

type PositionSolverManifold struct {
	Normal     Point
	Point      Point
	Separation float64
}

func MakePositionSolverManifold() PositionSolverManifold {
	return PositionSolverManifold{}
}

func (solvermanifold *PositionSolverManifold) Initialize(pc *ContactPositionConstraint, xfA Transform, xfB Transform, index int) {

	switch pc.Type {
	case ManifoldTypeCircles:
		{
			pointA := TransformPointMul(xfA, pc.LocalPoint)
			pointB := TransformPointMul(xfB, pc.LocalPoints[0])
			solvermanifold.Normal = PointSub(pointB, pointA)
			solvermanifold.Normal.Normalize()
			solvermanifold.Point = PointMulScalar(0.5, PointAdd(pointA, pointB))
			solvermanifold.Separation = PointDot(PointSub(pointB, pointA), solvermanifold.Normal) - pc.RadiusA - pc.RadiusB
		}
		break

	case ManifoldTypeFaceA:
		{
			solvermanifold.Normal = RotPointMul(xfA.Q, pc.LocalNormal)
			planePoint := TransformPointMul(xfA, pc.LocalPoint)

			clipPoint := TransformPointMul(xfB, pc.LocalPoints[index])
			solvermanifold.Separation = PointDot(PointSub(clipPoint, planePoint), solvermanifold.Normal) - pc.RadiusA - pc.RadiusB
			solvermanifold.Point = clipPoint
		}
		break

	case ManifoldTypeFaceB:
		{
			solvermanifold.Normal = RotPointMul(xfB.Q, pc.LocalNormal)
			planePoint := TransformPointMul(xfB, pc.LocalPoint)

			clipPoint := TransformPointMul(xfA, pc.LocalPoints[index])
			solvermanifold.Separation = PointDot(PointSub(clipPoint, planePoint), solvermanifold.Normal) - pc.RadiusA - pc.RadiusB
			solvermanifold.Point = clipPoint

			// Ensure normal points from A to B
			solvermanifold.Normal = solvermanifold.Normal.OperatorNegate()
		}
		break
	}
}

// Sequential solver.
func (solver *ContactSolver) SolvePositionConstraints() bool {

	minSeparation := 0.0

	for i := 0; i < solver.Count; i++ {
		pc := &solver.PositionConstraints[i]

		indexA := pc.IndexA
		indexB := pc.IndexB
		localCenterA := pc.LocalCenterA
		mA := pc.InvMassA
		iA := pc.InvIA
		localCenterB := pc.LocalCenterB
		mB := pc.InvMassB
		iB := pc.InvIB
		pointCount := pc.PointCount

		cA := solver.Positions[indexA].C
		aA := solver.Positions[indexA].A

		cB := solver.Positions[indexB].C
		aB := solver.Positions[indexB].A

		// Solve normal constraints
		for j := 0; j < pointCount; j++ {
			xfA := Transform{}
			xfB := Transform{}

			xfA.Q.Set(aA)
			xfB.Q.Set(aB)
			xfA.P = PointSub(cA, RotPointMul(xfA.Q, localCenterA))
			xfB.P = PointSub(cB, RotPointMul(xfB.Q, localCenterB))

			psm := MakePositionSolverManifold()
			psm.Initialize(pc, xfA, xfB, j)
			normal := psm.Normal

			point := psm.Point
			separation := psm.Separation

			rA := PointSub(point, cA)
			rB := PointSub(point, cB)

			// Track max constraint error.
			minSeparation = math.Min(minSeparation, separation)

			// Prevent large corrections and allow slop.
			C := FloatClamp(_baumgarte*(separation+_linearSlop), -_maxLinearCorrection, 0.0)

			// Compute the effective mass.
			rnA := PointCross(rA, normal)
			rnB := PointCross(rB, normal)
			K := mA + mB + iA*rnA*rnA + iB*rnB*rnB

			// Compute normal impulse
			impulse := 0.0
			if K > 0.0 {
				impulse = -C / K
			}

			P := PointMulScalar(impulse, normal)

			cA.OperatorMinusInplace(PointMulScalar(mA, P))
			aA -= iA * PointCross(rA, P)

			cB.OperatorPlusInplace(PointMulScalar(mB, P))
			aB += iB * PointCross(rB, P)
		}

		solver.Positions[indexA].C = cA
		solver.Positions[indexA].A = aA

		solver.Positions[indexB].C = cB
		solver.Positions[indexB].A = aB
	}

	// We can't expect minSpeparation >= -b2_linearSlop because we don't
	// push the separation above -b2_linearSlop.
	return minSeparation >= -3.0*_linearSlop
}

// Sequential position solver for position constraints.
func (solver *ContactSolver) SolveTOIPositionConstraints(toiIndexA int, toiIndexB int) bool {

	minSeparation := 0.0

	for i := 0; i < solver.Count; i++ {
		pc := &solver.PositionConstraints[i]

		indexA := pc.IndexA
		indexB := pc.IndexB
		localCenterA := pc.LocalCenterA
		localCenterB := pc.LocalCenterB
		pointCount := pc.PointCount

		mA := 0.0
		iA := 0.0
		if indexA == toiIndexA || indexA == toiIndexB {
			mA = pc.InvMassA
			iA = pc.InvIA
		}

		mB := 0.0
		iB := 0.0
		if indexB == toiIndexA || indexB == toiIndexB {
			mB = pc.InvMassB
			iB = pc.InvIB
		}

		cA := solver.Positions[indexA].C
		aA := solver.Positions[indexA].A

		cB := solver.Positions[indexB].C
		aB := solver.Positions[indexB].A

		// Solve normal constraints
		for j := 0; j < pointCount; j++ {
			xfA := Transform{}
			xfB := Transform{}

			xfA.Q.Set(aA)
			xfB.Q.Set(aB)
			xfB.P = PointSub(cB, RotPointMul(xfB.Q, localCenterB))
			xfA.P = PointSub(cA, RotPointMul(xfA.Q, localCenterA))

			psm := MakePositionSolverManifold()
			psm.Initialize(pc, xfA, xfB, j)
			normal := psm.Normal

			point := psm.Point
			separation := psm.Separation

			rA := PointSub(point, cA)
			rB := PointSub(point, cB)

			// Track max constraint error.
			minSeparation = math.Min(minSeparation, separation)

			// Prevent large corrections and allow slop.
			C := FloatClamp(_toiBaugarte*(separation+_linearSlop), -_maxLinearCorrection, 0.0)

			// Compute the effective mass.
			rnA := PointCross(rA, normal)
			rnB := PointCross(rB, normal)
			K := mA + mB + iA*rnA*rnA + iB*rnB*rnB

			// Compute normal impulse
			impulse := 0.0
			if K > 0.0 {
				impulse = -C / K
			}

			P := PointMulScalar(impulse, normal)

			cA.OperatorMinusInplace(PointMulScalar(mA, P))
			aA -= iA * PointCross(rA, P)

			cB.OperatorPlusInplace(PointMulScalar(mB, P))
			aB += iB * PointCross(rB, P)
		}

		solver.Positions[indexA].C = cA
		solver.Positions[indexA].A = aA

		solver.Positions[indexB].C = cB
		solver.Positions[indexB].A = aB
	}

	// We can't expect minSpeparation >= -b2_linearSlop because we don't
	// push the separation above -b2_linearSlop.
	return minSeparation >= -1.5*_linearSlop
}
