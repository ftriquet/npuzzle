package main

import "container/heap"

type PriorityQueue []*state

func (p PriorityQueue) Len() int { return len(p) }

func (p PriorityQueue) Less(i, j int) bool {
	if search == "greedy" {
		return p[i].cost < p[j].cost
	}
	return p[i].heuristic < p[j].heuristic
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *PriorityQueue) Push(x interface{}) {
	n := len(*p)
	node := x.(*state)
	node.index = n
	*p = append(*p, node)
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	item.index = -1
	*p = old[0 : n-1]
	return item
}

func (p *PriorityQueue) Update(node *state) {
	heap.Fix(p, node.index)
}

type queue struct {
	data  map[uint64][]*state
	costs []uint64
}

func (q *queue) Push(s *state) {
	if q.data[s.cost] != nil {
		q.data[s.cost] = append(q.data[s.cost], s)
		return
	}
	q.data[s.cost] = append(q.data[s.cost], s)
	q.costs = append(q.costs, s.cost)
}

func (q *queue) Pop() *state {
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
		/* q.data[mincost] = nil */
		delete(q.data, mincost)
		q.costs = append(q.costs[0:costindex], q.costs[costindex+1:]...)
	} else {
		q.data[mincost] = q.data[mincost][1:]
	}
	return tmp
}
