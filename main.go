package main

import (
	"flag"
	"fmt"
	"ztalloc/render"
	"ztalloc/ztalloc"
	"github.com/goccy/go-graphviz"
)

var flagRadius *int
var flagOrigin *int

func init() {
	flagOrigin = flag.Int("o", 16, "Root (start value)")
	flagRadius = flag.Int("r", 5, "Radius to explore ")
	flag.Parse()
}

func main() {
	_ = renderTreeGraph(*flagOrigin, *flagRadius)
}

func renderTreeGraph(origin, radius int) error {
	g := graphviz.New()
	g.SetLayout(graphviz.TWOPI)
	visitor, closer, err := render.Z54Renderer(g, nil, radius)
	if err != nil {
		return err
	}
	defer closer.Close()
	traverser := render.NewTraverser(ztalloc.ZBinaryNode(origin), radius)
	err = traverser.Traverse(visitor)
	if err != nil {
		return err
	}

	// var buf bytes.Buffer
	// if err := g.Render(t.Graph(), "png", &buf); err != nil {
	// 	log.Fatalln("Unable to renderTreeGraph: %w", err)
	// }
	//fmt.Println(buf.String())
	if err := g.RenderFilename(visitor.Graph(), graphviz.SVG, fmt.Sprintf("./graphs/z54-orig%d-rad%d.svg",origin,radius)); err != nil {
		fmt.Printf("Unable to renderTreeGraph: %v", err)
	}
	return nil
}
