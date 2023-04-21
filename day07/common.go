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
	min, max := ss[0], ss[0]
	// the target position has to be between min and max (inclusive)
	// because if it is outside these bounds every single trip can be
	// made cheaper by choosing the min (or max) target position
	for _, s := range ss {
		min = util.Min(min, s)
		max = util.Max(max, s)
	}
	fuel := getFuel(ss, max, consumption) // got to start somewhere
	for i := min; i < max; i++ {
		fuel = util.Min(fuel, getFuel(ss, i, consumption))
	}
	return fmt.Sprint(fuel)
}
