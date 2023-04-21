# Advent of Code 2021

Go 1.20 solutions to the Advent of Code problems. Check out <https://adventofcode.com/2021>

- Put `input.txt` as downloaded from the Advent of Code website into the folder corresponding to the day
- Run `go generate`, `go build`, `./aoc2021 day part`
- Run `./aoc2021 -h` for options

## Dependencies

For building:

- PowerShell (if using `build.ps1`)
- goimports (if using `goimports` instead of `go fmt`)
- modd (if using modd)

<https://docs.microsoft.com/en-us/powershell/scripting/install/installing-powershell>

```
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/cortesi/modd/cmd/modd@latest
```

For code:

- <https://github.com/zellyn/kooky> - for reading Advent of Code session cookie from browser cache
- <https://github.com/mattn/godown> - for converting Advent of Code puzzle description from html to markdown

# Notes

- Input is never validated, it is assumed to be correct
- Tested on PC (intel i9) and Raspberry Pi
- On the test PC each puzzle part did not take more then 100 milliseconds to run
- Downloaded puzzle texts is in the [puzzles](puzzles) folder. It is not always entirely accurate as [godown](https://github.com/mattn/godown) occasionally makes a mistake.
