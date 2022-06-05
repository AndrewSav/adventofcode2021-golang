package day13

import (
	"fmt"
)

func Part1(inputFile string) string {
	points, folds := readData(inputFile)
	return fmt.Sprint(len(fold(folds[0], points)))
}
