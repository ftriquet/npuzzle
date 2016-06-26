package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func getLineFromString(line string, size, l int) ([]int, error) {
	comment := strings.Split(line, "#")
	line = comment[0]
	reg := regexp.MustCompile("[ \t]+")
	words := reg.Split(line, -1)
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
	hasSize := false
	numberRegex := regexp.MustCompile("[0-9]+")
	var size int
	i := 0
	var lines []struct {
		line       string
		lineNumber int
	}
	for scanner.Scan() {
		i++
		tmp := strings.Trim(scanner.Text(), " \t")
		if tmp[0] != '#' {
			tmp = strings.Split(tmp, "#")[0]
			if !hasSize {
				if numberRegex.MatchString(tmp) {
					hasSize = true
					s, err := strconv.ParseInt(tmp, 10, 0)
					if err != nil {
						return nil, fmt.Errorf("Invalid line number: %s", tmp)
					}
					size = int(s)
				} else {
					return nil, fmt.Errorf("Size must be provided at begining of file")
				}
			} else {
				lines = append(lines, struct {
					line       string
					lineNumber int
				}{
					tmp,
					i,
				})
			}
		}
	}
	if size != len(lines) {
		return nil, fmt.Errorf("Wrong numbers of lines (expected: %d, found: %d", size, len(lines))
	}
	b = make(board, size)
	for i := range lines {
		b[i], e = getLineFromString(lines[i].line, size, lines[i].lineNumber)
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
