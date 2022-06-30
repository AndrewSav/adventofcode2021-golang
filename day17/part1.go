package day17

import (
	"aoc2021/util"
	"fmt"
	"regexp"
)

func Part1(inputFile string) string {
	lines := util.ReadInput(inputFile)
	r := regexp.MustCompile(`^target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)$`)
	match := r.FindStringSubmatch(lines[0])
	_ = util.MustAtoi(match[1])
	_ = util.MustAtoi(match[2])
	y1 := util.MustAtoi(match[3])
	_ = util.MustAtoi(match[4])
	return fmt.Sprint(y1 * (y1 + 1) / 2)
}
