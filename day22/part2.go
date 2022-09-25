package day22

import (
	"aoc2021/util"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

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

var mm = map[memo]int64{}

func bla(allCubes map[*cube]int, selection []*cube, dd []dimension) (count int64) {
	var b strings.Builder
	for _, v := range selection {
		fmt.Fprintf(&b, "%s|", allCubes[v])
	}
	mkey := memo{b.String(), len(dd)}
	if c, ok := mm[mkey]; ok {
		count = c
		return
	}

	d := dd[0]
	if len(dd) > 0 {
		dd = dd[1:]
	}
	var previousOn bool
	var previousMultiplier int64
	for i, v := range d.stopPoints {
		newSelection := []*cube{}
		for _, c := range selection {
			min, max := d.getMinMax(c)
			if v >= min && v <= max {
				newSelection = append(newSelection, c)
			}
		}
		multiplier := int64(1)
		if len(dd) > 0 {
			multiplier = bla(allCubes, newSelection, dd)
		}
		on := len(newSelection) > 0 && newSelection[len(newSelection)-1].on || len(dd) > 0
		if on && i == len(d.stopPoints)-1 {
			count += multiplier
		}
		if previousOn && i > 0 {
			count += previousMultiplier * int64(v-d.stopPoints[i-1])
		}
		previousOn = on
		previousMultiplier = multiplier
	}
	mm[mkey] = count
	return
}

func Part2(inputFile string) string {
	data := util.ReadInput(inputFile)
	r := regexp.MustCompile(`^(?P<mode>on|off) x=(?P<xmin>-?\d+)\.\.(?P<xmax>-?\d+),y=(?P<ymin>-?\d+)\.\.(?P<ymax>-?\d+),z=(?P<zmin>-?\d+)\.\.(?P<zmax>-?\d+)$`)
	cubes := []*cube{}
	var xx, yy, zz []int
	mxx := map[int]struct{}{}
	myy := map[int]struct{}{}
	mzz := map[int]struct{}{}
	for _, s := range data {
		m := findNamedMatches(r, s)
		c := loadCube(m)
		//if c.xmin < -50 || c.xmax > 50 || c.ymin < -50 || c.ymax > 50 || c.zmin < -50 || c.zmax > 50 {
		//	continue
		//}
		mxx[c.xmin] = struct{}{}
		mxx[c.xmax] = struct{}{}
		mxx[c.xmin+1] = struct{}{}
		mxx[c.xmax-1] = struct{}{}
		mxx[c.xmin-1] = struct{}{}
		mxx[c.xmax+1] = struct{}{}
		myy[c.ymin] = struct{}{}
		myy[c.ymax] = struct{}{}
		myy[c.ymin+1] = struct{}{}
		myy[c.ymax-1] = struct{}{}
		myy[c.ymin-1] = struct{}{}
		myy[c.ymax+1] = struct{}{}
		mzz[c.zmin] = struct{}{}
		mzz[c.zmax] = struct{}{}
		mzz[c.zmin+1] = struct{}{}
		mzz[c.zmax-1] = struct{}{}
		mzz[c.zmin-1] = struct{}{}
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

	allCubes := map[*cube]int{}
	for i, v := range cubes {
		allCubes[v] = i
	}

	count := bla(allCubes, cubes, []dimension{{getX, xx}, {getY, yy}, {getZ, zz}})

	return fmt.Sprint(count)
}
