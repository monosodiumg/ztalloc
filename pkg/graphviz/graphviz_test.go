package graphviz

import (
	"testing"
)

func TestGraphvizDraw(t *testing.T) {
	src := make(chan uint64, 10)
	dst := make(chan string, 10)

	tests := []struct {
		n uint64
		r string
	}{
		{4, "v4\n"},
		{16, "v16\n"},
		{22, "v22\n"},
		{34, "v34\n"},
	}

	for _, q := range tests {
		//fmt.Printf("writing src <-%d\n", q.n)
		src <- q.n
		//fmt.Printf("wrote src <-%d\n", q.n)
	}

	go GraphvizDraw(src, dst)
	var viz string
	for i, q := range tests {
		viz = <-dst
		if viz != q.r {
			t.Errorf("GraphvizDraw was incorrect, got: %s, want: %s for item at position %d with value %d.", viz, q.r, i, q.n)
		}

	}

}

func TestGraphvizNodeRender(t *testing.T) {
	var node uint64 = 4
	r := GraphvizNodeRender(node)
	//fmt.Printf("TestGraphvizNodeRender::r=%s",r)
	if r != "v4\n" {
		t.Errorf("got: %s, want: %s.\n", r, "4\n")
	}
}
