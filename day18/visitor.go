package day18

type treeVisitor struct {
	visitConst     visitor
	visitPairStart visitor
	visitPairMid   visitor
	visitPairEnd   visitor
}

type visitor interface {
	visit(t *term) bool
}

type visitHandler func(t *term) bool

func (f visitHandler) visit(t *term) bool {
	return f(t)
}

func visit(t *term, v treeVisitor) bool {
	if t.isConst() {
		if v.visitConst != nil && v.visitConst.visit(t) {
			return true
		}
	} else {
		if v.visitPairStart != nil && v.visitPairStart.visit(t) {
			return true
		}
		if t.left != nil && visit(t.left, v) { // t.left can be nil if the node is mutated by the visitor
			return true
		}
		if v.visitPairMid != nil && v.visitPairMid.visit(t) {
			return true
		}
		if t.right != nil && visit(t.right, v) { // t.right can be nil if the node is mutated by a visitor
			return true
		}
		if v.visitPairEnd != nil && v.visitPairEnd.visit(t) {
			return true
		}
	}
	return false
}
