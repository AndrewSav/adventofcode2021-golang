package day04

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

const width = 5
const height = 5

type point struct {
	x int
	y int
}

type bingo struct {
	columnFillCount [width]int    // count of called out numbers in each column
	rawFillCount    [height]int   // count of called out numbers in each row
	sum             int           // total sum of all number on the board not called out so far
	lookup          map[int]point // translates a number on the board to its coordinates
	won             bool          // if the board is won
}

func (b *bingo) add(x, y, val int) { // since in go objects start empty this is used to fill an empty board with numbers from input
	b.sum += val
	b.lookup[val] = point{x, y}
}

func (b *bingo) mark(val int) (bool, int) { // this is called every time a bingo number is called out
	if p, ok := b.lookup[val]; ok && !b.won {
		b.sum -= val
		b.columnFillCount[p.x]++
		b.rawFillCount[p.y]++
		if b.columnFillCount[p.x] == width || b.rawFillCount[p.y] == height {
			b.won = true
			return true, b.sum * val
		}
	}
	return false, 0
}

func solve(inputFile string, firstWin bool) string {
	lines := util.ReadInput(inputFile)
	seq := util.AtoiSlice(strings.Split(lines[0], ",")) // numbers to call out from the first line of the input
	boards := []bingo{}
	for c := 1; c < len(lines); c++ { // creating all boards from input
		y := (c - 1) % (height + 1)
		if y == 0 { // this represents a blank input line before each board. Create a new board and continue to the next line
			boards = append(boards, bingo{lookup: map[int]point{}})
			continue
		}
		y-- // vertical position on the board. Since we have a blank line at the top, we need to decrement so that the index becomes zero based
		for x := 0; x < width; x++ {
			boards[len(boards)-1].add(x, y, util.MustAtoi(lines[c][x*3:x*3+2])) // all numbers are fixed width of 2 with a single space between them
		}
	}
	lastWin := -1
	for _, i := range seq { // now we start calling out the numbers
		for j := range boards {
			if won, score := boards[j].mark(i); won {
				if firstWin {
					return fmt.Sprint(score)
				} else {
					lastWin = score
				}
			}
		}
	}
	return fmt.Sprint(lastWin)
}
