package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	search    string
	heuristic string
	nstates   int
)

func checkFlags(heuristics []string) {
	flag := false
	for i := range heuristics {
		if heuristics[i] == heuristic {
			flag = true
		}
	}
	if !flag {
		os.Exit(1)
	}
	flag = false
	if search != "uniform" && search != "greedy" {
		os.Exit(1)
	}
}

func main() {
	heuristicValues := []string{
		"distance",
		"difference",
	}
	flag.StringVar(&search, "search", "uniform", "A* search type. Use greedy for fast solution finding and uniform (default) for the best solution")
	flag.StringVar(&heuristic, "heuristic", "distance", "Heuristic used to find the solution")
	flag.Parse()
	checkFlags(heuristicValues)
	b, e := parseBoard(os.Stdin)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", e)
		os.Exit(1)
	}
	final := getFinalBoard(len(b))
	invB := b.nbInversions()
	invF := final.nbInversions()
	if len(b)%2 == 0 {
		x, _ := b.getPos(0)
		invB += x
		invF += x
	}
	if !(invB%2 == invF%2) {
		fmt.Println("Puzzle non solvable")
		os.Exit(1)
	}
	solve(b, final)
}
