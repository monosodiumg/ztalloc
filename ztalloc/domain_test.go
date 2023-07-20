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
		name    string
		n       ZNode
		want    ZNode
		want1   TransformName
		wantErr bool
	}{
		{
			name:    "zero",
			n:       0,
			want:    0,
			want1:   Z2,
			wantErr: false,
		},
		{
			name:    "one",
			n:       1,
			want:    2,
			want1:   Z2,
			wantErr: false,
		},
		{
			name:    "sixteen",
			n:       16,
			want:    32,
			want1:   Z2,
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.n.LeftChild()
			if (err != nil) != tt.wantErr {
				t.Errorf("ZNode.LeftChild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZNode.LeftChild() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ZNode.LeftChild() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestZNode_RightChild(t *testing.T) {
	tests := []struct {
		name    string
		n       ZNode
		want    ZNode
		want1   TransformName
		wantErr bool
	}{
		{
			name:    "mod zero too small",
			n:       3,
			want:    0,
			want1:   Z3,
			wantErr: true,
		},
		{
			name:    "mod one too small",
			n:       1,
			want:    0,
			want1:   Z3,
			wantErr: true,
		},
		{
			name:    "mode two too small",
			n:       2,
			want:    0,
			want1:   Z3,
			wantErr: true,
		},
		{
			name:    "mod zero",
			n:       6,
			want:    0,
			want1:   Z3,
			wantErr: true,
		},
		{
			name:    "mod one",
			n:       4,
			want:    1,
			want1:   Z3,
			wantErr: false,
		},
		{
			name:    "mode two",
			n:       5,
			want:    0,
			want1:   Z3,
			wantErr: true,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.n.RightChild()
			if (err != nil) != tt.wantErr {
				t.Errorf("ZNode.RightChild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ZNode.RightChild() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ZNode.RightChild() got1 = %v, want %v", got1, tt.want1)
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
		}, // TODO: Add test cases.
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
		name    string
		n       CNode
		want    Node
		want1   TransformName
		wantErr bool
	}{
		{
			name:    "zero",
			n:       CNode(0),
			want:    nil,
			want1:   C2,
			wantErr: true,
		},
		{
			name:    "one",
			n:       CNode(1),
			want:    nil,
			want1:   C2,
			wantErr: true,
		},
		{
			name:    "two",
			n:       CNode(2),
			want:    CNode(1),
			want1:   C2,
			wantErr: false,
		},
		{
			name:    "three",
			n:       CNode(3),
			want:    nil,
			want1:   C2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.n.LeftChild()
			if (err != nil) != tt.wantErr {
				t.Errorf("CNode.LeftChild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CNode.LeftChild() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CNode.LeftChild() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCNode_RightChild(t *testing.T) {
	tests := []struct {
		name    string
		n       CNode
		want    Node
		want1   TransformName
		wantErr bool
	}{
		{
			name:    "zero",
			n:       CNode(0),
			want:    nil,
			want1:   C3,
			wantErr: true,
		},
		{
			name:    "one",
			n:       CNode(1),
			want:    CNode(4),
			want1:   C3,
			wantErr: false,
		},
		{
			name:    "two",
			n:       CNode(2),
			want:    nil,
			want1:   C3,
			wantErr: true,
		},
		{
			name:    "three",
			n:       CNode(3),
			want:    CNode(10),
			want1:   C3,
			wantErr: false,
		},
		{
			name:    "four",
			n:       CNode(4),
			want:    nil,
			want1:   C3,
			wantErr: true,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.n.RightChild()
			if (err != nil) != tt.wantErr {
				t.Errorf("CNode.RightChild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CNode.RightChild() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CNode.RightChild() got1 = %v, want %v", got1, tt.want1)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Value(); got != tt.want {
				t.Errorf("CNode.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
