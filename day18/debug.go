package day18

import (
	"fmt"
	"strings"
)

func (t *term) print(printLevel bool) {
	var builder strings.Builder
	var level int
	v := treeVisitor{
		visitConst: visitHandler(func(t *term) bool {
			fmt.Fprintf(&builder, "%d", t.value)
			return false
		}),
		visitPairStart: visitHandler(func(t *term) bool {
			level++
			if t.left.isConst() && printLevel {
				fmt.Fprintf(&builder, "[(%d)", level)
			} else {
				fmt.Fprintf(&builder, "[")
			}
			return false
		}),
		visitPairMid: visitHandler(func(t *term) bool {
			fmt.Fprintf(&builder, ",")
			return false
		}),
		visitPairEnd: visitHandler(func(t *term) bool {
			level--
			fmt.Fprintf(&builder, "]")
			return false
		}),
	}
	visit(t, v)
	fmt.Println(builder.String())
}
