package day18

type explodeStage int // stages of explosion search

const (
	notExploded         explodeStage = iota // have not found a pair to explode (yet)
	waitingForNextRight                     // found a pair and exploded it, but have not found the const on the right (if any) to adjust
	doneExploding                           // finished exploding, so can ignore the rest of the snailfish number
)

type explodeContext struct {
	left  *term        // latest left const we saw (if any), to be modified by the right-adjasent exploding pair
	right int          // the right const of the exploded pair to be added to the next const on the right (if any)
	level int          // current nesting level
	stage explodeStage // indicates that we processed an exploding pair
}

func (t *term) reduce() {
	for t.explode(&explodeContext{}) || t.split() {
	}
}

func (t *term) explode(ctx *explodeContext) bool { // returns true if explosion finished in the sub-tree
	if ctx.stage == doneExploding {
		return true
	}
	if t.isConst() {
		if ctx.stage == waitingForNextRight {
			t.value += ctx.right
			ctx.stage = doneExploding
		} else {
			ctx.left = t
		}
	} else {
		if ctx.level >= 4 && t.isPlain() && ctx.stage == notExploded {
			if ctx.left != nil {
				ctx.left.value += t.left.value
			}
			ctx.right = t.right.value
			t.left = nil
			t.right = nil
			t.value = 0 // strinctly speaking it will always already be 0, but just for clarity
			ctx.stage = waitingForNextRight
		} else {
			ctx.level++
			t.left.explode(ctx)
			t.right.explode(ctx)
			ctx.level--
		}
	}
	return ctx.stage == doneExploding
}

func (t *term) split() bool { // returns true if split has happened in the subtree
	if t.isConst() {
		if t.value >= 10 {
			t.left = &term{value: t.value / 2}
			t.right = &term{value: t.value / 2}
			if t.value%2 == 1 {
				t.right.value++
			}
			t.value = 0
			return true
		} else {
			return false
		}
	} else {
		return t.left.split() || t.right.split()
	}
}
