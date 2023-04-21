package day13

import (
	"fmt"
)

func Part1(inputFile string) string {
	points, folds := readData(inputFile)
	// execute the first fold only and count the number of resulting points
	return fmt.Sprint(len(fold(folds[0], points)))
}
