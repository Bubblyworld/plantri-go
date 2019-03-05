package graphutil

import (
	"fmt"

	plantri "github.com/bubblyworld/plantri-go"
)

// IdsToIndex returns a map of vertex ids to their index in the list returned
// by graph.Vertices(). This map is valid until the graph is modified.
func IdsToIndex(graph plantri.Graph) map[int]int {
	res := make(map[int]int)
	for i, v := range graph.Vertices() {
		res[v.Id()] = i
	}

	return res
}

// IndexToIds returns a map of vertex ids indexed by their index in the list
// returned by graph.Vertices(). This map is valid until the graph is modified.
func IndexToIds(graph plantri.Graph) map[int]int {
	res := make(map[int]int)
	for i, v := range graph.Vertices() {
		res[i] = v.Id()
	}

	return res
}

// Neighbours returns a list of ids of neighbouring vertices of the vertex with
// the given id. If the vertex doesn't exist, an empty list is returned.
func Neighbours(graph plantri.Graph, id int) []int {
	idSet := make(map[int]bool)
	for _, e := range graph.Edges() {
		if e.Source.Id() == e.Dest.Id() {
			continue
		}

		if e.Source.Id() == id {
			idSet[e.Dest.Id()] = true
		}

		if e.Dest.Id() == id {
			idSet[e.Source.Id()] = true
		}
	}

	var res []int
	for nid := range idSet {
		res = append(res, nid)
	}

	return res
}

// panicUnexpected panics with a standard error message on errors that should
// never happen logically. This should only be called if there is a critical
// failure in the code.
func panicUnexpected(fn string, err error) {
	panic(fmt.Errorf("plantri/graphutil: unexpected error in %s: %v", fn, err))
}
