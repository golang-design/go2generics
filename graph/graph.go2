// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graph

// NodeConstraint is the type constraint for graph nodes:
// they must have an Edges method that returns the Edge's
// that connect to this Node.
type NodeConstraint[Edge any] interface {
	Edges() []Edge
}

// EdgeConstraint is the type constraint for graph edges:
// they must have a Nodes method that returns the two Nodes
// that this edge connects.
type EdgeConstraint[Node any] interface {
	Nodes() (from, to Node)
}

// Graph is a graph composed of nodes and edges.
type Graph[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]] struct {
}

// New returns a new graph given a list of nodes.
func New[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]] (
	nodes []Node) *Graph[Node, Edge] {
	// TODO:
	return nil
}

// ShortestPath returns the shortest path between two nodes,
// as a list of edges.
func (g *Graph[Node, Edge]) ShortestPath(from, to Node) []Edge {
	// TODO:
	return nil
}
