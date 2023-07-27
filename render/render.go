package render

import (
	"fmt"
	"math"
	"ztalloc/ztalloc"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"golang.org/x/exp/slices"
)

type (
	TreeRenderer interface {
		RenderNode(ztalloc.Node, *cgraph.Node)
		RenderEdge(ztalloc.TransformName, *cgraph.Edge)

		RenderNodeV(int, *cgraph.Node)
	}

	ZTreeRenderer struct{}
	CTreeRenderer struct{}

	treeRenderer struct{
		renderNode func(ztalloc.Node, *cgraph.Node)
		renderEdge func(ztalloc.TransformName, *cgraph.Edge)
	}
)

func (t *treeRenderer) RenderNode(dn ztalloc.Node, gn *cgraph.Node){
	t.renderNode(dn, gn)
}

func (t *treeRenderer) RenderEdge(transformName ztalloc.TransformName, edge *cgraph.Edge){
	t.renderEdge(transformName , edge)
}


func (r ZTreeRenderer) RenderNode(dNode ztalloc.Node, gNode *cgraph.Node) {
	label := fmt.Sprintf("%d(%d)", dNode.Value(), dNode.Value()%54)
	gNode.SetLabel(label)
	gNode.SetColorScheme("paired12")
	gNode.SetStyle(cgraph.FilledNodeStyle)
	gNode.SetShape(cgraph.OvalShape)
	gNode.SetPenWidth(0)

	var fillc, fontc string
	switch dNode.Value() % 54 {
	case 4:
		fillc = "2"
		fontc = "black"
	case 16:
		fillc = "3"
		fontc = "black"
	case 22:
		fillc = "4"
		fontc = "black"
	case 34:
		fillc = "5"
		fontc = "black"
	case 40:
		fillc = "6"
		fontc = "black"
	case 52:
		fillc = "8"
		fontc = "black"
	default: //not expected
		gNode.SetShape(cgraph.DiamondShape)
		gNode.SetColorScheme("")
		fillc = "firebrick"
		fontc = "white"
	}

	gNode.SetFillColor(fillc)
	gNode.SetFontColor(fontc)
}

func (r ZTreeRenderer) RenderEdge(t ztalloc.TransformName, gedge *cgraph.Edge) {
	gedge.SetLabel(t)
	// gedge.Set(cgraph.VeeArrow)
	var fontc string
	switch t {
	case ztalloc.Z22:
		gedge.SetColorScheme("purd9")
		fontc = "6"
	case ztalloc.Z2222:
		gedge.SetColorScheme("purd9")
		fontc = "8"
	case ztalloc.Z32:
		gedge.SetColorScheme("pubu9")
		fontc = "6"
	case ztalloc.Z322:
		gedge.SetColorScheme("pubu9")
		fontc = "7"
	case ztalloc.Z3222:
		gedge.SetColorScheme("pubu9")
		fontc = "8"
	case ztalloc.Z32222:
		gedge.SetColorScheme("pubu9")
		fontc = "9"
	default: //not expected
		gedge.SetColorScheme("")
		fontc = "firebrick"
	}
	gedge.SetFontColor(fontc)
	gedge.SetColor(fontc)
}

func (r CTreeRenderer) RenderNode(dNode ztalloc.Node, gNode *cgraph.Node) {
	label := fmt.Sprintf("%d(%d)", dNode.Value(), dNode.Value()%54)
	gNode.SetLabel(label)
	gNode.SetColorScheme("paired12")

	gNode.SetStyle(cgraph.FilledNodeStyle)
	gNode.SetShape(cgraph.OvalShape)

	var fillc, fontc string
	switch dNode.Value() % 54 {
	case 4:
		fillc = "2"
		fontc = "black"
	case 16:
		fillc = "3"
		fontc = "black"
	case 22:
		fillc = "4"
		fontc = "black"
	case 34:
		fillc = "5"
		fontc = "black"
	case 40:
		fillc = "6"
		fontc = "black"
	case 52:
		fillc = "7"
		fontc = "black"
	default: //not expected
		gNode.SetShape(cgraph.DiamondShape)
		gNode.SetColorScheme("")
		fillc = "firebrick"
		fontc = "white"
	}

	gNode.SetFillColor(fillc)
	gNode.SetFontColor(fontc)
}

type (
	Colouring struct {
		fill string
		font string
	}
	Palette struct {
		graphvizPaletteName string
		colours             []Colouring
	}
)

// Uses the first n colours of the supplied pallete.
// TODO: handle palette too small case, e.g. test and wrap around if it fails.
func modNodeRenderer(n int, palette Palette, except []int) func(dNode ztalloc.Node, gNode *cgraph.Node) {
	// TODO optimise except membership test by sorting or else use map
	return func(dNode ztalloc.Node, gNode *cgraph.Node) {
		label := fmt.Sprintf("%d(%d)", dNode.Value(), dNode.Value()%n)
		gNode.SetLabel(label)
		gNode.SetStyle(cgraph.FilledNodeStyle)
		gNode.SetShape(cgraph.OvalShape)
		index := dNode.Value() % n
		if slices.Contains(except, index) {
			gNode.SetShape(cgraph.DiamondShape)
			gNode.SetColorScheme("")
			gNode.SetFillColor("firebrick")
			gNode.SetFontColor("white")
			return
		}
		gNode.SetColorScheme(palette.graphvizPaletteName)
		gNode.SetFillColor(palette.colours[index].fill)
		gNode.SetFontColor(palette.colours[index].font)
		return
	}
}

func (r CTreeRenderer) RenderEdge(t ztalloc.TransformName, gedge *cgraph.Edge) {
	gedge.SetLabel(t)
	gedge.SetColorScheme("piyg7")
	// gedge.Set(cgraph.VeeArrow)

	var fontc string
	switch t {
	case ztalloc.C2:
		fontc = "1"
	case ztalloc.C3:
		fontc = "2"
	default: //not expected
		gedge.SetColorScheme("")
		fontc = "firebrick"
	}
	gedge.SetFontColor(fontc)
	gedge.SetColor(fontc)
}

// If the Graphviz and Graph pointers are nil then these entitites will be created with
// defaults that scale the image to the number of nodes. In that case the returned Closer
// should be called when the graph object is no longer needed.
func Z54Renderer(gv *graphviz.Graphviz, graph *cgraph.Graph, radius int) (*context, Closer, error) {
	var closer CloserFunc
	closer = func() error { return nil }

	if gv == nil {
		gv = graphviz.New()
		gv.SetLayout(graphviz.TWOPI)
	}

	if graph == nil {
		var err error
		graph, err = gv.Graph(graphviz.StrictDirected)

		if err != nil {
			return nil, closer, fmt.Errorf("Failed creating graph: %w", err)
		}
		size := math.Pow(2, float64(radius+1)) / float64(radius)
		graph.SetSize(2*size, size)

		graph.SetNodeSeparator(40)
		graph.SetRatio(cgraph.FillRatio)
		closer = func() error { return graph.Close() }
	}

	return &context{
		graph:    graph,
		renderer: ZTreeRenderer{},
	}, closer, nil
}

func ModZRenderer(gv *graphviz.Graphviz, graph *cgraph.Graph, radius int) (*context, Closer, error) {
	var closer CloserFunc
	closer = func() error { return nil }

	if gv == nil {
		gv = graphviz.New()
		gv.SetLayout(graphviz.TWOPI)
	}

	if graph == nil {
		var err error
		graph, err = gv.Graph(graphviz.StrictDirected)

		if err != nil {
			return nil, closer, fmt.Errorf("Failed creating graph: %w", err)
		}
		size := math.Pow(2, float64(radius+1)) / float64(radius)
		graph.SetSize(2*size, size)

		graph.SetNodeSeparator(40)
		graph.SetRatio(cgraph.FillRatio)
		closer = func() error { return graph.Close() }
	}

	return &context{
		graph:    graph,
		renderer: ZTreeRenderer{},
	}, closer, nil
}

type (
	context struct {
		graph    *cgraph.Graph
		root     *cgraph.Node
		current  *cgraph.Node
		renderer TreeRenderer
	}

	Closer interface {
		Close() error
	}
	CloserFunc func() error
)

func (c *context) drawEdge(start, end *cgraph.Node, dEdge string) (*cgraph.Edge, error) {
	edge, err := c.graph.CreateEdge(end.Name(), start, end)
	if err != nil {
		return nil, fmt.Errorf("Failed creating edge of type %s from %s to %s: %w", dEdge, c.current.Name(), end.Name(), err)
	}
	c.renderer.RenderEdge(dEdge, edge)
	return edge, nil
}

func (c *context) drawNode(v int) (*cgraph.Node, error) {
	node, err := c.graph.CreateNode(fmt.Sprintf("%d", v))
	if err != nil {
		return nil, fmt.Errorf("Failed to draw node %d: %w", v, err)
	}
	c.renderer.RenderNodeV(v, node)
	return node, nil
}

func (r ZTreeRenderer) RenderNodeV(v int, gNode *cgraph.Node) {
	label := fmt.Sprintf("%d(%d)", v, v%54)
	gNode.SetLabel(label)
	gNode.SetColorScheme("paired12")
	gNode.SetStyle(cgraph.FilledNodeStyle)
	gNode.SetShape(cgraph.OvalShape)
	gNode.SetPenWidth(0)

	var fillc, fontc string
	switch v % 54 {
	case 4:
		fillc = "2"
		fontc = "black"
	case 16:
		fillc = "3"
		fontc = "black"
	case 22:
		fillc = "4"
		fontc = "black"
	case 34:
		fillc = "5"
		fontc = "black"
	case 40:
		fillc = "6"
		fontc = "black"
	case 52:
		fillc = "8"
		fontc = "black"
	default: //not expected
		gNode.SetShape(cgraph.DiamondShape)
		gNode.SetColorScheme("")
		fillc = "firebrick"
		fontc = "white"
	}

	gNode.SetFillColor(fillc)
	gNode.SetFontColor(fontc)
}

func (f CloserFunc) Close() error { return f() }

func (c *context) Graph() *cgraph.Graph {
	return c.graph
}

func (c *context) Close() error {
	if err := c.graph.Close(); err != nil {
		return err
	}
	return nil
}

func (c *context) Visit(v int, dEdge ztalloc.TransformName, _ int) (Visitor, error) {
	node, err := c.drawNode(v)
	if err != nil {
		return c, err
	}
	if c.root == nil {
		c.root = node
	}
	if c.root != node {
		_, err := c.drawEdge(c.current, node, dEdge)
		if err != nil {
			return c, fmt.Errorf("Visit failed to creating edge: %w", err)
		}
	}
	return &context{
		graph:    c.graph,
		root:     c.root,
		current:  node,
		renderer: c.renderer,
	}, nil
}
