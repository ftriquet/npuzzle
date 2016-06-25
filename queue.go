package main

import "container/heap"

type PriorityQueue []*state

func (p PriorityQueue) Len() int { return len(p) }

func (p PriorityQueue) Less(i, j int) bool {
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
