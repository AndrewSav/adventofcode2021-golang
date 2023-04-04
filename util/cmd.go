package util

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Flags struct {
	Day                int
	Part               int
	Variant            string
	Verbose            bool
	All                bool
	InputFile          string
	DownloadInput      bool
	SessionCookie      string
	PrintSessionCooike bool
}

func ParseArgs() Flags {
	var flags = Flags{}

	fs := flag.NewFlagSet("aoc2021", flag.ExitOnError)

	fs.Usage = func() {
		fmt.Println("This program runs solutions for Advent Of Code 2021")
		fmt.Println("https://adventofcode.com/2021")
		fmt.Printf("Usage: %s [FLAGS...] [day [part [variant]]]\n", filepath.Base(os.Args[0]))
		fmt.Println("  day (1-25) - specifies the day of the problem to run solution for. If not specified, the last existing")
		fmt.Println("  part (1-2) - specifies the part of the problem on the given day to run solution for. If not specified, the last existing")
		fmt.Println("  variant (arbitrary name) - specifies the alternate solution for the part. If not specified, the last existing")
		fmt.Println("Flags:")
		fs.PrintDefaults()
	}

	fs.BoolVar(&flags.Verbose, "v", false, "in addition to the solution also print the day number, the part number and the time taken to run")
	fs.BoolVar(&flags.All, "a", false, "run solutions for all days and parts. day and part arguments are ignored, and -v is implied, if this is specified. inputFile is ignored and the default location for it is used for each day")

	fs.BoolVar(&flags.DownloadInput, "d", false, "downloads the problem input. if -s is not specicified, attempts to read cookie from cookie.txt in the current directory, and if that fails from browser user cookie file in the default location. Download is skipped if the file already exists")
	fs.BoolVar(&flags.PrintSessionCooike, "q", false, "prints out session cookie from browser user cookie file in the default location. If this is specified, all other arguments and flags are ignored")

	fs.StringVar(&flags.InputFile, "i", "", fmt.Sprintf("path to the problem input `/path/to/input.txt`. If not specified defaults to '%s', where XX is the day selected", filepath.FromSlash("dayXX/input.txt")))
	fs.StringVar(&flags.SessionCookie, "s", "", "used with -d to provide `session_cookie` to authenticate with the Advent Of Code web site. Ignored otherwise")

	fs.Parse(os.Args[1:])

	if fs.NArg() > 3 {
		fmt.Printf("want 2 or less arguments, have %d\n", fs.NArg())
		fs.Usage()
		os.Exit(2)
	}

	if fs.NArg() > 2 {
		flags.Variant = fs.Arg(2)
	}

	if fs.NArg() > 1 {
		s := fs.Arg(1)
		if i, err := strconv.Atoi(s); err != nil || i < 1 || i > 2 {
			fmt.Printf("want part to be 1 or 2, have %s\n", s)
			fs.Usage()
			os.Exit(2)
		} else {
			flags.Part = i
		}

	}

	if fs.NArg() > 0 {
		s := fs.Arg(0)
		if i, err := strconv.Atoi(s); err != nil || i < 1 || i > 25 {
			fmt.Printf("want day to be 1 to 25, have %s\n", s)
			fs.Usage()
			os.Exit(2)
		} else {
			flags.Day = i
		}
	}
	return flags
}
