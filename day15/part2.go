package day15

import (
	"aoc2021/util"
	"fmt"
)

func Part2(inputFile string) string {
	lines := util.ReadInput(inputFile)
	data := make([]*vertexPQ, 0, len(lines)*len(lines[0])*25)
	for dy := 0; dy < 5; dy++ {
		for y, l := range lines {
			for dx := 0; dx < 5; dx++ {
				for x := range l {
					v := vertexPQ{level: (int(l[x]-"0"[0])-1+dx+dy)%9 + 1, neighbors: make([]*vertexPQ, 0)}
					offset := (y+dy*len(lines))*len(l)*5 + x + dx*len(l)
					if x+dx*len(l) > 0 {
						data[offset-1].neighbors = append(data[offset-1].neighbors, &v)
						v.neighbors = append(v.neighbors, data[offset-1])
					}
					if y+dy*len(lines) > 0 {
						data[offset-len(l)*5].neighbors = append(data[offset-len(l)*5].neighbors, &v)
						v.neighbors = append(v.neighbors, data[offset-len(l)*5])
					}
					data = append(data, &v)
				}
			}
		}
	}
	return fmt.Sprint(search(data, data[0], data[len(data)-1]))
}
