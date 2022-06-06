package main

import (
	"aoc2021/util"
	"fmt"
	"log"
	"os"
	"time"
)

//go:generate go run ./gen
func main() {
	log.SetFlags(log.Lshortfile)

	flags := util.ParseArgs()

	if flags.PrintSessionCooike {
		cookie, err := util.GetChromeCookie()
		if err != nil {
			log.Fatalf("%v", err)
		} else {
			fmt.Println(cookie)
			os.Exit(0)
		}
	}

	latestDay, latestPart, latestVariant := getLatest()

	if flags.Day == 0 {
		flags.Day = latestDay
	}

	if flags.Part == 0 {
		flags.Part = latestPart
	}

	if flags.Variant == "" {
		flags.Variant = latestVariant
	}

	if flags.InputFile == "" {
		flags.InputFile = util.GetDefautInputFilePath(flags.Day)
	}

	if flags.DownloadInput && flags.SessionCookie == "" {
		flags.SessionCookie = util.TryGetCookie()
	}

	var solutions []func() (string, int, int, string)
	if flags.All {
		flags.Verbose = true
		solutions = getRunAll()
		if flags.DownloadInput {
			for d := 1; d <= latestDay; d++ {
				util.DownloadInput(flags.SessionCookie, d, util.GetDefautInputFilePath(d))
			}
		}
	} else {
		if flags.DownloadInput {
			util.DownloadInput(flags.SessionCookie, flags.Day, flags.InputFile)
		}
		solutions = append(solutions, func() (string, int, int, string) { return run(flags.Day, flags.Part, flags.Variant, flags.InputFile) })
	}

	for _, s := range solutions {
		if flags.Verbose {
			start := time.Now()
			result, eday, epart, evariant := s()
			duration := time.Since(start)
			fmt.Printf("Day %d, Part %d%s, %s (%v)\n", eday, epart, evariant, result, duration)
		} else {
			result, _, _, _ := s()
			fmt.Println(result)
		}

	}
}
