package day17

import (
	"fmt"
)

func Part1(inputFile string) string {
	_, _, y1, _ := getInput(inputFile)
	return fmt.Sprint(y1 * (y1 + 1) / 2)
}
