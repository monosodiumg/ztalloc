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
			wantTransform: NONE,
			wantOk:        false,
		},
		{
			name:          "one",
			n:             1,
			wantNode:      0,
			wantTransform: NONE,
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
				t.Errorf("ZNode.LeftChild() gotOk = %v, wantOk %v", gotOk, tt.wantOk)
				return
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("ZNode.LeftChild() gotNode = %v, wantNode %v", gotNode, tt.wantNode)
			}
			if !reflect.DeepEqual(gotTransform, tt.wantTransform) {
				t.Errorf("ZNode.LeftChild() gotTransform = %v, wantTransform %v", gotTransform, tt.wantTransform)
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
			wantTransform: NONE,
			wantOk:        false,
		},
		{
			name:          "mod one too small",
			n:             1,
			wantNode:      0,
			wantTransform: NONE,
			wantOk:        false,
		},
		{
			name:          "mode two too small",
			n:             2,
			wantNode:      0,
			wantTransform: NONE,
			wantOk:        false,
		},
		{
			name:          "mod zero",
			n:             6,
			wantNode:      0,
			wantTransform: NONE,
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
			wantTransform: NONE,
			wantOk:        false,
		},
		{
			name:          "odd mod 1",
			n:             7,
			wantNode:      0,
			wantTransform: NONE,
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
				t.Errorf("ZNode.RightChild()gotOk = %v, wantOk %v", gotOk, tt.wantOk)
				return
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("ZNode.RightChild() gotNode = %v, wantNode %v", gotNode, tt.wantNode)
			}
			if !reflect.DeepEqual(gotTransform, tt.wantTransform) {
				t.Errorf("ZNode.RightChild() gotTransform = %v, wantTransform %v", gotTransform, tt.wantTransform)
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
		wantNode      CNode
		wantTransform TransformName
		wantOk        bool
	}{
		{
			name:          "too small zero",
			n:             0,
			wantNode:      0,
			wantTransform: NONE,
			wantOk:        false,
		},
		{
			name:          "odd one",
			n:             CNode(1),
			wantNode:      0,
			wantTransform: NONE,
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
				t.Errorf("CNode.LeftChild()gotOk = %v, wantOk %v", gotOk, tt.wantOk)
				return
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("CNode.LeftChild() gotNode = %v, wantNode %v", gotNode, tt.wantNode)
			}
			if !reflect.DeepEqual(gotTransform, tt.wantTransform) {
				t.Errorf("CNode.LeftChild() gotTransform = %v, wantTransform %v", gotTransform, tt.wantTransform)
			}
		})
	}
}

func TestCNode_RightChild(t *testing.T) {
	tests := []struct {
		name          string
		n             CNode
		wantNode      CNode
		wantTransform TransformName
		wantOk        bool
	}{
		{
			name:          "zero",
			n:             0,
			wantNode:      0,
			wantTransform: NONE,
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
			wantNode:      0,
			wantTransform: NONE,
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
			wantNode:      0,
			wantTransform: NONE,
			wantOk:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNode, gotTransform, gotOk := tt.n.RightChild()
			if gotOk != tt.wantOk {
				t.Errorf("CNode.RightChild()gotOk = %v, wantOk %v", gotOk, tt.wantOk)
				return
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("CNode.RightChild() gotNode = %v, wantNode %v", gotNode, tt.wantNode)
			}
			if !reflect.DeepEqual(gotTransform, tt.wantTransform) {
				t.Errorf("CNode.RightChild() gotTransform = %v, wantTransform %v", gotTransform, tt.wantTransform)
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

func TestZBinaryNode_LeftChild(t *testing.T) {
	tests := []struct {
		name   string
		n      ZBinaryNode
		wantN  ZBinaryNode
		wantT  TransformName
		wantOk bool
	}{
		{
			name:   "zero",
			n:     0,
			wantN:  0,
			wantT:  NONE,
		},
		{
			name:   "5",
			n:      5,
			wantN:  40,
			wantT:  Z222,
			wantOk: true,
		},
		{
			name:   "10",
			n:      10,
			wantN:  40,
			wantT:  Z22,
			wantOk: true,
		},
		{
			name:   "32",
			n:      32,
			wantN:  256,
			wantT:  Z222,
			wantOk: true,
		},
		{
			name:   "64",
			n:      64,
			wantN:  256,
			wantT:  Z22,
			wantOk: true,
		},
		{
			name:   "256",
			n:      256,
			wantN:  1024,
			wantT:  Z22,
			wantOk: true,
		},
		{
			name:   "1024",
			n:      1024,
			wantN:  16384,
			wantT:  Z2222,
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, gotT, gotOk := tt.n.LeftChild()
			if !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("ZBinaryNode.LeftChild() gotN = %v, wantN %v", gotN, tt.wantN)
			}
			if !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("ZBinaryNode.LeftChild() gotT = %v, wantT %v", gotT, tt.wantT)
			}
			if gotOk != tt.wantOk {
				t.Errorf("ZBinaryNode.LeftChild() gotOk = %v, wantOk %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestZBinaryNode_RightChild(t *testing.T) {
	tests := []struct {
		name   string
		n      ZBinaryNode
		wantN  ZBinaryNode
		wantT  TransformName
		wantOk bool
	}{
		{
			name:   "zero",
			n:     0,
			wantN:  0,
			wantT:  NONE,
		},
		{
			name:   "5",
			n:      5,
			wantN:  0,
			wantT:  NONE,
			wantOk: false,
		},
		{
			name:   "10",
			n:      10,
			wantN:  0,
			wantT:  NONE,
			wantOk: false,
		},
		{
			name:   "16",
			n:      16,
			wantN:  40,
			wantT:  Z3222,
			wantOk: true,
		},		
		{
			name:   "32",
			n:      32,
			wantN:  0,
			wantT:  NONE,
			wantOk: false,
		},
		{
			name:   "52",
			n:      52,
			wantN:  34,
			wantT:  Z32,
			wantOk: true,
		},				
		{
			name:   "64",
			n:      64,
			wantN:  0,
			wantT:  NONE,
			wantOk: false,
		},
		{
			name:   "256",
			n:      256,
			wantN:  340,
			wantT:  Z322,
			wantOk: true,
		},
		{
			name:   "1024",
			n:      1024,
			wantN:  682,
			wantT:  Z32,
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, gotT, gotOk := tt.n.RightChild()
			if !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("ZBinaryNode.LeftChild() gotN = %v, wantN %v", gotN, tt.wantN)
			}
			if !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("ZBinaryNode.LeftChild() gotT = %v, wantT %v", gotT, tt.wantT)
			}
			if gotOk != tt.wantOk {
				t.Errorf("ZBinaryNode.LeftChild() gotOk = %v, wantOk %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestIsFertile(t *testing.T) {
	tests := []struct {
		name string
		arg int
		want bool
	}{
		{
			name: "one",
			arg: 1,
			want: false,
		},
		{
			name: "two",
			arg: 2,
			want: false,
		},		
		{
			name: "three",
			arg: 3,
			want: false,
		},		
		{
			name: "four",
			arg: 4,
			want: true,
		},		
		{
			name: "sixteen",
			arg: 16,
			want: true,
		},		
		{
			name: "twentyone",
			arg: 21,
			want: false,
		},	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFertile(tt.arg); got != tt.want {
				t.Errorf("IsFertile() = %v, want %v", got, tt.want)
			}
		})
	}
}
