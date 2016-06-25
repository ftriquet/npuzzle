package main

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Heuristic func(board, board) uint64

func Differences(final, current board) uint64 {
	var sum uint64
	for i := range final {
		for j := range final[i] {
			sum += uint64(abs(final[i][j] - current[i][j]))
		}
	}
	return sum
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

func Distances(final, current board) uint64 {
	var sum uint64
	for i := range current {
		for j := range current[i] {
			sum += distance(current, final, i, j)
		}
	}
	return sum
}
