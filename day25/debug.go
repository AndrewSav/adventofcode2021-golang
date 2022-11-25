package day25

import (
	"fmt"
)

func dump(sea [][]rune) {
	count := 0
	for _, l := range sea {
		for _, c := range l {
			if c == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%c", c)
				count++
			}
		}
		fmt.Println()
	}
	fmt.Printf("Count: %d\n", count)
}
