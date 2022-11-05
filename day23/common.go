package day23

import (
	"aoc2021/util"
	"container/heap"
)

var roomSlots int // Part 1 has 2 slots in each room and Part 2 has 4
var final state   // This is the final amphipod arrangement we are aiming to achieve

// Data structure based on https://github.com/devries/advent_of_code_2021/blob/main/day23_p1/main.go
type state struct {
	hallway [7]rune    // out 11 actual hallway positions amphipod can only stay in 7, since cannot stay in front of the 4 rooms
	room    [4][4]rune // 4 rooms, 4 slots  each, for Part 1 only [4][2]rune are used which is 4 rooms, 2 slots each
}

var weights = map[rune]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}

// The next two line are used to calculate horisonal distances between two amphipod positions
// translating index in state into x coorddiate
var hallwayToX = []int{1, 2, 4, 6, 8, 10, 11}
var roomToX = []int{3, 5, 7, 9}

// This is used to check if all hallway positions between amphipod start and end positions do not contain another amphipod
// e.g for halfway position 0 and room 2 we need to check all the hallway spaces from 0 to hallwayToRoom[0][2], which is 3
// that is we are checking positions 1, 2 and 3
var hallwayToRoom = [7][4]int{
	{1, 2, 3, 4},
	{1, 2, 3, 4},
	{2, 2, 3, 4},
	{2, 3, 3, 4},
	{2, 3, 4, 4},
	{2, 3, 4, 5},
	{2, 3, 4, 5},
}

func getDistance(hallway, room, roomSlot int) int {
	return util.Abs(hallwayToX[hallway]-roomToX[room]) + roomSlot + 1
}

func getRoomToRoomDistance(room, roomSlot, otherRoom, otherroomSlot int) int {
	return util.Abs(room-otherRoom)*2 + roomSlot + 1 + otherroomSlot + 1
}

// Room is ready if it is either empty or contains
// amphipods of only the correct destination type (and not full)
// Return the mumber of deepest empty slot to occoupy
func (s state) isRoomReady(room int) int {
	for i := roomSlots - 1; i >= 0; i-- {
		x := s.room[room][i]
		if x == 0 {
			return i // room is ready for move in and i is the next free slot
		}
		if x != 'A'+rune(room) {
			return -1 // room is not ready for move in
		}
	}
	return -2 // room is already settled
}

// Check if other amphipods in this room blocking specified
// amphipod exit
func (s state) canMoveOut(room, slot int) bool {
	if slot == 0 {
		return true
	}
	for i := slot - 1; i >= 0; i-- {
		if s.room[room][i] != 0 {
			return false
		}
	}
	return true
}

// Check that all hallway postion an amphipod will need to occupy
// while moving from room to room are not already occupied
func (s state) canMoveBetweenRooms(room, otherRoom int) bool {
	if room > otherRoom {
		room, otherRoom = otherRoom, room
	}
	for i := room + 2; i <= otherRoom+1; i++ {
		if s.hallway[i] != 0 {
			return false
		}
	}
	return true
}

// Check that all hallway postion an amphipod will need to occupy
// while moving from hallwat to room or the other way around are not already occupied
func (s state) canMoveBetweenHallwayAndRoom(hallway, room int) bool {
	target := hallwayToRoom[hallway][room]
	if target == hallway { // nothing in between so can move
		return true
	}
	// skip checking statring (moving from hallway to a room)
	// or destination (moving from a room to hallway) position:
	// starting postion is occupied by the amphipod so no point checking
	// and destination postions is already checked by the time this function is called
	if target > hallway {
		hallway++ // moving from left to right
	} else {
		hallway-- // moving from right to left
	}
	from := util.Min(hallway, target)
	to := util.Max(hallway, target)
	for i := from; i <= to; i++ {
		if s.hallway[i] != 0 {
			return false
		}
	}
	return true
}

func (s state) getAllPossibleMoves() (result []Item) {
	for hallway, amphipod := range s.hallway { // for all hallway spots
		if amphipod == 0 {
			continue
		}
		destinationRoom := int(amphipod - 'A') // this is where the amphipod wants to move based on it's kind. It won't move anywhere else
		if destinationSlot := s.isRoomReady(destinationRoom); destinationSlot >= 0 && s.canMoveBetweenHallwayAndRoom(hallway, destinationRoom) {
			new := s
			new.hallway[hallway] = 0
			new.room[destinationRoom][destinationSlot] = amphipod
			result = append(result, Item{new, weights[amphipod] * getDistance(hallway, destinationRoom, destinationSlot)})
		}
	}
	for room, r := range s.room { // for all rooms
		for slot, amphipod := range r { // for all slots in the room
			if amphipod == 0 {
				continue
			}
			if !s.canMoveOut(room, slot) { // if other amphipods blocking the room exit, this amphipod is not moving
				continue
			}
			destinationRoom := int(amphipod - 'A') // this is where the amphipod wants to move based on it's kind.
			if destinationRoom == room {
				if s.isRoomReady(destinationRoom) != -1 {
					continue // we are already settled
				}
			}
			// check if we can move directly to the desination room without stopping at the hallway
			if destinationSlot := s.isRoomReady(destinationRoom); destinationSlot >= 0 && s.canMoveBetweenRooms(room, destinationRoom) {
				new := s
				new.room[room][slot] = 0
				new.room[destinationRoom][destinationSlot] = amphipod
				result = append(result, Item{new, weights[amphipod] * getRoomToRoomDistance(room, slot, destinationRoom, destinationSlot)})
				continue // if we can go strait to the destination, let's not waste time on other possibilities
			}
			// let's check all the possible hallway destination spots
			for targetHallway, h := range s.hallway {
				if h != 0 {
					continue
				}
				if s.canMoveBetweenHallwayAndRoom(targetHallway, room) {
					new := s
					new.room[room][slot] = 0
					new.hallway[targetHallway] = amphipod
					result = append(result, Item{new, weights[amphipod] * getDistance(targetHallway, room, slot)})
				}
			}
		}
	}
	return
}

// Dijkstra's search using a priority queue
// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Using_a_priority_queue
func solve(start state) int {
	dist := map[state]int{start: 0}
	q := make(PriorityQueue, 1)
	q[0] = Item{start, 0}
	heap.Init(&q)
	for {
		u := heap.Pop(&q).(Item)
		if u.value == final {
			return dist[u.value]
		}
		for _, v := range u.value.getAllPossibleMoves() {
			alt := dist[u.value] + v.priority
			if i, ok := dist[v.value]; !ok || alt < i {
				dist[v.value] = alt
				heap.Push(&q, Item{v.value, alt})
			}
		}
	}
}

// parse a single input line into passed room *[4][4]rune
func parseLine(room *[4][4]rune, line int, input []rune) {
	for char := 0; char <= 6; char += 2 {
		room[(char)/2][line] = input[char]
	}
}
