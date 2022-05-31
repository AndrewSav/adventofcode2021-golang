package day10

import (
	"aoc2021/util"
	"fmt"
	"sort"

	"github.com/golang-collections/collections/stack"
)

func Part2(inputFile string) string {
	scores := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	lines := util.ReadInput(inputFile)
	results := []int{}
	for _, l := range lines {
		s := stack.New()
		broken := false
		for _, c := range l {
			switch c {
			case '(', '[', '{', '<':
				s.Push(c)
				continue
			}
			if pairs[s.Pop().(rune)] != c {
				broken = true
				break
			}
		}
		if !broken {
			count := 0
			for s.Len() > 0 {
				count = count*5 + scores[pairs[s.Pop().(rune)]]
			}
			results = append(results, count)
		}
	}
	sort.Ints(results)
	return fmt.Sprint(results[len(results)/2])
}
