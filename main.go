package main

import (
	"fmt"
	"ztalloc/pkg/core"
)

func main() {
	depth := uint8(24)

	visitedCount := 0
	v := func(a core.Node, parent core.Node, d uint8) {
		visitedCount++;
		s := fmt.Sprintf("%8d %3d: %31d <--- %d", visitedCount, depth-d, a.V, parent.V)
		println(s)

	}
	traverse := core.DfoGen(core.PREORDER, v)

	start := core.NodeFromInt(16)
	traverse(start, depth)

}
