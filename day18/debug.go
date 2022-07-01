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

type printer struct{}

func (p printer) visitConst(t *term, ctx *printContext) bool {
	fmt.Fprintf(&ctx.builder, "%d", t.value)
	return false
}
func (p printer) visitPairStart(t *term, ctx *printContext) bool {
	ctx.level++
	if t.left.isConst() && ctx.printLevel {
		fmt.Fprintf(&ctx.builder, "[(%d)", ctx.level)
	} else {
		fmt.Fprintf(&ctx.builder, "[")
	}
	return false
}
func (p printer) visitPairMid(t *term, ctx *printContext) bool {
	fmt.Fprintf(&ctx.builder, ",")
	return false
}
func (p printer) visitPairEnd(t *term, ctx *printContext) bool {
	ctx.level--
	fmt.Fprintf(&ctx.builder, "]")
	return false
}
func (t *term) print(printLevel bool) {
	ctx := &printContext{printLevel: printLevel}
	visit(t, visitor[printContext](printer{}), ctx)
	fmt.Println(ctx.builder.String())
}
