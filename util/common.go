package util

import (
	"log"
	"strconv"
	"strings"
)

func Abs[T int | int32 | int64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Min[T int | int32 | int64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T int | int32 | int64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MustAtoi(s string) int {
	if i, err := strconv.ParseInt(strings.Trim(s, " "), 10, 32); err != nil {
		log.Fatalf("cannot convert string '%s' to a number: %v", s, err)
	} else {
		return int(i)
	}
	panic("unexpected code path")
}

func GetPlot(inputFile string) [][]int {
	lines := ReadInput(inputFile)
	var data = make([][]int, len(lines))
	for i, l := range lines {
		data[i] = Atoi(strings.Split(l, ""))
	}
	return data
}
