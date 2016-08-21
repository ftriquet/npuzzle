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
	c := make(chan bool, 2)
	go func(s *state, op *queue, ch chan bool) {
		for cost, states := range op.data {
			if cost < s.heuristic {
				for _, state := range states {
					if state.b.equals(s.b) {
						ch <- true
						return
					}
				}
			}
		}
		ch <- false
	}(st, open, c)

	go func(s *state, cl *queue, ch chan bool) {
		for _, states := range cl.data {
			for _, state := range states {
				if state.b.equals(s.b) {
					ch <- true
					return
				}
			}
		}
		ch <- false
	}(st, close, c)
	res := <-c
	if res {
		return res
	}
	return <-c
}

func solvePuzzle2(b, final board) {
	open := &queue{
		make(map[uint64][]*state),
		nil,
		0,
	}
	close := &queue{
		make(map[uint64][]*state),
		nil,
		0,
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
	for open.size > 0 {
		if tmp := open.size + close.size; tmp > nStates {
			nStates = tmp
		}
		st := open.Pop()
		//fmt.Printf("Cost: %d, Heuristic: %d\n", st.cost, st.heuristic)
		if final.equals(st.b) {
			solutions := 0
			for s := st; s != nil; s = s.ancestor {
				solutions++
			}
			defer func(n, p, s int) {
				fmt.Printf("Size of solution: %d\n", s)
				fmt.Printf("Total number of states visited: %d\n", n)
				fmt.Printf("Max number of states at the same time in memory: %d\n", p)
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
