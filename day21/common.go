package day21

import (
	"regexp"
	"strconv"
)

func getPosition(s string) int {
	r := regexp.MustCompile(`Player \d+ starting position: (\d+)`)
	i, _ := strconv.Atoi(r.FindStringSubmatch(s)[1])
	return i
}
