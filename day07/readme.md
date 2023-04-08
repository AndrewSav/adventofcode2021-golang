## --- Day 7: The Treachery of Whales ---

When reading this puzzle one might think that a clever optimisation could be required,
but the brute force approach works very well: calculate fuel usage for each crab to each 
possible goal position, and take the minimum. Fuel usage uses a different formula in parts
one and two, and while in the part two it is described with words, rather by a given formula,
the formula (`n * (n + 1) / 2`) is quite easy to come up with. If you are not familiar with
this wikipedia has an [article](https://en.wikipedia.org/wiki/1_%2B_2_%2B_3_%2B_4_%2B_%E2%8B%AF)
on it, but I was able to quickly figure it out based on my school education ;)
