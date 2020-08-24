package graph

import "testing"

type Node struct {
	n int
}

func (n Node) Edges() []Edge {
	return nil
}

type Edge struct {
	n int
}

func (e Edge) Nodes() (from, to Node) {
	return Node{}, Node{}
}

func TestGraph(t *testing.T) {
	g := New[Node, Edge]([]Node{})
	g.ShortestPath(Node{}, Node{})
}
