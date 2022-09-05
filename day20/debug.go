package day20

import (
	"fmt"
)

func (im *image) print() {
	for y := im.minY; y < im.maxY; y++ {
		for x := im.minX; x < im.maxX; x++ {
			if im.GetValue(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
