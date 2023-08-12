package day14

// This implements how the score is calculated
// according to the puzzle description
func getScore[T int | int64](scores map[byte]T) T {
	var leastCommon, mostCommon T
	for _, v := range scores {
		leastCommon, mostCommon = v, v
		break
	}
	for _, v := range scores {
		mostCommon = max(mostCommon, v)
		leastCommon = min(leastCommon, v)
	}
	return mostCommon - leastCommon
}
