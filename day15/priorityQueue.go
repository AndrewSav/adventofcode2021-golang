package day15

import (
	"container/heap"
)

type vertexPQ struct {
	level     int
	neighbors []*vertexPQ
	priority  int
	index     int
}

type PriorityQueue []*vertexPQ

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	vertexPQ := x.(*vertexPQ)
	vertexPQ.index = n
	*pq = append(*pq, vertexPQ)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	vertexPQ := old[n-1]
	old[n-1] = nil
	vertexPQ.index = -1
	*pq = old[0 : n-1]
	return vertexPQ
}

func (pq *PriorityQueue) update(vertexPQ *vertexPQ, priority int) {
	vertexPQ.priority = priority
	heap.Fix(pq, vertexPQ.index)
}
