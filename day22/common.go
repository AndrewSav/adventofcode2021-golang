package day22

import (
	"aoc2021/util"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// https://stackoverflow.com/a/53587770/284111
func findNamedMatches(regex *regexp.Regexp, str string) map[string]string {
	match := regex.FindStringSubmatch(str)

	results := map[string]string{}
	for i, name := range match {
		results[regex.SubexpNames()[i]] = name
	}
	return results
}

type cube struct {
	xmin  int
	xmax  int
	ymin  int
	ymax  int
	zmin  int
	zmax  int
	on    bool
	index string
}

func loadCube(m map[string]string, index string) *cube {
	return &cube{
		xmin:  util.MustAtoi(m["xmin"]),
		xmax:  util.MustAtoi(m["xmax"]),
		ymin:  util.MustAtoi(m["ymin"]),
		ymax:  util.MustAtoi(m["ymax"]),
		zmin:  util.MustAtoi(m["zmin"]),
		zmax:  util.MustAtoi(m["zmax"]),
		on:    m["mode"] == "on",
		index: index,
	}
}

func getX(c *cube) (int, int) {
	return c.xmin, c.xmax
}
func getY(c *cube) (int, int) {
	return c.ymin, c.ymax
}
func getZ(c *cube) (int, int) {
	return c.zmin, c.zmax
}

type memo struct {
	selection string
	dd        int
}

var memoMap = map[memo]int64{}

func getStopPoints(selection []*cube, getMinMax func(c *cube) (int, int)) (stopPoints []int) {
	m := map[int]struct{}{}
	for _, c := range selection {
		min, max := getMinMax(c)
		m[min] = struct{}{}
		m[max+1] = struct{}{}
	}
	for k := range m {
		stopPoints = append(stopPoints, k)
	}
	sort.Ints(stopPoints)
	return
}

func getSelectionKey(selection []*cube) string {
	var b strings.Builder
	for _, v := range selection {
		b.WriteString(v.index + "|")
	}
	return b.String()
}

// https://work.njae.me.uk/2021/12/29/advent-of-code-2021-day-22/
func sweep(selection []*cube, getMinMaxs [](func(c *cube) (int, int))) int64 {
	memoKey := memo{getSelectionKey(selection), len(getMinMaxs)}
	if c, ok := memoMap[memoKey]; ok {
		return c
	}

	getMinMax := getMinMaxs[0]
	if len(getMinMaxs) > 0 {
		getMinMaxs = getMinMaxs[1:]
	}

	var (
		count              int64
		previousOn         bool
		previousMultiplier int64
		stopPoints         = getStopPoints(selection, getMinMax)
	)

	for i, v := range stopPoints {
		newSelection := []*cube{}
		for _, c := range selection {
			min, max := getMinMax(c)
			if v >= min && v <= max {
				newSelection = append(newSelection, c)
			}
		}
		multiplier := int64(1)
		if len(getMinMaxs) > 0 {
			multiplier = sweep(newSelection, getMinMaxs)
		}
		on := len(newSelection) > 0 && newSelection[len(newSelection)-1].on || len(getMinMaxs) > 0
		if on && i == len(stopPoints)-1 {
			count += multiplier
		}
		if previousOn && i > 0 {
			count += previousMultiplier * int64(v-stopPoints[i-1])
		}
		previousOn = on
		previousMultiplier = multiplier
	}
	memoMap[memoKey] = count
	return count
}

func solve(inputFile string, part1 bool) string {
	data := util.ReadInput(inputFile)
	r := regexp.MustCompile(`^(?P<mode>on|off) x=(?P<xmin>-?\d+)\.\.(?P<xmax>-?\d+),y=(?P<ymin>-?\d+)\.\.(?P<ymax>-?\d+),z=(?P<zmin>-?\d+)\.\.(?P<zmax>-?\d+)$`)
	cubes := []*cube{}
	for i, s := range data {
		m := findNamedMatches(r, s)
		c := loadCube(m, fmt.Sprint(i))
		if part1 && (c.xmin < -50 || c.xmax > 50 || c.ymin < -50 || c.ymax > 50 || c.zmin < -50 || c.zmax > 50) {
			continue
		}
		cubes = append(cubes, c)
	}

	count := sweep(cubes, [](func(c *cube) (int, int)){getX, getY, getZ})

	return fmt.Sprint(count)
}
