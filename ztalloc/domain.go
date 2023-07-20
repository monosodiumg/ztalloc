package ztalloc

type (
	// Left and Right children are reached with transforms starting with 3 and 2 respectively.
	Node interface {		
		// The bool indicates is a child could be returned. The returned Node is meaingful only if the bool is true.
		LeftChild() (Node, TransformName, bool)  // first step is a 3
		// The bool indicates is a child could be returned. The returned Node is meaingful only if the bool is true.
		RightChild() (Node, TransformName, bool) // first step is a 2
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

func (n ZNode) LeftChild() (Node, TransformName, bool) {
	if n < 2 {
		return ZNode(0), Z2, false
	}
	return ZNode(2 * n), Z2, true
}

func (n ZNode) RightChild() (Node, TransformName, bool) {
	if n%2 != 0 || n%3 != 1 || n < 4 {
		return ZNode(0), Z3, false
	}
	return ZNode((n - 1) / 3), Z3, true
}

func (n ZNode) Value() int {
	return int(n)
}

// LeftChild is the
func (n CNode) LeftChild() (Node, TransformName, bool) {
	if n%2 != 0 || n < 2 {
		return CNode(0), C2, false
	}
	return CNode(n / 2), C2, true
}

func (n CNode) RightChild() (Node, TransformName, bool) {
	if n%2 != 1 || n < 1 {
		return CNode(0), C3, false
	}
	return CNode(3*n + 1), C3, true
}

func (n CNode) Value() int {
	return int(n)
}
