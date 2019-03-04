package graphutil

import plantri "github.com/bubblyworld/plantri-go"

// LabelsToIndex returns a map of vertex labels to their index in the list
// returned by graph.Vertices(). This map is valid until the graph is modified.
func LabelsToIndex(graph plantri.Graph) map[int]int {
	res := make(map[int]int)
	for i, v := range graph.Vertices() {
		res[v.Label()] = i
	}

	return res
}

// IndexToLabels returns a map of vertex labels indexed by their index in the
// list returned by graph.Vertices(). This map is valid until the graph is
// modified.
func IndexToLabels(graph plantri.Graph) map[int]int {
	res := make(map[int]int)
	for i, v := range graph.Vertices() {
		res[i] = v.Label()
	}

	return res
}
