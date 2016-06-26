package main

import (
	"container/heap"
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

func contains(open, closed PriorityQueue, st *state) bool {
	for i := range closed {
		if st.heuristic >= closed[i].heuristic && closed[i].b.equals(st.b) {
			return true
		}
	}
	for i := range open {
		if st.heuristic >= open[i].heuristic && open[i].b.equals(st.b) {
			return true
		}
	}
	return false
}

func solve(b, final board) {
	open := make(PriorityQueue, 0)
	close := make(PriorityQueue, 0)
	heap.Init(&open)
	heap.Init(&close)
	initialState := &state{
		b,
		0,
		0,
		0,
		nil,
	}
	heap.Push(&open, initialState)
	for len(open) > 0 {
		st := heap.Pop(&open).(*state)
		if final.equals(st.b) {
			fmt.Println("==============")
			fmt.Println("FINI")
			fmt.Println("==============")
			for st != nil {
				printBoard(st.b, os.Stdout)
				st = st.ancestor
			}
			break
		} else if !contains(open, close, st) {
			voisins := st.getNextStates(final)
			for i := range voisins {
				heap.Push(&open, voisins[i])
				open.Update(voisins[i])
			}
			heap.Push(&close, st)
			close.Update(st)
		}
	}
}
