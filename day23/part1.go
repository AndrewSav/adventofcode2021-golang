package day23

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

func parseInputPart1(input []string) (result state) {
	parseLine(&result.room, 0, []rune(strings.Trim(input[2], " #")))
	parseLine(&result.room, 1, []rune(strings.Trim(input[3], " #")))
	return
}

func Part1(inputFile string) string {
	roomSlots = 2
	final = state{room: [4][4]rune{{'A', 'A'}, {'B', 'B'}, {'C', 'C'}, {'D', 'D'}}}
	return fmt.Sprint(solve(parseInputPart1(util.ReadInput(inputFile))))
}
