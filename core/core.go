package core

import (
	"errors"
	"fmt"
)

//TODO : func (a Node) Parent() Node

type Node struct {
	N18, N54, V int
	R18, R54    uint8
}

var _root = Node{
	N18: 0,
	N54: 0,
	V:   4,
	R18: 4,
	R54: 4,
}

func (a Node) ToString() string {
	return fmt.Sprintf("%d = 18*%d+%d = 54*%d+%d", a.V, a.N18, a.R18, a.N54, a.R54)
}

func NodeFromInt(v int) Node {
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

func Three(n int) int {
	r54 := n % 54
	var m int
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
func Two(n int) int {
	var c = n * (n % 18)
	return c
}

//pre-order inorder post-order:
type TraversalOrder int

const (
	PREORDER TraversalOrder = iota + 1
	INORDER
	POSTORDER
)

type NodeGen = func(Node) Node
type Visitor = func(Node, Node, uint8)
type DFTraverser = func(int, uint8) error
type _DFTraverser = func(Node, uint8)
type DFTraverserGen = func(NodeGen) DFTraverser

func guardedChildTraverserGen(g NodeGen, t *Visitor) _DFTraverser {
	return func(a Node, d uint8) {
		if d > 0 {
			(*t)(g(a), a, d)
		}
	}
}

func DfoGen(order TraversalOrder, v Visitor) DFTraverser {
	var trav Visitor
	three := guardedChildTraverserGen(Node.ThreeChild, &trav)
	two := guardedChildTraverserGen(Node.TwoChild, &trav)

	switch order {
	case PREORDER:
		trav = func(a Node, parent Node, d uint8) {
			v(a, parent, d)
			three(a, d-1)
			two(a, d-1)
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

	return func(n int, d uint8) error {
		r18 := n % 18
		if r18 != 4 && r18 != 16 {
			return errors.New("Not a fecund value. Must be congruent to 4 or 16 modulus 18")
		}
		if n <= 4 {
			return errors.New("Start value must be at least 16")
		}

		a := NodeFromInt(n)
		var zero Node
		trav(a, zero, d)
		return nil
	}
}
