package day18

import (
	"aoc2021/util"
	"fmt"
)

func Part2(inputFile string) string {
	lines := util.ReadInput(inputFile)
	var result int
	cache := make(map[int]*term) // gives about 15 times performance inporvement
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i != j {
				if cache[i] == nil {
					cache[i] = parse(lines[i])
				}
				if cache[j] == nil {
					cache[j] = parse(lines[j])
				}
				result = util.Max(result, add(cache[i], cache[j]).getMagnitude())
			}
		}
	}
	return fmt.Sprint(result)
}
