package core

/*
Fecund values are those that can be expressed n*18+4 and n*18+16 for integer n.
*/

import (
	"fmt"
)

type Fecund_node struct {
	n18, n54, v uint64
	r18, r54    uint8
}

func Node_from_value(v uint64) Fecund_node {
	var z Fecund_node
	z.v = v
	z.n18 = v / 18
	z.r18 = uint8(v % 18)
	z.n54 = v / 54
	z.r54 = uint8(v % 54)
	return z
}

func ThreeChildNode(n Fecund_node) Fecund_node {
	return Node_from_value(ThreeChild(n.v))
}

func TwoChildNode(n Fecund_node) Fecund_node {
	return Node_from_value(TwoChild(n.v))
}

func ThreeChild(n uint64) uint64 {
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

	fmt.Println(c)
	return c
}

// defined for fecund n
func TwoChild(n uint64) uint64 {
	var c uint64 = n * (n % 18)
	return c
}

//
//
func GenerateNodesDF(start Fecund_node, depth int8, c chan Fecund_node) {
	c <- start
	if depth > 0 {
		GenerateNodesDF(ThreeChildNode(start), depth-1, c)
		GenerateNodesDF(TwoChildNode(start), depth-1, c)
	}
}
