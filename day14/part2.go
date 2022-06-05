package day14

import (
	"aoc2021/util"
	"fmt"
)

func Part2(inputFile string) string {
	lines := util.ReadInput(inputFile)
	s := lines[0]
	rules := make(map[string][2]string)
	for i := 2; i < len(lines); i++ {
		first := string(lines[i][0])
		second := string(lines[i][1])
		last := string(lines[i][6])
		rules[first+second] = [2]string{first + last, last + second}
	}

	pairs := make(map[string]int64)
	scores := make(map[byte]int64)
	for i := 0; i < len(s)-1; i++ {
		pairs[s[i:i+2]]++
		scores[s[i]]++
	}
	scores[s[len(s)-1]]++

	for i := 0; i < 40; i++ {
		old := make(map[string]int64)
		for k, v := range pairs {
			old[k] = v
		}
		for k, v := range old {
			if r, ok := rules[k]; ok {
				pairs[r[0]] += v
				pairs[r[1]] += v
				pairs[k] -= v
				scores[r[0][1]] += v
			}
		}
	}
	return fmt.Sprint(getScore(scores))
}
