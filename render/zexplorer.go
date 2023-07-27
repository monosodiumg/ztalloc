package render

import (
	"fmt"
	"ztalloc/ztalloc"
)

type (
	Traverser interface {
		// Traverse treats the origin as having radius zero.
		Traverse(Visitor) error
	}

	traverser struct {
		origin ztalloc.Node
		radius int
	}
	Visitor interface {
		// Visit the node. edge is the incoming edge, radius is the distance from the origin.
		Visit(node int, edge string, radius int) (Visitor, error)
	}
)

func NewTraverser(origin ztalloc.Node, radius int) Traverser {
	return &traverser{
		origin: origin,
		radius: radius,
	}
}

// Calls Visit on on every node up to radius number of steps. The origin is at radius 0.
// Nodes may be visted concurrently. Traversal order is not be stable.
func (t *traverser) Traverse(v Visitor) error {
	rootVisitor, err := v.Visit(t.origin.Value(), "", 0)
	if err != nil {
		return fmt.Errorf("Error visting root %d: %w", t.origin.Value(), err)
	}
	if t.radius > 0 {
		t.traverseChildren(t.origin, rootVisitor, 1)
	}
	return nil
}

func (t *traverser) traverseChildren(parent ztalloc.Node, v Visitor, radius int) {
	if radius <= t.radius {
		if lChild, lEdge, lOk := parent.LeftChild(); lOk {
			lVisitor, lerr := v.Visit(lChild.Value(), lEdge, radius+1)
			if lerr == nil {
				t.traverseChildren(lChild, lVisitor, radius+1)
			}
		}
		if rChild, rEdge, rOk := parent.RightChild(); rOk {
			rVisitor, rerr := v.Visit(rChild.Value(), rEdge, radius+1)
			if rerr == nil {
				t.traverseChildren(rChild, rVisitor, radius+1)
			}
		}
	}
}
