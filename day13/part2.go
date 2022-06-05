package day13

import (
	"aoc2021/util"
	"fmt"
)

func Part2(inputFile string) string {
	points, folds := readData(inputFile)
	for _, f := range folds {
		points = fold(f, points)
	}
	var maxx, maxy int
	for _, p := range points {
		maxx = util.Max(maxx, p.x)
		maxy = util.Max(maxy, p.y)
	}
	plot := make([][]rune, maxy+1)
	for y := range plot {
		plot[y] = make([]rune, maxx+1)
		for x := range plot[y] {
			plot[y][x] = '.'
		}
	}
	for _, p := range points {
		plot[p.y][p.x] = '#'
	}
	return fmt.Sprint(util.OCR2021Day13Part2(plot))
}
