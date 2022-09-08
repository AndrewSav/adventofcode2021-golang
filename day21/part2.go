package day21

import (
	"aoc2021/util"
	"fmt"
	"sort"
	"strings"
)

type step struct {
	roll     int
	score    int
	position int
}

type state struct {
	position    int
	score       int
	rollHistory []step
	count       int
}

func formatRollHistory(initialPostion int, rollHistory []step) string {

	position := initialPostion
	score := 0
	var sb strings.Builder
	//fmt.Fprint(&sb, rollHistory)
	fmt.Fprintf(&sb, "%d(%d), ", position+1, score)
	for _, b := range rollHistory {
		fmt.Fprintf(&sb, "%d(%d), ", b.position+1, b.score)
	}
	return sb.String()[0 : sb.Len()-2]
}

var weights = map[int]int{3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

func addResults(m map[int]int) map[int]int {

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	result := map[int]int{}
	acc := 0
	for _, k := range keys {
		acc += m[k]
		result[k] = acc
	}
	return result

}

func doTheNumbers(j int) (map[int]int64, map[int]int64) {
	var win = map[int]int64{}
	var lose = map[int]int64{}
	//count := 0
	//count2 := int64(0)
	stack := []state{{position: j, score: 0, count: 1}}
	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//for i := 1; i <= 3; i++ {
		for i := 3; i <= 9; i++ {
			newPosition := (s.position + i) % 10
			newScore := s.score
			//if len(s.rollHistory)%3 == 2 {
			newScore += newPosition + 1
			newCount := s.count * weights[i]
			//}
			x := len(s.rollHistory)
			rhs := step{position: newPosition, score: newScore, roll: i}
			if newScore < 21 {
				stack = append(stack, state{position: newPosition, score: newScore, count: newCount, rollHistory: append(s.rollHistory[:x:x], rhs)})
				lose[x+1] += int64(newCount)
			} else {
				//count++
				//count2 += newCount
				win[x+1] += int64(newCount)
				//ss := formatRollHistory(j, append(s.rollHistory[:x:x], rhs))
				//fmt.Printf("%s | %d(%d)\n", ss, newPosition+1, newScore)
			}
		}
	}
	//fmt.Println(count)
	//fmt.Println(count2)
	return win, lose
}

func Part2(inputFile string) string {
	data := util.ReadInput(inputFile)

	first, second := getPosition(data[0]), getPosition(data[1])
	//first := 4
	//second := 8

	//for j := 0; j < 10; j++ {
	//	win, lose := doTheNumbers(first - 1)
	//	fmt.Println(win)
	//	fmt.Println(lose)
	//	fmt.Println()
	//}

	win1, lose1 := doTheNumbers(first - 1)
	win2, lose2 := doTheNumbers(second - 1)
	//fmt.Println(win1)
	//fmt.Println(lose1)
	//fmt.Println()
	//fmt.Println(win2)
	//fmt.Println(lose2)
	var firstWins, secondWins int64
	for i := 3; i <= 10; i++ {
		//fmt.Println(win1[i] * lose2[i-1])
		firstWins += win1[i] * lose2[i-1]
		secondWins += win2[i] * lose1[i]
	}
	return fmt.Sprint(util.Max(firstWins, secondWins))
}
