package plantri

import (
	"errors"
	"fmt"
)

var ErrAdjMatrixOutOfBounds = errors.New("plantri: adjacency matrix vertex index out of bounds")

// Graph is an abstraction of a simple graph with labelled vertices.
type Graph interface {
	// Size returns the number of edges in the graph.
	Size() int

	// Order returns the number of vertices in the graph.
	Order() int

	// Vertices returns the list of vertices in the graph.
	Vertices() []Vertex

	// Edges returns the list of edges in the graph.
	Edges() []Edge
}

// Vertex represents a vertex in a simple graph.
type Vertex interface {
	// Id returns the label associated with this vertex.
	Id() int
}

// Edge represents a directed edge between two vertices.
type Edge struct {
	Source Vertex
	Dest   Vertex
}

// AdjMatrix is a graph with edges represented by an adjacency matrix.
type AdjMatrix struct {
	n      int
	matrix [][]bool
}

func newAdjMatrix(n int) *AdjMatrix {
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

	return &AdjMatrix{
		n:      n,
		matrix: matrix,
	}
}
func AsAdjMatrix(g Graph) *AdjMatrix {
	if am, ok := g.(*AdjMatrix); ok {
		return am
	}

	var vlabels []int
	for _, v := range g.Vertices() {
		vlabels = append(vlabels, v.Id())
	}

	vindexes := make(map[int]int)
	for i, vlabel := range vlabels {
		vindexes[vlabel] = i
	}

	res := newAdjMatrix(g.Order())
	for _, e := range g.Edges() {
		err := res.addEdge(vindexes[e.Source.Id()], vindexes[e.Dest.Id()])
		if err != nil { // should never happen
			panicUnexpected("AsAdjMatrix", err)
		}
	}

	return res
}

func (am *AdjMatrix) inBounds(i int) bool {
	return i >= 0 && i < am.n
}

func (am *AdjMatrix) getVertex(i int) (*adjMatrixVertex, error) {
	if !am.inBounds(i) {
		return nil, ErrAdjMatrixOutOfBounds
	}

	return &adjMatrixVertex{}, nil
}

func (am *AdjMatrix) addEdge(i, j int) error {
	if !am.inBounds(i) || !am.inBounds(j) {
		return ErrAdjMatrixOutOfBounds
	}

	am.matrix[i][j] = true
	am.matrix[j][i] = true
	return nil
}

func (am *AdjMatrix) Size() int {
	return len(am.Edges())
}

func (am *AdjMatrix) Order() int {
	return am.n
}

func (am *AdjMatrix) Vertices() []Vertex {
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

func (am *AdjMatrix) Edges() []Edge {
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
	index int
}

func (amv *adjMatrixVertex) Id() int {
	return amv.index
}

// panicUnexpected panics with a standard error message on errors that should
// never happen logically. This should only be called if there is a critical
// failure in plantri's code.
func panicUnexpected(fn string, err error) {
	panic(fmt.Errorf("plantri: unexpected error in %s: %v", fn, err))
}

// Compile-time interface implementation checks.
var _ Graph = new(AdjMatrix)
var _ Vertex = new(adjMatrixVertex)
