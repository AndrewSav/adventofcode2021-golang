## --- Day 20: Trench Map ---

The main difficulty in this one was the "infinity" handling. Since the infinity ends up being the same value (0 or 1) in all the directions, we just need to figure out if we need to keep more in memory than the initial square. If we, for argument's sake call the top left corner (0,0) it is clear that (-1,-1) cell may be affected and (-2,-2) cell will not, so our initial square will expand by 1 cell in every direction in every enhance cycle. Note also, that while (-2,-2) cell in the example above won't change, it is a (diagonal) neighbour of the (-1,-1) cell, and thus is used in the enhanced algorithm calculation. That's why we need to keep track 2 cells more in each direction, then the largest square we will ever apply our enhancement algorithm to.

In my solution I start by allocating memory to fit that largest square + 2 in each direction, and I place the initial input smaller square at the centre of it. Since the cell value in the rest of the grid we are not keeping in memory can oscillate between the enhancement cycles we need to keep track of that too.

The solution itself, once we have this structure in place is not difficult.
