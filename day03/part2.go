package day03

import (
	"aoc2021/util"
	"fmt"
	"strconv"
)

// return qualified lines either for oxygen generator or for CO2 scrubber
func splitWtihRule(input []string, pos int, isOxy bool) []string {
	ones, zeroes := split(input, pos)
	if (len(ones) >= len(zeroes) && isOxy) || (len(ones) < len(zeroes) && !isOxy) {
		return ones
	} else {
		return zeroes
	}
}

func toNum(s string) int { // binary strting to integer
	i, _ := strconv.ParseInt(s, 2, 32)
	return int(i)
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
	return fmt.Sprint(toNum(oxySet[0]) * toNum(co2Set[0]))
}
