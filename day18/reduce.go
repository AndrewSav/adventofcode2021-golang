package day18

type redcucerContext struct {
	level           int
	left            *term
	hasJustExploded bool
	right           int
	reduced         bool
}

type exploder struct{}

func (p exploder) visitConst(t *term, ctx *redcucerContext) bool {
	if ctx.hasJustExploded {
		t.value += ctx.right
		return true
	} else {
		ctx.left = t
	}
	return false
}
func (p exploder) visitPairStart(t *term, ctx *redcucerContext) bool {
	ctx.level++
	if ctx.level > 4 && t.isPlain() && !ctx.hasJustExploded {
		if ctx.left != nil {
			ctx.left.value += t.left.value
		}
		ctx.hasJustExploded = true
		ctx.right = t.right.value
		t.left = nil
		t.right = nil
		ctx.reduced = true
	}
	return false
}
func (p exploder) visitPairMid(t *term, ctx *redcucerContext) bool {
	return false
}
func (p exploder) visitPairEnd(t *term, ctx *redcucerContext) bool {
	ctx.level--
	return false
}

type splitter struct{}

func (p splitter) visitConst(t *term, ctx *redcucerContext) bool {
	if t.value >= 10 {
		t.left = &term{value: t.value / 2}
		if t.value%2 == 0 {
			t.right = &term{value: t.value / 2}
		} else {
			t.right = &term{value: t.value/2 + 1}
		}
		t.value = 0
		ctx.reduced = true
		return true
	}
	return false
}
func (p splitter) visitPairStart(t *term, ctx *redcucerContext) bool {
	return false
}
func (p splitter) visitPairMid(t *term, ctx *redcucerContext) bool {
	return false
}
func (p splitter) visitPairEnd(t *term, ctx *redcucerContext) bool {
	return false
}

func (t *term) reduce() bool {
	ctx := &redcucerContext{}
	visit(t, visitor[redcucerContext](exploder{}), ctx)
	if ctx.reduced {
		return true
	}
	visit(t, visitor[redcucerContext](splitter{}), ctx)
	return ctx.reduced
}
