package main

type state struct {
	b         board
	index     int
	cost      uint64
	heuristic uint64
	ancestor  *state
}

func (s *state) getNextStates(finalBoard *board) []*state {
	sts := s.b.getNextStates()
	stop := make(chan *state)
	res := make([]*state, len(sts))
	for i := range res {
		go func(j int) {
			heuristick := uint64(0)
			switch heuristic {
			case "manhattan":
				heuristick = sts[j].getCost(manhattanDistance, *finalBoard)
			case "euclidian":
				heuristick = sts[j].getCost(euclidianDistance, *finalBoard)
			case "conflict":
				heuristick = sts[j].getCost(manhattanDistance, *finalBoard) + linearConflict(*finalBoard, sts[j])
			case "all":
				heuristick = sts[j].getCost(manhattanDistance, *finalBoard) +
					sts[j].getCost(hammingDistance, *finalBoard) +
					linearConflict(*finalBoard, sts[j])
			default:
				heuristick = sts[j].getCost(manhattanDistance, *finalBoard) +
					sts[j].getCost(hammingDistance, *finalBoard) +
					linearConflict(*finalBoard, sts[j])
			}
			var newCost uint64
			if search == "greedy" {
				newCost = heuristick
			} else {
				newCost = s.heuristic + 1 + heuristick
			}
			tmp := &state{
				sts[j],
				0,
				newCost,
				s.heuristic + 1,
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
