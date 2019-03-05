package graphutil

import (
	plantri "github.com/bubblyworld/plantri-go"
)

const (
	maxuint = ^uint(0)
	maxint  = int(maxuint >> 2)
)

// Eccentricities returns a map of vertex eccentricity indexed by vertex id.
// Vertex eccentricities are calculated using the Floyd-Warshall algorithm
// followed by a column scan to find the maximum distance between a given
// vertex and any other in the graph.
func Eccentricities(graph plantri.Graph) map[int]int {
	var dists [][]int
	for i := 0; i < graph.Order(); i++ {
		dists = append(dists, []int{})

		for j := 0; j < graph.Order(); j++ {
			dists[i] = append(dists[i], maxint)
		}
	}

	// Initial distances.
	idsToIndex := IdsToIndex(graph)
	for _, v := range graph.Vertices() {
		i := idsToIndex[v.Id()]
		dists[i][i] = 0
	}
	for _, e := range graph.Edges() {
		i1 := idsToIndex[e.Source.Id()]
		i2 := idsToIndex[e.Dest.Id()]

		dists[i1][i2] = 1
		dists[i2][i1] = 1
	}

	// Floyd-Warshall algorithm - provides minimum distance between all
	// pairs of vertices of the grpah in O(n^3) time.
	for k := 0; k < graph.Order(); k++ {
		for i := 0; i < graph.Order(); i++ {
			for j := 0; j < graph.Order(); j++ {
				if dists[i][k]+dists[k][j] < dists[i][j] {
					dists[i][j] = dists[i][k] + dists[k][j]
				}
			}
		}
	}

	res := make(map[int]int)
	indexToIds := IndexToIds(graph)
	for i := 0; i < graph.Order(); i++ {
		ecc := -maxint
		for j := 0; j < graph.Order(); j++ {
			if dists[i][j] > ecc {
				ecc = dists[i][j]
			}
		}

		res[indexToIds[i]] = ecc
	}

	return res
}
