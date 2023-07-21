package render_test

import (
	"github.com/goccy/go-graphviz"
	"reflect"
	"testing"
	"ztalloc/render"
	"ztalloc/ztalloc"
)

func TestNewTreeGraph(t *testing.T) {
	type args struct {
		gv       *graphviz.Graphviz
		renderer render.TreeRenderer
		start    ztalloc.Node
		depth    int
	}
	tests := []struct {
		name    string
		args    args
		want    render.TreeGraph
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := render.NewTreeGraph(tt.args.gv, tt.args.renderer, tt.args.start, tt.args.depth)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTreeGraph() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTreeGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTreeGraph_Draw_ZBinary(t *testing.T) {
	tests := []struct {
		name     string
		start    ztalloc.Node
		depth    int
		renderer render.ZTreeRenderer
		wantErr  bool
	}{
		{
			name:     "sixteen",
			start:    ztalloc.ZBinaryNode(16),
			depth:    4,
			renderer: render.ZTreeRenderer{},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr, _ := render.NewTreeGraph(graphviz.New(), tt.renderer, tt.start, tt.depth)
			if err := tr.Draw(); (err != nil) != tt.wantErr {
				t.Errorf("TreeGraph.Draw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
