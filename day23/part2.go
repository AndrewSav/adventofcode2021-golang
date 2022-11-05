package day23

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

func parseInputPart2(input []string) (result state) {
	parseLine(&result.room, 0, []rune(strings.Trim(input[2], " #")))
	parseLine(&result.room, 1, []rune("D#C#B#A"))
	parseLine(&result.room, 2, []rune("D#B#A#C"))
	parseLine(&result.room, 3, []rune(strings.Trim(input[3], " #")))
	return
}

func Part2(inputFile string) string {
	roomSlots = 4
	final = state{room: [4][4]rune{{'A', 'A', 'A', 'A'}, {'B', 'B', 'B', 'B'}, {'C', 'C', 'C', 'C'}, {'D', 'D', 'D', 'D'}}}
	return fmt.Sprint(solve(parseInputPart2(util.ReadInput(inputFile))))
}
