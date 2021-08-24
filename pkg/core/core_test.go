package core

import (
	"testing"
)

func TestNodeFromInt(t *testing.T) {
	tests := []Fecund_node{{v: 4, n18: 0, r18: 4, n54: 0, r54: 4},
		{v: 16, n18: 0, r18: 16, n54: 0, r54: 16},
		{v: 22, n18: 1, r18: 4, n54: 0, r54: 22},
		{v: 34, n18: 1, r18: 16, n54: 0, r54: 34},
		{v: 40, n18: 2, r18: 4, n54: 0, r54: 40},
		{v: 52, n18: 2, r18: 16, n54: 0, r54: 52},
		{v: 58, n18: 3, r18: 4, n54: 1, r54: 4},
		{v: 70, n18: 3, r18: 16, n54: 1, r54: 16},
		{v: 76, n18: 4, r18: 4, n54: 1, r54: 22},
		{v: 88, n18: 4, r18: 16, n54: 1, r54: 34},
		{v: 94, n18: 5, r18: 4, n54: 1, r54: 40},
		{v: 106, n18: 5, r18: 16, n54: 1, r54: 52}}
	for _, q := range tests {
		r := NodeFromInt(q.v)
		if r != q {
			t.Errorf("NodeFromInt was incorrect for value %d", q.v)
		}
	}
}

func TestTwoChild(t *testing.T) {

	tests := []struct {
		n uint64
		c uint64
	}{
		{4, 16},
		{16, 256},
		{22, 88},
		{34, 544},
	}

	for _, q := range tests {
		r := TwoChild(q.n)
		if r != q.c {
			t.Errorf("TwoChild was incorrect, got: %d, want: %d.", q.n, q.c)
		}
	}

}

func TestThreeChild(t *testing.T) {

	tests := []struct {
		n uint64
		c uint64
	}{
		{4, 4},
		{16, 40},
		{22, 112},
		{34, 22},
		{40, 52},
		{52, 34},
		{58, 76},
	}

	for _, q := range tests {
		r := ThreeChild(q.n)
		if r != q.c {
			t.Errorf("ThreeChild was incorrect, got: %d, want: %d.", q.n, q.c)
		}
	}

}

func TestGenerateNodesDF(t *testing.T) {
	type args struct {
		start Fecund_node
		depth int8
		c     chan Fecund_node
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateNodesDF(tt.args.start, tt.args.depth, tt.args.c)
		})
	}
}
