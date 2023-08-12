## --- Day 15: Chiton ---

This was rather interesting one, because I learned a lot. When you see a graph search for the shortest path you immediately think about Dijkstra or A\*. I decided to go with Dijkstra for this one. My [initial](https://github.com/AndrewSav/adventofcode2021-golang/blob/2e1de5de449d92eba2a265b71fca90298659aee7/day15/part1.go) straight-forward implementation took more than a second for Part 1, and I had a suspicion that it might not cut at all for Part 2. The next attempt was using a priority queue, which was supposed to give some performance improvement. Golang does not have built-in priority queue, but it has [this](https://pkg.go.dev/container/heap#example-package-PriorityQueue). The implementation looked like [this](https://github.com/AndrewSav/adventofcode2021-golang/blob/8bed435a871b0451c0909ac66da1a97fb7bd45f0/day15/part1.go). This only took 4 milliseconds so that was a definite improvement. And yet I felt dissatisfied, because I did not particularly like the priority queue implementation. I went through a few alternative algorithms that the Wikipedia page on Dijkstra algorithm mentions, and Bucket Queue seemed to fit the bill. As Wikipedia puts it: "it particularly suited to applications in which the priorities have a small range", and this is exactly our case. The implementation also looks quite simple, which was the main appeal to me. The bucket queue gave marginally better results than the priority queue, and that was my final solution.

Part 2 was just applying the work done in part 1 to the bigger cave.