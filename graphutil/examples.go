package graphutil

import plantri "github.com/bubblyworld/plantri-go"

// Path returns the path graph on n vertices.
func Path(n int) plantri.Graph {
	res := plantri.NewAdjMatrix(n)
	for i := 0; i < n-1; i++ {
		err := res.AddEdge(i, i+1)
		if err != nil {
			panicUnexpected("Path", err)
		}
	}

	return res
}

// Cycle returns the cycle graph on n vertices.
func Cycle(n int) plantri.Graph {
	res := plantri.NewAdjMatrix(n)
	for i := 0; i < n; i++ {
		err := res.AddEdge(i, (i+1)%n)
		if err != nil {
			panicUnexpected("Cycle", err)
		}
	}

	return res
}

// Complete returns the complete graph on n vertices.
func Complete(n int) plantri.Graph {
	res := plantri.NewAdjMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			err := res.AddEdge(i, j)
			if err != nil {
				panicUnexpected("Complete", err)
			}
		}
	}

	return res
}
