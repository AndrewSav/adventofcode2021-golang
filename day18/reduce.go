package day18

type explodeStage int // stages of explosion search

const (
	notExploded         explodeStage = iota // have not found a pair to explode (yet)
	waitingForNextRight                     // found a pair and exploded it, but have not found the const on the right (if any) to adjust
	doneExploding                           // finished exploding, so can ignore the rest of the snailfish number
)

type explodeContext struct {
	left  *term        // latest left const we saw (if any), to be modified by the right-adjacent exploding pair
	right int          // the right const of the exploded pair to be added to the next const on the right (if any)
	level int          // current nesting level
	stage explodeStage // indicates that we processed an exploding pair
}

func (t *term) reduce() {
	for t.explode(&explodeContext{}) || t.split() {
	}
}

func (t *term) explode(ctx *explodeContext) bool { // returns true if explosion happened in the sub-tree
	if ctx.stage == doneExploding {
		return true
	}
	if !t.isConst() { // if pair
		if ctx.level >= 4 && t.isPlain() && ctx.stage == notExploded { // we just found a pair that will explode
			// the pair's left value is added to the first regular number to the left of the exploding pair (if any)
			if ctx.left != nil {
				ctx.left.value += t.left.value
			}
			// We need to add the pair's right value to the first regular number to the right of the exploding pair (if any)
			// but we do not know yet what this number would be, so we save it to the context until we know
			ctx.right = t.right.value
			// the entire exploding pair is replaced with the regular number 0
			t.left = nil
			t.right = nil
			t.value = 0 // strictly speaking it will always already be 0, but just for clarity
			// This means that we are now looking to finish exploding by adding to the first regular number to the right
			ctx.stage = waitingForNextRight
		} else {
			ctx.level++
			t.left.explode(ctx)
			t.right.explode(ctx)
			ctx.level--
		}
	} else {
		if ctx.stage == waitingForNextRight {
			// The pair's right value is added to the first regular number to the right of the exploding pair (if any)
			t.value += ctx.right
			ctx.stage = doneExploding
		} else { // ctx.stage == notExploded
			// keep track of the last number on the left for the potential explosion
			ctx.left = t
		}
	}
	return ctx.stage != notExploded // if there is no number on the right of the exploding pair, the stage will remain `waitingForNextRight`
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
