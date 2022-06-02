package day09

import (
	"aoc2021/util"
	"fmt"
	"sort"
)

const boundaryColor = 9
const fillColor = -1

type plot [][]int

func (p plot) boundaryFill(x, y, fillColor, boundaryColor int) (result int) {

	color := p[y][x]
	if color != boundaryColor && color != fillColor {
		p[y][x] = fillColor
		result++
		if x < len(p[y])-1 {
			result += p.boundaryFill(x+1, y, fillColor, boundaryColor)
		}
		if y < len(p)-1 {
			result += p.boundaryFill(x, y+1, fillColor, boundaryColor)
		}
		if x > 0 {
			result += p.boundaryFill(x-1, y, fillColor, boundaryColor)
		}
		if y > 0 {
			result += p.boundaryFill(x, y-1, fillColor, boundaryColor)
		}
	}
	return
}

func Part2(inputFile string) string {
	data := plot(util.GetPlot(inputFile))
	var results []int
	for y, dd := range data {
		for x := range dd {
			if data[y][x] != boundaryColor && data[y][x] != fillColor {
				results = append(results, data.boundaryFill(x, y, fillColor, boundaryColor))
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(results)))
	return fmt.Sprint(results[0] * results[1] * results[2])
}
