package graphutil

import (
	"testing"

	plantri "github.com/bubblyworld/plantri-go"
	"github.com/stretchr/testify/assert"
)

func TestEccentricities(t *testing.T) {
	testCases := []struct {
		name  string
		graph plantri.Graph

		// Expected eccentricities indexed by vertex id.
		expectedEccs map[int]int
	}{
		{
			name:  "trivial graph",
			graph: Path(1),
			expectedEccs: map[int]int{
				0: 0,
			},
		},
		{
			name:  "path of order 2",
			graph: Path(2),
			expectedEccs: map[int]int{
				0: 1,
				1: 1,
			},
		},
		{
			name:  "path of order 3",
			graph: Path(3),
			expectedEccs: map[int]int{
				0: 2,
				1: 1,
				2: 2,
			},
		},
		{
			name:  "path of order 4",
			graph: Path(4),
			expectedEccs: map[int]int{
				0: 3,
				1: 2,
				2: 2,
				3: 3,
			},
		},
		{
			name:  "cycle of order 3",
			graph: Cycle(3),
			expectedEccs: map[int]int{
				0: 1,
				1: 1,
				2: 1,
			},
		},
		{
			name:  "cycle of order 4",
			graph: Cycle(4),
			expectedEccs: map[int]int{
				0: 2,
				1: 2,
				2: 2,
				3: 2,
			},
		},
		{
			name:  "complete of order 5",
			graph: Complete(5),
			expectedEccs: map[int]int{
				0: 1,
				1: 1,
				2: 1,
				3: 1,
				4: 1,
			},
		},
	}

	for _, tc := range testCases {
		eccs := Eccentricities(tc.graph)
		assert.Equal(t, tc.expectedEccs, eccs, tc.name)
	}
}
