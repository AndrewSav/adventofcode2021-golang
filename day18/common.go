package day18

import (
	"aoc2021/util"
	"unicode"
)

type term struct {
	left  *term
	right *term
	value int
}

func (t *term) isConst() bool {
	return t.left == nil || t.right == nil
}

func (t *term) isPlain() bool {
	return t.left != nil && t.right != nil && t.left.isConst() && t.right.isConst()
}

func add(left, right *term) (result *term) {
	result = &term{left: left, right: right}
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
		return &term{left: left, right: right}, reminder[1:]
	} else {
		i := 0
		for ; unicode.IsDigit(rune(s[i])); i++ {
		}
		return &term{value: util.MustAtoi(s[:i])}, s[i:]
	}

}
