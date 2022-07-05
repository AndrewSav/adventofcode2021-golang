package day19

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

type point struct {
	x, y, z int
}

// https://stackoverflow.com/a/58471362/284111
var rotations = [...](func(p point) point){
	point.roll, point.cw, point.cw, point.cw,
	point.roll, point.acw, point.acw, point.acw,
	point.roll, point.cw, point.cw, point.cw,
	point.roll, point.acw, point.acw, point.acw,
	point.roll, point.cw, point.cw, point.cw,
	point.roll, point.acw, point.acw, point.acw,
}

func (p point) roll() point { // the face facing you is now rotated downwards
	return point{p.x, p.z, -p.y}
}

func (p point) cw() point { // the face facing you is turned clockwise, but still facing you
	return point{-p.z, p.y, p.x}
}

func (p point) acw() point { //the face facing you is turned anti-clockwise, but still facing you
	return point{p.z, p.y, -p.x}
}

func (p point) diff(other point) point {
	return point{p.x - other.x, p.y - other.y, p.z - other.z}
}

func rotateSlice(pp []point, rotate func(p point) point) (result []point) {
	for _, p := range pp {
		result = append(result, rotate(p))
	}
	return
}

func checkTweleve(other []point, current []point) (result []point) {
	counters := make(map[point]int)
	maxCounter := 0
	var offset point
	for _, i := range current {
		for _, j := range other {
			diff := i.diff(j)
			counters[diff]++
			if counters[diff] > maxCounter {
				maxCounter = counters[diff]
				offset = diff
			}
		}
	}
	if maxCounter < 12 {
		return
	}
	for _, i := range current {
		result = append(result, i.diff(offset))
	}
	return
}

func Part1(inputFile string) string {
	lines := util.ReadInput(inputFile)
	scanners := [][]point{}
	for _, l := range lines {
		if strings.HasPrefix(l, "---") {
			scanners = append(scanners, []point{})
			continue
		}
		if len(l) == 0 {
			continue
		}
		coords := util.AtoiSlice(strings.Split(l, ","))
		scanners[len(scanners)-1] = append(scanners[len(scanners)-1], point{coords[0], coords[1], coords[2]})
	}

	normalized := [][]point{scanners[0]}
	scanners = scanners[1:]
	//count := 0
	for len(scanners) > 0 {
		for j := len(scanners) - 1; j >= 0; j-- {
			for _, other := range normalized {
				done := false
				//count++
				currentOrientation := scanners[j]
				for _, rotation := range rotations {
					shifted := checkTweleve(other, currentOrientation)
					if shifted != nil {
						done = true
						normalized = append(normalized, shifted)
						scanners[j] = scanners[len(scanners)-1]
						scanners = scanners[:len(scanners)-1]
						break
					}
					currentOrientation = rotateSlice(currentOrientation, rotation)
				}
				if done {
					break
				}
			}
		}
	}

	unique := make(map[point]struct{})
	for _, s := range normalized {
		for _, p := range s {
			unique[p] = struct{}{}
		}
	}
	//fmt.Println(count)
	return fmt.Sprint(len(unique))
}
