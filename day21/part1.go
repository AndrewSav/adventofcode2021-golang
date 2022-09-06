package day21

import (
	"aoc2021/util"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func getPosition(s string) int {
	r := regexp.MustCompile(`Player \d+ starting position: (\d+)`)
	match := r.FindStringSubmatch(s)
	if match == nil {
		log.Fatalf("line '%s' cannot be matched", s)
	}
	i, err := strconv.Atoi(match[1])
	if err != nil {
		panic(err)
	}
	return i
}

type die struct {
	state int
	count int
}

func (d *die) next() int {
	if d.state < 1 || d.state > 100 {
		d.state = 1
	}
	result := int(d.state)
	d.state++
	d.count++
	return result
}

func (d *die) next3() int {
	return d.next() + d.next() + d.next()
}

func Part1(inputFile string) string {
	data := util.ReadInput(inputFile)

	positions := [2]int{getPosition(data[0]), getPosition(data[1])}
	scores := [2]int{0, 0}

	d := die{}
	for i := 0; ; i = (i + 1) % 2 {
		positions[i] = (positions[i]-1+d.next3())%10 + 1
		scores[i] += positions[i]
		if scores[i] >= 1000 {
			return fmt.Sprint(scores[(i+1)%2] * d.count)
		}
	}
}
