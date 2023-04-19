package day19

import (
	"aoc2021/util"
	"strings"
)

type point struct {
	x, y, z int
}

const (
	scannerThreashold = 12                                              // beacons to overlap
	distanceThreshold = scannerThreashold * (scannerThreashold - 1) / 2 // number of pairwise distances between `scannerThreashold` beacons
)

// https://stackoverflow.com/a/58471362/284111
// There are 24 possible rotations of a kube. If we start from a certain rotation and then
// apply ones from this array we will cycle through all 24 of them
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

func (p point) dist(other point) int {
	return (p.x-other.x)*(p.x-other.x) + (p.y-other.y)*(p.y-other.y) + (p.z-other.z)*(p.z-other.z)
}

func (p point) manhatten(other point) int {
	return util.Abs((p.x - other.x)) + util.Abs((p.y - other.y)) + util.Abs((p.z - other.z))
}

func rotateSlice(pp []point, rotate func(p point) point) (result []point) {
	for _, p := range pp {
		result = append(result, rotate(p))
	}
	return
}

func checkAlignment(other []point, current []point) (result []point, offset point) {
	counters := make(map[point]int)
	maxCounter := 0
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
	if maxCounter < scannerThreashold {
		return
	}
	for _, i := range current {
		result = append(result, i.diff(offset))
	}
	return
}

func getDistances(pp []point) (result map[int]int) {
	result = make(map[int]int)
	for i := 0; i < len(pp); i++ {
		for j := i + 1; j < len(pp); j++ {
			result[pp[i].dist(pp[j])]++
		}
	}
	return
}

func checkDistances(other map[int]int, current map[int]int) bool {
	count := 0
	for i, v := range current {
		if j, ok := other[i]; ok {
			count += util.Min(j, v)
			if count >= distanceThreshold {
				return true
			}
		}
	}
	return false
}

type scanner struct {
	index  int
	points []point
	offset point
}

func getNormalized(inputFile string) []scanner {
	lines := util.ReadInput(inputFile)
	scanners := []scanner{}
	for _, l := range lines {
		if strings.HasPrefix(l, "---") {
			scanners = append(scanners, scanner{len(scanners), []point{}, point{}})
			continue
		}
		if len(l) == 0 {
			continue
		}
		coords := util.AtoiSlice(strings.Split(l, ","))
		scanners[len(scanners)-1].points = append(scanners[len(scanners)-1].points, point{coords[0], coords[1], coords[2]})
	}

	distances := []map[int]int{}
	for _, s := range scanners {
		distances = append(distances, getDistances(s.points))
	}

	normalized := []scanner{scanners[0]}
	scanners = scanners[1:]
	for len(scanners) > 0 {
		for j := len(scanners) - 1; j >= 0; j-- {
			for _, other := range normalized {
				// This check is not strictly nescessary but it gives about x20 speed up
				// If not enough distances match up across the two scanners there is no point bothering with rotations and alignment checks
				if !checkDistances(distances[scanners[j].index], distances[other.index]) {
					continue
				}
				done := false
				currentOrientation := scanners[j].points
				for _, rotation := range rotations {
					shifted, offset := checkAlignment(other.points, currentOrientation)
					if shifted != nil {
						done = true
						normalized = append(normalized, scanner{scanners[j].index, shifted, offset})
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

	return normalized
}
