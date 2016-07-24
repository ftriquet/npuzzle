package main

type board [][]int

func (b board) dup() board {
	size := len(b)
	r := make(board, size)
	for i := range r {
		r[i] = make([]int, size)
		for j := range b[i] {
			r[i][j] = b[i][j]
		}
	}
	return r
}

func (b board) size() int {
	return len(b)
}

func (b board) equals(b2 board) bool {
	/* return Distances(b, b2) == 0 */
	for i := range b {
		for j := range b[i] {
			if b[i][j] != b2[i][j] {
				return false
			}
		}
	}
	return true
}

func (b board) getCost(h Heuristic, final board) uint64 {
	return h(b, final)
}

func (b board) toArray() []int {
	var r []int
	for i := range b {
		r = append(r, b[i]...)
	}
	return r
}

func boardFromArray(b []int, size int) board {
	board := make(board, size)
	for i := range board {
		board[i] = b[size*i : size*(i+1)]
	}
	return board
}

func (b board) nbInversions() int {
	t := b.toArray()
	count := 0
	for i, v := range t {
		for _, w := range t[i+1:] {
			if v > w {
				count++
			}
		}
	}
	return count
}
