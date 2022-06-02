package day11

import (
	"aoc2021/util"
	"fmt"
)

type octopi [][]int
type point struct {
	x int
	y int
}

func (o octopi) cycle() int {
	flashes := []point{}
	for y, ll := range o {
		for x, l := range ll {
			if l == 9 {
				flashes = append(flashes, point{x: x, y: y})
			}
			o[y][x]++
		}
	}
	count := 0
	for _, f := range flashes {
		count += o.flash(f.x, f.y)
	}
	for y, ll := range o {
		for x, l := range ll {
			if l > 9 {
				o[y][x] = 0
			}
		}
	}
	return count
}

func (o octopi) propagateFlash(x, y int) int {
	if y < 0 || x < 0 || y >= len(o) || x >= len(o[y]) {
		return 0
	}
	o[y][x]++
	if o[y][x] != 10 {
		return 0
	}
	return o.flash(x, y)
}

func (o octopi) flash(x, y int) int {
	count := 1
	count += o.propagateFlash(x-1, y-1)
	count += o.propagateFlash(x-1, y)
	count += o.propagateFlash(x-1, y+1)
	count += o.propagateFlash(x, y-1)
	count += o.propagateFlash(x, y+1)
	count += o.propagateFlash(x+1, y-1)
	count += o.propagateFlash(x+1, y)
	count += o.propagateFlash(x+1, y+1)
	return count
}

func Part1(inputFile string) string {
	data := octopi(util.GetPlot(inputFile))
	count := 0
	for i := 0; i < 100; i++ {
		count += data.cycle()
	}
	return fmt.Sprint(count)
}
