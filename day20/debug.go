package day20

import (
	"fmt"
)

func (im *image) print() {
	for y := 0; y < len(im.data); y++ {
		for x := 0; x < len(im.data); x++ {
			if im.data[y][x] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
