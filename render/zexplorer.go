package render

import (
	"fmt"
	"ztalloc/ztalloc"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type (
	TreeGraph struct {
		start    ztalloc.Node
		renderer TreeRenderer
		graph    *cgraph.Graph
		depth    int
	}
	chirality int
)

const (
	left chirality = iota
	right
)

func NewTreeGraph(gv *graphviz.Graphviz, renderer TreeRenderer, start ztalloc.Node, depth int) (TreeGraph, error) {
	if depth < 0 {
		return TreeGraph{}, fmt.Errorf("Illegal depth %d. Must be non-negative", depth)
	}
	graph, err := gv.Graph(graphviz.StrictDirected)
	if err != nil {
		return TreeGraph{}, fmt.Errorf("Error creating TreeGraph: %w", err)
	}
	return TreeGraph{
		start:    start,
		renderer: renderer,
		graph:    graph,
		depth:    depth,
	}, nil
}

func (t TreeGraph) Graph() *cgraph.Graph { return t.graph }

func (t TreeGraph) Close() error {
	if err := t.graph.Close(); err != nil {
		return err
	}
	return nil
}

func (t TreeGraph) Draw() error {
	gRoot, err := t.drawNode(t.start)
	if err != nil {
		return fmt.Errorf("Error drawing root %d: %w", t.start.Value(), err)
	}
	err = t.drawChildren(t.start, gRoot, t.depth)
	if err != nil {
		return fmt.Errorf("Error drawing tree rooted at %d: %w", t.start.Value(), err)
	}
	return nil
}

func (g TreeGraph) drawChildren(dNode ztalloc.Node, gNode *cgraph.Node, depth int) error {
	if depth == 0 {
		// draw the node and no traversal (only happens for a singleton node)
		return nil
	}
	ldNode, lgNode, lerr := g.drawChild(dNode, gNode, left)
	rdNode, rgNode, rerr := g.drawChild(dNode, gNode, right)

	if depth > 1 {
		// iterate over children
		if lerr == nil && lgNode != nil {
			lerr = g.drawChildren(ldNode, lgNode, depth-1)
		}
		if rerr == nil && rgNode != nil {
			rerr = g.drawChildren(rdNode, rgNode, depth-1)
		}
	}
	if lerr != nil {
		return fmt.Errorf("Error drawing left descendents of %d: %w", dNode.Value(), lerr)
	}
	if rerr != nil {
		return fmt.Errorf("Error drawing right descendents of %d: %w", dNode.Value(), rerr)
	}

	return nil
}

// Add a node and an edge to the graph, representing the child and the parent-child respectively. The returned child data and graph nodes will be nil if the child cannot be produced.
func (g TreeGraph) drawChild(dParent ztalloc.Node, gParent *cgraph.Node, child chirality) (ztalloc.Node, *cgraph.Node, error) {
	var dChild ztalloc.Node
	// var dEdge ztalloc.TransformName
	var suffix string
	var ok bool
	var dEdge ztalloc.TransformName

	if child == left {
		dChild, dEdge, ok = dParent.LeftChild()
		suffix = "l"
	} else {
		dChild, dEdge, ok = dParent.RightChild()
		suffix = "r"
	}
	if !ok {
		return nil, nil, nil
	}

	gChild, err := g.drawNode(dChild)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to draw %s child of %d: %w", suffix, dParent.Value(), err)
	}
	gEdge, err := g.graph.CreateEdge(gChild.Name(), gParent, gChild)
	if err != nil {
		return nil, nil, nil
	}

	g.renderer.RenderEdge(dEdge, gEdge)
	return dChild, gChild, nil
}

func (g TreeGraph) drawNode(dNode ztalloc.Node) (*cgraph.Node, error) {
	var gNode *cgraph.Node
	gNode, err := g.graph.CreateNode(fmt.Sprintf("%d", dNode.Value()))
	if err != nil {
		return nil, fmt.Errorf("Failed to draw node %d: %w", dNode.Value(), err)
	}
	g.renderer.RenderNode(dNode, gNode)
	return gNode, nil
}
