package render

import (
	"fmt"
	"ztalloc/ztalloc"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type (
	TreeGraph struct {
		start ztalloc.Node
		g     *graphviz.Graphviz
		graph *cgraph.Graph
	}
	chirality int
)

const (
	left chirality = iota
	right
)

func NewTreeGraph(root ztalloc.Transform) (TreeGraph, error) {
	g := graphviz.New()
	graph, err := g.Graph(graphviz.StrictDirected)
	if err != nil {
		return TreeGraph{}, err
	}
	return TreeGraph{
		g:     g,
		graph: graph,
	}, nil
}

func (t TreeGraph) Close() error {
	if err := t.graph.Close(); err != nil {
		return err
	}
	t.g.Close()
	return nil
}

func (g TreeGraph) Draw() error {
	gRoot, err := g.drawNode(g.start)
	if err != nil {
		return fmt.Errorf("Error drawing root %d: %w", g.start.Value(), err)
	}
	err = g.drawChildren(g.start, gRoot, 5)
	if err != nil {
		return fmt.Errorf("Error drawing tree rooted at %d: %w", g.start.Value(), err)
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
		if lerr == nil {
			lerr = g.drawChildren(ldNode, lgNode, depth-1)
		}
		if rerr == nil {
			rerr = g.drawChildren(rdNode, rgNode, depth-1)
		}
	}
	if lerr != nil || rerr != nil {
		return fmt.Errorf("Error drawing descendents of %d: left: %w, right: %w", dNode.Value(), lerr, rerr)
	}
	return nil
}

func (g TreeGraph) drawChild(dParent ztalloc.Node, gParent *cgraph.Node, child chirality) (ztalloc.Node, *cgraph.Node, error) {
	var dChild ztalloc.Node
	// var dEdge ztalloc.TransformName
	var suffix string
	var err error
	var dEdge ztalloc.TransformName

	if child == left {
		dChild, dEdge, err = dParent.LeftChild()
		suffix = "l"
	} else {
		dChild, dEdge, err = dParent.RightChild()
		suffix = "r"
	}
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to draw %s child of %d: %w", suffix, dParent.Value(), err)
	}

	gChild, err := g.drawNode(dChild)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to draw %s child of %d: %w", suffix, dParent.Value(), err)
	}
	gEdge, err := g.graph.CreateEdge(gChild.Name(), gParent, gChild)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to draw %s child of %d: %w", suffix, dParent.Value(), err)
	}

	RenderEdge(dEdge, gEdge)
	return dChild, gChild, nil
}

func (g TreeGraph) drawNode(dNode ztalloc.Node) (*cgraph.Node, error) {
	var gNode *cgraph.Node
	gNode, err := g.graph.CreateNode(fmt.Sprintf("%d", dNode.Value()))
	if err != nil {
		return nil, fmt.Errorf("Failed to draw node %d: %w", dNode.Value(), err)
	}
	RenderNode(dNode, gNode)
	return gNode, nil
}
