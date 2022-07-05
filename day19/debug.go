package day19

import (
	"fmt"
	"sort"
)

func (p point) print() {
	fmt.Printf("%d, %d, %d", p.x, p.y, p.z)
}

func (p point) println() {
	p.print()
	fmt.Println()
}

func printPointSlice(pp []point) {
	sort.Slice(pp, func(i, j int) bool {
		return pp[i].x < pp[j].x || pp[i].x == pp[j].x && pp[i].y < pp[j].y || pp[i].x == pp[j].x && pp[i].y == pp[j].y && pp[i].z <= pp[j].z
	})
	for _, p := range pp {
		p.println()
	}

}
