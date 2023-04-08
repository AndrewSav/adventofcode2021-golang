package day02

import (
	"aoc2021/util"
	"fmt"
	"regexp"
	"strconv"
)

type data struct {
	hpos  int
	depth int
	aim   int //aim is only used in part 2
}

type mutator func(*data, int)

func solve(inputFile string, up, down, forward mutator) string {
	var data data
	lines := util.ReadInput(inputFile)
	r := regexp.MustCompile(`^(forward|down|up) (\d)$`)
	for _, l := range lines {
		match := r.FindStringSubmatch(l)
		i, _ := strconv.Atoi(match[2])
		switch match[1] {
		case "up":
			up(&data, i)
		case "down":
			down(&data, i)
		case "forward":
			forward(&data, i)
		default:
			panic("unexpected code path")
		}
	}
	return fmt.Sprint(data.hpos * data.depth)
}
