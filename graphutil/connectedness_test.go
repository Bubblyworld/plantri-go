package graphutil

import (
	"sort"
	"testing"

	plantri "github.com/bubblyworld/plantri-go"
	"github.com/stretchr/testify/assert"
)

func TestConnectedMonochromeSubset(t *testing.T) {
	testCases := []struct {
		name   string
		graph  plantri.Graph
		colors map[int]int

		// Expected monochrome connected subsets by vertex id.
		expectedSubsets [][]int
	}{
		{
			name:  "Path of order 5 with 1 subsets",
			graph: Path(5),
			colors: map[int]int{
				0: 1,
				1: 1,
				2: 1,
				3: 1,
				4: 1,
			},
			expectedSubsets: [][]int{
				{0, 1, 2, 3, 4},
			},
		},
		{
			name:  "Path of order 5 with 4 subsets",
			graph: Path(5),
			colors: map[int]int{
				0: 1,
				1: 2,
				2: 2,
				3: 3,
				4: 2,
			},
			expectedSubsets: [][]int{
				{0},
				{1, 2},
				{3},
				{4},
			},
		},
		{
			name:  "Path of order 5 with 5 subsets",
			graph: Path(5),
			colors: map[int]int{
				0: 1,
				1: 2,
				2: 3,
				3: 4,
				4: 5,
			},
			expectedSubsets: [][]int{
				{0},
				{1},
				{2},
				{3},
				{4},
			},
		},
	}

	for _, tc := range testCases {
		subsetsV := ConnectedMonochromeSubsets(tc.graph, tc.colors)

		var subsets [][]int
		for _, sv := range subsetsV {
			var subset []int
			for _, v := range sv {
				subset = append(subset, v.Id())
			}

			sort.Ints(subset) // for deterministic tests
			subsets = append(subsets, subset)
		}

		assert.ElementsMatch(t, tc.expectedSubsets, subsets, tc.name)
	}
}
