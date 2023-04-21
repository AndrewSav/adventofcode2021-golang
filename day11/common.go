package day11

type octopi [][]int
type point struct {
	x int
	y int
}

func (o octopi) cycle() int {
	flashes := []point{}
	for y, ll := range o { // first, increment all points, and collect all potential flashes
		for x, l := range ll {
			if l == 9 {
				flashes = append(flashes, point{x: x, y: y})
			}
			o[y][x]++
		}
	}
	count := 0
	for _, f := range flashes { // then execute each collect flash, add number of flashed points to count
		count += o.flash(f.x, f.y)
	}
	for y, ll := range o { // finally reset flashed point back to zero
		for x, l := range ll {
			if l > 9 {
				o[y][x] = 0
			}
		}
	}
	return count
}

func (o octopi) propagateFlash(x, y int) int {
	if y < 0 || x < 0 || y >= len(o) || x >= len(o[y]) { // if we are outside just return
		return 0
	}
	o[y][x]++          // increase energy from the calling neighbour flash
	if o[y][x] != 10 { // we only flash once, on 10, if it's greater than 10 we already flashed here
		return 0
	}
	return o.flash(x, y) // we are 10, so let's flash
}

func (o octopi) flash(x, y int) int {
	count := 1 // point itself flashes
	// each 8 directions from the point can potentially flash
	count += o.propagateFlash(x-1, y-1)
	count += o.propagateFlash(x-1, y)
	count += o.propagateFlash(x-1, y+1)
	count += o.propagateFlash(x, y-1)
	count += o.propagateFlash(x, y+1)
	count += o.propagateFlash(x+1, y-1)
	count += o.propagateFlash(x+1, y)
	count += o.propagateFlash(x+1, y+1)
	return count
}
