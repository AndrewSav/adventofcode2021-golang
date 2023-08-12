package day16

import (
	"fmt"
)

func (l *literal) getValue() int64 {
	return int64(l.value)
}

func (o *operator) getValue() int64 {
	var c []int64
	for _, s := range o.subPackets {
		c = append(c, s.getValue())
	}
	switch o.typeId {
	case 0:
		result := int64(0)
		for _, v := range c {
			result += v
		}
		return result
	case 1:
		result := int64(1)
		for _, v := range c {
			result *= v
		}
		return result
	case 2:
		result := c[0]
		for _, v := range c {
			result = min(result, v)
		}
		return result
	case 3:
		result := c[0]
		for _, v := range c {
			result = max(result, v)
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
