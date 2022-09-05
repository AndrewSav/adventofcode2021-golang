package day20

import (
	"fmt"
)

func Part1(inputFile string) string {
	imageEnhancementAlgorithm, image := loadInput(inputFile)

	image.Enhance(imageEnhancementAlgorithm)
	image.Enhance(imageEnhancementAlgorithm)

	return fmt.Sprintf("%d", image.onCount)
}
