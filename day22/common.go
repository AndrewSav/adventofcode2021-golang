package day22

import (
	"aoc2021/util"
	"fmt"
	"sort"
	"strings"
)

// represents limits from input for single dimension of an input line
// each line has 3 dimensions: x, y and z
type dimension struct {
	min int
	max int
}

type cuboid struct {
	dimensions [3]dimension
	on         bool   // mode: true if "on", false if "off"
	index      string // this is used for memoisation, it is faster to store it as string so we do not need to convert each time it's used
}

// returns ordered list of points where to stop while sweeping
// those points are where the voxel state (on/off) potentially changes during a sweep
func getStopPoints(selection []*cuboid, dimensionIndex int) (stopPoints []int) {
	m := map[int]struct{}{}
	for _, c := range selection {
		m[c.dimensions[dimensionIndex].min] = struct{}{}
		// note "+1": voxel state won't change on the last point of a dimension (.max)
		// because the cuboid border is *after* that voxel in the sweep direction
		m[c.dimensions[dimensionIndex].max+1] = struct{}{}
	}
	for k := range m { // for weeding out unique values only
		stopPoints = append(stopPoints, k)
	}
	sort.Ints(stopPoints)
	return
}

// for memoisation we need to convert array of pointers to something that we can use in a map key, such as a string
// I tried to use xxhash here instead, it gave about 10% improvement at the cost of introducing 3-rd party dependency
func getSelectionKey(selection []*cuboid, dimensionIndex int) string {
	var b strings.Builder
	for _, v := range selection {
		b.WriteString(v.index + "|")
	}
	b.WriteString(fmt.Sprint(dimensionIndex))
	return b.String()
}

var memoMap = map[string]int64{} // maps sweep function parameters to the function result

// https://work.njae.me.uk/2021/12/29/advent-of-code-2021-day-22/
func sweep(selection []*cuboid, dimensionIndex int) int64 {
	// first we check if we already have a result for the parameters passed
	// and if we do, return it. It runs about 3 times as fast due to memoisation
	memoKey := getSelectionKey(selection, dimensionIndex)
	if c, ok := memoMap[memoKey]; ok {
		return c
	}

	var (
		count      int64 // the count of "on" voxels
		stopPoints = getStopPoints(selection, dimensionIndex)
	)

	// we are looping through intervals between stop points so "-1"
	// this is know as "fencepost issue"
	for i := 0; i < len(stopPoints)-1; i++ {
		stopPoint := stopPoints[i]
		// filter out cuboids that do not include the current stop point
		newSelection := []*cuboid{}
		for _, c := range selection {
			if stopPoint >= c.dimensions[dimensionIndex].min && stopPoint <= c.dimensions[dimensionIndex].max {
				newSelection = append(newSelection, c)
			}
		}

		// the "exit" stop point of a cuboid is just outside of the cuboid so it is possible
		// that it belongs to no cuboids in the current selection. This means we are now on an
		// interval that is not inside of any cuboid
		if len(newSelection) == 0 {
			continue
		}

		intervalLength := int64(stopPoints[i+1] - stopPoint)
		// is this the last dimension to sweep (dot along a line)?
		if dimensionIndex == 0 {
			// last applied cuboid determines if the voxel is on or off
			if newSelection[len(newSelection)-1].on {
				count += intervalLength // just add the number of voxel on the line between the stop points
			}
		} else {
			// get the number of on voxels from the next dimension, and multiply it on the
			// interval length to get the number in this dimension for this interval
			count += sweep(newSelection, dimensionIndex-1) * intervalLength
		}
	}
	memoMap[memoKey] = count // remember results for memoisation
	return count
}

func solve(inputFile string, part1 bool) string {
	data := util.ReadInput(inputFile)
	cuboids := []*cuboid{}
	for i, s := range data {
		var xMin, xMax, yMin, yMax, zMin, zMax int
		tokens := strings.Split(s, " ")
		fmt.Sscanf(tokens[1], "x=%d..%d,y=%d..%d,z=%d..%d", &xMin, &xMax, &yMin, &yMax, &zMin, &zMax)
		c := &cuboid{
			dimensions: [3]dimension{
				{min: zMin, max: zMax},
				{min: yMin, max: yMax},
				{min: xMin, max: xMax},
			},
			on:    tokens[0] == "on",
			index: fmt.Sprint(i),
		}
		if part1 && isOutOfBound(c) {
			continue
		}
		cuboids = append(cuboids, c)
	}
	count := sweep(cuboids, len(cuboids[0].dimensions)-1)
	return fmt.Sprint(count)
}
