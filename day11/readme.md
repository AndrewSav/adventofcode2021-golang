## --- Day 11: Dumbo Octopus ---

This required some mental gymnastics, but is not too difficult. The code closely follows the puzzle description: first, we increment all the points, then we flash those greater than 9, and finally we reset to 0 those greater than 9.

When we flash, we also are considering neighbours, increment them and flash them too on the value of 10. If the value is more than 10 it means that this particular neighbour already flashed on that cycle so we will not flash it again.
