package day19

import (
	"aoc2021/util"
	"fmt"
)

func Part2(inputFile string) string {
	normalized := getNormalized(inputFile)
	max := 0
	for i := 0; i < len(normalized); i++ {
		for j := i + 1; j < len(normalized); j++ {
			max = util.Max(max, normalized[i].offset.manhatten(normalized[j].offset))
		}
	}
	return fmt.Sprint(max)
}
