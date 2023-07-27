package collatz

import "testing"

func TestTwoChild(t *testing.T) {

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
		r := TwoChild(q.n)
		if r != q.c {
			t.Errorf("TwoChild was incorrect, got: %d, want: %d.", q.n, q.c)
		}
	}

}

func TestThreeChild(t *testing.T) {

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
		r := ThreeChild(q.n)
		if r != q.c {
			t.Errorf("ThreeChild was incorrect, got: %d, want: %d.", q.n, q.c)
		}
	}

}
