package day18

type simpleVisitor[T any] func(t *term, ctx *T)

func visitSimple[T any](t *term, v simpleVisitor[T], ctx *T) {
	v(t, ctx)
	if !t.isConst() {
		visitSimple(t.left, v, ctx)
		visitSimple(t.right, v, ctx)
	}
}

type visitor[T any] interface {
	visitConst(t *term, ctx *T) bool
	visitPairStart(t *term, ctx *T) bool
	visitPairMid(t *term, ctx *T) bool
	visitPairEnd(t *term, ctx *T) bool
}

func visit[T any](t *term, v visitor[T], ctx *T) bool {
	if t.isConst() {
		if v.visitConst(t, ctx) {
			return true
		}
	} else {
		if v.visitPairStart(t, ctx) {
			return true
		}
		if t.left != nil && visit(t.left, v, ctx) {
			return true
		}
		if v.visitPairMid(t, ctx) {
			return true
		}
		if t.right != nil && visit(t.right, v, ctx) {
			return true
		}
		if v.visitPairEnd(t, ctx) {
			return true
		}
	}
	return false
}
