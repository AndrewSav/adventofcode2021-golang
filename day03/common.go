package day03

// Given an array of lines and a position, returns an array of lines where given position is occupied by 1
// and an array of line where given poistion is ocupied by something else (0)
func split(input []string, pos int) (ones []string, zeroes []string) {
	for _, l := range input {
		if l[pos] == '1' {
			ones = append(ones, l)
		} else {
			zeroes = append(zeroes, l)
		}
	}
	return
}
