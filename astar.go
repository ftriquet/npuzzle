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

func contains2(open, close *queue, st *state) bool {
	if open.data[st.cost] != nil {
		for _, state := range open.data[st.cost] {
			if state.b.equals(st.b) {
				return true
			}
		}
	}
	if close.data[st.cost] != nil {
		for _, state := range close.data[st.cost] {
			if state.b.equals(st.b) {
				return true
			}
		}
	}
	return false
}

func solvePuzzle(b, final board) {
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
	fmt.Println("Initial:")
	printBoard(b, os.Stdout)
	fmt.Println("Final:")
	printBoard(final, os.Stdout)
	for len(open) > 0 {
		st := heap.Pop(&open).(*state)
		if final.equals(st.b) {
			fmt.Println("FINI")
			solution := 0
			for st != nil {
				solution++
				defer func(st *state) {
					printBoard(st.b, os.Stdout)
				}(st)
				st = st.ancestor
			}
			fmt.Printf("Size of solution: %d\n", solution)
			break
		} else {
			voisins := st.getNextStates(&final)
			for i := range voisins {
				if !contains(open, close, voisins[i]) {
					heap.Push(&open, voisins[i])
					open.Update(voisins[i])
					/* fmt.Println("State not in opened or closed") */
				} else {
					/* fmt.Println("State already in opened or closed") */
					/* printBoard(voisins[i].b, os.Stdout) */
				}
			}
			heap.Push(&close, st)
			close.Update(st)
		}
	}
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
	open.Push(initialState)
	fmt.Println("Initial:")
	printBoard(b, os.Stdout)
	fmt.Println("Final:")
	printBoard(final, os.Stdout)
	/* fmt.Println(open) */
	/* fmt.Println(open.costs) */
	for len(open.costs) > 0 {
		st := open.Pop()
		if final.equals(st.b) {
			fmt.Println("FINI")
			solution := 0
			for st != nil {
				solution++
				defer func(st *state) {
					printBoard(st.b, os.Stdout)
				}(st)
				st = st.ancestor
			}
			fmt.Printf("Size of solution: %d\n", solution)
			break
		} else {
			voisins := st.getNextStates(&final)
			for i := range voisins {
				if !contains2(open, close, voisins[i]) {
					open.Push(voisins[i])
					/* fmt.Println("State not in opened or closed") */
				} else {
					/* fmt.Println("State already in opened or closed") */
					/* printBoard(voisins[i].b, os.Stdout) */
				}
			}
			close.Push(st)
		}
	}
}
