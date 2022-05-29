package day07

func Part1(inputFile string) string {
	return solve(inputFile, func(n int) int { return n })
}
