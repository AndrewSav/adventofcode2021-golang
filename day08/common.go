package day08

import (
	"aoc2021/util"
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	ss := []rune(s)
	sort.Slice(ss, func(i int, j int) bool { return ss[i] < ss[j] })
	return string(ss)
}

func containsAll(s string, ee string) bool {
	for _, e := range ee {
		if !strings.Contains(s, string(e)) {
			return false
		}
	}
	return true
}

func eliminate(source *[]string, selector string, reverseSelector ...bool) string {
	reverse := false
	if len(reverseSelector) > 0 {
		reverse = reverseSelector[0]
	}
	for i, s := range *source {
		if containsAll(s, selector) != reverse {
			(*source)[i] = (*source)[len(*source)-1]
			*source = (*source)[:len(*source)-1]
			return s
		}
	}
	panic("unexpected code path")
}

func detect(data []string) (result map[string]int) {
	result = make(map[string]int)
	lookup := make(map[int]*[]string)
	for _, s := range data {
		if lookup[len(s)] == nil {
			lookup[len(s)] = new([]string)
		}
		*lookup[len(s)] = append(*lookup[len(s)], sortString(s))
	}
	result[(*lookup[2])[0]] = 1
	result[(*lookup[4])[0]] = 4
	result[(*lookup[3])[0]] = 7
	result[(*lookup[7])[0]] = 8

	result[eliminate(lookup[5], (*lookup[2])[0])] = 3

	seg := (*lookup[4])[0]
	for _, s := range (*lookup[2])[0] {
		seg = strings.ReplaceAll(seg, string(s), "")
	}

	result[eliminate(lookup[5], seg)] = 5
	result[(*lookup[5])[0]] = 2

	result[eliminate(lookup[6], (*lookup[2])[0], true)] = 6
	result[eliminate(lookup[6], (*lookup[4])[0])] = 9
	result[(*lookup[6])[0]] = 0
	return
}

type addNextDigit func(acc *int, v int)

func solve(inputFile string, part addNextDigit) string {
	lines := util.ReadInput(inputFile)
	count := 0
	for _, l := range lines {
		parts := strings.Split(l, "|")
		digits := strings.Split(strings.Trim(parts[0], " "), " ")
		val := detect(digits)
		display := strings.Split(strings.Trim(parts[1], " "), " ")
		acc := 0
		for _, n := range display {
			part(&acc, val[sortString(n)])
		}
		count += acc
	}
	return fmt.Sprint(count)
}
