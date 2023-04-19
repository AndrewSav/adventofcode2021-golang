package day18

import (
	"aoc2021/util"
)

type term struct { // either 'value' (const) or 'left' and 'right' (pair) are used in each instance
	left  *term
	right *term
	value int
}

func (t *term) isConst() bool {
	return t.left == nil || t.right == nil // they either are both nils or both not nils
}

func (t *term) isPlain() bool { // determines if this pair has no nested pairs and can potentially explode
	return t.left != nil && t.right != nil && t.left.isConst() && t.right.isConst()
}

func add(left, right *term) (result *term) {
	result = &term{left: left, right: right}
	result.reduce()
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
	result, _ = parseInternal(s)
	return
}

func parseInternal(s string) (*term, string) { // returns the parsed term and the unparsed string reminder
	if s[0] == '[' {
		left, reminder := parseInternal(s[1:])
		right, reminder := parseInternal(reminder[1:])
		return &term{left: left, right: right}, reminder[1:]
	} else {
		return &term{value: util.MustAtoi(s[:1])}, s[1:]
	}
}
