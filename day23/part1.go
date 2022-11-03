package day23

import (
	"aoc2021/util"
	"fmt"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

// #############
// #123456789AB#
// ###.#.#.#.###
//
//	#.#.#.#.#
//	#########
//
// where A - 10 and B - 11 for x in floorNode
type floorNode struct {
	x int // 1,2,4,6,8,10,11 - hallway, 3,5,7,9 - rooms
	y int // 0 - hallway, 1 - room front, 2 - room rear
}

type amphipod struct {
	kind   rune
	number int
}

var weights = map[rune]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}
var destinations = map[rune]int{'A': 3, 'B': 5, 'C': 7, 'D': 9}

var rooms = []floorNode{
	{x: 3, y: 1},
	{x: 3, y: 2},
	{x: 5, y: 1},
	{x: 5, y: 2},
	{x: 7, y: 1},
	{x: 7, y: 2},
	{x: 9, y: 1},
	{x: 9, y: 2},
}

var hallways = []floorNode{
	{x: 1},
	{x: 2},
	{x: 4},
	{x: 6},
	{x: 8},
	{x: 10},
	{x: 11},
}

func makeHallwayFirst(first, second floorNode) (floorNode, floorNode) {
	if first.y == 0 {
		return first, second
	} else {
		return second, first
	}
}

func between(start, finish floorNode) (result []floorNode) {
	result = append(result, finish)
	hallway, room := makeHallwayFirst(start, finish)
	if room.y == 2 {
		result = append(result, floorNode{x: room.x, y: 1})
	}
	from := util.Min(hallway.x, room.x)
	to := util.Max(hallway.x, room.x)
	for i := from + 1; i < to; i++ {
		if i != 3 && i != 5 && i != 7 && i != 9 {
			result = append(result, floorNode{x: i, y: 0})
		}
	}
	return
}

type moveInfo struct {
	amphipod amphipod
	target   floorNode
	length   int
}

func getDistance(from, to floorNode) int {
	return util.Abs(from.x-to.x) + util.Abs(from.y-to.y)
}

func (s *state) getPossibleMoves(a amphipod) (result []moveInfo) {
	var targets []floorNode
	start := s.amphipodToFloorNode[a]
	if start.y == 0 {
		targets = rooms
	} else {
		targets = hallways
	}
	for _, t := range targets {
		if t.y == 1 { // cannot move into front room if rear room is not occupied, or occupied by different type
			if neighbour, ok := s.floorNodeToAmphipod[floorNode{x: t.x, y: 2}]; !ok || a.kind != neighbour.kind {
				continue
			}
		}
		if t.y != 0 { // cannot move into a room which is not the destination room
			if destinations[a.kind] != t.x {
				continue
			}
		} else {
			if start.x == destinations[a.kind] {
				if start.y == 2 || s.floorNodeToAmphipod[floorNode{x: start.x, y: 2}].kind == a.kind {
					continue
				}
			}
		}
		bb := between(start, t)
		blocked := false
		for _, b := range bb {
			if _, ok := s.floorNodeToAmphipod[b]; ok {
				blocked = true
				break
			}
		}
		if !blocked {
			result = append(result, moveInfo{target: t, length: getDistance(start, t), amphipod: a})
		}
	}
	return
}

func (s *state) getAllPossibleMoves() (result []moveInfo) {
	for a := range s.amphipodToFloorNode {
		result = append(result, s.getPossibleMoves(a)...)
	}
	//fmt.Printf("moves: %d\n", len(result))
	return
}

func (s *state) move(m moveInfo) (result *state) {
	energy := m.length * weights[m.amphipod.kind]
	previousCumulativeEnergy := 0

	if len(s.steps) > 0 {
		previousCumulativeEnergy = s.steps[len(s.steps)-1].cumulativeEnergy
	}

	stp := step{
		moveInfo:         m,
		source:           s.amphipodToFloorNode[m.amphipod],
		energy:           energy,
		cumulativeEnergy: energy + previousCumulativeEnergy,
	}
	result = &state{
		steps:               make([]step, len(s.steps)),
		amphipodToFloorNode: make(map[amphipod]floorNode),
		floorNodeToAmphipod: make(map[floorNode]amphipod),
	}
	copy(result.steps, s.steps)
	result.steps = append(result.steps, stp)
	//TODO: review
	start := s.amphipodToFloorNode[m.amphipod]
	for k, v := range s.floorNodeToAmphipod {
		if k != start {
			result.floorNodeToAmphipod[k] = v
		}
	}
	result.floorNodeToAmphipod[m.target] = m.amphipod
	for k, v := range s.amphipodToFloorNode {
		result.amphipodToFloorNode[k] = v
	}
	result.amphipodToFloorNode[m.amphipod] = m.target
	return
}

func (s *state) areWeThereYet() bool {
	for k, v := range s.amphipodToFloorNode {
		if destinations[k.kind] != v.x {
			return false
		}
	}
	return true
}

func (s *state) dump() {
	fmt.Println("#############")
	fmt.Print("#")
	for i := 1; i <= 11; i++ {
		if a, ok := s.floorNodeToAmphipod[floorNode{i, 0}]; ok {
			fmt.Printf("%c", a.kind)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("#")
	fmt.Print("###")
	for i := 3; i <= 9; i++ {
		if i%2 == 0 {
			fmt.Print("#")
			continue
		}
		if a, ok := s.floorNodeToAmphipod[floorNode{i, 1}]; ok {
			fmt.Printf("%c", a.kind)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("###")
	fmt.Print("  #")
	for i := 3; i <= 9; i++ {
		if i%2 == 0 {
			fmt.Print("#")
			continue
		}
		if a, ok := s.floorNodeToAmphipod[floorNode{i, 2}]; ok {
			fmt.Printf("%c", a.kind)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("#")
	fmt.Println("  #########")

}

type step struct {
	moveInfo
	source           floorNode
	energy           int
	cumulativeEnergy int
}

type state struct {
	steps               []step
	floorNodeToAmphipod map[floorNode]amphipod
	amphipodToFloorNode map[amphipod]floorNode
}

//var floorNodeToAmphipod = map[floorNode]amphipod{}
//var amphipodToFloorNode = map[amphipod]floorNode{}

func bla(start *state) {

	best := MaxInt

	stack := [](*state){start}
	for len(stack) > 0 {
		//fmt.Printf("stack: %d\n", len(stack))
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//fmt.Printf("steps: %d\n", len(s.steps))
		//s.dump()
		moves := s.getAllPossibleMoves()
		if len(moves) == 0 {
			//fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		}
		for _, move := range moves {
			next := s.move(move)
			if next.areWeThereYet() {
				en := next.steps[len(next.steps)-1].cumulativeEnergy
				min := util.Min(best, en)
				if min < best {
					fmt.Printf("found %d, steps %d\n", en, len(next.steps))
					best = min
					temp := start
					temp.dump()
					for _, z := range next.steps {
						temp = temp.move(z.moveInfo)
						temp.dump()
						fmt.Printf("%d %d\n", z.energy, z.cumulativeEnergy)
					}
				}

			} else {
				stack = append(stack, next)
			}
		}
	}
}

func Part1(inputFile string) string {

	a1 := amphipod{'A', 1}
	a2 := amphipod{'A', 2}
	b1 := amphipod{'B', 1}
	b2 := amphipod{'B', 2}
	c1 := amphipod{'C', 1}
	c2 := amphipod{'C', 2}
	d1 := amphipod{'D', 1}
	d2 := amphipod{'D', 2}

	s := &state{
		amphipodToFloorNode: make(map[amphipod]floorNode),
		floorNodeToAmphipod: make(map[floorNode]amphipod),
	}

	s.amphipodToFloorNode[a1] = floorNode{5, 1}
	s.amphipodToFloorNode[a2] = floorNode{5, 2}
	s.amphipodToFloorNode[b1] = floorNode{3, 1}
	s.amphipodToFloorNode[b2] = floorNode{7, 1}
	s.amphipodToFloorNode[c1] = floorNode{9, 1}
	s.amphipodToFloorNode[c2] = floorNode{9, 2}
	s.amphipodToFloorNode[d1] = floorNode{3, 2}
	s.amphipodToFloorNode[d2] = floorNode{7, 2}

	//s.amphipodToFloorNode[a1] = floorNode{3, 2}
	//s.amphipodToFloorNode[a2] = floorNode{9, 2}
	//s.amphipodToFloorNode[b1] = floorNode{3, 1}
	//s.amphipodToFloorNode[b2] = floorNode{7, 1}
	//s.amphipodToFloorNode[c1] = floorNode{5, 1}
	//s.amphipodToFloorNode[c2] = floorNode{7, 2}
	//s.amphipodToFloorNode[d1] = floorNode{5, 2}
	//s.amphipodToFloorNode[d2] = floorNode{9, 1}
	for k, v := range s.amphipodToFloorNode {
		s.floorNodeToAmphipod[v] = k
	}

	bla(s)

	return inputFile
}
