package ds

import (
	"fmt"
	"io"
	"os"
)

type Vertex struct {
	y      int
	weight int
	next   *Vertex
}

type Graph struct {
	v   int
	e   int
	adj []*Vertex
	directed bool
}

func NewGraph(v int, directed bool) *Graph {
	adj := make([]*Vertex, v)

	for i := 0; i < v; i++ {
		adj[i] = nil
	}

	return &Graph{v, 0, adj, directed}
}

// V returns the number of vertices in the graph.
func (g *Graph) V() int {
	return g.v
}

// E returns the number of edges in the graph.
func (g *Graph) E() int {
	return g.e
}

// AddEdge adds an edge between vertices x and y with the given weight.
func (g *Graph) AddEdge(x, y, weight int) {
	g.adj[x] = &Vertex{y, weight, g.adj[x]}
	g.e++

	if !g.directed {
		g.adj[y] = &Vertex{x, weight, g.adj[y]}
		g.e++
	}
}

// Write writes the graph to the given writer.
func (g *Graph) Write(w io.Writer) {
	fmt.Fprintf(w, "%d %d\n", g.v, g.e)

	for i := 0; i < g.v; i++ {
		fmt.Fprintf(w, "%d: ", i)
		for v := g.adj[i]; v != nil; v = v.next {
			fmt.Fprintf(w, "(%d, %d) ", v.y, v.weight)
		}
		fmt.Fprintln(w)
	}
}

// Print prints the graph.
func (g *Graph) Print() {
	g.Write(os.Stdout)
}
