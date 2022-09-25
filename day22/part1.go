package day22

import (
	"aoc2021/util"
	"fmt"
	"regexp"
)

func Part1(inputFile string) string {
	data := util.ReadInput(inputFile)
	r := regexp.MustCompile(`^(?P<mode>on|off) x=(?P<xmin>-?\d+)\.\.(?P<xmax>-?\d+),y=(?P<ymin>-?\d+)\.\.(?P<ymax>-?\d+),z=(?P<zmin>-?\d+)\.\.(?P<zmax>-?\d+)$`)
	cubes := [101 * 101 * 101]bool{}
	for _, s := range data {
		m := findNamedMatches(r, s)
		c := loadCube(m)
		if c.xmin < -50 || c.xmax > 50 || c.ymin < -50 || c.ymax > 50 || c.zmin < -50 || c.zmax > 50 {
			continue
		}
		for x := c.xmin; x <= c.xmax; x++ {
			for y := c.ymin; y <= c.ymax; y++ {
				for z := c.zmin; z <= c.zmax; z++ {
					val := m["mode"] == "on"
					i := 101*101*(z+50) + 101*(y+50) + (x + 50)
					cubes[i] = val
				}
			}
		}
	}

	count := 0
	for _, v := range cubes {
		if v {
			count++
		}
	}

	return fmt.Sprint(count)
}
