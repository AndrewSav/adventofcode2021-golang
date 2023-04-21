## --- Day 12: Passage Pathing ---

Ah, good old graph search! I used [breadth-first](https://en.wikipedia.org/wiki/Breadth-first_search) for this one. We need to define the minimal necessary data structures, such as node and path. Path is used for keeping track of small caves we've been to, so that we do not violate the puzzle conditions with regards to them.

Once the data structures and the breadth-first implementation is in place, we have the solution.
