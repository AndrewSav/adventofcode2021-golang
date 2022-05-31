package day10

import (
	"aoc2021/util"
	"fmt"
	"sort"

	"github.com/golang-collections/collections/stack"
)

func solveInner(inputFile string) (int, int) {
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
	resultsAuto := []int{}
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
				resultSyntax += syntaxScores[c]
				broken = true
				break
			}
		}
		if !broken {
			count := 0
			for s.Len() > 0 {
				count = count*5 + autoScores[pairs[s.Pop().(rune)]]
			}
			resultsAuto = append(resultsAuto, count)
		}
	}
	sort.Ints(resultsAuto)
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
