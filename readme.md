# Advent of Code 2021

Go 1.18 solutions to the Advent of Code problems. Check out <https://adventofcode.com/2021>

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
