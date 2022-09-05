package day20

import (
	"fmt"
)

func Part2(inputFile string) string {
	imageEnhancementAlgorithm, image := loadInput(inputFile)

	for i := 0; i < 50; i++ {
		image.Enhance(imageEnhancementAlgorithm)
	}

	return fmt.Sprintf("%d", image.onCount)
}
