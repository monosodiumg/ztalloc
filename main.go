package main

import (
	"fmt"
	"log"

	"ztalloc/pkg/core"
	"ztalloc/pkg/render"
	"ztalloc/pkg/traversal"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func main() {
	render()
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
	// "ztalloc/core"
}

func render() {
	g := graphviz.New()
	graph, err := g.Graph(graphviz.StrictDirected)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()
	graph.SetStyle(cgraph.RoundedGraphStyle)

	ga, err := graph.CreateNode("a")
	// n.SetStyle()()
	if err != nil {
		log.Fatal(err)
	}
	za, _ := core.Get(5400000000004)
	render.RenderNode(za, ga)

	gb, err := graph.CreateNode("b")
	// n.SetStyle()()
	if err != nil {
		log.Fatal(err)
	}
	zb, _ := core.Get(58)
	render.RenderNode(zb, gb)

	gab, err := graph.CreateEdge("e", ga, gb)
	if err != nil {
		log.Fatal(err)
	}

	// e.SetLabel(core.E22)
	render.RenderEdge(core.E22, gab)
	var buf bytes.Buffer
	if err := g.Render(graph, "dot", &buf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	// 	image, err := g.RenderImage(graph)
	// if err != nil {
	//   log.Fatal(err)
	// }

	// 3. write to file directly
	if err := g.RenderFilename(graph, graphviz.PNG, "./graph.png"); err != nil {
		log.Fatal(err)
	}

}
