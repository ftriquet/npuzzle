package main

import "math"

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Heuristic func(board, board) uint64

var heuristics map[string]Heuristic

func setHeuristics() {
	heuristics = map[string]Heuristic{
		"euclidian": euclidianDistance,
		"manhattan": manhattanDistance,
		"hamming":   hammingDistance,
		"conflict":  linearConflict,
	}
}

func euclidianDistance(final, current board) uint64 {
	sum := float64(0.0)
	for i := range current {
		for j, val := range current[i] {
			x, y := final.getPos(val)
			dx := float64(abs(x - i))
			dy := float64(abs(y - j))
			sum += math.Sqrt(dx*dx + dy*dy)
		}
	}
	return uint64(sum)
}

func manhattanDistance(final, current board) uint64 {
	var sum uint64
	for i := range current {
		for j, val := range current[i] {
			x, y := final.getPos(val)
			sum += uint64(abs(x-i) + abs(y-j))
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
