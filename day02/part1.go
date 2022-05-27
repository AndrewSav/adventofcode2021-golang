package day02

import (
	"aoc2021/util"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func Part1(inputFile string) string {
	var hpos, depth int
	lines := util.ReadInput(inputFile)
	r := regexp.MustCompile(`^(forward|down|up) (\d)$`)
	for index, l := range lines {
		match := r.FindStringSubmatch(l)
		if match == nil {
			log.Fatalf("line %d '%s' cannot be matched", index, l)
		}
		i, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		switch match[1] {
		case "up":
			depth -= i
		case "down":
			depth += i
		case "forward":
			hpos += i
		default:
			panic("unexpected code path")
		}
	}
	return fmt.Sprint(hpos * depth)
}
