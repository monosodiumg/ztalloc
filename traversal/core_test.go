package traversal

import (
	"reflect"
	"testing"
)

//left is the three child, right is the two child
var nodesWithChildren = []struct{
	node Node
	left, right int}{
	{Node{V: 4, N18: 0, R18: 4, N54: 0, R54: 4}, 4 , 16},
	{Node{V: 16, N18: 0, R18: 16, N54: 0, R54: 16}, 40 , 256},
	{Node{V: 22, N18: 1, R18: 4, N54: 0, R54: 22}, 122 , 88},
	{Node{V: 34, N18: 1, R18: 16, N54: 0, R54: 34}, 22 , 544},
	{Node{V: 40, N18: 2, R18: 4, N54: 0, R54: 40}, 52 , 160},
	{Node{V: 52, N18: 2, R18: 16, N54: 0, R54: 52}, 34 , 832},
	{Node{V: 58, N18: 3, R18: 4, N54: 1, R54: 4}, 76 , 232},
	{Node{V: 70, N18: 3, R18: 16, N54: 1, R54: 16}, 184 , 1120},
	{Node{V: 76, N18: 4, R18: 4, N54: 1, R54: 22}, 400 , 304},
	{Node{V: 88, N18: 4, R18: 16, N54: 1, R54: 34}, 58 , 1408},
	{Node{V: 94, N18: 5, R18: 4, N54: 1, R54: 40}, 124 , 376},
	{NodeV: 106, N18: 5, R18: 16, N54: 1, R54: 52}, 70 , 1696},
}

func TestNodeFromInt(t *testing.T) {
	tests := nodes
	for _, q := range tests {
		r := NodeFromInt(q.V)
		if r != q {
			t.Errorf("NodeFromInt was incorrect for value %d", q.V)
		}
	}
}

func TestTwo(t *testing.T) {

	tests := []struct {
		n int
		c int
	}{
		{4, 16},
		{16, 256},
		{22, 88},
		{34, 544},
	}

	for _, q := range tests {
		r := Two(q.n)
		if r != q.c {
			t.Errorf("TwoChild was incorrect, got: %d, want: %d.", q.n, q.c)
		}
	}

}

func TestThree(t *testing.T) {

	tests := []struct {
		n int
		c int
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
		r := Three(q.n)
		if r != q.c {
			t.Errorf("ThreeChild was incorrect, got: %d, want: %d.", q.n, q.c)
		}
	}

}

func TestNode_TwoChild(t *testing.T) {
	tests := nodes
	for _, q := range tests {
		
		if r != q.v {
			t.Errorf("ThreeChild was incorrect, got: %d, want: %d.", q.n, q.c)
		}
	}
	}
}
