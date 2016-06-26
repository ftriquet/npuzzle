package main

type queue map[uint64][]*state

type state struct {
	b         board
	index     int
	cost      uint64
	heuristic uint64
	ancestor  *state
}

func (s *state) getNextStates(finalBoard board) []*state {
	sts := s.b.getNextStates()
	stop := make(chan *state)
	res := make([]*state, len(sts))
	for i := range res {
		go func(j int) {
			cost := sts[j].getCost(Distances, finalBoard)
			tmp := &state{
				sts[j],
				0,
				cost,
				s.heuristic + cost,
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
