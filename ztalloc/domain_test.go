package ztalloc

import (
	"reflect"
	"testing"
)

func TestRoot(t *testing.T) {
	tests := []struct {
		name string
		want Node
	}{
		{name: "root is 16", want: node(16)},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Root(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Root() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTransformType(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTransformType(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransformType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetTransformType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		want    int
		wantErr bool
	}{
		{
			name:    "n=4",
			n:       4,
			want:    4,
			wantErr: false,
		},
		{
			name:    "n=5",
			n:       5,
			want:    5,
			wantErr: true,
		},
		{
			name:    "n=6",
			n:       6,
			want:    6,
			wantErr: true,
		},
		{
			name:    "n=16",
			n:       16,
			want:    16,
			wantErr: false,
		},
		{
			name:    "n=22",
			n:       22,
			want:    22,
			wantErr: false,
		},
		{
			name:    "n=22",
			n:       22,
			want:    22,
			wantErr: false,
		},
		{
			name:    "n=28",
			n:       28,
			want:    28,
			wantErr: false,
		},
		{
			name:    "n=34",
			n:       34,
			want:    34,
			wantErr: false,
		},
		{
			name:    "n=40",
			n:       40,
			want:    40,
			wantErr: false,
		},
		{
			name:    "n=52",
			n:       52,
			want:    52,
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, err := Get(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err == nil)			 {
				got := node.Value()
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Get() = %v, want %v", got, tt.want)
				}
	
			}
			
		})
	}
}

func Test_node_Parent(t *testing.T) {
	tests := []struct {
		name    string
		n       node
		want    Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.Parent()
			if (err != nil) != tt.wantErr {
				t.Errorf("node.Parent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.Parent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_LeftChild(t *testing.T) {
	tests := []struct {
		name    string
		n       node
		want    Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.LeftChild()
			if (err != nil) != tt.wantErr {
				t.Errorf("node.LeftChild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.LeftChild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_RightChild(t *testing.T) {
	tests := []struct {
		name    string
		n       node
		want    Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.RightChild()
			if (err != nil) != tt.wantErr {
				t.Errorf("node.RightChild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.RightChild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Value(t *testing.T) {
	tests := []struct {
		name string
		n    node
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Value(); got != tt.want {
				t.Errorf("node.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
