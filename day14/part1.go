package day14

import (
	"aoc2021/util"
	"fmt"
	"math"
)

func add(result []byte, index *int, value byte, scores map[byte]int) {
	result[*index] = value
	*index++
	scores[value]++
}

func cycle(data []byte, rules map[byte]map[byte]byte) ([]byte, int) {
	out := 0
	result := make([]byte, len(data)*2)
	scores := make(map[byte]int)
	for in := 0; in < len(data)-1; in++ {
		add(result, &out, data[in], scores)
		if m, ok := rules[data[in]]; ok {
			if c, ok := m[data[in+1]]; ok {
				add(result, &out, c, scores)
			}
		}
	}
	add(result, &out, data[len(data)-1], scores)
	min := math.MaxInt
	max := 0
	for _, v := range scores {
		max = util.Max(max, v)
		min = util.Min(min, v)
	}
	printScores(scores)
	fmt.Println()
	return result[:out], max - min
}

func Part1(inputFile string) string {
	lines := util.ReadInput(inputFile)
	s := []byte(lines[0])
	rules := make(map[byte]map[byte]byte)
	for i := 2; i < len(lines); i++ {
		if _, ok := rules[lines[i][0]]; !ok {
			rules[lines[i][0]] = make(map[byte]byte)
		}
		rules[lines[i][0]][lines[i][1]] = lines[i][6]
	}
	score := 0
	for i := 0; i < 10; i++ {
		s, score = cycle(s, rules)
	}
	fmt.Println(string(s))
	return fmt.Sprint(score)
}
