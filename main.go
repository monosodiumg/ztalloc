package main

import (
	"fmt"

	"ztalloc/pkg/core"
)

func main() {
	err := runCore()
	if err != nil {
		fmt.Println(err)
	}
}

func runCore() error {
	depth := uint8(2)

	visitedCount := 0
	v := func(a core.Node, parent core.Node, d uint8) {
		visitedCount++
		s := fmt.Sprintf("%8d %3d: %31d <--- %d", visitedCount, depth-d, a.V, parent.V)
		println(s)

	}
	traverse := core.DfoGen(core.PREORDER, v)

	//

	start := int(106)
	return traverse(start, depth)

}
