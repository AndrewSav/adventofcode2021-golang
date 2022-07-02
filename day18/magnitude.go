package day18

func (t *term) getMagnitude2() int {
	var current int
	var left []int
	v := treeVisitor{
		visitConst: visitHandler(func(t *term) bool {
			current = t.value
			return false
		}),
		visitPairMid: visitHandler(func(t *term) bool {
			left = append(left, current)
			return false
		}),
		visitPairEnd: visitHandler(func(t *term) bool {
			pop := left[len(left)-1]
			left = left[:len(left)-1]
			current = pop*3 + current*2
			return false
		}),
	}
	visit(t, v)
	return current
}
