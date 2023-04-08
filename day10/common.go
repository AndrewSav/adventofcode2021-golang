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
	autoCompleteScores := map[rune]int{
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
	resultsAutoComplete := []int64{}
	for _, l := range lines {
		s := stack{}
		broken := false
		for _, c := range l {
			switch c { // if it's an opening character push it onto stack and go to the next one
			case '(', '[', '{', '<':
				s = append(s, c)
				continue
			}
			if pairs[s.Pop()] != c { // see if the closing character that we got matches the openning character that we expect
				resultSyntax += syntaxScores[c]
				broken = true
				break
			}
		}
		if !broken { // if there was no syntax error, now pop all the remaining unclosed character and calculate closing score
			count := int64(0)
			for len(s) > 0 {
				count = count*5 + int64(autoCompleteScores[pairs[s.Pop()]])
			}
			resultsAutoComplete = append(resultsAutoComplete, count)
		}
	}
	// getting the middle score
	sort.Slice(resultsAutoComplete, func(i, j int) bool { return resultsAutoComplete[i] < resultsAutoComplete[j] })
	return resultSyntax, resultsAutoComplete[len(resultsAutoComplete)/2]
}

func solve(inputFile string, syntax bool) string {
	resultSyntax, resultAutoComplete := solveInner(inputFile)
	if syntax {
		return fmt.Sprint(resultSyntax)
	} else {
		return fmt.Sprint(resultAutoComplete)
	}
}
