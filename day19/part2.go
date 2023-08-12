package day19

import (
	"fmt"
)

func Part2(inputFile string) string {
	normalized := getNormalized(inputFile)
	top := 0
	for i := 0; i < len(normalized); i++ {
		for j := i + 1; j < len(normalized); j++ {
			top = max(top, normalized[i].offset.manhattan(normalized[j].offset))
		}
	}
	return fmt.Sprint(top)
}
