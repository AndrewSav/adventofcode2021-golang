package day03

import (
	"aoc2021/util"
	"fmt"
	"log"
	"strconv"
)

func splitWtihRule(input []string, pos int, isOxy bool) []string {
	ones, zeroes := split(input, pos)
	if (len(ones) >= len(zeroes) && isOxy) || (len(ones) < len(zeroes) && !isOxy) {
		return ones
	} else {
		return zeroes
	}
}

func toNum(s string) int {
	if i, err := strconv.ParseInt(s, 2, 32); err != nil {
		log.Fatalf("cannot convert strin '%s' to a binary number: %v", s, err)
	} else {
		return int(i)
	}
	panic("unexpected code path")
}

func Part2(inputFile string) string {
	lines := util.ReadInput(inputFile)
	oxySet, co2Set := lines, lines
	dim := len(lines[0])
	for pos := 0; pos < dim && (len(oxySet) > 1 || len(co2Set) > 1); pos++ {
		if len(oxySet) > 1 {
			oxySet = splitWtihRule(oxySet, pos, true)
		}
		if len(co2Set) > 1 {
			co2Set = splitWtihRule(co2Set, pos, false)
		}
	}
	if len(oxySet) != 1 || len(co2Set) != 1 {
		panic(fmt.Errorf("want 1 of each oxySet, co2Set, have oxySet: %s, co2Set: %s\n", len(oxySet), len(co2Set)))
	}
	return fmt.Sprint(toNum(oxySet[0]) * toNum(co2Set[0]))
}
