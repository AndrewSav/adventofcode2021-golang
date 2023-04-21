package day13

import (
	"aoc2021/util"
	"fmt"
)

func Part2(inputFile string) string {
	points, folds := readData(inputFile)
	// execute all folds
	for _, f := range folds {
		points = fold(f, points)
	}
	// find the bounds of the resulting "image"
	var maxX, maxY int
	for _, p := range points {
		maxX = util.Max(maxX, p.x)
		maxY = util.Max(maxY, p.y)
	}
	// initialize "image" array
	plot := make([][]rune, maxY+1)
	for y := range plot {
		plot[y] = make([]rune, maxX+1)
		for x := range plot[y] {
			plot[y][x] = '.'
		}
	}
	// plot the "image" into the array
	for _, p := range points {
		plot[p.y][p.x] = '#'
	}
	// OCR the image and return the result
	return fmt.Sprint(util.OCR2021Day13Part2(plot))
}
