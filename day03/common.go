package day03

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
