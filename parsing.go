package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func getLineFromString(line string, size, l int) ([]int, error) {
	words := strings.Split(line, " ")
	lineSize := len(words)
	if lineSize != size {
		return nil, fmt.Errorf("Line %d is invalid (size should be %d, %d found)", l, size, lineSize)
	}
	res := make([]int, lineSize)
	for i := range res {
		tmp, err := strconv.ParseInt(words[i], 10, 0)
		if err != nil {
			return nil, fmt.Errorf("Invalid syntax at line %d: %s", l, words[i])
		}
		res[i] = int(tmp)
	}
	return res, nil
}

func parseBoard(in io.Reader) (b board, e error) {
	scanner := bufio.NewScanner(in)
	e = nil
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	size := len(lines)
	b = make(board, size)
	for i := range lines {
		b[i], e = getLineFromString(lines[i], size, i+1)
		if e != nil {
			return nil, e
		}
	}
	if !b.isValid() {
		return nil, fmt.Errorf("Invalid game board")
	}
	return
}

func (b board) isValid() bool {
	max := 0
	min := 0
	visited := make(map[int]struct{})
	for i := range b {
		for j := range b[i] {
			if b[i][j] > max {
				max = b[i][j]
			}
			if b[i][j] < min {
				min = b[i][j]
			}
			visited[b[i][j]] = struct{}{}
		}
	}
	if len(visited) != b.size()*b.size() {
		return false
	}
	if min != 0 {
		return false
	}
	if max != b.size()*b.size()-1 {
		return false
	}
	return true
}
