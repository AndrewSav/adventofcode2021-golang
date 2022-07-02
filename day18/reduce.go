package day18

func (t *term) reduce() bool {
	var level, right int
	var left *term
	var hasJustExploded, reduced bool
	exploder := treeVisitor{
		visitConst: visitHandler(func(t *term) {
			if reduced && !hasJustExploded {
				return
			}
			if hasJustExploded {
				t.value += right
				hasJustExploded = false
			} else {
				left = t
			}
		}),
		visitPairStart: visitHandler(func(t *term) {
			level++
			if reduced {
				return
			}
			if level > 4 && t.isPlain() && !hasJustExploded {
				if left != nil {
					left.value += t.left.value
				}
				hasJustExploded = true
				right = t.right.value
				t.left = nil
				t.right = nil
				reduced = true
			}
		}),
		visitPairEnd: visitHandler(func(t *term) {
			level--
		}),
	}
	splitter := treeVisitor{
		visitConst: visitHandler(func(t *term) {
			if reduced {
				return
			}
			if t.value >= 10 {
				t.left = &term{value: t.value / 2}
				t.right = &term{value: t.value / 2}
				if t.value%2 == 1 {
					t.right.value++
				}
				t.value = 0
				reduced = true
			}
		}),
	}
	visit(t, exploder)
	if reduced {
		return true
	}
	visit(t, splitter)
	return reduced
}
