package core

import (
	"testing"
)

func TestNodeFromInt(t *testing.T) {
	tests := []Node{{V: 4, N18: 0, R18: 4, N54: 0, R54: 4},
		{V: 16, N18: 0, R18: 16, N54: 0, R54: 16},
		{V: 22, N18: 1, R18: 4, N54: 0, R54: 22},
		{V: 34, N18: 1, R18: 16, N54: 0, R54: 34},
		{V: 40, N18: 2, R18: 4, N54: 0, R54: 40},
		{V: 52, N18: 2, R18: 16, N54: 0, R54: 52},
		{V: 58, N18: 3, R18: 4, N54: 1, R54: 4},
		{V: 70, N18: 3, R18: 16, N54: 1, R54: 16},
		{V: 76, N18: 4, R18: 4, N54: 1, R54: 22},
		{V: 88, N18: 4, R18: 16, N54: 1, R54: 34},
		{V: 94, N18: 5, R18: 4, N54: 1, R54: 40},
		{V: 106, N18: 5, R18: 16, N54: 1, R54: 52}}
	for _, q := range tests {
		r := NodeFromInt(q.V)
		if r != q {
			t.Errorf("NodeFromInt was incorrect for value %d", q.V)
		}
	}
}

func TestTwo(t *testing.T) {

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
		r := Two(q.n)
		if r != q.c {
			t.Errorf("TwoChild was incorrect, got: %d, want: %d.", q.n, q.c)
		}
	}

}

func TestThree(t *testing.T) {

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
		r := Three(q.n)
		if r != q.c {
			t.Errorf("ThreeChild was incorrect, got: %d, want: %d.", q.n, q.c)
		}
	}

}
