package main

import (
	"fmt"
	"os"
)

func (b board) getPos(val int) (int, int) {
	for i := range b {
		for j := range b[i] {
			if val == b[i][j] {
				return i, j
			}
		}
	}
	return -1, -1
}

func (b board) getNextStates() []board {
	i, j := b.getPos(0)
	var states []board
	if i != 0 {
		states = append(states, b.getPermutation(i, j, i-1, j))
	}
	if i != b.size()-1 {
		states = append(states, b.getPermutation(i, j, i+1, j))
	}
	if j != 0 {
		states = append(states, b.getPermutation(i, j, i, j-1))
	}
	if j != b.size()-1 {
		states = append(states, b.getPermutation(i, j, i, j+1))
	}
	return states
}

func (b board) getPermutation(i, j, x, y int) board {
	newBoard := b.dup()
	newBoard[i][j], newBoard[x][y] = newBoard[x][y], newBoard[i][j]
	return newBoard
}

func contains2(open, close *queue, st *state) bool {
	for _, states := range open.data {
		for _, state := range states {
			if state.heuristic <= st.heuristic && state.b.equals(st.b) {
				return true
			}
		}
	}
	for _, states := range close.data {
		for _, state := range states {
			if state.heuristic <= st.heuristic && state.b.equals(st.b) {
				return true
			}
		}
	}
	return false
}

func solvePuzzle2(b, final board) {
	open := &queue{
		make(map[uint64][]*state),
		nil,
	}
	close := &queue{
		make(map[uint64][]*state),
		nil,
	}
	initialState := &state{
		b,
		0,
		0,
		0,
		nil,
	}
	if final.equals(initialState.b) {
		fmt.Println("The puzzle is already solved")
		return
	}
	nStates := 1
	statesInOpen := 0
	open.Push(initialState)
	for len(open.costs) > 0 {
		tmp := 0
		for k := range open.data {
			tmp += len(open.data[k])
		}
		if tmp > nStates {
			nStates = tmp
		}
		st := open.Pop()
		if final.equals(st.b) {
			solutions := 0
			for s := st; s != nil; s = s.ancestor {
				solutions++
			}
			defer func(n, p, s int) {
				fmt.Printf("Size of solution: %d\n", s)
				fmt.Printf("Total number of states visited: %d\n", n)
				fmt.Printf("Max number of states at the same time in open set: %d\n", p)
			}(statesInOpen, nStates, solutions)
			for st != nil {
				defer func(st *state) {
					printBoard(st.b, os.Stdout)
				}(st)
				st = st.ancestor
			}
			break
		} else {
			voisins := st.getNextStates(&final)
			for i := range voisins {
				if !contains2(open, close, voisins[i]) {
					statesInOpen++
					open.Push(voisins[i])
				}
			}
			close.Push(st)
		}
	}
}
