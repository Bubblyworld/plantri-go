package graphutil

import (
	plantri "github.com/bubblyworld/plantri-go"
)

// ConnectedMonochromeSubsets takes a coloured graph (a graph along with an
// map representing the "color" of each vertex, indexed by vertex id) and
// decomposes it into maximal connected subsets of vertices of the same color.
// TODO(guy): "Components" is the terminology here.
func ConnectedMonochromeSubsets(graph plantri.Graph,
	colors map[int]int) [][]plantri.Vertex {

	visitedIds := make(map[int]bool)
	idToVertex := make(map[int]plantri.Vertex)
	for _, v := range graph.Vertices() {
		idToVertex[v.Id()] = v
	}

	var res [][]plantri.Vertex
	for id := range idToVertex {
		subset := floodfill(graph, colors, visitedIds, colors[id], id)
		if len(subset) == 0 {
			continue
		}

		var subsetV []plantri.Vertex
		for _, sid := range subset {
			subsetV = append(subsetV, idToVertex[sid])
		}

		res = append(res, subsetV)
	}

	return res
}

func floodfill(graph plantri.Graph, colors map[int]int,
	visitedIds map[int]bool, color int, id int) []int {

	if visitedIds[id] {
		return nil
	}

	if color != colors[id] {
		return nil
	}

	res := []int{id}
	visitedIds[id] = true
	for _, nid := range Neighbours(graph, id) {
		res = append(res, floodfill(graph, colors, visitedIds, color, nid)...)
	}

	return res
}
