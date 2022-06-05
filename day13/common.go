package day13

import (
	"aoc2021/util"
	"strings"
)

type point struct {
	x int
	y int
}

func newPointFromSlice(parts []int) *point {
	return &point{x: parts[0], y: parts[1]}
}

func removeDuplicates(points []point) []point {
	processed := map[point]struct{}{}
	w := 0
	for _, s := range points {
		if _, exists := processed[s]; !exists {
			processed[s] = struct{}{}
			points[w] = s
			w++
		}
	}
	return points[:w]
}
func foldx(c int, points []point) []point {
	for i, p := range points {
		if p.x > c {
			points[i].x = 2*c - p.x
		}
	}
	return removeDuplicates(points)
}

func foldy(c int, points []point) []point {
	for i, p := range points {
		if p.y > c {
			points[i].y = 2*c - p.y
		}
	}
	return removeDuplicates(points)
}

func fold(fold point, points []point) (result []point) {
	result = points
	if fold.x != 0 {
		result = foldx(fold.x, result)
	}
	if fold.y != 0 {
		result = foldy(fold.y, result)
	}
	return
}

func readData(inputFile string) (points, folds []point) {
	lines := util.ReadInput(inputFile)
	coordsSection := true
	for _, l := range lines {
		if l == "" {
			coordsSection = false
			continue
		}
		if coordsSection {
			points = append(points, *newPointFromSlice(util.Atoi(strings.Split(l, ","))))
		} else {
			parts := strings.Split(l[len("fold along "):], "=")
			if parts[0] == "x" {
				folds = append(folds, point{util.MustAtoi(parts[1]), 0})
			} else {
				folds = append(folds, point{0, util.MustAtoi(parts[1])})
			}
		}
	}
	return
}
