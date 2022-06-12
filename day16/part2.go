package day16

import (
	"aoc2021/util"
	"fmt"
)

func (l *literal) getValue() int {
	return l.value
}

func (o *operator) getValue() int {
	var c []int
	for _, s := range o.subPackets {
		c = append(c, s.getValue())
	}
	switch o.typeId {
	case 0:
		result := 0
		for _, v := range c {
			result += v
		}
		return result
	case 1:
		result := 1
		for _, v := range c {
			result *= v
		}
		return result
	case 2:
		result := c[0]
		for _, v := range c {
			result = util.Min(result, v)
		}
		return result
	case 3:
		result := c[0]
		for _, v := range c {
			result = util.Max(result, v)
		}
		return result
	case 5:
		if c[0] > c[1] {
			return 1
		} else {
			return 0
		}
	case 6:
		if c[0] < c[1] {
			return 1
		} else {
			return 0
		}
	case 7:
		if c[0] == c[1] {
			return 1
		} else {
			return 0
		}
	}
	panic("unexpected code path")
}

func Part2(inputFile string) string {
	return fmt.Sprint(loadAndParse(inputFile).getValue())
}
