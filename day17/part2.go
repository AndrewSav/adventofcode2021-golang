package day17

import (
	"fmt"
)

type set map[int]struct{} // we will accumulate all horizontal velocities that will land us in the zone for a specific vertical velocity

func (s set) add(targetX, inertX map[int][]int, step int) { // add eligible x velocities for the given step
	for _, i := range targetX[step] { // first add all velocities that lands us in the zone x-wise for the given step
		s[i] = struct{}{}
	}
	for k, v := range inertX { // now also add velocities when we became inert on any step after we became inert, inclusive
		if step >= k {
			for _, x := range v {
				s[x] = struct{}{}
			}
		}
	}
}

func Part2(inputFile string) string {
	x1, x2, y1, y2 := getInput(inputFile)

	targetX := make(map[int][]int) // list of initial x velocities that will land us in the zone (x) on a given step
	inertX := make(map[int][]int)  // list of initial x velocities where we come to stop in the zone (x), so we remain there for a given and all subsequent steps

	for v := 1; v <= x2; v++ { // first let's see which x velocities end us up in the zone
		x := 0
		for dx := v; dx > 0; dx-- {
			x = x + dx
			if x >= x1 && x <= x2 {
				step := v - dx
				if dx == 1 {
					inertX[step] = append(inertX[step], v)
				} else {
					targetX[step] = append(targetX[step], v)
				}
			}
		}
	}

	result := 0

	for v := y1; v <= -y1; v++ { // now let's iterate through y velocities, and cross-reference those that land us in the zone with x velocities we collected earlier
		y := 0
		eligibleX := make(set) // for the current y velocity v we will collect x velocities here that lands us in the zone both x and y wise
		for dy := v; y >= y1; dy-- {
			y = y + dy
			if y >= y1 && y <= y2 {
				step := v - dy
				eligibleX.add(targetX, inertX, step)
			}
		}
		result += len(eligibleX) // adding the number of distinct x velocities for current y velocity that lands us in the zone to get the final answer
	}

	return fmt.Sprint(result)
}
