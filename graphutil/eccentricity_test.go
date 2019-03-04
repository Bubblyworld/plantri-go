package graphutil

import (
	"log"
	"testing"

	plantri "github.com/bubblyworld/plantri-go"
)

func TestEccentricities(t *testing.T) {
	testCases := []struct {
		graph plantri.Graph
	}{}

	for _, tc := range testCases {
		eccs := Eccentricities(tc.graph)
		log.Print(eccs)
	}
}
