package main

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Heuristic func(board, board) uint64

func difference(final, current board) uint64 {
	var sum uint64
	for i := range final {
		for j := range final[i] {
			sum += uint64(abs(final[i][j] - current[i][j]))
		}
	}
	return sum
}

func euclidianDistance(final, current board) uint64 {
	res := uint64(0)
	for i := range current {
		for j := range current[i] {

		}
	}
}

func distance(current, final board, x, y int) uint64 {
	for i := range final {
		for j := range final[i] {
			if final[i][j] == current[x][y] {
				return uint64(abs(x-i) + abs(y-j))
			}
		}
	}
	return 0
}

func manhattanDistance(final, current board) uint64 {
	var sum uint64
	for i := range current {
		for j := range current[i] {
			sum += distance(current, final, i, j)
		}
	}
	return sum
}

func hammingDistance(final, current board) uint64 {
	res := uint64(0)
	for i := range current {
		for j, val := range current[i] {
			x, y := final.getPos(val)
			if x != i || y != j {
				res++
			}
		}
	}
	return res
}

func conflict(b board, line, i, k int) bool {
	ivalue := b[line][i]
	kvalue := b[line][k]
	ix, iy := b.getPos(ivalue)
	kx, ky := b.getPos(kvalue)
	if ix == kx && iy > ky {
		return true
	}
	return false
}

func linearConflict(final, current board) uint64 {
	res := uint64(0)
	for j := 0; j < len(current); j++ {
		for i := 0; i < len(current)-1; i++ {
			for k := i + 1; k < len(current); k++ {
				if conflict(current, j, i, k) {
					res += 2
				}
			}
		}
	}
	return res
}
