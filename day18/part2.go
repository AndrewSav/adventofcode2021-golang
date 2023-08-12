package day18

import (
	"aoc2021/util"
	"fmt"
)

func Part2(inputFile string) string {
	lines := util.ReadInput(inputFile)
	var result int
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i != j {
				result = max(result, add(parse(lines[i]), parse(lines[j])).getMagnitude())
			}
		}
	}
	return fmt.Sprint(result)
}
