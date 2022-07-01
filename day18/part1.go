package day18

import (
	"aoc2021/util"
	"fmt"
)

func Part1(inputFile string) string {
	lines := util.ReadInput(inputFile)
	var result *term
	for i, l := range lines {
		if i == 0 {
			result = parse(l)
		} else {
			result = add(result, parse(l))
		}
	}
	return fmt.Sprint(result.getMagnitude())
}
