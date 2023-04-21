package day02

func Part2(inputFile string) string {
	return solve(
		inputFile, func(d *data, i int) {
			d.aim -= i
		}, func(d *data, i int) {
			d.aim += i
		}, func(d *data, i int) {
			d.horizontalPosition += i
			d.depth += d.aim * i
		})
}
