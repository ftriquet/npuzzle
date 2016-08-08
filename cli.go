package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	puzzleSize uint
	unsolvable bool
	iterations uint
	inputFile  string
	random     bool
	search     string
	heuristic  string
)

func main() {
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
				cli.BoolFlag{
					Name:        "unsolvable, u",
					Usage:       "Generates a unsolvable puzzle",
					Destination: &unsolvable,
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
					Value:       "distance",
					Usage:       "Heuristic used to find the solution",
					Destination: &search,
				},
			},
		},
	}
	app.Run(os.Args)
}

func solve(cxt *cli.Context) error {
	var puzzle board
	var e error
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
		puzzle = generateBoard(int(puzzleSize), false, int(iterations))
	}
	final := getFinalBoard(len(puzzle))
	invB := puzzle.nbInversions()
	invF := final.nbInversions()
	if len(puzzle)%2 == 0 {
		x, _ := puzzle.getPos(0)
		invB += x
		invF += x
	}
	if !(invB%2 == invF%2) {
		return cli.NewExitError("Unsolvable puzzle", 1)
	}
	solvePuzzle2(puzzle, final)
	return nil
}

func generate(cxt *cli.Context) error {
	b := generateBoard(int(puzzleSize), !unsolvable, int(iterations))
	printBoard(b, os.Stdout)
	return nil
}
