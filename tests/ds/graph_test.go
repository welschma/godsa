package ds_test

import (
	"bytes"
	"testing"

	"github.com/welschma/godsa/ds"
)

func TestGraph(t *testing.T) {
	g := ds.NewGraph(3, false) // assuming NewGraph now takes a third argument for the directed flag

	if g.V() != 3 {
		t.Errorf("Expected 3 vertices, got %d", g.V())
	}
	if g.E() != 0 {
		t.Errorf("Expected 0 edges, got %d", g.E())
	}

	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 1)

	if g.E() != 4 { // if the graph is undirected, each edge is counted twice
		t.Errorf("Expected 4 edges, got %d", g.E())
	}

	var buffer bytes.Buffer
	g.Write(&buffer)

	expected := "3 4\n0: (1, 1) \n1: (2, 1) (0, 1) \n2: (1, 1) \n" // if the graph is undirected, each edge appears twice
	if buffer.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, buffer.String())
	}
}
