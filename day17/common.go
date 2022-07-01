package day17

import (
	"aoc2021/util"
	"regexp"
)

func getInput(inputFile string) (x1 int, x2 int, y1 int, y2 int) {
	lines := util.ReadInput(inputFile)
	r := regexp.MustCompile(`^target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)$`)
	match := r.FindStringSubmatch(lines[0])
	x1 = util.MustAtoi(match[1])
	x2 = util.MustAtoi(match[2])
	y1 = util.MustAtoi(match[3])
	y2 = util.MustAtoi(match[4])
	return
}
