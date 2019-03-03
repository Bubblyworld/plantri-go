package plantri

import (
	"errors"
	"fmt"
)

var ErrAdjMatrixOutOfBounds = errors.New("plantri: adjacency matrix vertex index out of bounds")

// Graph is an abstraction of a simple graph.
type Graph interface {
	// Vertices returns the list of vertices in the graph.
	Vertices() []Vertex

	// Edges returns the list of edges in the graph.
	Edges() []Edge
}

// Vertex represents a vertex in a simple graph.
type Vertex interface {
	// Neighbours returns all neighbouring vertices of this vertex.
	Neighbours() []Vertex
}

// Edge represents a directed edge between two vertices.
type Edge struct {
	Source Vertex
	Dest   Vertex
}

// adjMatrix is a graph with edges represented by an adjacency matrix.
type adjMatrix struct {
	n      int
	matrix [][]bool
}

func newAdjMatrix(n int) *adjMatrix {
	var matrix [][]bool
	for i := 0; i < n; i++ {
		matrix = append(matrix, []bool{})
		for j := 0; j < n; j++ {
			eq := false
			if i == j {
				eq = true
			}

			matrix[i] = append(matrix[i], eq)
		}
	}

	return &adjMatrix{
		n:      n,
		matrix: matrix,
	}
}

func (am *adjMatrix) inBounds(i int) bool {
	return i >= 0 && i < am.n
}

func (am *adjMatrix) getVertex(i int) (*adjMatrixVertex, error) {
	if !am.inBounds(i) {
		return nil, ErrAdjMatrixOutOfBounds
	}

	return &adjMatrixVertex{
		parent: am,
		index:  i,
	}, nil
}

func (am *adjMatrix) addEdge(i, j int) error {
	if !am.inBounds(i) || !am.inBounds(j) {
		return ErrAdjMatrixOutOfBounds
	}

	am.matrix[i][j] = true
	am.matrix[j][i] = true
	return nil
}

func (am *adjMatrix) Vertices() []Vertex {
	var res []Vertex

	for i := 0; i < am.n; i++ {
		v, err := am.getVertex(i)
		if err != nil { // should never happen
			panicUnexpected("adjMatrix.Vertices", err)
		}

		res = append(res, v)
	}

	return res
}

func (am *adjMatrix) Edges() []Edge {
	var res []Edge

	for i := 0; i < am.n; i++ {
		for j := i + 1; j < am.n; j++ {
			if !am.matrix[i][j] {
				continue
			}

			v1, err := am.getVertex(i)
			if err != nil {
				panicUnexpected("adjMatrix.Edges", err)
			}

			v2, err := am.getVertex(j)
			if err != nil {
				panicUnexpected("adjMatrix.Edges", err)
			}

			res = append(res, Edge{
				Source: v1,
				Dest:   v2,
			})
		}
	}

	return res
}

// adjMatrixVertex is a vertex of an adjMatrix graph.
type adjMatrixVertex struct {
	index  int
	parent *adjMatrix
}

func (amv *adjMatrixVertex) Neighbours() []Vertex {
	var res []Vertex

	for i := 0; i < amv.parent.n; i++ {
		if i == amv.index {
			continue
		}

		if !amv.parent.matrix[amv.index][i] {
			continue
		}

		v, err := amv.parent.getVertex(i)
		if err != nil { // should never happen
			panicUnexpected("adjMatrixVertex.Neighbours", err)
		}

		res = append(res, v)
	}

	return res
}

// panicUnexpected panics with a standard error message on errors that should
// never happen logically. This should only be called if there is a critical
// failure in plantri's code.
func panicUnexpected(fn string, err error) {
	panic(fmt.Errorf("plantri: unexpected error in %s: %v", fn, err))
}

// Compile-time interface implementation checks.
var _ Graph = new(adjMatrix)
var _ Vertex = new(adjMatrixVertex)
