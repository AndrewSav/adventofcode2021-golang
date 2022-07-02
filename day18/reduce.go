package day18

func (t *term) reduce() bool {
	var level, right int
	var left *term
	var hasJustExploded, reduced bool
	exploder := treeVisitor{
		visitConst: visitHandler(func(t *term) bool {
			if reduced && !hasJustExploded {
				return false
			}
			if hasJustExploded {
				t.value += right
				hasJustExploded = false
				return true
			} else {
				left = t
			}
			return false
		}),
		visitPairStart: visitHandler(func(t *term) bool {
			level++
			if reduced {
				return false
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
			return false
		}),
		visitPairEnd: visitHandler(func(t *term) bool {
			level--
			return false
		}),
	}
	splitter := treeVisitor{
		visitConst: visitHandler(func(t *term) bool {
			if reduced {
				return false
			}
			if t.value >= 10 {
				t.left = &term{value: t.value / 2}
				t.right = &term{value: t.value / 2}
				if t.value%2 == 1 {
					t.right.value++
				}
				t.value = 0
				reduced = true
				return true
			}
			return false
		}),
	}
	visit(t, exploder)
	if reduced {
		return true
	}
	visit(t, splitter)
	return reduced
}
