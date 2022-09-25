package day22

func isOutOfBound(c *cuboid) (outOfBound bool) {
	for _, d := range c.dimensions {
		if d.min < -50 || d.max > 50 {
			outOfBound = true
			break
		}
	}
	return
}

func Part1(inputFile string) string {
	return solve(inputFile, true)
}
