package day22

import (
	"aoc2021/util"
	"regexp"
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
	xmin int
	xmax int
	ymin int
	ymax int
	zmin int
	zmax int
	on   bool
}

func loadCube(m map[string]string) *cube {
	return &cube{
		xmin: util.MustAtoi(m["xmin"]),
		xmax: util.MustAtoi(m["xmax"]),
		ymin: util.MustAtoi(m["ymin"]),
		ymax: util.MustAtoi(m["ymax"]),
		zmin: util.MustAtoi(m["zmin"]),
		zmax: util.MustAtoi(m["zmax"]),
		on:   m["mode"] == "on",
	}
}
