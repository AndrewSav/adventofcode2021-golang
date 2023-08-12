## --- Day 24: Arithmetic Logic Unit ---

Many puzzles in Advent of Code have some unstated invariants that we have to assume. I particularly disliked this puzzle because, in my opinion it takes it a notch too far, and I was unable to solve it. Many people were though, so there is that. My main gripe is that while you always can solve a puzzle
for your particular input, you never know if your solution is generic enough to work for all people inputs, and this just rubs me in a wrong way.

Anyway, when I first saw the puzzle I thought that all we need to do is to just code up the given assembly interpreter, which I gleefully did. Of course that did not work, because the puzzle is, at first sight, is to iterate over all 14 digit numbers, run the program given in the input on all of them and then find the number that gives a particular result. Brute forcing this was designed to be infeasible.

So once I found how other people solved it, for me the two challenges remained is to understand the idea, which did take some time, and then port the python code to golang so it looks more or less idiomatic.

The idea is that the input is always 14 almost identical blocks, one, per the 14-digit number digit, only 3 numbers of which can vary in each of the blocks. Analysing the block we will see that each block depending on the parameters either a) pushes a digit based on a digit from the 14-digit number on a stack, b) pops a digit from the said stack c) does something else. There are exactly 7 pushes, and exactly 7 opportunities to execute the option b) or c). It is possible to manipulate the 14-digit number so that option b) is always chosen when there is a choice between b) and c). The end result of the input problem is 0 (which is required) only if the number of the pushes and the pops are balanced. If option c) is ever selected the end result will never get to 0.

The above is a very simplified explanation, and "stack" in it is an abstraction. In reality base-26 digits are added or removed from an integer with `z*26 + digit` and `z/26` operations. We also skip over the explanation how to manipulate 14-digits to get the desired result, the explanation is given in  comments to the code.
