package day15

import (
	"aoc2021/util"
	"fmt"
)

type vertex struct {
	level     int
	neighbors []*vertex
}

type bucketQueue []vertex

func (q *bucketQueue) add(v *vertex, val int) {
	bucketOffset := val - (*q)[0].level
	if len(*q) < bucketOffset+1 { // if the bucket is not in there yet create it and all buckets in-between
		*q = append(*q, make([]vertex, bucketOffset+1-len(*q))...)
	}
	(*q)[bucketOffset].level = val                                         // for brivety, only needs to be set for new buckets in a loop insude the above if
	(*q)[bucketOffset].neighbors = append((*q)[bucketOffset].neighbors, v) //finally add the value to the bucket
}

// Dijkstra's search using Bucket queue
// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Using_a_priority_queue
// https://en.wikipedia.org/wiki/Bucket_queue
func search(entry *vertex, exit *vertex) int {
	dist := map[*vertex]int{entry: 0}
	q := bucketQueue{vertex{level: 0, neighbors: []*vertex{entry}}}
	for {
		for len(q[0].neighbors) == 0 { // find first non-empty bucket
			q = q[1:]
		}
		u := q[0].neighbors[0]
		q[0].neighbors = q[0].neighbors[1:] // remove from backet
		if u == exit {
			return dist[u]
		}
		for _, v := range u.neighbors {
			alt := dist[u] + v.level
			if i, ok := dist[v]; !ok || alt < i {
				dist[v] = alt
				q.add(v, alt) // we do not remove old (if any) because practically it does not seem to matter
			}
		}
	}
}

func solve(inputFile string, multx, multy int) string {
	lines := util.ReadInput(inputFile)
	data := make([]*vertex, 0, len(lines)*multy*len(lines[0])*multx)
	for dy := 0; dy < multy; dy++ {
		for y, l := range lines {
			for dx := 0; dx < multx; dx++ {
				for x := range l {
					v := vertex{level: (int(l[x]-"0"[0])-1+dx+dy)%9 + 1, neighbors: make([]*vertex, 0, 2)}
					offset := (y+dy*len(lines))*len(l)*multx + x + dx*len(l)
					if x+dx*len(l) > 0 {
						data[offset-1].neighbors = append(data[offset-1].neighbors, &v)
						v.neighbors = append(v.neighbors, data[offset-1])
					}
					if y+dy*len(lines) > 0 {
						data[offset-len(l)*multx].neighbors = append(data[offset-len(l)*multx].neighbors, &v)
						v.neighbors = append(v.neighbors, data[offset-len(l)*multx])
					}
					data = append(data, &v)
				}
			}
		}
	}
	return fmt.Sprint(search(data[0], data[len(data)-1]))
}
