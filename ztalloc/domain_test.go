package ztalloc_test

import (
	"reflect"
	"testing"
	. "ztalloc/ztalloc"
)

func TestNewZTree(t *testing.T) {
	v := 123
	t.Run("Root value", func(t *testing.T) {
		if got := ZNode(v); got.Value() != v {
			t.Errorf("NewZTree() = %v, want %v", got, v)
		}
	})
}

func TestZNode_LeftChild(t *testing.T) {
	tests := []struct {
		name          string
		n             ZNode
		wantNode      ZNode
		wantTransform TransformName
		wantOk        bool
	}{
		{
			name:          "zero",
			n:             0,
			wantNode:      0,
			wantTransform: Z2,
			wantOk:        false,
		},
		{
			name:          "one",
			n:             1,
			wantNode:      0,
			wantTransform: Z2,
			wantOk:        false,
		},
		{
			name:          "sixteen",
			n:             16,
			wantNode:      32,
			wantTransform: Z2,
			wantOk:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNode, gotTransform, gotOk := tt.n.LeftChild()
			if gotOk != tt.wantOk {
				t.Errorf("ZNode.LeftChild()want = %v, wantErr %v", gotOk, tt.wantOk)
				return
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("ZNode.LeftChild() got = %v, want %v", gotNode, tt.wantNode)
			}
			if !reflect.DeepEqual(gotTransform, tt.wantTransform) {
				t.Errorf("ZNode.LeftChild() got1 = %v, want %v", gotTransform, tt.wantTransform)
			}
		})
	}
}

func TestZNode_RightChild(t *testing.T) {
	tests := []struct {
		name          string
		n             ZNode
		wantNode      ZNode
		wantTransform TransformName
		wantOk        bool
	}{
		{
			name:          "mod zero too small",
			n:             3,
			wantNode:      0,
			wantTransform: Z3,
			wantOk:        false,
		},
		{
			name:          "mod one too small",
			n:             1,
			wantNode:      0,
			wantTransform: Z3,
			wantOk:        false,
		},
		{
			name:          "mode two too small",
			n:             2,
			wantNode:      0,
			wantTransform: Z3,
			wantOk:        false,
		},
		{
			name:          "mod zero",
			n:             6,
			wantNode:      0,
			wantTransform: Z3,
			wantOk:        false,
		},
		{
			name:          "mod one",
			n:             4,
			wantNode:      1,
			wantTransform: Z3,
			wantOk:        true,
		},
		{
			name:          "mod two",
			n:             5,
			wantNode:      0,
			wantTransform: Z3,
			wantOk:        false,
		},
		{
			name:          "odd mod 1",
			n:             7,
			wantNode:      0,
			wantTransform: Z3,
			wantOk:        false,
		},
		{
			name:          "even mod 1",
			n:             10,
			wantNode:      3,
			wantTransform: Z3,
			wantOk:        true,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNode, gotTransform, gotOk := tt.n.RightChild()
			if gotOk != tt.wantOk {
				t.Errorf("ZNode.RightChild()want = %v, wantErr %v", gotOk, tt.wantOk)
				return
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("ZNode.RightChild() got = %v, want %v", gotNode, tt.wantNode)
			}
			if !reflect.DeepEqual(gotTransform, tt.wantTransform) {
				t.Errorf("ZNode.RightChild() got1 = %v, want %v", gotTransform, tt.wantTransform)
			}
		})
	}
}

func TestZNode_Value(t *testing.T) {
	tests := []struct {
		name string
		n    ZNode
		want int
	}{
		{
			name: "zero",
			n:    0,
			want: 0,
		},
		{
			name: "13",
			n:    13,
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Value(); got != tt.want {
				t.Errorf("ZNode.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCNode_LeftChild(t *testing.T) {
	tests := []struct {
		name          string
		n             CNode
		wantNode      Node
		wantTransform TransformName
		wantOk        bool
	}{
		{
			name:          "too small zero",
			n:             CNode(0),
			wantNode:      CNode(0),
			wantTransform: C2,
			wantOk:        false,
		},
		{
			name:          "odd one",
			n:             CNode(1),
			wantNode:      CNode(0),
			wantTransform: C2,
			wantOk:        false,
		},
		{
			name:          "even two",
			n:             CNode(2),
			wantNode:      CNode(1),
			wantTransform: C2,
			wantOk:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNode, gotTransform, gotOk := tt.n.LeftChild()
			if gotOk != tt.wantOk {
				t.Errorf("CNode.LeftChild()want = %v, wantErr %v", gotOk, tt.wantOk)
				return
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("CNode.LeftChild() got = %v, want %v", gotNode, tt.wantNode)
			}
			if !reflect.DeepEqual(gotTransform, tt.wantTransform) {
				t.Errorf("CNode.LeftChild() got1 = %v, want %v", gotTransform, tt.wantTransform)
			}
		})
	}
}

func TestCNode_RightChild(t *testing.T) {
	tests := []struct {
		name          string
		n             CNode
		wantNode      Node
		wantTransform TransformName
		wantOk        bool
	}{
		{
			name:          "zero",
			n:             CNode(0),
			wantNode:      CNode(0),
			wantTransform: C3,
			wantOk:        false,
		},
		{
			name:          "one",
			n:             CNode(1),
			wantNode:      CNode(4),
			wantTransform: C3,
			wantOk:        true,
		},
		{
			name:          "two",
			n:             CNode(2),
			wantNode:      CNode(0),
			wantTransform: C3,
			wantOk:        false,
		},
		{
			name:          "three",
			n:             CNode(3),
			wantNode:      CNode(10),
			wantTransform: C3,
			wantOk:        true,
		},
		{
			name:          "four",
			n:             CNode(4),
			wantNode:      CNode(0),
			wantTransform: C3,
			wantOk:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNode, gotTransform, gotOk := tt.n.RightChild()
			if gotOk != tt.wantOk {
				t.Errorf("CNode.RightChild()want = %v, wantErr %v", gotOk, tt.wantOk)
				return
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("CNode.RightChild() got = %v, want %v", gotNode, tt.wantNode)
			}
			if !reflect.DeepEqual(gotTransform, tt.wantTransform) {
				t.Errorf("CNode.RightChild() got1 = %v, want %v", gotTransform, tt.wantTransform)
			}
		})
	}
}

func TestCNode_Value(t *testing.T) {
	tests := []struct {
		name string
		n    CNode
		want int
	}{
		{
			name: "zero",
			n:    0,
			want: 0,
		},
		{
			name: "13",
			n:    13,
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Value(); got != tt.want {
				t.Errorf("CNode.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
