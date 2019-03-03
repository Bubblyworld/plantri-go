package plantri

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHeader(t *testing.T) {
	testCases := []struct {
		name           string
		path           string
		expectedError  error
		expectedHeader string
	}{
		{
			name:          "invalid header - eof",
			path:          "testdata/invalid_header_1",
			expectedError: ErrNoHeader,
		},
		{
			name:          "invalid header - beginning",
			path:          "testdata/invalid_header_2",
			expectedError: ErrNoHeader,
		},
		{
			name:           "planar code",
			path:           "testdata/degree_4.planar",
			expectedHeader: headerPlanarCode,
		},
	}

	for _, tc := range testCases {
		data, err := ioutil.ReadFile(tc.path)
		if !assert.NoError(t, err, tc.name) {
			continue
		}

		header, err := getHeader(data)
		assert.Equal(t, tc.expectedError, err, tc.name)
		assert.Equal(t, tc.expectedHeader, header, tc.name)
	}
}

func TestLoadNumber(t *testing.T) {
	// Number of maximal planar graphs indexed by order. We will use this as a
	// smoke test to ensure plantri is loading the correct number of graphs.
	//   see http://oeis.org/A000109.
	numberMpgs := []int{0, 0, 0, 1, 1, 1, 2, 5, 14, 50, 233}

	// A classical result is that a maximal planar graph of order n has exactly
	// 3n - 6 edges in all cases.
	numberEdges := []int{0, 0, 0, 3, 6, 9, 12, 15, 18, 21, 24}

	for i := 4; i <= 10; i++ {
		name := fmt.Sprintf("order_%d", i)

		graphs, err := Load(fmt.Sprintf("testdata/degree_%d.planar", i))
		if !assert.NoError(t, err, name) {
			continue
		}

		assert.Equal(t, numberMpgs[i], len(graphs), name)
		for _, graph := range graphs {
			assert.Equal(t, i, len(graph.Vertices()), name)
			assert.Equal(t, numberEdges[i], len(graph.Edges()), name)
		}
	}
}
