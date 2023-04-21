package day02

func Part1(inputFile string) string {
	return solve(
		inputFile, func(d *data, i int) {
			d.depth -= i
		}, func(d *data, i int) {
			d.depth += i
		}, func(d *data, i int) {
			d.horizontalPosition += i
		})
}
