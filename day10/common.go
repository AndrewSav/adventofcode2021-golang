package day10

import (
	"aoc2021/util"
	"fmt"
	"sort"
)

type stack []rune

func (s *stack) Pop() (result rune) {
	result, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return
}

func solveInner(inputFile string) (int, int64) {
	syntaxScores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	autoScores := map[rune]int{
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
	resultSyntax := 0
	resultsAuto := []int64{}
	for _, l := range lines {
		s := stack{}
		broken := false
		for _, c := range l {
			switch c {
			case '(', '[', '{', '<':
				s = append(s, c)
				continue
			}
			if pairs[s.Pop()] != c {
				resultSyntax += syntaxScores[c]
				broken = true
				break
			}
		}
		if !broken {
			count := int64(0)
			for len(s) > 0 {
				count = count*5 + int64(autoScores[pairs[s.Pop()]])
			}
			resultsAuto = append(resultsAuto, count)
		}
	}
	sort.Slice(resultsAuto, func(i, j int) bool { return resultsAuto[i] < resultsAuto[j] })
	return resultSyntax, resultsAuto[len(resultsAuto)/2]
}

func solve(inputFile string, syntax bool) string {
	resultSyntax, resultAuto := solveInner(inputFile)
	if syntax {
		return fmt.Sprint(resultSyntax)
	} else {
		return fmt.Sprint(resultAuto)
	}
}
