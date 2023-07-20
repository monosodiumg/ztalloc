package ztalloc

import "fmt"

type (
	// Left and Right children are reached with transforms starting with 3 and 2 respectively.
	Node interface {
		// Parent() (Node, TransformName, error)
		LeftChild() (Node, TransformName, error)  // first step is a 3
		RightChild() (Node, TransformName, error) // first step is a 2
		Value() int
	}
	ZNode int
	CNode int

	TransformName = string
	Transform     = func(int) int

	// ZTree   Node
	// Z54Tree Node
	// CTree   Node
)

const (
	C2 TransformName = "2"
	C3 TransformName = "3"

	Z2    TransformName = "2"
	Z3    TransformName = "3"
	Z22   TransformName = "22"
	Z2222 TransformName = "2222"
	Z32   TransformName = "32"
	Z322  TransformName = "322"
	Z3222 TransformName = "3222"
)

var CTransforms = map[TransformName]Transform{
	C2: func(n int) int { return n / 2 },
	C3: func(n int) int { return 3*n + 1 },
}

var ZTransforms = map[TransformName]Transform{
	Z2: func(n int) int { return n * 2 },
	Z3: func(n int) int { return (n - 1) / 3 },
}

var Z54Transforms = map[TransformName]Transform{
	Z22:   func(n int) int { return n * 4 },
	Z2222: func(n int) int { return n * 16 },
	Z32:   func(n int) int { return (n - 1) / 3 * 2 },
	Z322:  func(n int) int { return (n - 1) / 3 * 4 },
	Z3222: func(n int) int { return (n - 1) / 3 * 8 },
}

// func NewZNode(v int) ZNode {
// 	return ZNode(v)
// }

// LeftChild is the
func (n ZNode) LeftChild() (ZNode, TransformName, error) {
	return 2 * n, Z2, nil
}

func (n ZNode) RightChild() (ZNode, TransformName, error) {
	if n%3 != 1 || n < 4 {
		return ZNode(0), Z3, fmt.Errorf("ZNode %d has no RightChild", n)
	}
	return (n - 1) / 3, Z3, nil
}

func (n ZNode) Value() int {
	return int(n)
}

// LeftChild is the
func (n CNode) LeftChild() (Node, TransformName, error) {
	if n%2 != 0 || n < 2 {
		return nil, C2, fmt.Errorf("CNode %d has no LeftChild", n)
	}
	return n / 2, C2, nil
}

func (n CNode) RightChild() (Node, TransformName, error) {
	if n%2 != 1 || n < 0 {
		return nil, C3, fmt.Errorf("CNode %d has no RightChild", n)
	}
	return 3*n + 1, C3, nil
}

func (n CNode) Value() int {
	return int(n)
}
