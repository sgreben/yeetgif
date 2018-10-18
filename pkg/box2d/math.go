package box2d

import (
	"image"
	"math"
)

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func AbsInt(v int) int {
	if v < 0 {
		return v * -1
	}

	return v
}

/// This function is used to ensure that a floating point number is not a NaN or infinity.
func IsValid(x float64) bool {
	return !math.IsNaN(x) && !math.IsInf(x, 0)
}

/// This is a approximate yet fast inverse square-root.
func InvSqrt(x float64) float64 {
	// https://groups.google.com/forum/#!topic/golang-nuts/8vaZ1ERYIQ0
	// Faster with math.Sqrt
	return 1.0 / math.Sqrt(x)
}

type Point struct {
	X, Y float64
}

/// Set this vector to all zeros.
func (v *Point) SetZero() {
	v.X = 0.0
	v.Y = 0.0
}

func (v *Point) SetImagePoint(p *image.Point) {
	v.X = float64(p.X)
	v.Y = float64(p.Y)
}

func (v Point) ImagePoint() image.Point {
	return image.Point{
		X: int(v.X),
		Y: int(v.Y),
	}
}

/// Set this vector to some specified coordinates.
func (v *Point) Set(x, y float64) {
	v.X = x
	v.Y = y
}

/// Negate this vector.
func (v Point) OperatorNegate() Point {
	return Point{X: -v.X, Y: -v.Y}
}

/// Read from and indexed element.
func (v Point) OperatorIndexGet(i int) float64 {
	if i == 0 {
		return v.X
	}

	return v.Y
}

/// Write to an indexed element.
func (v *Point) OperatorIndexSet(i int, value float64) {
	if i == 0 {
		v.X = value
	}

	v.Y = value
}

/// Add a vector to this vector.
func (v *Point) OperatorPlusInplace(other Point) {
	v.X += other.X
	v.Y += other.Y
}

/// Subtract a vector from this vector.
func (v *Point) OperatorMinusInplace(other Point) {
	v.X -= other.X
	v.Y -= other.Y
}

/// Multiply this vector by a scalar.
func (v *Point) OperatorScalarMulInplace(a float64) {
	v.X *= a
	v.Y *= a
}

/// Get the length of this vector (the norm).
func (v Point) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/// Get the length squared. For performance, use this instead of
/// b2Point::Length (if possible).
func (v Point) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

/// Convert this vector into a unit vector. Returns the length.
func (v *Point) Normalize() float64 {

	length := v.Length()

	if length < _epsilon {
		return 0.0
	}

	invLength := 1.0 / length
	v.X *= invLength
	v.Y *= invLength

	return length
}

/// Does this vector contain finite coordinates?
func (v Point) IsValid() bool {
	return IsValid(v.X) && IsValid(v.Y)
}

func (v Point) Clone() Point {
	return Point{X: v.X, Y: v.Y}
}

// Mat22 is a 2-by-2 matrix. Stored in column-major order.
type Mat22 struct {
	Ex, Ey Point
}

/// Construct this matrix using columns.
func MakeMat22FromColumns(c1, c2 Point) Mat22 {
	return Mat22{
		Ex: c1,
		Ey: c2,
	}
}

/// Initialize this matrix using columns.
func (m *Mat22) Set(c1 Point, c2 Point) {
	m.Ex = c1
	m.Ey = c2
}

/// Set this to the identity matrix.
func (m *Mat22) SetIdentity() {
	m.Ex.X = 1.0
	m.Ey.X = 0.0
	m.Ex.Y = 0.0
	m.Ey.Y = 1.0
}

/// Set this matrix to all zeros.
func (m *Mat22) SetZero() {
	m.Ex.X = 0.0
	m.Ey.X = 0.0
	m.Ex.Y = 0.0
	m.Ey.Y = 0.0
}

func (m Mat22) GetInverse() Mat22 {

	a := m.Ex.X
	b := m.Ey.X
	c := m.Ex.Y
	d := m.Ey.Y

	B := Mat22{}

	det := a*d - b*c
	if det != 0.0 {
		det = 1.0 / det
	}

	B.Ex.X = det * d
	B.Ey.X = -det * b
	B.Ex.Y = -det * c
	B.Ey.Y = det * a

	return B
}

/// Solve A * x = b, where b is a column vector. This is more efficient
/// than computing the inverse in one-shot cases.
func (m Mat22) Solve(b Point) Point {

	a11 := m.Ex.X
	a12 := m.Ey.X
	a21 := m.Ex.Y
	a22 := m.Ey.Y
	det := a11*a22 - a12*a21

	if det != 0.0 {
		det = 1.0 / det
	}

	return Point{
		X: det * (a22*b.X - a12*b.Y),
		Y: det * (a11*b.Y - a21*b.X),
	}
}

///////////////////////////////////////////////////////////////////////////////
/// Rotation
///////////////////////////////////////////////////////////////////////////////
type Rot struct {
	/// Sine and cosine
	S, C float64
}

/// Set using an angle in radians.
func (r *Rot) Set(anglerad float64) {
	r.S = math.Sin(anglerad)
	r.C = math.Cos(anglerad)
}

/// Set to the identity rotation
func (r *Rot) SetIdentity() {
	r.S = 0.0
	r.C = 1.0
}

// Transform is a translation and rotation. It is used to represent
// the position and orientation of rigid frames.
type Transform struct {
	P Point
	Q Rot
}

/// Set this to the identity transform.
func (t *Transform) SetIdentity() {
	t.P.SetZero()
	t.Q.SetIdentity()
}

/// Set this based on the position and angle.
func (t *Transform) Set(position Point, anglerad float64) {
	t.P = position
	t.Q.Set(anglerad)
}

// Sweep is a structure that describes the motion of a body/shape for TOI computation.
// Shapes are defined with respect to the body origin, which may
// no coincide with the center of mass. However, to support dynamics
// we must interpolate the center of mass position.
type Sweep struct {
	LocalCenter Point   ///< local center of mass position
	C0, C       Point   ///< center world positions
	A0, A       float64 ///< world angles

	/// Fraction of the current time step in the range [0,1]
	/// c0 and a0 are the positions at alpha0.
	Alpha0 float64
}

/// Perform the dot product on two vectors.
func PointDot(a, b Point) float64 {
	return a.X*b.X + a.Y*b.Y
}

/// Perform the cross product on two vectors. In 2D this produces a scalar.
func PointCross(a, b Point) float64 {
	return a.X*b.Y - a.Y*b.X
}

/// Perform the cross product on a vector and a scalar. In 2D this produces
/// a vector.
func PointCrossVectorScalar(a Point, s float64) Point {
	return Point{X: s * a.Y, Y: -s * a.X}
}

/// Perform the cross product on a scalar and a vector. In 2D this produces
/// a vector.
func PointCrossScalarVector(s float64, a Point) Point {
	return Point{X: -s * a.Y, Y: s * a.X}
}

/// Multiply a matrix times a vector. If a rotation matrix is provided,
/// then this transforms the vector from one frame to another.
func PointMat22Mul(A Mat22, v Point) Point {
	return Point{X: A.Ex.X*v.X + A.Ey.X*v.Y, Y: A.Ex.Y*v.X + A.Ey.Y*v.Y}
}

/// Add two vectors component-wise.
func PointAdd(a, b Point) Point {
	return Point{X: a.X + b.X, Y: a.Y + b.Y}
}

// Subtract two vectors component-wise.
func PointSub(a, b Point) Point {
	return Point{X: a.X - b.X, Y: a.Y - b.Y}
}

func PointMulScalar(s float64, a Point) Point {
	return Point{X: s * a.X, Y: s * a.Y}
}

func PointDistance(a, b Point) float64 {
	return PointSub(a, b).Length()
}

func PointDistanceSquared(a, b Point) float64 {
	c := PointSub(a, b)
	return PointDot(c, c)
}

/// Multiply two rotations: q * r
func RotMul(q, r Rot) Rot {
	return Rot{
		S: q.S*r.C + q.C*r.S,
		C: q.C*r.C - q.S*r.S,
	}
}

/// Transpose multiply two rotations: qT * r
func RotMulT(q, r Rot) Rot {
	return Rot{
		S: q.C*r.S - q.S*r.C,
		C: q.C*r.C + q.S*r.S,
	}
}

/// Rotate a vector
func RotPointMul(q Rot, v Point) Point {
	return Point{
		X: q.C*v.X - q.S*v.Y,
		Y: q.S*v.X + q.C*v.Y,
	}
}

/// Inverse rotate a vector
func RotPointMulT(q Rot, v Point) Point {
	return Point{
		X: q.C*v.X + q.S*v.Y,
		Y: -q.S*v.X + q.C*v.Y,
	}
}

func TransformPointMul(T Transform, v Point) Point {
	return Point{
		X: (T.Q.C*v.X - T.Q.S*v.Y) + T.P.X,
		Y: (T.Q.S*v.X + T.Q.C*v.Y) + T.P.Y,
	}
}

func TransformPointMulT(T Transform, v Point) Point {
	px := v.X - T.P.X
	py := v.Y - T.P.Y
	x := (T.Q.C*px + T.Q.S*py)
	y := (-T.Q.S*px + T.Q.C*py)

	return Point{X: x, Y: y}
}

func TransformMul(A, B Transform) Transform {
	q := RotMul(A.Q, B.Q)
	p := PointAdd(RotPointMul(A.Q, B.P), A.P)

	return Transform{P: p, Q: q}
}

func TransformMulT(A, B Transform) Transform {
	q := RotMulT(A.Q, B.Q)
	p := RotPointMulT(A.Q, PointSub(B.P, A.P))

	return Transform{P: p, Q: q}
}

func PointAbs(a Point) Point {
	return Point{X: math.Abs(a.X), Y: math.Abs(a.Y)}
}

func PointMin(a, b Point) Point {
	return Point{
		X: math.Min(a.X, b.X),
		Y: math.Min(a.Y, b.Y),
	}
}

func PointMax(a, b Point) Point {
	return Point{
		X: math.Max(a.X, b.X),
		Y: math.Max(a.Y, b.Y),
	}
}

func FloatClamp(a, low, high float64) float64 {
	return math.Max(
		low,
		math.Min(a, high),
	)
}

func (sweep Sweep) GetTransform(xf *Transform, beta float64) {

	xf.P = PointAdd(
		PointMulScalar(1.0-beta, sweep.C0),
		PointMulScalar(beta, sweep.C),
	)

	angle := (1.0-beta)*sweep.A0 + beta*sweep.A
	xf.Q.Set(angle)

	// Shift to origin
	xf.P.OperatorMinusInplace(RotPointMul(xf.Q, sweep.LocalCenter))
}

func (sweep *Sweep) Advance(alpha float64) {
	beta := (alpha - sweep.Alpha0) / (1.0 - sweep.Alpha0)
	sweep.C0.OperatorPlusInplace(PointMulScalar(beta, PointSub(sweep.C, sweep.C0)))
	sweep.A0 += beta * (sweep.A - sweep.A0)
	sweep.Alpha0 = alpha
}

/// Normalize an angle in radians to be between -pi and pi
func (sweep *Sweep) Normalize() {
	twoPi := 2.0 * math.Pi
	d := twoPi * math.Floor(sweep.A0/twoPi)
	sweep.A0 -= d
	sweep.A -= d
}
