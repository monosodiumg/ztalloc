package ztalloc

import "fmt"

type (
	Node interface {
		Parent() (Node, error)
		LeftChild() (Node, error)
		RightChild() (Node, error)
		Value() int
	}
	node     int
	TransformName = string
	Transform = func(int) int
)

const (
	E22   TransformName = "22"
	E2222 TransformName = "2222"
	E32   TransformName = "32"
	E322  TransformName = "322"
	E3222 TransformName = "3222"
)

var Transforms = map[TransformName]Transform{
	E22:   func(n int) int { return n * 4 },
	E2222: func(n int) int { return n * 16 },
	E32:   func(n int) int { return (n - 1) / 3 * 2 },
	E322:  func(n int) int { return (n - 1) / 3 * 4 },
	E3222: func(n int) int { return (n - 1) / 3 * 8 },
}

func Root() Node {
	return node(16)
}

func GetTransformType(a, b int) (string, error) {
	return "", fmt.Errorf("Not implemented")
}

func Get(n int) (Node, error) {
	return node(n), nil
}

func (n node) Parent() (Node, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (n node) LeftChild() (Node, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (n node) RightChild() (Node, error) {
	return nil, fmt.Errorf("Not implemented")
}
func (n node) Value() int {
	return int(n)
}
