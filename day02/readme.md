## --- Day 2: Dive! ---

This puzzle also is quite straight-forward. We have different definitions of
"forward", "up" and "down" between parts one and two, so it makes sense to
abstract those away as a function type, and then call the appropriate function
when doing the main calculation.
