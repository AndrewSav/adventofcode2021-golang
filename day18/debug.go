package day18

import "fmt"

func (t *term) print() {
	t.printInternal(0)
	fmt.Println()

}

func (t *term) printInternal(level int) {
	if t.isConst() {
		fmt.Printf("%d: %d;", t.value, t.level)
	} else {
		t.left.printInternal(level + 1)
		t.right.printInternal(level + 1)
	}
}
