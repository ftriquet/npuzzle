package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	puzzleSize   uint
	solvable     bool
	iterations   uint
	inputFile    string
	random       bool
	search       string
	heuristic    string
	gblHeuristic Heuristic
)

func main() {
	setHeuristics()
	app := cli.NewApp()
	app.Usage = "Solve n-puzzle game"
	app.Name = "npuzzle"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		cli.Command{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "Generate puzzle",
			Action:  generate,
			Flags: []cli.Flag{
				cli.UintFlag{
					Name:        "size, s",
					Value:       3,
					Usage:       "Choose the `SIZE` of the puzzle",
					Destination: &puzzleSize,
				},
				cli.BoolTFlag{
					Name:        "solvable",
					Usage:       "Generates a solvable or unsolvable puzzle",
					Destination: &solvable,
				},
				cli.UintFlag{
					Name:        "iterations, i",
					Value:       10,
					Usage:       "Choose the number of `ITERATIONS` of the puzzle generation",
					Destination: &iterations,
				},
			},
		},
		cli.Command{
			Name:    "solve",
			Aliases: []string{"s"},
			Usage:   "Solve a puzzle",
			Action:  solve,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "file, f",
					Value:       "",
					Usage:       "The puzzle file to solve",
					Destination: &inputFile,
				},
				cli.BoolFlag{
					Name:        "random, r",
					Usage:       "Solve a random generated puzzle",
					Destination: &random,
				},
				cli.UintFlag{
					Name:        "size",
					Value:       3,
					Usage:       "Choose the `SIZE` of the puzzle",
					Destination: &puzzleSize,
				},
				cli.UintFlag{
					Name:        "iterations, i",
					Value:       10,
					Usage:       "Choose the number of `ITERATIONS` of the puzzle generation",
					Destination: &iterations,
				},
				cli.StringFlag{
					Name:        "search, s",
					Value:       "uniform",
					Usage:       "Type of search algorithm",
					Destination: &search,
				},
				cli.StringFlag{
					Name:        "heuristic",
					Value:       "",
					Usage:       "Heuristic used to find the solution",
					Destination: &heuristic,
				},
			},
		},
	}
	app.Run(os.Args)
}

func solve(cxt *cli.Context) error {
	var puzzle board
	var e error
	gblHeuristic, e = getHeuristic(heuristic)
	if e != nil {
		return cli.NewExitError(e.Error(), 1)
	}

	if search != "greedy" && search != "uniform" {
		return cli.NewExitError("Error: Invalid value for parameter search", 1)
	}
	if random && inputFile != "" {
		return cli.NewExitError("Error: Can't solve both random and file puzzle", 1)
	}
	if !random {
		var inputStream *os.File
		if inputFile == "" {
			inputStream = os.Stdin
		} else {
			f, e := os.Open(inputFile)
			if e != nil {
				return cli.NewExitError(fmt.Sprintf("Error: %v", e), 1)
			}
			inputStream = f
		}
		puzzle, e = parseBoard(inputStream)
		if e != nil {
			return cli.NewExitError(fmt.Sprintf("Error: %v", e), 1)
		}
	} else {
		puzzle = generateBoard(int(puzzleSize), true, int(iterations))
	}
	final := getFinalBoard(len(puzzle))
	if !puzzle.Solvable(final) {
		return cli.NewExitError("Unsolvable puzzle", 1)
	}
	solvePuzzle2(puzzle, final)
	return nil
}

func generate(cxt *cli.Context) error {
	b := generateBoard(int(puzzleSize), solvable, int(iterations))

	fmt.Println(puzzleSize)
	printBoard(b, os.Stdout)
	return nil
}
