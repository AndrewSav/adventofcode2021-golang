## --- Day 8: Seven Segment Search ---

This is where things stating to get a bit more complex. Since part 1 is (in a certain sense)
just a special case of part 2, I will write about part 2 here. One thing that the puzzle does not
tell us is the algorythm how to deduce which segment string represent each digit. It is not that
difficult to figure out though. For digits 1,7,4 and 8 the puzzle (implicitly) tells use to use
the segment strings of 2,3,4 and 7 respectively, so we should start with that. What will remain
after that is 2, 3 and 5 with segment string of 5 and 0,6 and 9 with segment string of 6.

Since we alredy identified the segment string represnting 1, we can find the segment string
representing 3 by searching for the segment strings of 5, for one that contains all the segments
from the segment string represnting 1. Neither 2 nor 5 would qualify. Then, we notice that if
we remove segments that 1 contains from segments that 4 contains we will get this (letters will differ):

```
 ....
b    .
b    .
 dddd
.    .
.    .
 ....
```

And this configuration is conained in digit 5 but not contained in digint 2. Thus we took care of all
length 5 segment strings. Similarily we can deal with the segment strings of length 6. 1 is entirely
contained in 0 and 9, so the only segment string of 6 that does not contain it is 6. And then whe can
detect 9, because it fully contains 4, whereas 0 does not.

The rest is easy.
