package day25

import (
	"aoc2021/util"
	"fmt"
)

func Part1(inputFile string) string {
	lines := util.ReadInput(inputFile)
	sea := make([][]rune, len(lines))
	for i, l := range lines {
		sea[i] = []rune(l)
		for j := range sea[i] {
			if sea[i][j] == '.' {
				sea[i][j] = 0
			}
		}
	}
	steps := 0
	for moved := true; moved; steps++ {
		moved = false

		// Make next sea state array first
		nextSea := make([][]rune, len(sea))
		for i := range nextSea {
			nextSea[i] = make([]rune, len(sea[i]))
		}

		// First right-facing cucumbers move if they can
		for i, l := range sea {
			for j, c := range l {
				if c == '>' && l[(j+1)%len(l)] == 0 {
					moved = true
					nextSea[i][(j+1)%len(l)] = c
				} else if c != 0 {
					nextSea[i][j] = c
				}
			}
		}
		sea = nextSea

		// Make next sea state array again
		nextSea = make([][]rune, len(sea))
		for i := range nextSea {
			nextSea[i] = make([]rune, len(sea[i]))
		}

		// Then down-facing cucumbers move if they can
		for i, l := range sea {
			for j, c := range l {
				if c == 'v' && sea[(i+1)%len(sea)][j] == 0 {
					moved = true
					nextSea[(i+1)%len(sea)][j] = c
				} else if c != 0 {
					nextSea[i][j] = c
				}
			}
		}
		sea = nextSea
	}
	return fmt.Sprint(steps)
}
