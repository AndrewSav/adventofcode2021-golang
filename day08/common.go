package day08

import (
	"aoc2021/util"
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string { // since all segment values are shuffled randomly we need to sort them so we could compare them confidently
	ss := []rune(s)
	sort.Slice(ss, func(i int, j int) bool { return ss[i] < ss[j] })
	return string(ss)
}

// return true if s contains every characters from ee
// and false if there is a character in ee that is not in s
func containsAll(s string, ee string) bool {
	for _, e := range ee {
		if !strings.Contains(s, string(e)) {
			return false
		}
	}
	return true
}

// return the string from source that contains all the segments from selector
// (or does NOT contain all the segments from the selector if reverseSelector is true)
// removes the returned string from source
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
	result = make(map[string]int)     // this will keep alphabetically sorted the segment string and the corresponding digits to each
	lookup := make(map[int]*[]string) // this lookup maps the segment string length to an array of segment strings of that length
	for _, s := range data {          // fill in the lookup table from all the ten segment strings
		if lookup[len(s)] == nil {
			lookup[len(s)] = new([]string)
		}
		*lookup[len(s)] = append(*lookup[len(s)], sortString(s))
	}
	// first the "easy" digits
	result[(*lookup[2])[0]] = 1 // if the length is 2, it must be 1, since only 1 has 2 segments
	result[(*lookup[4])[0]] = 4 // if the length is 4, it must be 4, since only 4 has 4 segments
	result[(*lookup[3])[0]] = 7 // if the length is 3, it must be 7, since only 7 has 3 segments
	result[(*lookup[7])[0]] = 8 // if the length is 7, it must be 8, since only 8 has 7 segments

	result[eliminate(lookup[5], (*lookup[2])[0])] = 3 // 2,3 and 5 have a segment string of 5. The only one of those that contains all segments from 1 is 3

	seg := (*lookup[4])[0]              // this will represent the top left vertical segment and the middle horisontal segment
	for _, s := range (*lookup[2])[0] { // remove from segment string of 4 all the segments of 1
		seg = strings.ReplaceAll(seg, string(s), "")
	}

	result[eliminate(lookup[5], seg)] = 5 // out of 2 and 5 only 5 has top left vertical segment and the middle horisontal segment
	result[(*lookup[5])[0]] = 2           // the only remaining digit with a segement string of 5 is now 2, once we eliminated 3 and 5

	result[eliminate(lookup[6], (*lookup[2])[0], true)] = 6 // 0,9 and 6 have a segment string of 6. The only one of those that DOES NOT contain all segments from 1 is 6
	result[eliminate(lookup[6], (*lookup[4])[0])] = 9       // out of 0 and 9 only 9 contains all segments from 4
	result[(*lookup[6])[0]] = 0                             // the only remaining digit with a segement string of 6 is now 0, once we eliminated 6 and 9
	return
}

type addNextDigit func(acc *int, v int)

func solve(inputFile string, part addNextDigit) string {
	lines := util.ReadInput(inputFile)
	count := 0
	for _, l := range lines {
		parts := strings.Split(l, "|")
		digits := strings.Split(strings.Trim(parts[0], " "), " ")
		stringToDigit := detect(digits)
		display := strings.Split(strings.Trim(parts[1], " "), " ")
		acc := 0
		for _, n := range display {
			part(&acc, stringToDigit[sortString(n)])
		}
		count += acc
	}
	return fmt.Sprint(count)
}
