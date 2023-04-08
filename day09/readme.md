## --- Day 9: Smoke Basin ---

Another puzzle where part 1 and part 2 do not have much in common. Part 1 is easily solved by
naive brute force. When I saw part 2, the first though I had was [Flood Fill](https://en.wikipedia.org/wiki/Flood_fill)
which led to boundary fill, which is a variant of flood fill, and the implementation.
We just iterate through each "cell" and run boundary fill algorythm starting from that cell
if this cell is neither a boundary nor has alreadt been filled. As we run boundart fill
algorythm we also count the number of "cells" filled, since this is what the puzzle is
asking us about.
