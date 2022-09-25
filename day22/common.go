package day22

import (
	"aoc2021/util"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

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

type dimension struct {
	getMinMax  func(c *cube) (int, int)
	stopPoints []int
}

type memo struct {
	selection string
	dd        int
}

var memoMap = map[memo]int64{}

func sweep(selection []*cube, dd []dimension) int64 {
	var b strings.Builder
	for _, v := range selection {
		b.WriteString(v.index + "|")
	}
	memoKey := memo{b.String(), len(dd)}
	if c, ok := memoMap[memoKey]; ok {
		return c
	}
	d := dd[0]
	if len(dd) > 0 {
		dd = dd[1:]
	}

	var count int64
	var previousOn bool
	var previousMultiplier int64
	stopPoints := d.stopPoints
	for i, v := range stopPoints {
		newSelection := []*cube{}
		for _, c := range selection {
			min, max := d.getMinMax(c)
			if v >= min && v <= max {
				newSelection = append(newSelection, c)
			}
		}
		multiplier := int64(1)
		if len(dd) > 0 {
			multiplier = sweep(newSelection, dd)
		}
		on := len(newSelection) > 0 && newSelection[len(newSelection)-1].on || len(dd) > 0
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
	var xx, yy, zz []int
	mxx := map[int]struct{}{}
	myy := map[int]struct{}{}
	mzz := map[int]struct{}{}
	for i, s := range data {
		m := findNamedMatches(r, s)
		c := loadCube(m, fmt.Sprint(i))
		if part1 && (c.xmin < -50 || c.xmax > 50 || c.ymin < -50 || c.ymax > 50 || c.zmin < -50 || c.zmax > 50) {
			continue
		}
		mxx[c.xmin] = struct{}{}
		mxx[c.xmax+1] = struct{}{}
		myy[c.ymin] = struct{}{}
		myy[c.ymax+1] = struct{}{}
		mzz[c.zmin] = struct{}{}
		mzz[c.zmax+1] = struct{}{}
		cubes = append(cubes, c)
	}

	for k := range mxx {
		xx = append(xx, k)
	}
	for k := range myy {
		yy = append(yy, k)
	}
	for k := range mzz {
		zz = append(zz, k)
	}

	sort.Ints(xx)
	sort.Ints(yy)
	sort.Ints(zz)

	count := sweep(cubes, []dimension{{getX, xx}, {getY, yy}, {getZ, zz}})

	return fmt.Sprint(count)
}
