package day24

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

// Based mostly on https://www.keiruaprod.fr/blog/2021/12/29/a-comprehensive-guide-to-aoc-2021-day-24.html

/*
The input consists of 14 processing blocks - each for one of the inputs.
Each of them has the same structure and is parametrized by three
variables, let's call them "div" (line 5 of an input block), "chk" (line 6),
and "add" (line 16). Written in pseudocode and using "inp" for the input value
each block performs the following computation:

if z % 26 + chk == inp {
  z = z / div;
} else {
  z = (z / div)*26 + inp + add;
}

The value of div is always either 1 or 26. It looks like each block either adds
or replaces a single digit on the right of a 26 base number, or removes a single
digit from the right (or does neither). We can also verify that inp + add is
never 26 or over. When div = 1, chk is always greater than 9, so
`z % 26 + chk == inp` check is always false in this case. The above block of code
can thus be rewritten this way:

if div == 1 {
  z = z*26 + inp + add;
} else {
  if (z % 26 + chk == inp) {
    z = z / 26;
  } else {
    z = (z / 26)*26 + inp + add;
  }
}

We have exactly 7 blocks with div = 1, and exactly 7 blocks with div = 26,
each block with div == 1 adds a base-26 digit to z, as described above, the only way the
final result can come to zero is if each block with div == 26 removes one digit. For
that to happen `z % 26 + chk == inp` must be true.

To arrange this start with an arbitrary array of 14 inputs denoted
by [inp_0, inp_1, ..., inp_13]. Let shf be 0 if div is 1 and be 1 if div is 26.
If the first two instruction blocks have shf_0 == 0 and shf_1 == 0 then after
the first two inputs two digit will have been pushed to the stack:

    z_stack = [inp_0 + add_0, inp_1 + add_1]

If then shf_2 == 1 we want inp_2 to have a value that would cause the last digit to pop.
The last digit is popped if

        z.last_digit + chk_2 == inp_2
    =>  inp_1 + add_1 + chk_2 == inp_2

So we changed inp_2 to be  (inp_2 = inp_1 + add_1 + chk_2). It can now happen that the
condition 1 <= inp_2 <= 9 is violated. In this case we can add an
arbitrary value to inp_2 to restore this condition. We will need to
add the same value to inp_1 too in order to maintain the previous
equality. We need to be careful that after these adjustments we also
maintain 1 <= inp_1 <= 9. The least we can do is for cases where
inp_2 < 1 to choose the value so that inp_2 = 1 and for cases with
inp_2 > 9 to choose the value so that inp_2 = 9. If this still doesn't
work for inp_1, then no other value will work for both either. (And
we know the puzzle has a solution so it will work)

This strategy can be used to take any input sequence and correct
it so that it passes the test. So for part 1 we'll want to start with
the highest possible input, 99999999999999, and for part 2 with the
lowest, 11111111111111.

*/

// line numbers with the (div, chk, add) parameters.
const divOffset = 4
const chkOffset = 5
const addOffset = 15

const subLen = 18 // length of each of the 14 processing blocks

// get div, chk or add parameter
func getVal(lines []string, index, offset int) int {
	return util.MustAtoi(strings.Split(lines[index*subLen+offset], " ")[2])
}

func solve(lines []string, inp [14]int) string {
	stack := []int{}
	for i := 0; i < 14; i++ {
		div := getVal(lines, i, divOffset)
		if div == 1 {
			// we just push index on the stack, not the actual digit
			// becase with the index we can easily access both the digit
			// inp[j] and the parameter getVal(lines, j, addOffset)
			stack = append(stack, i)
		} else {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			add := getVal(lines, j, addOffset)
			chk := getVal(lines, i, chkOffset)
			inp[i] = inp[j] + add + chk
			if inp[i] > 9 {
				inp[j] = inp[j] - (inp[i] - 9)
				inp[i] = 9
			}
			if inp[i] < 1 {
				inp[j] = inp[j] + (1 - inp[i])
				inp[i] = 1
			}
		}
	}
	var sb strings.Builder
	for _, b := range inp {
		fmt.Fprintf(&sb, "%d", b)
	}
	return sb.String()
}
