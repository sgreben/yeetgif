package box2d

type DestructionListenerInterface interface {
	/// Called when any fixture is about to be destroyed due
	/// to the destruction of its parent body.
	SayGoodbyeToFixture(fixture *Fixture)
}

type ContactFilterInterface interface {
	ShouldCollide(fixtureA *Fixture, fixtureB *Fixture) bool
}

/// Contact impulses for reporting. Impulses are used instead of forces because
/// sub-step forces may approach infinity for rigid body collisions. These
/// match up one-to-one with the contact points in b2Manifold.
type ContactImpulse struct {
	NormalImpulses  [_maxManifoldPoints]float64
	TangentImpulses [_maxManifoldPoints]float64
	Count           int
}

func MakeContactImpulse() ContactImpulse {
	return ContactImpulse{}
}

type ContactListenerInterface interface {
	/// Called when two fixtures begin to touch.
	BeginContact(contact ContactInterface) // contact has to be backed by a pointer

	/// Called when two fixtures cease to touch.
	EndContact(contact ContactInterface) // contact has to be backed by a pointer

	/// This is called after a contact is updated. This allows you to inspect a
	/// contact before it goes to the solver. If you are careful, you can modify the
	/// contact manifold (e.g. disable contact).
	/// A copy of the old manifold is provided so that you can detect changes.
	/// Note: this is called only for awake bodies.
	/// Note: this is called even when the number of contact points is zero.
	/// Note: this is not called for sensors.
	/// Note: if you set the number of contact points to zero, you will not
	/// get an EndContact callback. However, you may get a BeginContact callback
	/// the next step.
	PreSolve(contact ContactInterface, oldManifold Manifold) // contact has to be backed by a pointer

	/// This lets you inspect a contact after the solver is finished. This is useful
	/// for inspecting impulses.
	/// Note: the contact manifold does not include time of impact impulses, which can be
	/// arbitrarily large if the sub-step is small. Hence the impulse is provided explicitly
	/// in a separate data structure.
	/// Note: this is only called for contacts that are touching, solid, and awake.
	PostSolve(contact ContactInterface, impulse *ContactImpulse) // contact has to be backed by a pointer
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// WorldCallbacks.cpp
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

type BroadPhaseQueryCallback func(fixture *Fixture) bool

type ContactFilter struct {
}

// Return true if contact calculations should be performed between these two shapes.
// If you implement your own collision filter you may want to build from this implementation.
func (cf *ContactFilter) ShouldCollide(fixtureA *Fixture, fixtureB *Fixture) bool {
	filterA := fixtureA.GetFilterData()
	filterB := fixtureB.GetFilterData()

	if filterA.GroupIndex == filterB.GroupIndex && filterA.GroupIndex != 0 {
		return filterA.GroupIndex > 0
	}

	collide := (filterA.MaskBits&filterB.CategoryBits) != 0 && (filterA.CategoryBits&filterB.MaskBits) != 0
	return collide
}

/// Called for each fixture found in the query. You control how the ray cast
/// proceeds by returning a float:
/// return -1: ignore this fixture and continue
/// return 0: terminate the ray cast
/// return fraction: clip the ray to this point
/// return 1: don't clip the ray and continue
/// @param fixture the fixture hit by the ray
/// @param point the point of initial intersection
/// @param normal the normal vector at the point of intersection
/// @return -1 to filter, 0 to terminate, fraction to clip the ray for
/// closest hit, 1 to continue
type RaycastCallback func(fixture *Fixture, point Point, normal Point, fraction float64) float64
