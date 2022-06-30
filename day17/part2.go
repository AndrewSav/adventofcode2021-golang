package day17

import (
	"aoc2021/util"
	"fmt"
	"regexp"
)

type set map[int]struct{}

func (s set) Add(new []int) {
	for _, i := range new {
		s[i] = struct{}{}
	}
}

func (s set) AddInert(m map[int][]int, i int) {
	for k, v := range m {
		if i >= k {
			for _, x := range v {
				s[x] = struct{}{}
			}
		}
	}
}

func Part2(inputFile string) string {
	lines := util.ReadInput(inputFile)
	r := regexp.MustCompile(`^target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)$`)
	match := r.FindStringSubmatch(lines[0])
	x1 := util.MustAtoi(match[1])
	x2 := util.MustAtoi(match[2])
	y1 := util.MustAtoi(match[3])
	y2 := util.MustAtoi(match[4])

	targetx := make(map[int][]int)
	inertx := make(map[int][]int)

	for v := 1; v <= x2; v++ {
		x := 0
		for dx := v; dx > 0; dx-- {
			x = x + dx
			if x >= x1 && x <= x2 {
				step := v - dx + 1
				if dx == 1 {
					inertx[step] = append(inertx[step], v)
				} else {
					targetx[step] = append(targetx[step], v)
				}
			}
		}
	}

	result := 0

	for v := y1; v <= -y1; v++ {
		y := 0
		elidgiblex := make(set)
		for dy := v; y >= y1; dy-- {
			y = y + dy
			if y >= y1 && y <= y2 {
				step := v - dy + 1
				elidgiblex.Add(targetx[step])
				elidgiblex.AddInert(inertx, step)
			}
		}
		result += len(elidgiblex)
	}

	return fmt.Sprint(result)
}
