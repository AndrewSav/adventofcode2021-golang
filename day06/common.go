package day06

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

func tick(m [9]int64) (result [9]int64) {
	for i, n := range m {
		if i > 0 {
			result[i-1] += n
		} else {
			result[6] += n
			result[8] += n
		}
	}
	return
}

func solve(inputFile string, days int) string {
	lines := util.ReadInput(inputFile)
	ff := util.AtoiSlice(strings.Split(lines[0], ","))
	// the trick is not to track each individual fish, but each group of fish
	// with the same cycle offset. 0-8 are all possible offsets
	m := [9]int64{}
	for _, f := range ff {
		m[f]++
	}
	for i := 0; i < days; i++ {
		m = tick(m)
	}
	count := int64(0)
	for _, f := range m {
		count += f
	}
	return fmt.Sprint(count)
}
