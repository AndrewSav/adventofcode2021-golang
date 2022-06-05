package day15

import (
	"aoc2021/util"
	"fmt"
	"math"
)

type vertex struct {
	level     byte
	neighbors []*vertex
}

func removeMin(q map[*vertex]*struct{}, dist map[*vertex]int) (result *vertex) {
	min := math.MaxInt
	for k := range q {
		newmin := util.Min(min, dist[k])
		if newmin < min {
			result = k
			min = newmin
		}
	}
	delete(q, result)
	return
}

// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Pseudocode
func search(data []*vertex, entry *vertex, exit *vertex) int {
	dist := make(map[*vertex]int)
	q := make(map[*vertex]*struct{})
	for _, d := range data {
		dist[d] = math.MaxInt
		q[d] = &struct{}{}
	}
	dist[entry] = 0
	for {
		u := removeMin(q, dist)
		if u == exit {
			return dist[u]
		}
		for _, v := range u.neighbors {
			if q[v] == nil {
				continue
			}
			alt := dist[u] + int(v.level)
			if alt < dist[v] {
				dist[v] = alt
			}
		}
	}
}

func Part1(inputFile string) string {
	lines := util.ReadInput(inputFile)
	data := make([]*vertex, 0, len(lines)*len(lines[0]))
	for y, l := range lines {
		for x := range l {
			v := vertex{level: l[x] - "0"[0], neighbors: make([]*vertex, 0)}
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
