package day18

type treeVisitor struct {
	visitConst     visitor
	visitPairStart visitor
	visitPairMid   visitor
	visitPairEnd   visitor
}

type visitor interface {
	visit(t *term)
}

type visitHandler func(t *term)

func (f visitHandler) visit(t *term) {
	f(t)
}

func visit(t *term, v treeVisitor) {
	if t.isConst() {
		if v.visitConst != nil {
			v.visitConst.visit(t)
		}
	} else {
		if v.visitPairStart != nil {
			v.visitPairStart.visit(t)
		}
		if t.left != nil { // t.left can be nil if the node is mutated by the visitor
			visit(t.left, v)
		}
		if v.visitPairMid != nil {
			v.visitPairMid.visit(t)
		}
		if t.right != nil { // t.right can be nil if the node is mutated by a visitor
			visit(t.right, v)
		}
		if v.visitPairEnd != nil {
			v.visitPairEnd.visit(t)
		}
	}
}
