package main

import (
	"fmt"
	"strings"
)

type state struct {
	b         board
	index     int
	cost      uint64
	heuristic uint64
	ancestor  *state
}

func getHeuristic(heuristicsArg string) (func(board, board) uint64, error) {
	split := strings.Split(heuristicsArg, "+")
	var activeHeuristics []Heuristic

	if len(split) == 1 && split[0] == "" {
		return func(b, final board) uint64 {
			return manhattanDistance(b, final) +
				hammingDistance(b, final) +
				linearConflict(b, final)
		}, nil
	}

	for _, h := range split {
		if heuristics[h] == nil {
			return nil, fmt.Errorf("Invalid heuristic parameter: '%s'", h)
		}
		activeHeuristics = append(activeHeuristics, heuristics[h])
	}

	return func(b, final board) uint64 {
		var res uint64
		for _, h := range activeHeuristics {
			res += h(b, final)
		}
		return res
	}, nil
}

func (s *state) getNextStates(finalBoard *board) []*state {
	sts := s.b.getNextStates()
	stop := make(chan *state)
	res := make([]*state, len(sts))
	for i := range res {
		go func(j int) {
			heuristick := uint64(0)
			// switch heuristic {
			// case "manhattan":
			// 	heuristick = sts[j].getCost(manhattanDistance, *finalBoard)
			// case "euclidian":
			// 	heuristick = sts[j].getCost(euclidianDistance, *finalBoard)
			// case "conflict":
			// 	heuristick = sts[j].getCost(manhattanDistance, *finalBoard) + linearConflict(*finalBoard, sts[j])
			// case "all":
			// 	heuristick = sts[j].getCost(manhattanDistance, *finalBoard) +
			// 		sts[j].getCost(hammingDistance, *finalBoard) +
			// 		linearConflict(*finalBoard, sts[j])
			// default:
			// 	heuristick = sts[j].getCost(manhattanDistance, *finalBoard) +
			// 		sts[j].getCost(hammingDistance, *finalBoard) +
			// 		linearConflict(*finalBoard, sts[j])
			// }
			heuristick = gblHeuristic(sts[j], *finalBoard)
			var newCost uint64
			if search == "greedy" {
				newCost = heuristick
			} else {
				newCost = s.cost + 1 + heuristick
			}
			tmp := &state{
				sts[j],
				0,
				s.cost + 1,
				newCost,
				s,
			}
			stop <- tmp
		}(i)
	}
	i := 0
	for i < len(sts) {
		res[i] = <-stop
		i++
	}
	close(stop)
	return res
}
