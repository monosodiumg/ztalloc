package core

import (
	"fmt"
)

//TODO : func (a Node) Parent() Node

type Node struct {
	N18, N54, V uint64
	R18, R54    uint8
}

func (a Node) ToString() string {
	return fmt.Sprintf("%d = 18*%d+%d = 54*%d+%d", a.V, a.N18, a.R18, a.N54, a.R54)
}

func NodeFromInt(v uint64) Node {
	//TODO: error if not fecund value
	var z Node
	z.V = v
	z.N18 = v / 18
	z.R18 = uint8(v % 18)
	z.N54 = v / 54
	z.R54 = uint8(v % 54)
	return z
}

func (a Node) TwoChild() Node {
	c := Two(a.V)
	return NodeFromInt(c)
}

func (a Node) ThreeChild() Node {
	c := Three(a.V)
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

type Visitor = func(Node, Node, uint8)
type NodeGen = func(Node) Node
type _DFTraverser = func(Node, Node, uint8)
type DFTraverser = func(Node, uint8)
type DFTraverserGen = func(NodeGen) DFTraverser

func DfoGen(order TraversalOrder, v Visitor) DFTraverser {
	var trav _DFTraverser
	three := func(a Node, d uint8) {
		if d > 0 {
			trav(a.ThreeChild(), a, d -1)
		}
	}
	two := func(a Node, d uint8) {
		if d > 0 {
			trav(a.TwoChild(), a, d -1)
		}
	}

	switch order {
	case PREORDER:
		trav = func(a Node, parent Node, d uint8) {
			v(a, parent, d)
			three(a, d)
			two(a, d)
		}
	case INORDER:
		trav = func(a Node, parent Node, d uint8) {
			three(a, d-1)
			v(a, parent, d)
			two(a, d-1)
		}
	case POSTORDER:
		trav = func(a Node, parent Node, d uint8) {
			three(a, d-1)
			two(a, d-1)
			v(a, parent, d)

		}
	}

	return func(a Node, d uint8) {
		var zero Node
		trav(a, zero, d)
	}
}
