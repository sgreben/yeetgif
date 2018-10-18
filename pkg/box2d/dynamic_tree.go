package box2d

import (
	"math"
)

type TreeQueryCallback func(nodeId int) bool

type TreeRayCastCallback func(input RayCastInput, nodeId int) float64

const _nullNode = -1

type TreeNode struct {

	/// Enlarged AABB
	Aabb AABB

	UserData interface{}

	// union
	// {
	Parent int
	Next   int
	//};

	Child1 int
	Child2 int

	// leaf = 0, free node = -1
	Height int
}

func (node TreeNode) IsLeaf() bool {
	return node.Child1 == _nullNode
}

/// A dynamic AABB tree broad-phase, inspired by Nathanael Presson's btDbvt.
/// A dynamic tree arranges data in a binary tree to accelerate
/// queries such as volume queries and ray casts. Leafs are proxies
/// with an AABB. In the tree we expand the proxy AABB by b2_fatAABBFactor
/// so that the proxy AABB is bigger than the client object. This allows the client
/// object to move by small amounts without triggering a tree update.
///
/// Nodes are pooled and relocatable, so we use node indices rather than pointers.
type DynamicTree struct {

	// Public members:
	// None

	// Private members:
	Root int

	Nodes        []TreeNode
	NodeCount    int
	NodeCapacity int

	FreeList int

	/// This is used to incrementally traverse the tree for re-balancing.
	Path int

	InsertionCount int
}

func (tree DynamicTree) GetUserData(proxyId int) interface{} {
	return tree.Nodes[proxyId].UserData
}

func (tree DynamicTree) GetFatAABB(proxyId int) AABB {
	return tree.Nodes[proxyId].Aabb
}

func (tree *DynamicTree) Query(queryCallback TreeQueryCallback, aabb AABB) {
	stack := &GrowableStack{}
	stack.Push(tree.Root)

	for stack.GetCount() > 0 {
		nodeId := stack.Pop().(int)
		if nodeId == _nullNode {
			continue
		}

		node := &tree.Nodes[nodeId]

		if TestOverlapBoundingBoxes(node.Aabb, aabb) {
			if node.IsLeaf() {
				proceed := queryCallback(nodeId)
				if proceed == false {
					return
				}
			} else {
				stack.Push(node.Child1)
				stack.Push(node.Child2)
			}
		}
	}
}

func (tree DynamicTree) RayCast(rayCastCallback TreeRayCastCallback, input RayCastInput) {

	p1 := input.P1
	p2 := input.P2
	r := PointSub(p2, p1)
	r.Normalize()

	// v is perpendicular to the segment.
	v := PointCrossScalarVector(1.0, r)
	abs_v := PointAbs(v)

	// Separating axis for segment (Gino, p80).
	// |dot(v, p1 - c)| > dot(|v|, h)

	maxFraction := input.MaxFraction

	// Build a bounding box for the segment.
	segmentAABB := AABB{}
	{
		t := PointAdd(p1, PointMulScalar(maxFraction, PointSub(p2, p1)))
		segmentAABB.Min = PointMin(p1, t)
		segmentAABB.Max = PointMax(p1, t)
	}

	stack := &GrowableStack{}
	stack.Push(tree.Root)

	for stack.GetCount() > 0 {
		nodeId := stack.Pop().(int)
		if nodeId == _nullNode {
			continue
		}

		node := &tree.Nodes[nodeId]

		if TestOverlapBoundingBoxes(node.Aabb, segmentAABB) == false {
			continue
		}

		// Separating axis for segment (Gino, p80).
		// |dot(v, p1 - c)| > dot(|v|, h)
		c := node.Aabb.GetCenter()
		h := node.Aabb.GetExtents()

		separation := math.Abs(PointDot(v, PointSub(p1, c))) - PointDot(abs_v, h)
		if separation > 0.0 {
			continue
		}

		if node.IsLeaf() {
			subInput := RayCastInput{}
			subInput.P1 = input.P1
			subInput.P2 = input.P2
			subInput.MaxFraction = maxFraction

			value := rayCastCallback(subInput, nodeId)

			if value == 0.0 {
				// The client has terminated the ray cast.
				return
			}

			if value > 0.0 {
				// Update segment bounding box.
				maxFraction = value
				t := PointAdd(p1, PointMulScalar(maxFraction, PointSub(p2, p1)))
				segmentAABB.Min = PointMin(p1, t)
				segmentAABB.Max = PointMax(p1, t)
			}
		} else {
			stack.Push(node.Child1)
			stack.Push(node.Child2)
		}
	}
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// DynamicTree.cpp
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func MakeDynamicTree() DynamicTree {

	tree := DynamicTree{}
	tree.Root = _nullNode

	tree.NodeCapacity = 16
	tree.NodeCount = 0
	tree.Nodes = make([]TreeNode, tree.NodeCapacity)

	// Build a linked list for the free list.
	for i := 0; i < tree.NodeCapacity-1; i++ {
		tree.Nodes[i].Next = i + 1
		tree.Nodes[i].Height = -1
	}

	tree.Nodes[tree.NodeCapacity-1].Next = _nullNode
	tree.Nodes[tree.NodeCapacity-1].Height = -1
	tree.FreeList = 0

	tree.Path = 0

	tree.InsertionCount = 0

	return tree
}

// func (tree *DynamicTree) ~b2DynamicTree() {
// 	// This frees the entire tree in one shot.
// 	b2Free(tree.Nodes);
// }

// Allocate a node from the pool. Grow the pool if necessary.
func (tree *DynamicTree) AllocateNode() int {

	// Expand the node pool as needed.
	if tree.FreeList == _nullNode {

		// The free list is empty. Rebuild a bigger pool.
		tree.Nodes = append(tree.Nodes, make([]TreeNode, tree.NodeCapacity)...)
		tree.NodeCapacity *= 2

		// Build a linked list for the free list. The parent
		// pointer becomes the "next" pointer.
		for i := tree.NodeCount; i < tree.NodeCapacity-1; i++ {
			tree.Nodes[i].Next = i + 1
			tree.Nodes[i].Height = -1
		}

		tree.Nodes[tree.NodeCapacity-1].Next = _nullNode
		tree.Nodes[tree.NodeCapacity-1].Height = -1
		tree.FreeList = tree.NodeCount
	}

	// Peel a node off the free list.
	nodeId := tree.FreeList
	tree.FreeList = tree.Nodes[nodeId].Next
	tree.Nodes[nodeId].Parent = _nullNode
	tree.Nodes[nodeId].Child1 = _nullNode
	tree.Nodes[nodeId].Child2 = _nullNode
	tree.Nodes[nodeId].Height = 0
	tree.Nodes[nodeId].UserData = nil
	tree.NodeCount++

	return nodeId
}

// Return a node to the pool.
func (tree *DynamicTree) FreeNode(nodeId int) {
	tree.Nodes[nodeId].Next = tree.FreeList
	tree.Nodes[nodeId].Height = -1
	tree.FreeList = nodeId
	tree.NodeCount--
}

// Create a proxy in the tree as a leaf node. We return the index
// of the node instead of a pointer so that we can grow
// the node pool.
func (tree *DynamicTree) CreateProxy(aabb AABB, userData interface{}) int {

	proxyId := tree.AllocateNode()

	// Fatten the aabb.
	r := Point{X: _aabbExtension, Y: _aabbExtension}
	tree.Nodes[proxyId].Aabb.Min = PointSub(aabb.Min, r)
	tree.Nodes[proxyId].Aabb.Max = PointAdd(aabb.Max, r)
	tree.Nodes[proxyId].UserData = userData
	tree.Nodes[proxyId].Height = 0

	tree.InsertLeaf(proxyId)

	return proxyId
}

func (tree *DynamicTree) DestroyProxy(proxyId int) {

	tree.RemoveLeaf(proxyId)
	tree.FreeNode(proxyId)
}

func (tree *DynamicTree) MoveProxy(proxyId int, aabb AABB, displacement Point) bool {

	if tree.Nodes[proxyId].Aabb.Contains(aabb) {
		return false
	}

	tree.RemoveLeaf(proxyId)

	// Extend AABB.
	b := aabb.Clone()
	r := Point{X: _aabbExtension, Y: _aabbExtension}
	b.Min = PointSub(b.Min, r)
	b.Max = PointAdd(b.Max, r)

	// Predict AABB displacement.
	d := PointMulScalar(_aabbMultiplier, displacement)

	if d.X < 0.0 {
		b.Min.X += d.X
	} else {
		b.Max.X += d.X
	}

	if d.Y < 0.0 {
		b.Min.Y += d.Y
	} else {
		b.Max.Y += d.Y
	}

	tree.Nodes[proxyId].Aabb = b

	tree.InsertLeaf(proxyId)

	return true
}

func (tree *DynamicTree) InsertLeaf(leaf int) {
	tree.InsertionCount++

	if tree.Root == _nullNode {
		tree.Root = leaf
		tree.Nodes[tree.Root].Parent = _nullNode
		return
	}

	// Find the best sibling for this node
	leafAABB := tree.Nodes[leaf].Aabb
	index := tree.Root
	for tree.Nodes[index].IsLeaf() == false {
		child1 := tree.Nodes[index].Child1
		child2 := tree.Nodes[index].Child2

		area := tree.Nodes[index].Aabb.GetPerimeter()

		combinedAABB := &AABB{}
		combinedAABB.CombineTwoInPlace(tree.Nodes[index].Aabb, leafAABB)
		combinedArea := combinedAABB.GetPerimeter()

		// Cost of creating a new parent for this node and the new leaf
		cost := 2.0 * combinedArea

		// Minimum cost of pushing the leaf further down the tree
		inheritanceCost := 2.0 * (combinedArea - area)

		// Cost of descending into child1
		cost1 := 0.0
		if tree.Nodes[child1].IsLeaf() {
			aabb := &AABB{}
			aabb.CombineTwoInPlace(leafAABB, tree.Nodes[child1].Aabb)
			cost1 = aabb.GetPerimeter() + inheritanceCost
		} else {
			aabb := &AABB{}
			aabb.CombineTwoInPlace(leafAABB, tree.Nodes[child1].Aabb)
			oldArea := tree.Nodes[child1].Aabb.GetPerimeter()
			newArea := aabb.GetPerimeter()
			cost1 = (newArea - oldArea) + inheritanceCost
		}

		// Cost of descending into child2
		cost2 := 0.0
		if tree.Nodes[child2].IsLeaf() {
			aabb := &AABB{}
			aabb.CombineTwoInPlace(leafAABB, tree.Nodes[child2].Aabb)
			cost2 = aabb.GetPerimeter() + inheritanceCost
		} else {
			aabb := &AABB{}
			aabb.CombineTwoInPlace(leafAABB, tree.Nodes[child2].Aabb)
			oldArea := tree.Nodes[child2].Aabb.GetPerimeter()
			newArea := aabb.GetPerimeter()
			cost2 = newArea - oldArea + inheritanceCost
		}

		// Descend according to the minimum cost.
		if cost < cost1 && cost < cost2 {
			break
		}

		// Descend
		if cost1 < cost2 {
			index = child1
		} else {
			index = child2
		}
	}

	sibling := index

	// Create a new parent.
	oldParent := tree.Nodes[sibling].Parent
	newParent := tree.AllocateNode()
	tree.Nodes[newParent].Parent = oldParent
	tree.Nodes[newParent].UserData = nil
	tree.Nodes[newParent].Aabb.CombineTwoInPlace(leafAABB, tree.Nodes[sibling].Aabb)
	tree.Nodes[newParent].Height = tree.Nodes[sibling].Height + 1

	if oldParent != _nullNode {
		// The sibling was not the root.
		if tree.Nodes[oldParent].Child1 == sibling {
			tree.Nodes[oldParent].Child1 = newParent
		} else {
			tree.Nodes[oldParent].Child2 = newParent
		}

		tree.Nodes[newParent].Child1 = sibling
		tree.Nodes[newParent].Child2 = leaf
		tree.Nodes[sibling].Parent = newParent
		tree.Nodes[leaf].Parent = newParent
	} else {
		// The sibling was the root.
		tree.Nodes[newParent].Child1 = sibling
		tree.Nodes[newParent].Child2 = leaf
		tree.Nodes[sibling].Parent = newParent
		tree.Nodes[leaf].Parent = newParent
		tree.Root = newParent
	}

	// Walk back up the tree fixing heights and AABBs
	index = tree.Nodes[leaf].Parent
	for index != _nullNode {
		index = tree.Balance(index)

		child1 := tree.Nodes[index].Child1
		child2 := tree.Nodes[index].Child2

		tree.Nodes[index].Height = 1 + MaxInt(tree.Nodes[child1].Height, tree.Nodes[child2].Height)
		tree.Nodes[index].Aabb.CombineTwoInPlace(tree.Nodes[child1].Aabb, tree.Nodes[child2].Aabb)

		index = tree.Nodes[index].Parent
	}

	//Validate();
}

func (tree *DynamicTree) RemoveLeaf(leaf int) {
	if leaf == tree.Root {
		tree.Root = _nullNode
		return
	}

	parent := tree.Nodes[leaf].Parent
	grandParent := tree.Nodes[parent].Parent
	sibling := 0
	if tree.Nodes[parent].Child1 == leaf {
		sibling = tree.Nodes[parent].Child2
	} else {
		sibling = tree.Nodes[parent].Child1
	}

	if grandParent != _nullNode {
		// Destroy parent and connect sibling to grandParent.
		if tree.Nodes[grandParent].Child1 == parent {
			tree.Nodes[grandParent].Child1 = sibling
		} else {
			tree.Nodes[grandParent].Child2 = sibling
		}
		tree.Nodes[sibling].Parent = grandParent
		tree.FreeNode(parent)

		// Adjust ancestor bounds.
		index := grandParent
		for index != _nullNode {
			index = tree.Balance(index)

			child1 := tree.Nodes[index].Child1
			child2 := tree.Nodes[index].Child2

			tree.Nodes[index].Aabb.CombineTwoInPlace(tree.Nodes[child1].Aabb, tree.Nodes[child2].Aabb)
			tree.Nodes[index].Height = 1 + MaxInt(tree.Nodes[child1].Height, tree.Nodes[child2].Height)

			index = tree.Nodes[index].Parent
		}
	} else {
		tree.Root = sibling
		tree.Nodes[sibling].Parent = _nullNode
		tree.FreeNode(parent)
	}

	// //Validate();
}

// Perform a left or right rotation if node A is imbalanced.
// Returns the new root index.
func (tree *DynamicTree) Balance(iA int) int {

	A := &tree.Nodes[iA]
	if A.IsLeaf() || A.Height < 2 {
		return iA
	}

	iB := A.Child1
	iC := A.Child2

	B := &tree.Nodes[iB]
	C := &tree.Nodes[iC]

	balance := C.Height - B.Height

	// Rotate C up
	if balance > 1 {
		iF := C.Child1
		iG := C.Child2
		F := &tree.Nodes[iF]
		G := &tree.Nodes[iG]

		// Swap A and C
		C.Child1 = iA
		C.Parent = A.Parent
		A.Parent = iC

		// A's old parent should point to C
		if C.Parent != _nullNode {
			if tree.Nodes[C.Parent].Child1 == iA {
				tree.Nodes[C.Parent].Child1 = iC
			} else {
				tree.Nodes[C.Parent].Child2 = iC
			}
		} else {
			tree.Root = iC
		}

		// Rotate
		if F.Height > G.Height {
			C.Child2 = iF
			A.Child2 = iG
			G.Parent = iA
			A.Aabb.CombineTwoInPlace(B.Aabb, G.Aabb)
			C.Aabb.CombineTwoInPlace(A.Aabb, F.Aabb)

			A.Height = 1 + MaxInt(B.Height, G.Height)
			C.Height = 1 + MaxInt(A.Height, F.Height)
		} else {
			C.Child2 = iG
			A.Child2 = iF
			F.Parent = iA
			A.Aabb.CombineTwoInPlace(B.Aabb, F.Aabb)
			C.Aabb.CombineTwoInPlace(A.Aabb, G.Aabb)

			A.Height = 1 + MaxInt(B.Height, F.Height)
			C.Height = 1 + MaxInt(A.Height, G.Height)
		}

		return iC
	}

	// Rotate B up
	if balance < -1 {
		iD := B.Child1
		iE := B.Child2

		D := &tree.Nodes[iD]
		E := &tree.Nodes[iE]

		// Swap A and B
		B.Child1 = iA
		B.Parent = A.Parent
		A.Parent = iB

		// A's old parent should point to B
		if B.Parent != _nullNode {
			if tree.Nodes[B.Parent].Child1 == iA {
				tree.Nodes[B.Parent].Child1 = iB
			} else {
				tree.Nodes[B.Parent].Child2 = iB
			}
		} else {
			tree.Root = iB
		}

		// Rotate
		if D.Height > E.Height {
			B.Child2 = iD
			A.Child1 = iE
			E.Parent = iA
			A.Aabb.CombineTwoInPlace(C.Aabb, E.Aabb)
			B.Aabb.CombineTwoInPlace(A.Aabb, D.Aabb)

			A.Height = 1 + MaxInt(C.Height, E.Height)
			B.Height = 1 + MaxInt(A.Height, D.Height)
		} else {
			B.Child2 = iE
			A.Child1 = iD
			D.Parent = iA
			A.Aabb.CombineTwoInPlace(C.Aabb, D.Aabb)
			B.Aabb.CombineTwoInPlace(A.Aabb, E.Aabb)

			A.Height = 1 + MaxInt(C.Height, D.Height)
			B.Height = 1 + MaxInt(A.Height, E.Height)
		}

		return iB
	}

	return iA
}

func (tree DynamicTree) GetHeight() int {
	if tree.Root == _nullNode {
		return 0
	}

	return tree.Nodes[tree.Root].Height
}

//
func (tree DynamicTree) GetAreaRatio() float64 {
	if tree.Root == _nullNode {
		return 0.0
	}

	root := &tree.Nodes[tree.Root]
	rootArea := root.Aabb.GetPerimeter()

	totalArea := 0.0
	for i := 0; i < tree.NodeCapacity; i++ {
		node := &tree.Nodes[i]
		if node.Height < 0 {
			// Free node in pool
			continue
		}

		totalArea += node.Aabb.GetPerimeter()
	}

	return totalArea / rootArea
}

// Compute the height of a sub-tree.
func (tree DynamicTree) ComputeHeight(nodeId int) int {
	node := &tree.Nodes[nodeId]

	if node.IsLeaf() {
		return 0
	}

	height1 := tree.ComputeHeight(node.Child1)
	height2 := tree.ComputeHeight(node.Child2)
	return 1 + MaxInt(height1, height2)
}

func (tree DynamicTree) ComputeTotalHeight() int {
	return tree.ComputeHeight(tree.Root)
}

func (tree DynamicTree) ValidateStructure(index int) {
	if index == _nullNode {
		return
	}

	if index == tree.Root {
	}

	node := &tree.Nodes[index]

	child1 := node.Child1
	child2 := node.Child2

	if node.IsLeaf() {
		return
	}

	tree.ValidateStructure(child1)
	tree.ValidateStructure(child2)
}

func (tree DynamicTree) Validate() {
}

func (tree DynamicTree) GetMaxBalance() int {
	maxBalance := 0
	for i := 0; i < tree.NodeCapacity; i++ {
		node := &tree.Nodes[i]
		if node.Height <= 1 {
			continue
		}

		child1 := node.Child1
		child2 := node.Child2
		balance := AbsInt(tree.Nodes[child2].Height - tree.Nodes[child1].Height)
		maxBalance = MaxInt(maxBalance, balance)
	}

	return maxBalance
}

func (tree *DynamicTree) RebuildBottomUp() {
	//int* nodes = (int*)b2Alloc(m_nodeCount * sizeof(int));
	nodes := make([]int, tree.NodeCount)
	count := 0

	// Build array of leaves. Free the rest.
	for i := 0; i < tree.NodeCapacity; i++ {
		if tree.Nodes[i].Height < 0 {
			// free node in pool
			continue
		}

		if tree.Nodes[i].IsLeaf() {
			tree.Nodes[i].Parent = _nullNode
			nodes[count] = i
			count++
		} else {
			tree.FreeNode(i)
		}
	}

	for count > 1 {
		minCost := math.MaxFloat64
		iMin := -1
		jMin := -1

		for i := 0; i < count; i++ {
			aabbi := tree.Nodes[nodes[i]].Aabb

			for j := i + 1; j < count; j++ {
				aabbj := tree.Nodes[nodes[j]].Aabb
				b := &AABB{}
				b.CombineTwoInPlace(aabbi, aabbj)
				cost := b.GetPerimeter()
				if cost < minCost {
					iMin = i
					jMin = j
					minCost = cost
				}
			}
		}

		index1 := nodes[iMin]
		index2 := nodes[jMin]
		child1 := &tree.Nodes[index1]
		child2 := &tree.Nodes[index2]

		parentIndex := tree.AllocateNode()
		parent := &tree.Nodes[parentIndex]
		parent.Child1 = index1
		parent.Child2 = index2
		parent.Height = 1 + MaxInt(child1.Height, child2.Height)
		parent.Aabb.CombineTwoInPlace(child1.Aabb, child2.Aabb)
		parent.Parent = _nullNode

		child1.Parent = parentIndex
		child2.Parent = parentIndex

		nodes[jMin] = nodes[count-1]
		nodes[iMin] = parentIndex
		count--
	}

	tree.Root = nodes[0]
	//b2Free(nodes)

	tree.Validate()
}

func (tree *DynamicTree) ShiftOrigin(newOrigin Point) {
	// Build array of leaves. Free the rest.
	for i := 0; i < tree.NodeCapacity; i++ {
		tree.Nodes[i].Aabb.Min.OperatorMinusInplace(newOrigin)
		tree.Nodes[i].Aabb.Max.OperatorMinusInplace(newOrigin)
	}
}
