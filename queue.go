package main

type queue struct {
	data  map[uint64][]*state
	costs []uint64
	size  int
}

func (q *queue) Push(s *state) {
	q.size++
	if q.data[s.heuristic] != nil {
		q.data[s.heuristic] = append(q.data[s.heuristic], s)
		return
	}
	q.data[s.heuristic] = append(q.data[s.heuristic], s)
	q.costs = append(q.costs, s.heuristic)
}

func (q *queue) Pop() *state {
	q.size--
	mincost := q.costs[0]
	costindex := 0
	for i, v := range q.costs {
		if v < mincost {
			mincost = v
			costindex = i
		}
	}
	tmp := q.data[mincost][0]
	if len(q.data[mincost]) == 1 {
		delete(q.data, mincost)
		q.costs = append(q.costs[0:costindex], q.costs[costindex+1:]...)
	} else {
		q.data[mincost] = q.data[mincost][1:]
	}
	return tmp
}
