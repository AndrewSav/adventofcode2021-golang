package day17

import (
	"aoc2021/util"
	"fmt"
	"regexp"
)

func Part1(inputFile string) string {
	lines := util.ReadInput(inputFile)
	r := regexp.MustCompile(`^target area: x=(\d+)..(\d+), y=-(\d+)..-(\d+)$`)
	match := r.FindStringSubmatch(lines[0])
	a := util.MustAtoi(match[3])
	return fmt.Sprint(a * (a - 1) / 2)
}
