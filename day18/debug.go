package day18

import (
	"fmt"
	"strings"
)

func (t *term) print(printLevel bool) {
	var builder strings.Builder
	var level int
	v := treeVisitor{
		visitConst: visitHandler(func(t *term) {
			fmt.Fprintf(&builder, "%d", t.value)
		}),
		visitPairStart: visitHandler(func(t *term) {
			level++
			if t.left.isConst() && printLevel {
				fmt.Fprintf(&builder, "[(%d)", level)
			} else {
				fmt.Fprintf(&builder, "[")
			}
		}),
		visitPairMid: visitHandler(func(t *term) {
			fmt.Fprintf(&builder, ",")
		}),
		visitPairEnd: visitHandler(func(t *term) {
			level--
			fmt.Fprintf(&builder, "]")
		}),
	}
	visit(t, v)
	fmt.Println(builder.String())
}
