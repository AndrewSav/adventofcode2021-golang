package day08

func getEntireNumbers(acc *int, v int) {
	*acc = *acc*10 + v
}

func Part2(inputFile string) string {
	return solve(inputFile, getEntireNumbers)
}
