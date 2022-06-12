package day16

import (
	"aoc2021/util"
	"encoding/hex"
	"fmt"
	"strings"
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
		//fmt.Println(result)
		return result
	case 1:
		result := 1
		for _, v := range c {
			result *= v
		}
		//fmt.Println(result)
		return result
	case 2:
		result := c[0]
		for _, v := range c {
			result = util.Min(result, v)
		}
		//fmt.Println(result)
		return result
	case 3:
		result := c[0]
		for _, v := range c {
			result = util.Max(result, v)
		}
		//fmt.Println(result)
		return result
	case 5:
		if c[0] > c[1] {
			//fmt.Println(1)
			return 1
		} else {
			//fmt.Println(0)
			return 0
		}
	case 6:
		if c[0] < c[1] {
			//fmt.Println(1)
			return 1
		} else {
			//fmt.Println(0)
			return 0
		}
	case 7:
		if c[0] == c[1] {
			//fmt.Println(1)
			return 1
		} else {
			///fmt.Println(0)
			return 0
		}
	}
	panic("unexpected code path")
}

func Part2(inputFile string) string {
	lines := util.ReadInput(inputFile)
	bytes, _ := hex.DecodeString(lines[0])
	var sb strings.Builder
	for _, b := range bytes {
		fmt.Fprintf(&sb, "%08b", b)
	}
	input := sb.String()
	v, _ := parse(input)
	//fmt.Println(reminder)
	//fmt.Println(v.string())
	return fmt.Sprint(v.getValue())
}
