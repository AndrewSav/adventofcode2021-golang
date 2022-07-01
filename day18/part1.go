package day18

import (
	"aoc2021/util"
	"fmt"
	"unicode"
)

type term struct {
	left  *term
	right *term
	value int
	level byte
}

func (t *term) isConst() bool {
	return t.left == nil || t.right == nil
}

func (t *term) isPlain() bool {
	return t.left != nil && t.right != nil && t.left.isConst() && t.right.isConst()
}

type context struct {
	left            *term
	hasJustExploded bool
	right           int
}

func (t *term) reduce() bool {
	return t.reduceExplode(&context{}) || t.reduceSplit()
}

func (t *term) reduceExplode(ctx *context) bool {
	if ctx.hasJustExploded {
		if t.isConst() {
			t.value = t.value + ctx.right
			ctx.hasJustExploded = false
			return true
		} else {
			return t.left.reduceExplode(ctx)
		}
	} else {
		if t.level >= 4 && t.isPlain() {
			if ctx.left != nil {
				ctx.left.value = ctx.left.value + t.left.value
			}
			ctx.hasJustExploded = true
			ctx.right = t.right.value
			t.left = nil
			t.right = nil
			t.value = 0
			return true
		} else {
			if t.isConst() {
				ctx.left = t
				return false
			} else {
				return t.left.reduceExplode(ctx) && !ctx.hasJustExploded || t.right.reduceExplode(ctx)
			}
		}
	}
}

func (t *term) reduceSplit() bool {
	if t.isConst() {
		if t.value >= 10 {
			t.left = &term{value: t.value / 2, level: t.level + 1}
			if t.value%2 == 0 {
				t.right = &term{value: t.value / 2, level: t.level + 1}
			} else {
				t.right = &term{value: t.value/2 + 1, level: t.level + 1}
			}
			t.value = 0
			return true
		} else {
			return false
		}
	} else {
		return t.left.reduceSplit() || t.right.reduceSplit()
	}
}

func (t *term) setLevel(level byte) {
	t.level = level
	if !t.isConst() {
		t.left.setLevel(level + 1)
		t.right.setLevel(level + 1)
	}
}

func add(left, right *term) (result *term) {
	result = &term{left: left, right: right}
	result.setLevel(0)
	for result.reduce() {
	}
	return
}

func (t *term) getMagnitude() int {
	if t.isConst() {
		return t.value
	} else {
		return t.left.getMagnitude()*3 + t.right.getMagnitude()*2
	}
}

func parse(s string) (result *term) {
	result, _ = parseInternal(s, 0)
	return
}

func parseInternal(s string, level byte) (*term, string) {
	if s[0] == '[' {
		left, reminder := parseInternal(s[1:], level+1)
		right, reminder := parseInternal(reminder[1:], level+1)
		return &term{left: left, right: right, level: level}, reminder[1:]
	} else {
		i := 0
		for ; unicode.IsDigit(rune(s[i])); i++ {
		}
		return &term{value: util.MustAtoi(s[:i]), level: level}, s[i:]
	}

}

func Part1(inputFile string) string {
	lines := util.ReadInput(inputFile)
	var result *term
	for i, l := range lines {
		if i == 0 {
			result = parse(l)
		} else {
			result = add(result, parse(l))
		}
	}

	return fmt.Sprint(result.getMagnitude())
}
