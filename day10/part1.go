package day10

import (
	"aoc2021/util"
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func Part1(inputFile string) string {
	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	lines := util.ReadInput(inputFile)
	count := 0
	for _, l := range lines {
		s := stack.New()
		for _, c := range l {
			switch c {
			case '(', '[', '{', '<':
				s.Push(c)
				continue
			}
			if s.Pop().(rune) != pairs[c] {
				count += scores[c]
				break
			}
		}
	}
	return fmt.Sprint(count)
}
