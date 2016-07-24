package main

import "math/rand"

func getIndex(num int, tab []int) int {
	for i, v := range tab {
		if num == v {
			return i
		}
	}
	return -1
}

func swapEmpty(board []int, size int) {
	// swap empty piece with another
	index := getIndex(0, board)
	poss := []int{}

	if index%size > 0 {
		poss = append(poss, index-1)
	}
	if index%size < size-1 {
		poss = append(poss, index+1)
	}
	if index/size > 0 {
		poss = append(poss, index-size)
	}
	if index/size < size-1 {
		poss = append(poss, index+size)
	}
	i := rand.Intn(len(poss))
	board[index] = board[i]
	board[i] = 0
}

func generateBoard(size int, solvable bool, iterations int) board {
	b := getFinalBoard(size).toArray()

	for i := 0; i < iterations; i++ {
		swapEmpty(b, size)
	}
	if !solvable {
		// swap two normal pieces
		if b[0] == 0 || b[1] == 0 {
			b[size-1], b[size-2] = b[size-2], b[size-1]
		} else {
			b[0], b[1] = b[1], b[0]
		}
	}
	return boardFromArray(b, size)
}