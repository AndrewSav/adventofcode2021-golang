package day15

import (
	"container/heap"
)

type vertex struct {
	level     byte
	neighbors []*vertex
	priority  int
	index     int
}

type PriorityQueue []*vertex

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
	vertex := x.(*vertex)
	vertex.index = n
	*pq = append(*pq, vertex)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	vertex := old[n-1]
	old[n-1] = nil
	vertex.index = -1
	*pq = old[0 : n-1]
	return vertex
}

func (pq *PriorityQueue) update(vertex *vertex, priority int) {
	vertex.priority = priority
	heap.Fix(pq, vertex.index)
}
