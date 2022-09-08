package day21

import (
	"log"
	"regexp"
	"strconv"
)

func getPosition(s string) int {
	r := regexp.MustCompile(`Player \d+ starting position: (\d+)`)
	match := r.FindStringSubmatch(s)
	if match == nil {
		log.Fatalf("line '%s' cannot be matched", s)
	}
	i, err := strconv.Atoi(match[1])
	if err != nil {
		panic(err)
	}
	return i
}
