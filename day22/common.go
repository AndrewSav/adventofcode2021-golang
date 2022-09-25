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

// represents limits from input for single dimention of an input line
// each line has 3 dimensions: x, y and z
type dimension struct {
	min int
	max int
}

type cuboid struct {
	dimensions [3]dimension
	on         bool   // mode: true if "on", false if "off"
	index      string // this is used for memoisation later
}

func loadCuboid(m map[string]string, index string) *cuboid {
	return &cuboid{
		dimensions: [3]dimension{
			dimension{min: util.MustAtoi(m["zmin"]), max: util.MustAtoi(m["zmax"])},
			dimension{min: util.MustAtoi(m["ymin"]), max: util.MustAtoi(m["ymax"])},
			dimension{min: util.MustAtoi(m["xmin"]), max: util.MustAtoi(m["xmax"])},
		},
		on:    m["mode"] == "on",
		index: index,
	}
}

// returns ordered list of points where to stop while sweeping
// those points are where the voxel state (on/off) potentialy changes during a sweep
func getStopPoints(selection []*cuboid, dimensionIndex int) (stopPoints []int) {
	m := map[int]struct{}{}
	for _, c := range selection {
		m[c.dimensions[dimensionIndex].min] = struct{}{}
		m[c.dimensions[dimensionIndex].max+1] = struct{}{} // note "+1": voxel state won't change on the last point of a dimension (.max)
	}
	for k := range m { // for weeding out unique values only
		stopPoints = append(stopPoints, k)
	}
	sort.Ints(stopPoints)
	return
}

// for memoisation we need to convert array of pointers to something that we can use, such in a map key, such as a sting
func getSelectionKey(selection []*cuboid) string {
	var b strings.Builder
	for _, v := range selection {
		b.WriteString(v.index + "|")
	}
	return b.String()
}

// this represnts the input parameters of the the sweep function for memoisation
type memo struct {
	selectionKey   string
	dimensionIndex int
}

var memoMap = map[memo]int64{} // maps sweep function parameters to the function result

// https://work.njae.me.uk/2021/12/29/advent-of-code-2021-day-22/
func sweep(selection []*cuboid, dimensionIndex int) int64 {
	// first we check if we already have a result for the parameters passed
	// and if we do, return it
	memoKey := memo{getSelectionKey(selection), dimensionIndex}
	if c, ok := memoMap[memoKey]; ok {
		return c
	}

	var (
		count              int64 // the count of "on" voxels
		previousOn         bool  // whether previous stop point voxel was "on"
		previousMultiplier int64 // count of voxel from the next dimension from previous stop point
		stopPoints         = getStopPoints(selection, dimensionIndex)
	)

	for i, v := range stopPoints {
		// filter out cuboids that do not include the current stop point
		newSelection := []*cuboid{}
		for _, c := range selection {
			if v >= c.dimensions[dimensionIndex].min && v <= c.dimensions[dimensionIndex].max {
				newSelection = append(newSelection, c)
			}
		}
		multiplier := int64(1)  // if this is the last dimension
		if dimensionIndex > 0 { // otherwise get number of voxels that are on from the next dimension
			multiplier = sweep(newSelection, dimensionIndex-1)
		}
		// if this is not the last dimension, it should be on so that voxels from the next dimension are counted
		// otherwise the last cuboid applied will tell us if it's off or on
		on := len(newSelection) > 0 && newSelection[len(newSelection)-1].on || dimensionIndex > 0
		// add voxels from the previous stop point till current one (excluding current one - it will be added on the next iteration)
		if previousOn && i > 0 {
			count += previousMultiplier * int64(v-stopPoints[i-1])
		}
		previousOn = on
		previousMultiplier = multiplier
	}
	memoMap[memoKey] = count // remeber results for memoisation
	return count
}

func solve(inputFile string, part1 bool) string {
	data := util.ReadInput(inputFile)
	r := regexp.MustCompile(`^(?P<mode>on|off) x=(?P<xmin>-?\d+)\.\.(?P<xmax>-?\d+),y=(?P<ymin>-?\d+)\.\.(?P<ymax>-?\d+),z=(?P<zmin>-?\d+)\.\.(?P<zmax>-?\d+)$`)
	cuboids := []*cuboid{}
	for i, s := range data {
		m := findNamedMatches(r, s)
		c := loadCuboid(m, fmt.Sprint(i))
		if part1 && isOutOfBound(c) {
			continue
		}
		cuboids = append(cuboids, c)
	}
	count := sweep(cuboids, len(cuboids[0].dimensions)-1)
	return fmt.Sprint(count)
}
