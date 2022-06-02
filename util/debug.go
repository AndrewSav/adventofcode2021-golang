package util

import "fmt"

func printPlot(p [][]int) {
	for y, ll := range p {
		for x := range ll {
			fmt.Print(p[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func printPlotWithSeparator(p [][]int, separator string) {
	for y, ll := range p {
		for x := range ll {
			fmt.Print(p[y][x])
			fmt.Print(separator)
		}
		fmt.Println()
	}
	fmt.Println()
}
