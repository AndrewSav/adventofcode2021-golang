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

	if flags.PrintSessionCooike && flags.DownloadDescriptions {
		log.Fatalf("flags -q and -p are incompatible")
	}

	if flags.PrintSessionCooike {
		cookie := util.GetBrowserCookie()
		if cookie == "" {
			log.Fatalf("failed to get cookie from browser")
		} else {
			fmt.Println(cookie)
			os.Exit(0)
		}
	}

	if flags.DownloadDescriptions {
		cookie := flags.SessionCookie
		if cookie == "" {
			cookie, _ = util.TryGetCookie()
		}
		util.DownloadDescriptions(cookie)
		os.Exit(0)
	}

	if flags.Day == 0 {
		flags.Day = days[len(days)-1].Day
	}

	if flags.Part == 0 {
		for _, day := range days {
			if day.Day == flags.Day {
				flags.Part = day.Parts[len(day.Parts)-1].Part
				break
			}
		}
	}

	if flags.Variant == "" {
		for _, day := range days {
			if day.Day == flags.Day {
				partFound := false
				for _, part := range day.Parts {
					if !partFound && flags.Part == part.Part {
						partFound = true
					}
					if partFound {
						if flags.Part == part.Part {
							flags.Variant = part.Variant
						} else {
							break
						}
					}
				}
				break
			}
		}
	}

	if flags.InputFile == "" {
		flags.InputFile = util.GetDefautInputFilePath(flags.Day)
	}

	if flags.DownloadInput && flags.SessionCookie == "" {
		cookie, err := util.TryGetCookie()
		if err != nil {
			log.Fatalf("%v", err)
		}
		flags.SessionCookie = cookie
	}

	var solutions []func() (string, int, int, string)
	if flags.All {
		flags.Verbose = true
		solutions = getRunAll()
		if flags.DownloadInput {
			for d := 1; d <= days[len(days)-1].Day; d++ {
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

func getRunAll() (result []func() (string, int, int, string)) {
	for _, day := range days {
		for _, part := range day.Parts {
			day, part, variant := day.Day, part.Part, part.Variant // prevent variable capture in the closure
			result = append(result, func() (string, int, int, string) {
				return run(day, part, variant, util.GetDefautInputFilePath(day))
			})
		}
	}
	return
}
