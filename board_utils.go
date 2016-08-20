package main

import (
	"fmt"
	"io"
)

func newBoard(size int) board {
	b := make(board, size)
	for i := 0; i < size; i++ {
		b[i] = make([]int, size)
	}
	return b
}

func getFinalBoard(size int) board {
	b := newBoard(size)
	fillBoard(b, size, 1)
	return b
}

func fillBoard(b board, size, first int) {
	if size == 1 {
		b[0][0] = 0
		return
	}
	if size == 2 {
		b[0][0] = first
		b[0][1] = first + 1
		b[1][0] = 0
		b[1][1] = first + 2
		return
	}
	count, i, j := 0, 0, 0
	for ; j < size; j++ {
		b[i][j] = first + count
		count++
	}
	j = size - 1
	i = 1
	for ; i < size; i++ {
		b[i][j] = first + count
		count++
	}
	i = size - 1
	j = size - 2
	for ; j >= 0; j-- {
		b[i][j] = first + count
		count++
	}
	j = 0
	i = size - 2
	for ; i > 0; i-- {
		b[i][j] = first + count
		count++
	}
	sub, _ := getSubBoard(b)
	fillBoard(sub, size-2, first+getBandSize(size))
}

func getBandSize(boardSize int) int {
	return 4*boardSize - 4
}

func printBoard(brd board, out io.Writer) {
	for i := range brd {
		for j := range brd[i] {
			fmt.Fprintf(out, "%-2d ", brd[i][j])
		}
		fmt.Fprintf(out, "\n")
	}
	fmt.Fprintf(out, "\n")
}

func getSubBoard(brd board) (res board, finish bool) {
	if len(brd) <= 2 {
		res, finish = brd, true
	} else {
		finish = false
		res = newBoard(len(brd) - 2)
		subLines := brd[1 : len(brd)-1]
		for i := range subLines {
			res[i] = subLines[i][1 : len(subLines[i])-1]
		}
	}
	return
}
