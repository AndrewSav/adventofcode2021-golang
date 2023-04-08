## --- Day 5: Hydrothermal Venture ---

We do not know the dimensions of the plane we are working with so we need to figure it out from input data first.
I was able to find something I could use for that in the golang standard library: [Rectangle.Union](https://pkg.go.dev/image#Rectangle.Union)
method in the [image](https://pkg.go.dev/image) module. In order to use that we need to represent our lines as
`Rectangle` objects. Since this is not an exact abstraction it causes a few problems. One is that a line is in
fact a "empty" rectangle, so `Union` code ignores it. Thankfully we can use [Rectangle.Inset](https://pkg.go.dev/image#Rectangle.Inset)
in order to make a rectangle that encloses the line and use `Union` on that. Another problem is that for some
operation on rectangle objects we need [canoniclal](https://pkg.go.dev/image#Rectangle.Canon) representation of
a rectangle, where the `min` coordinates are less (or equal) than `max` coordinates. Unfortunately, for diagonals,
converting to canonical rectangle may change the rectangle diagonal! Each rectangle has two of them and we need
a specific one that is given by the input. We can work around this, by only using `Rectangle.Canon` right before
we are doing a calculation that would require it, but keeping the old non-canonical values intact for future
calculations.

Then we need a data structure to keep the number of line going through each point within our dimensions. Since we already
using the `image` module we can use [image.Gray](https://pkg.go.dev/image#Gray) for that, where the number that normaly
represents the colour of a point (`Y` field) is used to keep track of the number of lines going through the point.

Once the above is taken care of, the rest is easy.