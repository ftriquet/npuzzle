package main

import "os"

func main() {
	/*
		b := board{
			{1, 4, 2},
			{0, 8, 3},
			{6, 7, 5},
		}
	*/
	b, _ := parseBoard(os.Stdin)
	final := getFinalBoard(3)
	solve(b, final)

	/*
		tmp := board{
			{1, 0, 3},
			{8, 2, 5},
			{7, 4, 6},
		}
		for _, v := range tmp.getNextStates() {
			printBoard(v, os.Stdout)
		}
	*/
}
