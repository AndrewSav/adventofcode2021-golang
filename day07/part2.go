package day07

func Part2(inputFile string) string {
	return solve(inputFile, func(n int) int {
		return n * (n + 1) / 2
	})
}
