## --- Day 19: Beacon Scanner ---

The input gives us 26 scanners with roughly 26 beacons
scanned by each. What we are going to do is to "normilise" the scanners. Let's assume that the first one has the true
facing and orientation, and let convert coordiantes of the remaining scanners into the coordinates system of the
first scanner. This is how we are goint to do this. Until we have no non-normalised scanners left, loop through all
the remaining non-normalised scanners. Try to match the scanner with each of the already normalised scanners until
we either find a match, that aligns or run out of normalised scanners. Check if a non-normalised and normalised
scanners match in the following way. For each of possible 24 rotations of the non-normalised scanner find the offset
(or difference) between every beacon of the non-normalised scanner and every becadon on the normalised scanner (that
will be about `26*26=676` offsets). If we find 12 (as per puzzle desctiption) offsets that align we've found a match.
Now we translate the non-normalised scanner coordinates to the rotation that matched, and then shift the coordinates
by the offset that matched, which makes it a normalised scanner.

Once all scanners are normalised all the beacons coordinates are in the same coordinate system. With this achived,
both part 1 and part 2 are easy.

I was not satisfied with this solution, because it would take about one second to run both for part 1 and 2, and I
felt, that it is possible to do better. So, once again I resorted to Google. People more clever than me, figured out
that there is an optimisation, that you can apply to the process above. After processing the puzzle input, for each
scanner calculate pairwise distances between each beacon of the scanner. Before processing the 24 rotations, look if
there are matches in the pairwise distances of the normalised scanner and the pairwise distances of the non-normalised
scanner. If there are less then `12*(12-1)/2 = 66` matches, do not bother checking the rotations, there will be no match.

This optimisation gave x20 speed increase which I was satisfied with.
