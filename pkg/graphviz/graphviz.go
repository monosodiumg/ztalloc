package graphviz

import (
	"fmt"
)
/*
func main() {
	node_source := make(chan uint64, 100)
	render_dst := make(chan string, 10)
	GraphvizDraw(node_source, render_dst)
}
*/
func GraphvizDraw(src chan uint64, dst chan string) {
	
	for n := range src {
		fmt.Printf("GraphvizDraw src <-%d\n", n)
		dst <- GraphvizNodeRender(n)
	}
}

func GraphvizNodeRender(node uint64) string {
	r := fmt.Sprintf("v%d\n", node)
	//fmt.Printf("GraphvizNodeRender of %d is %s\n", node, r)
	return r
}
