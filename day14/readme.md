## --- Day 14: Extended Polymerization ---

Part 1 here implements the naive approach where we keep the string, and inserting the letters as prescribed. That apporach works for 10 iterations,
but does not work for 40 iterations: it is too slow, I left it running overnight and then it ran out of memory.

I was not able to figure out the second part on my own so I had to google it. The trick is, similar to Day 6's Lanternfish, to keep track of all
possible pairs count, instead of trying to maintain the entire string. At first we collect all pairs (that is every two ajacent letters) in the
initial string, and assign each of them a counter equal to how many times they occur in the initial string. Then we notice that every time we apply
a rule, the counters change quite regularily, for example, if we have a rule `SV -> O`, the counter for SV will go down one, and the countres for
`SO` and `OV` get up one each. This is because `SV` was changed to `SOV` and instead a `SV` pair we now have an `SO` and `OV` pairs. This model of
calculation allows us to finish in reasonable time for Part 2.

The data structure we are using is a map which represent the rules. Each rule mapping a pair, to the two pairs it produces, so `SV -> O` rule is 
represnted by a map entry with the key `SV` and the value `{"SO", "OV"}`. We keep the counters in another map, with they key of the letter pair,
e.g. `SV` and the value of the counter. Finally, we also keep a map of counters for each letter for score purposes. Each rule increment exactly one
of those by one, for example the `SV -> O` increases the `O` counter by one.