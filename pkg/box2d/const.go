package box2d

import (
	"math"
)

const (

	// A small length used as a collision and constraint tolerance. Usually it is
	// chosen to be numerically significant, but visually insignificant.
	_linearSlop = 0.005

	// A small angle used as a collision and constraint tolerance. Usually it is
	// chosen to be numerically significant, but visually insignificant.
	_angularSlop = 2.0 / 180.0 * math.Pi
	_epsilon     = math.SmallestNonzeroFloat64

	// The radius of the polygon/edge shape skin. This should not be modified. Making
	// this smaller means polygons will have an insufficient buffer for continuous collision.
	// Making it larger may create artifacts for vertex collision.
	_polygonRadius = 2.0 * _linearSlop

	// The maximum number of contact points between two convex shapes. Do
	// not change this value.
	_maxManifoldPoints = 2

	// The maximum number of vertices on a convex polygon.
	_maxPolygonVertices = 64

	/// This is used to fatten AABBs in the dynamic tree. This allows proxies
	/// to move by a small amount without triggering a tree adjustment.
	/// This is in meters.
	_aabbExtension = 0.1

	// This is used to fatten AABBs in the dynamic tree. This is used to predict
	// the future position based on the current displacement.
	// This is a dimensionless multiplier.
	_aabbMultiplier = 2.0

	// Maximum number of sub-steps per contact in continuous physics simulation.
	_maxSubSteps = 128

	// Maximum number of contacts to be handled to solve a TOI impact.
	_maxTOIContacts = 128

	// A velocity threshold for elastic collisions. Any collision with a relative linear
	// velocity below this threshold will be treated as inelastic.
	_velocityThreshold = 5.0

	// The maximum linear position correction used when solving constraints. This helps to
	// prevent overshoot.
	_maxLinearCorrection = 32.0

	// The maximum linear velocity of a body. This limit is very large and is used
	// to prevent numerical problems. You shouldn't need to adjust this.
	_maxTranslation        = 500.0
	_maxTranslationSquared = (_maxTranslation * _maxTranslation)

	// The maximum angular velocity of a body. This limit is very large and is used
	// to prevent numerical problems. You shouldn't need to adjust this.
	_maxRotation        = (2.0 * math.Pi)
	_maxRotationSquared = (_maxRotation * _maxRotation)

	// This scale factor controls how fast overlap is resolved. Ideally this would be 1 so
	// that overlap is removed in one time step. However using values close to 1 often lead
	// to overshoot.
	_baumgarte   = 0.2
	_toiBaugarte = 0.75

	// The time that a body must be still before it will go to sleep.
	_timeToSleep = 0.5

	// A body cannot sleep if its linear velocity is above this tolerance.
	_linearSleepTolerance = 0.5

	// A body cannot sleep if its angular velocity is above this tolerance.
	_angularSleepTolerance = (2.0 / 180.0 * math.Pi)
)
