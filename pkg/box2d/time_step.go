package box2d

/// This is an internal structure.
type TimeStep struct {
	Dt                 float64 // time step
	InverseDt          float64 // inverse time step (0 if dt == 0).
	DtRatio            float64 // dt * inv_dt0
	VelocityIterations int
	PositionIterations int
	WarmStarting       bool
}

/// This is an internal structure.
type Position struct {
	C Point
	A float64
}

/// This is an internal structure.
type Velocity struct {
	V Point
	W float64
}

/// Solver Data
type SolverData struct {
	Step       TimeStep
	Positions  []Position
	Velocities []Velocity
}
