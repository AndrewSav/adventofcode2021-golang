package day15

import (
	"aoc2021/util"
	"fmt"
)

type vertex struct {
	level     int
	neighbors []*vertex
}

func add(f []vertex, v *vertex, val int) []vertex {
	offset := val - f[0].level
	if len(f) < offset+1 {
		f = append(f, make([]vertex, offset+1-len(f))...)
	}
	f[offset].level = val
	f[offset].neighbors = append(f[offset].neighbors, v)
	return f
}

func update(f []vertex, v *vertex, old, new int) {
	offset := old - f[0].level
	for i, c := range f[offset].neighbors {
		if c == v {
			f[offset].neighbors[i] = f[offset].neighbors[len(f[offset].neighbors)-1]
			f[offset].neighbors = f[offset].neighbors[:len(f[offset].neighbors)-1]
			break
		}
	}
	add(f, v, new)
}

// https://en.wikipedia.org/wiki/Bucket_queue
func searchBucket(data []*vertex, entry *vertex, exit *vertex) int {
	dist := make(map[*vertex]int)
	dist[entry] = 0
	f := make([]vertex, 1)
	f[0] = vertex{level: 0, neighbors: []*vertex{entry}}
	for {
		for len(f[0].neighbors) == 0 {
			f = f[1:]
		}
		u := f[0].neighbors[0]
		f[0].neighbors = f[0].neighbors[1:]
		if u == exit {
			return dist[u]
		}
		for _, v := range u.neighbors {
			alt := dist[u] + v.level
			if i, ok := dist[v]; !ok {
				dist[v] = alt
				f = add(f, v, alt)
			} else {
				if alt < i {
					old := dist[v]
					dist[v] = alt
					update(f, v, old, alt)
				}
			}
		}
	}
}

func Part1a(inputFile string) string {
	lines := util.ReadInput(inputFile)
	data := make([]*vertex, 0, len(lines)*len(lines[0]))
	for y, l := range lines {
		for x := range l {
			v := vertex{level: int(l[x] - "0"[0]), neighbors: make([]*vertex, 0)}
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

	return fmt.Sprint(searchBucket(data, data[0], data[len(data)-1]))
}
