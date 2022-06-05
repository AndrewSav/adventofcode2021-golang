package day14

import (
	"fmt"
	"sort"
)

func printScores(scores map[byte]int) {
	keys := make([]byte, 0, len(scores))

	for k := range scores {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, k := range keys {
		fmt.Println(string(k), scores[k])
	}

	fmt.Println()
	sort.SliceStable(keys, func(i, j int) bool {
		return scores[keys[i]] < scores[keys[j]]
	})

	for _, k := range keys {
		fmt.Println(string(k), scores[k])
	}
}
