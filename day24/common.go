package day24

import (
	"aoc2021/util"
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

if div == 1 { // push
  z = z*26 + inp + add;
} else {
  if (z % 26 + chk == inp) { // pop
    z = z / 26;
  } else { // we do not want this code path
    z = (z / 26)*26 + inp + add;
  }
}

We have exactly 7 blocks with div = 1, and exactly 7 blocks with div = 26,
each block with div == 1 adds a base-26 digit to z, as described above, the only way the
final result can come to zero is if each block with div == 26 removes one digit. For
that to happen `z % 26 + chk == inp` must be true.

To arrange this let's break up all 14 block on pairs, where the first block in the pair
(denoted by index j) pushes an input digit (denoted ny inp_j) and the second block
(denoted by index i) in the pair we want to pop. Since we have exactly 7 of each they will
nicely pair up. In order for `z % 26 + chk == inp` to hold, the following should be true:

        z.last_digit + chk_i == inp_i
    =>  inp_j + add_j + chk_i == inp_i
    =>  inp_i - inp_j == add_j + chk_i

add_j and chk_i is given to us by the procesing blocks as described above, and inp_j and inp_i
we can manipulate. add_j + chk_i can be either negative or non negative, let's consider the
non-negative case. If we want the min result we should set inp_j to 1 and inp_i to 1 + add_j + chk_i,
note that this will never go over 9 because otherwise the puzzle would have no solution. If we
want the max result, we should set inp_i to 9 then, and inp_j to 9 - (add_j + chk_i).

If add_j + chk_i is negative, we simply use the same logic above with abs(add_j + chk_i) and then swap
inp_i and inp_j.

Once we sorted out all the 7 pairs we get our final 14 digit number.
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

func solve(lines []string, max bool) string {
	inp := [14]rune{}
	stack := []int{}
	for i := 0; i < 14; i++ {
		div := getVal(lines, i, divOffset)
		if div == 1 {
			stack = append(stack, i)
		} else {
			// j is the index where we pushed the number currently (i) being popped
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// expected difference between inp[i] - inp[j]
			diff := getVal(lines, j, addOffset) + getVal(lines, i, chkOffset)
			if max {
				// if we are looking for max, one of the two digits will be 9
				inp[i] = '9'
				// and the other will be |diff| less
				inp[j] = inp[i] - rune(util.Abs(diff))
			} else {
				// if we are looking for min, one of the two digits will be 1
				inp[j] = '1'
				// and the other will be |diff| more
				inp[i] = inp[j] + rune(util.Abs(diff))
			}
			// what we just did above keeps inp[j] <= inp[i]
			// swap the digits around to maintain inp[i] - inp[j] = add + chk
			// inp[i] should be less than inp[j] if add + chk is negative
			if diff < 0 {
				inp[i], inp[j] = inp[j], inp[i]
			}
		}
	}
	return string(inp[:])
}
