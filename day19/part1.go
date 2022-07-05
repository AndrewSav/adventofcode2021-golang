package day19

import (
	"fmt"
)

func Part1(inputFile string) string {
	normalized := getNormalized(inputFile)

	unique := make(map[point]struct{})
	for _, s := range normalized {
		for _, p := range s.points {
			unique[p] = struct{}{}
		}
	}

	return fmt.Sprint(len(unique))
}
