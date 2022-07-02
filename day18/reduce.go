package day18

type redcucerContext struct {
	level           int
	left            *term
	hasJustExploded bool
	right           int
	reduced         bool
}

func (t *term) reduce() bool {
	ctx := &redcucerContext{}
	exploder := treeVisitor[redcucerContext]{
		visitConst: visitHandler[redcucerContext](func(t *term, ctx *redcucerContext) bool {
			if ctx.hasJustExploded {
				t.value += ctx.right
				return true
			} else {
				ctx.left = t
			}
			return false
		}),
		visitPairStart: visitHandler[redcucerContext](func(t *term, ctx *redcucerContext) bool {
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
		}),
		visitPairEnd: visitHandler[redcucerContext](func(t *term, ctx *redcucerContext) bool {
			ctx.level--
			return false
		}),
	}
	splitter := treeVisitor[redcucerContext]{
		visitConst: visitHandler[redcucerContext](func(t *term, ctx *redcucerContext) bool {
			if t.value >= 10 {
				t.left = &term{value: t.value / 2}
				t.right = &term{value: t.value / 2}
				if t.value%2 == 1 {
					t.right.value++
				}
				t.value = 0
				ctx.reduced = true
				return true
			}
			return false
		}),
	}
	visit(t, exploder, ctx)
	if ctx.reduced {
		return true
	}
	visit(t, splitter, ctx)
	return ctx.reduced
}
