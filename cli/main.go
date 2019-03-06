package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	plantri "github.com/bubblyworld/plantri-go"
	graphutil "github.com/bubblyworld/plantri-go/graphutil"
)

var input = flag.String("input", "", "Plantri inut file")
var testEccentricity = flag.Bool("test_eccentricity_conjecture", false,
	"Test the connected eccentricity level set conjecture for each input graph.")

func main() {
	flag.Parse()

	if err := validateFlags(); err != nil {
		fmt.Printf("Flag error: %v\n", err)
		printHelpPrompt()
		os.Exit(1)
	}

	graphs, err := plantri.Load(*input)
	if err != nil {
		fmt.Printf("Error loading input: %v\n", err)
		os.Exit(1)
	}

	if *testEccentricity {
		testEccentricityFn(graphs)
	}
}

func validateFlags() error {
	if *input == "" {
		return errors.New("input file required")
	}

	return nil
}

func printHelpPrompt() {
	fmt.Print("To see the available flags and get help, use -h.\n")
	fmt.Print("Depending on the problem, you may need to call a doctor too.\n")
}

func testEccentricityFn(graphs []plantri.Graph) {
	fmt.Printf("Testing the connected eccentricity level set conjecture for all graphs in %s...\n",
		*input)

	// TODO(guy): For full robustness, we should do a cheap maximal
	//            planarity check. Currently the user has to be sure they are
	//            testing the right kinds of graph.

	for i, graph := range graphs {
		// We calculate the eccentricities of each vertex in the graph, and
		// then subdivide the graph into maximal connected subsets with the
		// same eccentricity.
		eccs := graphutil.Eccentricities(graph)
		comps := graphutil.ConnectedMonochromeSubsets(graph, eccs)

		// If the conjecture is true, there should not be any disconnected
		// subsets with the same eccentricity.
		conjHolds := true
		eccsSet := make(map[int]bool)
		for _, comp := range comps {
			if len(comp) == 0 {
				continue
			}

			// Every vertex in the component will have the same eccentricity,
			// as this is a postcondition of ConnectedMonochromeSubsets().
			ecc := eccs[comp[0].Id()]
			if eccsSet[ecc] {
				conjHolds = false
			}
			eccsSet[ecc] = true
		}

		fmt.Printf("  conjecture holds for graph %2d: %v\n", i, conjHolds)
		if !conjHolds {
			fmt.Printf("%s", graph)
		}
	}
}
