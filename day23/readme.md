## --- Day 23: Amphipod ---

I feel a bit embarrassed about this one, because while it felt easy enough, I could not get the runtime down under 100ms (which was my self-imposed goal). I used the Dijkstra search with priority queue, unlike Day 15 I could not avoid the priority queue here and that's fine. And yet my solution ran for about 4 seconds, that was too bad. I started looking at other people's solution, wondering what the trick is, but there were no trick: turned out I chose my data structures very badly, and that made all the difference. I had to rewrite half of the code to use the data structures I saw elsewhere, and that got me where I wanted.

Initially I did it like this:

```go
// #############
// #123456789AB#
// ###.#.#.#.###
//   #.#.#.#.#
//   #########
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

type state struct {
	floorNodeToAmphipod map[floorNode]amphipod
	amphipodToFloorNode map[amphipod]floorNode
	hash                string
}
```

The `state` would go on the Queue, and I needed to have a way to index it, that is I needed to be able to tell from the state if it was already visited or not. Since states are created by transitions from other states, obviously I cannot mark them as visited, I needed to compare, and that's why I used a hash function for them, that is the only way in golang to make them keys in a map. Much better structure is:

```go
// 01 2 3 4 56 hallway
//   0 1 2 3   room
//   0 1 2 3   room
//   0 1 2 3   room
//   0 1 2 3   room
type state struct {
	hallway [7]rune    // out 11 actual hallway positions amphipod can only stay in 7, since cannot stay in front of the 4 rooms
	room    [4][4]rune // 4 rooms, 4 slots  each, for Part 1 only [4][2]rune are used which is 4 rooms, 2 slots each
}
```

Here we have a "flat" structure which can easily be a map key.

All the problems in coding the solution also come from unusual data structure. For Dijkstra we need a way for a state to generate adjacent states. This is roughly how it works:

First, check each hallway spot, if there is an amphipod there, check if it can move to the destination room and settle there. According to the puzzle description, if an amphipod cannot settle, it will not move.

Second, check each room from closer to hallway to deeper inside room slots. We only need to check the first amphipod in each room found this way, since the rest of them will be blocked by this one and unable to move. If this, and all the other amphipods in the room are already in the right place the amphipod won't move. Otherwise we check if there is a direct path to the destination room, where the amphipod can settle for good. If there is, that's a move, and we can skip checking any hallway moves for this amphipod, because it cannot be any more energy efficient. If there is no such a move, then we check each of the hallway spot and see if the amphipod can move into each of them.

In the end I had to battle a bit with golang, since figuring out how to make the solution for parts 1 & 2 generic was not obvious. I saw people just copying the entire thing and changing `room [4][2]rune` to `room [4][4]rune`. I tried a few approaches and ended up just not using the two unused slots in that array for Part 1.
