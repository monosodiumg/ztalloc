package core

import (
	"fmt"
)

//TODO : func (a Node) Parent() Node

type Node struct {
	n18, n54, v uint64
	r18, r54    uint8
}

func (a Node) ToString() string {
	return fmt.Sprintf("%d = 18*%d+%d = 54*%d+%d", a.v, a.n18, a.r18, a.n54, a.r54)
}

func NodeFromInt(v uint64) Node {
	//TODO: error if not fecund value
	var z Node
	z.v = v
	z.n18 = v / 18
	z.r18 = uint8(v % 18)
	z.n54 = v / 54
	z.r54 = uint8(v % 54)
	return z
}

func (a Node) TwoChild() Node {
	c := Two(a.v)
	return NodeFromInt(c)
}

func (a Node) ThreeChild() Node {
	c := Three(a.v)
	return NodeFromInt(c)
}

func Three(n uint64) uint64 {
	r54 := n % 54
	var m uint64
	switch r54 {
	case 34:
		m = 2
	case 52:
		m = 2
	case 4:
		m = 4
	case 40:
		m = 4
	case 16:
		m = 8
	default: // case 4
		m = 16
	}
	c := m * (n - 1) / 3
	return c
}

// defined for fecund n
func Two(n uint64) uint64 {
	var c = n * (n % 18)
	return c
}

func (a Node) TraverseDepthFirst(depth int8, f func(*Node)) {
	f(&a)
	if depth > 0 {
		a.ThreeChild().TraverseDepthFirst(depth-1, f)
		a.TwoChild().TraverseDepthFirst(depth-1, f)
	}
}

//pre-order inorder post-order:
type TraversalOrder int

const (
	PREORDER TraversalOrder = iota + 1
	INORDER
	POSTORDER
)

type Visitor = func(Node)
type NodeGen = func(Node) Node
type DFTraverser = func(Node, int8)
type DFTraverserGen = func(NodeGen) DFTraverser

func DfoGen(order TraversalOrder, v Visitor) DFTraverser {
	three := Node.ThreeChild
	two := Node.TwoChild

	var trav DFTraverser
	switch order {
	case PREORDER:
		trav = func(a Node, d int8) {
			v(a)
			if d > 0 {
				trav(three(a), d-1)
			}
			if d > 0 {
				trav(two(a), d-1)
			}
		}

	}

	return trav
}

