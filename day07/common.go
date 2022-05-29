package day07

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

type fuelFunc func(int) int

func getFuel(ss []int, goal int, consumption fuelFunc) (fuel int) {
	for _, s := range ss {
		fuel += consumption(abs(s - goal))
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(inputFile string, consumption fuelFunc) string {
	lines := util.ReadInput(inputFile)
	ss := util.Atoi(strings.Split(lines[0], ","))
	mins, maxs := ss[0], ss[0]
	for _, s := range ss {
		mins = min(mins, s)
		maxs = max(maxs, s)
	}
	fuel := getFuel(ss, maxs, consumption)
	for i := mins; i < maxs; i++ {
		fuel = min(fuel, getFuel(ss, i, consumption))
	}
	return fmt.Sprint(fuel)
}
