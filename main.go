package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	search    string
	heuristic string
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
	final := getFinalBoard(3)
	solve(b, final)
}
