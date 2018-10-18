package box2d

import (
	"sort"
)

type BroadPhaseAddPairCallback func(userDataA interface{}, userDataB interface{})

type Pair struct {
	ProxyIdA int
	ProxyIdB int
}

const nullProxyID = -1

type BroadPhase struct {
	Tree DynamicTree

	ProxyCount int

	MoveBuffer   []int
	MoveCapacity int
	MoveCount    int

	PairBuffer   []Pair
	PairCapacity int
	PairCount    int

	QueryProxyId int
}

type PairByLessThan []Pair

func (a PairByLessThan) Len() int      { return len(a) }
func (a PairByLessThan) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PairByLessThan) Less(i, j int) bool {
	return PairLessThan(a[i], a[j])
}

/// This is used to sort pairs.
func PairLessThan(pair1 Pair, pair2 Pair) bool {
	if pair1.ProxyIdA < pair2.ProxyIdA {
		return true
	}

	if pair1.ProxyIdA == pair2.ProxyIdA {
		return pair1.ProxyIdB < pair2.ProxyIdB
	}

	return false
}

func (bp BroadPhase) GetUserData(proxyId int) interface{} {
	return bp.Tree.GetUserData(proxyId)
}

func (bp BroadPhase) TestOverlap(proxyIdA int, proxyIdB int) bool {
	return TestOverlapBoundingBoxes(
		bp.Tree.GetFatAABB(proxyIdA),
		bp.Tree.GetFatAABB(proxyIdB),
	)
}

func (bp BroadPhase) GetFatAABB(proxyId int) AABB {
	return bp.Tree.GetFatAABB(proxyId)
}

func (bp BroadPhase) GetProxyCount() int {
	return bp.ProxyCount
}

func (bp BroadPhase) GetTreeHeight() int {
	return bp.Tree.GetHeight()
}

func (bp BroadPhase) GetTreeBalance() int {
	return bp.Tree.GetMaxBalance()
}

func (bp BroadPhase) GetTreeQuality() float64 {
	return bp.Tree.GetAreaRatio()
}

func (bp *BroadPhase) UpdatePairs(addPairCallback BroadPhaseAddPairCallback) {
	// Reset pair buffer
	bp.PairCount = 0

	// Perform tree queries for all moving proxies.
	for i := 0; i < bp.MoveCount; i++ {
		bp.QueryProxyId = bp.MoveBuffer[i]
		if bp.QueryProxyId == nullProxyID {
			continue
		}

		// We have to query the tree with the fat AABB so that
		// we don't fail to create a pair that may touch later.
		fatAABB := bp.Tree.GetFatAABB(bp.QueryProxyId)

		// Query tree, create pairs and add them pair buffer.
		bp.Tree.Query(bp.QueryCallback, fatAABB)
	}

	// Reset move buffer
	bp.MoveCount = 0

	// Sort the pair buffer to expose duplicates.
	sort.Sort(PairByLessThan(bp.PairBuffer[:bp.PairCount]))

	// Send the pairs back to the client.
	i := 0
	for i < bp.PairCount {
		primaryPair := bp.PairBuffer[i]
		userDataA := bp.Tree.GetUserData(primaryPair.ProxyIdA)
		userDataB := bp.Tree.GetUserData(primaryPair.ProxyIdB)

		addPairCallback(userDataA, userDataB)
		i++

		// Skip any duplicate pairs.
		for i < bp.PairCount {
			pair := bp.PairBuffer[i]
			if pair.ProxyIdA != primaryPair.ProxyIdA || pair.ProxyIdB != primaryPair.ProxyIdB {
				break
			}
			i++
		}
	}

	// // Try to keep the tree balanced.
	// //m_tree.Rebalance(4);
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// BroadPhase.cpp
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func MakeBroadPhase() BroadPhase {

	pairCapacity := 16
	moveCapacity := 16

	tree := MakeDynamicTree()

	return BroadPhase{
		Tree:       tree,
		ProxyCount: 0,

		PairCapacity: pairCapacity,
		PairCount:    0,
		PairBuffer:   make([]Pair, pairCapacity),

		MoveCapacity: moveCapacity,
		MoveCount:    0,
		MoveBuffer:   make([]int, moveCapacity),
	}
}

func (bp *BroadPhase) CreateProxy(aabb AABB, userData interface{}) int {
	proxyId := bp.Tree.CreateProxy(aabb, userData)
	bp.ProxyCount++
	bp.BufferMove(proxyId)
	return proxyId
}

func (bp *BroadPhase) DestroyProxy(proxyId int) {
	bp.UnBufferMove(proxyId)
	bp.ProxyCount--
	bp.Tree.DestroyProxy(proxyId)
}

func (bp *BroadPhase) MoveProxy(proxyId int, aabb AABB, displacement Point) {
	buffer := bp.Tree.MoveProxy(proxyId, aabb, displacement)
	if buffer {
		bp.BufferMove(proxyId)
	}
}

func (bp *BroadPhase) TouchProxy(proxyId int) {
	bp.BufferMove(proxyId)
}

func (bp *BroadPhase) BufferMove(proxyId int) {
	if bp.MoveCount == bp.MoveCapacity {
		bp.MoveBuffer = append(bp.MoveBuffer, make([]int, bp.MoveCapacity)...)
		bp.MoveCapacity *= 2
	}

	bp.MoveBuffer[bp.MoveCount] = proxyId
	bp.MoveCount++
}

func (bp *BroadPhase) UnBufferMove(proxyId int) {
	for i := 0; i < bp.MoveCount; i++ {
		if bp.MoveBuffer[i] == proxyId {
			bp.MoveBuffer[i] = nullProxyID
		}
	}
}

// This is called from b2DynamicTree::Query when we are gathering pairs.
func (bp *BroadPhase) QueryCallback(proxyId int) bool {

	// A proxy cannot form a pair with itself.
	if proxyId == bp.QueryProxyId {
		return true
	}

	// Grow the pair buffer as needed.
	if bp.PairCount == bp.PairCapacity {
		bp.PairBuffer = append(bp.PairBuffer, make([]Pair, bp.PairCapacity)...)
		bp.PairCapacity *= 2
	}

	bp.PairBuffer[bp.PairCount].ProxyIdA = MinInt(proxyId, bp.QueryProxyId)
	bp.PairBuffer[bp.PairCount].ProxyIdB = MaxInt(proxyId, bp.QueryProxyId)
	bp.PairCount++

	return true
}

func (bp *BroadPhase) Query(callback TreeQueryCallback, aabb AABB) {
	bp.Tree.Query(callback, aabb)
}

func (bp *BroadPhase) RayCast(callback TreeRayCastCallback, input RayCastInput) {
	bp.Tree.RayCast(callback, input)
}

func (bp *BroadPhase) ShiftOrigin(newOrigin Point) {
	bp.Tree.ShiftOrigin(newOrigin)
}
