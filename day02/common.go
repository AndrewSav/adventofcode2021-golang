package day02

import (
	"aoc2021/util"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type data struct {
	hpos  int
	depth int
	aim   int
}

type mutator func(*data, int)

func solve(inputFile string, up, down, forward mutator) string {
	var data data
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
