package day05

import "fmt"

func (img *myImage) print() {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			fmt.Printf("%d", img.GrayAt(x, y).Y)
		}
		fmt.Println()
	}

}
