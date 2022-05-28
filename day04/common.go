package day04

import (
	"aoc2021/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const width = 5
const height = 5

type point struct {
	x int
	y int
}

type bingo struct {
	x      [width]int
	y      [height]int
	sum    int
	lookup map[int]point
	won    bool
}

func (b *bingo) add(x, y, val int) {
	b.sum += val
	b.lookup[val] = point{x, y}
}

func (b *bingo) mark(val int) (bool, int) {
	if p, ok := b.lookup[val]; ok && !b.won {
		b.sum -= val
		b.x[p.x]++
		b.y[p.y]++
		if b.x[p.x] == width || b.y[p.y] == height {
			b.won = true
			return true, b.sum * val
		}
	}
	return false, 0
}

func toNum(s string) int {
	if i, err := strconv.ParseInt(strings.Trim(s, " "), 10, 32); err != nil {
		log.Fatalf("cannot convert string '%s' to a number: %v", s, err)
	} else {
		return int(i)
	}
	panic("unexpected code path")
}

func solve(inputFile string, firstWin bool) string {
	lines := util.ReadInput(inputFile)
	seq := util.Atoi(strings.Split(lines[0], ","))
	mm := []bingo{}
	for c := 1; c < len(lines); c++ {
		y := (c - 1) % (height + 1)
		if y == 0 {
			mm = append(mm, bingo{lookup: map[int]point{}})
			continue
		}
		y--
		for x := 0; x < width; x++ {
			mm[len(mm)-1].add(x, y, toNum(lines[c][x*3:x*3+2]))
		}
	}
	lastWin := -1
	for _, i := range seq {
		for j, _ := range mm {
			if won, win := mm[j].mark(i); won {
				if firstWin {
					return fmt.Sprint(win)
				} else {
					lastWin = win
				}
			}
		}
	}
	return fmt.Sprint(lastWin)
}
