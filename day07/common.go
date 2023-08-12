package day07

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

type fuelFunc func(int) int

func getFuel(ss []int, goal int, consumption fuelFunc) (fuel int) {
	for _, s := range ss {
		fuel += consumption(util.Abs(s - goal))
	}
	return
}

func solve(inputFile string, consumption fuelFunc) string {
	lines := util.ReadInput(inputFile)
	ss := util.AtoiSlice(strings.Split(lines[0], ","))
	minpos, maxpos := ss[0], ss[0]
	// the target position has to be between minpos and mmaxpos (inclusive)
	// because if it is outside these bounds every single trip can be
	// made cheaper by choosing the minpos (or maxpos) target position
	for _, s := range ss {
		minpos = min(minpos, s)
		maxpos = max(maxpos, s)
	}
	fuel := getFuel(ss, maxpos, consumption) // got to start somewhere
	for i := minpos; i < maxpos; i++ {
		fuel = min(fuel, getFuel(ss, i, consumption))
	}
	return fmt.Sprint(fuel)
}
