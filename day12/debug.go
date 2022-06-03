package day12

import (
	"fmt"
	"strings"
)

func printPath(p *nodePath) {
	result := []string{}
	for _, n := range p.array {
		result = append(result, n.name)
	}
	fmt.Println(strings.Join(result, ","))
}
