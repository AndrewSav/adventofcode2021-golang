## --- Day 17: Trick Shot ---

The part 1 here is a bit of a trick, because you can calculate the result with a simple formula, and it does not really require a program. There are a few thing to note:

- The `x` velocity always decreases and eventually the probe stops moving along the `X` axis
- `x` and `y` velocity values are independent, which means that for each simulation of probe trajectory if we change `x` velocity only (or `y` velocity only), no `Y` (or `X`) coordinate will change in the entire simulation for any point plotted
- For the part 1 the value of x does not really matter. We are assuming that there is such an `x` that will eventually land us in the target zone (otherwise the puzzle would not have a solution), so if we fix that we can now concentrate on `y` only.
- For the Y coordinate, the probe will go up `y`, then `y-1`, `y-2`,  until 2 and 1, which will be the top point. Given `y`, that top point is obviously `y * (y+1) / 2` (see Day 7: The Treachery of Whales).
- Once the probe is at the top, it will repeat the journey along the `Y` axis in exactly the same points as on the way up, that is 1, 2, `y-2`, `y-1`, `y`. At this point the `Y` coordinate will be `0` again, and the next point `Y` coordinate will be `-y-1`. If this is beyond the bottom `Y` coordinate of the target area we overshoot. If it is equal to it, we found the velocity for our top point. Now we can find the top point itself using the aforementioned formula. (We need to consider that the `y` will become negative once it falls under the `X` axis, and also that the absolute value of the final `Y` velocity before it leaves the target zone will be `y+1`, if `y` is the initiall velociy. Those two facts cancel each other out though, and in the end give the expected formula)

We are assuming here that the target zone is always `Y` negative and `X` positive. This is not specified in the puzzle description but it  seems to hold true for all the puzzles.

For part 2 we will do some real simulations. Since `x` and `y` are independent let's consider possible `x` values first. We will consider all `x` values in a loop starting with 1 until we reach the end of the target zone. For each step we store which initial velocities ends us in the zone for that step. Since the probe stops moving along the `X` axis eventually we also separately store for each step which initial velocities end up both in the zone and with current velocity of 0. Those velocities will have the probe on this coordinate for all subsequent steps too (since the `X` velocity reached zero by that step).

Then we do similar simulation for all possible initial `y` values, and on each step that is in the zone `y` wise we also check the results from the previous paragraph if it also in the zone `x` wise. If it is, we count it as a result.
