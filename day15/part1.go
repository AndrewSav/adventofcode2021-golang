package day15

import (
	"aoc2021/util"
	"container/heap"
	"fmt"
)

// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Using_a_priority_queue
func search(data []*vertexPQ, entry *vertexPQ, exit *vertexPQ) int {
	dist := make(map[*vertexPQ]int)
	dist[entry] = 0
	f := PriorityQueue{entry}
	heap.Init(&f)
	for {
		u := heap.Pop(&f).(*vertexPQ)
		if u == exit {
			return dist[u]
		}
		for _, v := range u.neighbors {
			alt := dist[u] + int(v.level)
			if i, ok := dist[v]; !ok {
				dist[v] = alt
				heap.Push(&f, v)
				f.update(v, alt)
			} else {
				if alt < i {
					dist[v] = alt
					f.update(v, alt)
				}
			}
		}
	}
}

func Part1(inputFile string) string {
	lines := util.ReadInput(inputFile)
	data := make([]*vertexPQ, 0, len(lines)*len(lines[0]))
	for y, l := range lines {
		for x := range l {
			v := vertexPQ{level: int(l[x] - "0"[0]), neighbors: make([]*vertexPQ, 0)}
			offset := y*len(l) + x
			if x > 0 {
				data[offset-1].neighbors = append(data[offset-1].neighbors, &v)
				v.neighbors = append(v.neighbors, data[offset-1])
			}
			if y > 0 {
				data[offset-len(l)].neighbors = append(data[offset-len(l)].neighbors, &v)
				v.neighbors = append(v.neighbors, data[offset-len(l)])
			}
			data = append(data, &v)
		}
	}

	return fmt.Sprint(search(data, data[0], data[len(data)-1]))
}
