## --- Day 18: Snailfish ---

This one ended up being quite mind-bending. I spent quite some time on. While the puzzle description explains what "reduce" is, implementing it is another matter. My implementation of reduce is mutating, which means that when you add two snailfish numbers, the input is modified during the reduce process and you can no longer reuse that input. Of course it can be easily solved by making copies of the input first, but since it was not necessary in my case, I opted for this slightly "unclean" implementation, so that the codebase remained smaller. There is, of course also memory and CPU overhead (which I never measured) that is associated with copying.

The core of the solution is the implementation of the reduce operation. After I had my first implementation I actually took some time to try something else (as you can see in the commit history) and implement the same solution using the visitor patter and new shiny go generics. In the end it did not make the solution faster, more clear or better, and I reverted back to the original solution.

The reduce process in essence is applying the two steps in the description, and finishing if none of them succeeds. We apply explode first and if it does not succeed we apply split:

```go
type term struct { // either 'value' (const) or 'left' and 'right' (pair) are used in each instance
	left  *term
	right *term
	value int
}
func (t *term) reduce() {
	for t.explode(&explodeContext{}) || t.split() {
	}
}
```

The explode function is recursive. It keeps track of explosion context:

```go
const (
	notExploded         explodeStage = iota // have not found a pair to explode (yet)
	waitingForNextRight                     // found a pair and exploded it, but have not found the const on the right (if any) to adjust
	doneExploding                           // finished exploding, so can ignore the rest of the snailfish number
)

type explodeContext struct {
	left  *term        // pointer to the latest left const we saw (if any), to be modified by the right-adjacent exploding pair
	right int          // the right const of the exploded pair to be added to the next const on the right (if any)
	level int          // current nesting level
	stage explodeStage // indicates that we processed an exploding pair
}
```

We start with the top pair and call explode recursively for the left and then the right element of the pair, while keeping track of the current nesting level. We always store the latest constant seen in the `left` field of the context. If we are deep enough to satisfy the explode condition, we replace the current pair with the zero constant, add the left element of that pair to latest seen constant from `left` field of the context and store the right element of the pair in the `right` field of the context to add to then next seen constant (if any). Once we found the next constant and added the `right` value to it, or if we finished with the entire snailfish number we are done.
