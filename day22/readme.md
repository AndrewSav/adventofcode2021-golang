## --- Day 22: Reactor Reboot ---

For Part 1 the naïve approach of keeping track of every single cube has worked, but there was no way it was going to work for Part2. I was not able to figure out how to make it work, so I googled how other people did it. Most people used an approach of splitting each of any two intersecting cuboids in up to 27 other cuboids, so that each cube was contained in a single cuboid. This approach is quite easy to visualise, but there are a lot of different ways how two cuboids can intersect. I chose to implement a different approach that I saw mentioned in a single article. It uses [Sweep line algorithm](https://en.wikipedia.org/wiki/Sweep_line_algorithm) which actually is not an algorithm but an overarching principle a class of algorithms is based upon.

The idea, is that individual cubes (which we will call voxels, as per terminology of the article this implementation is based upon) state can only change around the borders of the cuboids listed in the input. If between two voxels (in each of the 3 dimensions) there is no input cuboids edges, all voxels will be in the same state. This allows us to skip all other points in-between, since we know what they would be. The fact is that some cuboids turn *on* and some turn *off* is easily dealt with by making sure that we apply to each group of voxels the same state that the last cuboid in the input order that contains that group of voxels prescribes.

We chose first dimension (z) and a sweep direction (from negative to positive) and determine the list of *stop points*. A stop point is a voxel that can potentially have different state than the previous voxel in the sweep direction. If a cuboid has z coordinates from, say -50 to 50, then the stop points will be -50, and 51. This is because -50 is the first voxel inside the cuboid in the sweep direction and 51 is the first voxel outside the cuboid in the sweep direction.

We are going to fist sweep a plain (z stop points) along z axis, for each stop point there, we are going to sweep a line along that plain for each of the y stop points in qualified cuboids, and finally we sweep a dot along the line for each of the x stop points in qualified cuboids.

We are moving from stop point to stop point in the sweep direction and at each point but the very first we add up all the "on" voxel between the previous point and the current (the previous included, the current excluded).

At each stop point we first get all the cuboids in the input order, that include the current point on the current dimension. If this is not the last dimension (that is, it is not x, but either z or y), we recursively call the sweep operation on the next dimension (y or x, respectively), passing only the cuboids that we just selected. The sweep operation returns the number of voxels on the passed dimension, and we multiply that by the length of the current interval between the stop points to get the number of voxels on the current dimension in the current interval. We keep iterating this way from stop point to stop point until we summed up all the voxels in each of the intervals.

If this is the last dimension (x), that is if we are sweeping a dot along a line, we can simply get the number of voxels within each interval where they are turned on, sum them up, and that will give us the number of on voxels on that entire line (for z and y fixed by the caller(s)).

In the end the top level sweep will get us the number of voxels the puzzle is asking us to determine.

Even this algorithm was somewhat slow, so I tried to memoise the results of the sweep function calls, and this got me under my 100ms goal.
