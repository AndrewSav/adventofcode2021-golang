package day23

import (
	"aoc2021/util"
	"container/heap"
	"fmt"
)

type state struct {
	hallway [7]rune
	room    [4][2]rune
}

var weights = map[rune]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}
var hallwayToX = []int{1, 2, 4, 6, 8, 10, 11}
var roomToX = []int{3, 5, 7, 9}

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

func (s *state) isRoomReady(room int) int {
	for i := 1; i >= 0; i-- {
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

func (s *state) canMoveOut(room, slot int) bool {
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

func (s *state) canMoveBetweenRooms(room, otherRoom int) bool {
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

func (s *state) canMoveBetweenHallwayAndRoom(hallway, room int) bool {
	target := hallwayToRoom[hallway][room]
	if target == hallway {
		return true
	}
	if target > hallway {
		hallway++
	} else {
		hallway--
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

func (s *state) getAllPossibleMoves() (result []Item) {
	for hallway, amphipod := range s.hallway {
		if amphipod == 0 {
			continue
		}
		destinationRoom := int(amphipod - 'A')
		if destinationSlot := s.isRoomReady(destinationRoom); destinationSlot >= 0 && s.canMoveBetweenHallwayAndRoom(hallway, destinationRoom) {
			new := *s
			new.hallway[hallway] = 0
			new.room[destinationRoom][destinationSlot] = amphipod
			result = append(result, Item{new, weights[amphipod] * getDistance(hallway, destinationRoom, destinationSlot)})
		}
	}
	for room, r := range s.room {
		for slot, amphipod := range r {
			if amphipod == 0 {
				continue
			}
			if !s.canMoveOut(room, slot) {
				continue
			}
			destinationRoom := int(amphipod - 'A')
			if destinationRoom == room {
				if s.isRoomReady(destinationRoom) != -1 {
					continue // we are already settled
				}
			}
			if destinationSlot := s.isRoomReady(destinationRoom); destinationSlot >= 0 && s.canMoveBetweenRooms(room, destinationRoom) {
				new := *s
				new.room[room][slot] = 0
				new.room[destinationRoom][destinationSlot] = amphipod
				result = append(result, Item{new, weights[amphipod] * getRoomToRoomDistance(room, slot, destinationRoom, destinationSlot)})
				continue // if we can go strait to the destination, let's not waste time on other possibilities
			}
			for targetHallway, h := range s.hallway {
				if h != 0 {
					continue
				}
				if s.canMoveBetweenHallwayAndRoom(targetHallway, room) {
					new := *s
					new.room[room][slot] = 0
					new.hallway[targetHallway] = amphipod
					result = append(result, Item{new, weights[amphipod] * getDistance(targetHallway, room, slot)})
				}
			}
		}
	}
	return
}

func solve(start state) int {
	final := state{room: [4][2]rune{{'A', 'A'}, {'B', 'B'}, {'C', 'C'}, {'D', 'D'}}}
	dist := map[state]int{start: 0}
	q := make(PriorityQueue, 1)
	q[0] = &Item{start, 0}
	heap.Init(&q)
	for {
		u := heap.Pop(&q).(*Item)
		if u.value == final {
			return dist[u.value]
		}
		for _, v := range u.value.getAllPossibleMoves() {
			alt := dist[u.value] + v.priority
			if i, ok := dist[v.value]; !ok || alt < i {
				dist[v.value] = alt
				heap.Push(&q, &Item{v.value, alt})
			}
		}
	}
}

func parseInput(input []string) (result state) {
	for line := 2; line <= 3; line++ {
		l := []rune(input[line])
		for char := 3; char <= 9; char += 2 {
			result.room[(char-3)/2][line-2] = l[char]
		}
	}
	return
}

func Part1(inputFile string) string {
	//defer profile.Start(profile.ProfilePath(".")).Stop()
	return fmt.Sprint(solve(parseInput(util.ReadInput(inputFile))))
}
