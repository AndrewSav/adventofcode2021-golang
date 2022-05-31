package day09

import (
	"aoc2021/util"
	"fmt"
	"sort"
	"strings"
)

type plot [][]int

const boundaryColor = 9
const fillColor = -1

func (p *plot) boundaryFill(x, y, fillColor, boundaryColor int) (result int) {

	color := (*p)[y][x]
	if color != boundaryColor && color != fillColor {
		(*p)[y][x] = fillColor
		result++
		if x < len((*p)[y])-1 {
			result += p.boundaryFill(x+1, y, fillColor, boundaryColor)
		}
		if y < len((*p))-1 {
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
	lines := util.ReadInput(inputFile)
	var data = make(plot, len(lines))
	for i, l := range lines {
		data[i] = util.Atoi(strings.Split(l, ""))
	}
	var results []int
	for y, dd := range data {
		for x := range dd {
			if data[y][x] != boundaryColor && data[y][x] != fillColor {
				results = append(results, data.boundaryFill(x, y, fillColor, boundaryColor))
			}
		}
	}
	sort.Ints(results)
	return fmt.Sprint(results[len(results)-1] * results[len(results)-2] * results[len(results)-3])

}
