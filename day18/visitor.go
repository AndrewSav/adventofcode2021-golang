package day18

type treeVisitor[T any] struct {
	visitConst     visitor[T]
	visitPairStart visitor[T]
	visitPairMid   visitor[T]
	visitPairEnd   visitor[T]
}

type visitor[T any] interface {
	visit(t *term, ctx *T) bool
}

type visitHandler[T any] func(t *term, ctx *T) bool

func (f visitHandler[T]) visit(t *term, ctx *T) bool {
	return f(t, ctx)
}

func visit[T any](t *term, v treeVisitor[T], ctx *T) bool {
	if t.isConst() {
		if v.visitConst != nil && v.visitConst.visit(t, ctx) {
			return true
		}
	} else {
		if v.visitPairStart != nil && v.visitPairStart.visit(t, ctx) {
			return true
		}
		if t.left != nil && visit(t.left, v, ctx) { // t.left can be nil if the node is mutated by the visitor
			return true
		}
		if v.visitPairMid != nil && v.visitPairMid.visit(t, ctx) {
			return true
		}
		if t.right != nil && visit(t.right, v, ctx) { // t.right can be nil if the node is mutated by a visitor
			return true
		}
		if v.visitPairEnd != nil && v.visitPairEnd.visit(t, ctx) {
			return true
		}
	}
	return false
}
