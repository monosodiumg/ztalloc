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
	znode int
	cnode int

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

func NewZTree(root int) znode {
	return znode(root)
}

func GetTransformType(a, b int) (string, error) {
	return "", fmt.Errorf("Not implemented")
}

// func Get(n int) (Node, error) {
// 	switch n % 54 {
// 	case 4, 16, 22, 28, 34, 40, 52:
// 		return node(n), nil
// 	default:
// 		return nil, fmt.Errorf("Cannot create Node from illegal value %d.", n)
// 	}
// }

func (n znode) Parent() (znode, error) {
	return 0, fmt.Errorf("Not implemented")
}

// LeftChild is the
func (n znode) LeftChild() (znode, error) {
	return 2 * n, nil
}

func (n znode) RightChild() (znode, error) {
	if n%3 != 1 {
		return n, fmt.Errorf("znode %d has no RightChild", n)
	}
	return (n - 1) / 3, nil
}

func (n znode) Value() int {
	return int(n)
}

// func (n cnode) Parent() (Node, error) {
// 	return nil, fmt.Errorf("Not implemented")
// }

// LeftChild is the
func (n cnode) LeftChild() (Node, TransformName, error) {
	if n%2 != 1 {
		return nil, C2, fmt.Errorf("cnode ")
	}
	return 2 * n, C2, nil
}

func (n cnode) RightChild() (Node, TransformName, error) {
	if n%3 != 1 {
		return n, C3, fmt.Errorf("cnode %d has no RightChild", n)
	}
	return (n - 1) / 3, C3, nil
}

func (n cnode) Value() int {
	return int(n)
}

