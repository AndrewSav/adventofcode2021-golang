package day18

func (t *term) getMagnitude2() int {
	var current int
	var left []int
	v := treeVisitor{
		visitConst: visitHandler(func(t *term) {
			current = t.value
		}),
		visitPairMid: visitHandler(func(t *term) {
			left = append(left, current)
		}),
		visitPairEnd: visitHandler(func(t *term) {
			pop := left[len(left)-1]
			left = left[:len(left)-1]
			current = pop*3 + current*2
		}),
	}
	visit(t, v)
	return current
}
