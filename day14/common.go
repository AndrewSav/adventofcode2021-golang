package day14

import (
	"aoc2021/util"
)

// This implements how the score is calculated
// according to the puzzle description
func getScore[T int | int64](scores map[byte]T) T {
	var min, max T
	for _, v := range scores {
		min, max = v, v
		break
	}
	for _, v := range scores {
		max = util.Max(max, v)
		min = util.Min(min, v)
	}
	return max - min
}
