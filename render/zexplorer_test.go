package render_test

import (
	"testing"
	"ztalloc/render"
	"ztalloc/ztalloc"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type (
	NodeVisit struct {
		Node   int
		Edge   string
		Radius int
	}
	NodeVisits struct {
		Visited []NodeVisit
	}

	// visitImpl struct {
	// 	visit func(node int, edge string, radius int) (render.Visitor, error)
	// }
)


// func (vi visitImpl) Visit(node int, edge string, radius int) (render.Visitor, error) {
// 	visitor, err := vi.Visit(node, edge, radius)
// 	return visitImpl{visit: visitor.Visit}, err
// }

func (v *NodeVisits) Visit(node int, edge string, radius int) (render.Visitor, error) {
	if v.Visited == nil {
		v.Visited = make([]NodeVisit, 0)
	}
	v.Visited = append(v.Visited, NodeVisit{
		Node:   node,
		Edge:   edge,
		Radius: radius,
	})
	return v, nil
}

func nodeVisitLess(a, b NodeVisit) bool { return a.Node < b.Node }

func TestNewTraverser(t *testing.T) {
	t.Run("working Traverser", func(t *testing.T) {
		wantVisits := NodeVisits{
			Visited: []NodeVisit{
				{
					Node:   2,
					Edge:   "",
					Radius: 0,
				},
			},
		}
		got := render.NewTraverser(ztalloc.ZNode(2), 0)
		v := NodeVisits{}
		err := got.Traverse(&v)

		if !cmp.Equal(v.Visited, wantVisits.Visited, cmpopts.SortSlices(nodeVisitLess)) {
			t.Errorf("NewTraverser() failed Traverse with simple Visitor, gotVisit %v, wantVisit %v", v.Visited[0].Node, wantVisits.Visited[0].Node)
		}
		if err != nil {
			t.Errorf("NewTraverser() failed Traverse with simple Visitor, got Err %v", err)
		}
	})
}

// func Test_traverser_Traverse(t *testing.T) {
// 	type fields struct {
// 		origin ztalloc.Node
// 		radius int
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		visitor render.Visitor
// 		wantErr bool
// 	}{
// 		{
// 			name: "Zero radius",
// 			fields: fields{
// 				origin: ztalloc.CNode(13),
// 				radius: 0,
// 			},
// 			visitor: make(nodeVisitSlice, 10),
// 			wantErr: false,
// 		},
// 		//TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tr := render.NewTraverser(tt.fields.origin, tt.fields.radius)

// 			if err := tr.Traverse(tt.visitor); (err != nil) != tt.wantErr {
// 				t.Errorf("traverser.Traverse() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
