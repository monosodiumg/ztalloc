package main

import (
	"flag"
	"github.com/goccy/go-graphviz"
	"log"
	"ztalloc/render"
	"ztalloc/ztalloc"
)

var flagDepth *int
var flagStart *int

func init() {
	// fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	flagDepth = flag.Int("d", 2, "Depth of expansion")
	flagStart = flag.Int("r", 4, "Root (start value)")
	flag.Parse()
}

func main() {
	// n.SetStyle()()
	// n.SetStyle()()
	// e.SetLabel(ztalloc.E22)
	// 	image, err := g.RenderImage(graph)
	// if err != nil {
	//   log.Fatal(err)
	// }
	// 3. write to file directly
	// renderGraphvizOld()

	renderTreeGraph(*flagStart, *flagDepth)
}

func renderTreeGraph(start, depth int) {
	g := graphviz.New()
	g.SetLayout(graphviz.TWOPI)
	
	t, err := render.NewTreeGraph(g, render.ZTreeRenderer{}, ztalloc.ZBinaryNode(start), depth)
	if err != nil {
		log.Fatalln("Unable to renderTreeGraph: %w", err)
	}

	if err := t.Draw(); err != nil {
		log.Fatalln("Unable to renderTreeGraph: %w", err)
	}

	// var buf bytes.Buffer
	// if err := g.Render(t.Graph(), "png", &buf); err != nil {
	// 	log.Fatalln("Unable to renderTreeGraph: %w", err)
	// }
	//fmt.Println(buf.String())
	if err := g.RenderFilename(t.Graph(), graphviz.SVG, "./graph.svg"); err != nil {
		log.Fatalln("Unable to renderTreeGraph: %w", err)
	}
}

// func renderGraphvizOld() {
// 	g := graphviz.New()
// 	graph, err := g.Graph(graphviz.StrictDirected)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer func() {
// 		if err := graph.Close(); err != nil {
// 			log.Fatal(err)
// 		}
// 		g.Close()
// 	}()
// 	graph.SetStyle(cgraph.RoundedGraphStyle)

// 	ga, err := graph.CreateNode("a")
// 	ga.N

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	za, _ := ztalloc.Get(5400000000004)
// 	render.RenderNode(za, ga)

// 	gb, err := graph.CreateNode("b")

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	zb, _ := ztalloc.Get(58)
// 	render.RenderNode(zb, gb)

// 	gab, err := graph.CreateEdge("e", ga, gb)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	render.RenderEdge(ztalloc.Z22, gab)
// 	var buf bytes.Buffer
// if err := g.Render(graph, "dot", &buf); err != nil {
// 	log.Fatal(err)
// }
// fmt.Println(buf.String())

// if err := g.RenderFilename(graph, graphviz.PNG, "./graph.png"); err != nil {
// 	log.Fatal(err)
// }
// }
