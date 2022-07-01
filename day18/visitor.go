package day18

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
		if t.left != nil && visit(t.left, v, ctx) { // t.left can be nil if the node is mutated by the visitor
			return true
		}
		if v.visitPairMid(t, ctx) {
			return true
		}
		if t.right != nil && visit(t.right, v, ctx) { // t.right can be nil if the node is mutated by a visitor
			return true
		}
		if v.visitPairEnd(t, ctx) {
			return true
		}
	}
	return false
}
