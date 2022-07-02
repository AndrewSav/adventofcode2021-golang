package day18

import (
	"fmt"
	"strings"
)

type printContext struct {
	builder    strings.Builder
	level      int
	printLevel bool
}

func (t *term) print(printLevel bool) {
	ctx := &printContext{printLevel: printLevel}
	v := treeVisitor[printContext]{
		visitConst: visitHandler[printContext](func(t *term, ctx *printContext) bool {
			fmt.Fprintf(&ctx.builder, "%d", t.value)
			return false
		}),
		visitPairStart: visitHandler[printContext](func(t *term, ctx *printContext) bool {
			ctx.level++
			if t.left.isConst() && ctx.printLevel {
				fmt.Fprintf(&ctx.builder, "[(%d)", ctx.level)
			} else {
				fmt.Fprintf(&ctx.builder, "[")
			}
			return false
		}),
		visitPairMid: visitHandler[printContext](func(t *term, ctx *printContext) bool {
			fmt.Fprintf(&ctx.builder, ",")
			return false
		}),
		visitPairEnd: visitHandler[printContext](func(t *term, ctx *printContext) bool {
			ctx.level--
			fmt.Fprintf(&ctx.builder, "]")
			return false
		}),
	}
	visit(t, v, ctx)
	fmt.Println(ctx.builder.String())
}
