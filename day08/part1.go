package day08

func getSimpleDigits(acc *int, v int) {
	if v == 1 || v == 4 || v == 7 || v == 8 {
		*acc++
	}
}

func Part1(inputFile string) string {
	return solve(inputFile, getSimpleDigits)
}
