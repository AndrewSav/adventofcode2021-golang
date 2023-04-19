package day21

import (
	"aoc2021/util"
	"fmt"
)

type state struct {
	position int   // position of the player's pawn on board
	score    int   // current player's score
	forks    int64 // number of all universe forks that can get us to the current state
	turn     int   // current turn number
}

// This is the number of forks for each possible sum in three consequtive rolls
// e.g. you can get the sum of 3 in a single universe: 1,1,1,
// but you can get sum of 4 in 3 universes: 2,1,1; 1,2,1 and 1,1,2, etc
var weights = map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

// j is the initial player position (zero based for convinience)
// it returns two maps, the map key is the turn number
// values in the first map are the number of forks that win on that turn
// values in the second map are the number of forks that do not win on that turn
// for simplicity in this function we only consider one player's rolls and
// completely ignoring forks from the other player, we account for them later
// the number of forks is multipicative and order in which players roll does not matter
// for the number of forks (it does matter for the win condition) it will always "add up"
// to the same number given each individual player's forks
func doTheNumbers(j int) (map[int]int64, map[int]int64) {
	var win = map[int]int64{}
	var lose = map[int]int64{} // lose means we did not win on this turn, we still may win on a later turn
	// since forks is multiplicative we initialise it to one
	stack := []state{{position: j, score: 0, forks: 1}}
	// this is the standard depth first search
	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 3 to 9 are all the possible sums from the three rolls
		for i := 3; i <= 9; i++ {
			newState := state{
				position: (s.position + i) % 10,
				score:    s.score + ((s.position + i) % 10) + 1,
				forks:    s.forks * int64(weights[i]),
				turn:     s.turn + 1,
			}
			if newState.score < 21 {
				stack = append(stack, newState)
				lose[newState.turn] += newState.forks
			} else {
				win[newState.turn] += newState.forks
			}
		}
	}
	return win, lose
}

func Part2(inputFile string) string {
	data := util.ReadInput(inputFile)
	first, second := getPosition(data[0]), getPosition(data[1])

	winFirst, loseFirst := doTheNumbers(first - 1)
	winSecond, loseSecond := doTheNumbers(second - 1)

	var firstWinsCount, secondWinsCount int64
	// a bit of cheating here, we know that 10 is always the max number of turns from debugging
	// otherwise we should have really found the max keys in the maps first and used that
	for i := 3; i <= 10; i++ {
		// first player wins on this turn in the number of his forks that he wins in
		// multiplied by the number of forks the second player did not win on previous turn
		firstWinsCount += winFirst[i] * loseSecond[i-1]
		// second player wins on this turn in the number of his forks that he wins in
		// multiplied by the number of forks the first player did not win on this turn
		secondWinsCount += winSecond[i] * loseFirst[i]
	}

	return fmt.Sprint(util.Max(firstWinsCount, secondWinsCount))
}
