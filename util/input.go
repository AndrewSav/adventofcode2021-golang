package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func ReadInput(inputFile string) (result []string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func Atoi(data []string) (result []int) {
	for index, s := range data {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("util.Atoi, index %d: %v", index, err)
		}
		result = append(result, i)
	}
	return
}

func GetDefautInputFilePath(day int) string {
	return filepath.FromSlash(fmt.Sprintf("day%02d/input.txt", day))
}
