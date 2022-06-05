package util

import (
	"fmt"
	"strings"
	"sync"
)

const fontSmallLetters = "ABCEFGHIJKLOPRSUYZ"

var smallLetterWidths = map[rune]int{
	'A': 4,
	'B': 4,
	'C': 4,
	'E': 4,
	'F': 4,
	'G': 4,
	'H': 4,
	'I': 3,
	'J': 4,
	'K': 4,
	'L': 4,
	'O': 4,
	'P': 4,
	'R': 4,
	'S': 4,
	'U': 4,
	'Y': 5,
	'Z': 4,
}

const fontSmall = `
.##..###...##..####.####..##..#..#.###...##.#..#.#.....##..###..###...###.#..#.#...#.####
#..#.#..#.#..#.#....#....#..#.#..#..#.....#.#.#..#....#..#.#..#.#..#.#....#..#.#...#....#
#..#.###..#....###..###..#....####..#.....#.##...#....#..#.#..#.#..#.#....#..#..#.#....#.
####.#..#.#....#....#....#.##.#..#..#.....#.#.#..#....#..#.###..###...##..#..#...#....#..
#..#.#..#.#..#.#....#....#..#.#..#..#..#..#.#.#..#....#..#.#....#.#.....#.#..#...#...#...
#..#.###...##..####.#.....###.#..#.###..##..#..#.####..##..#....#..#.###...##....#...####`

const fontLargeLetters = "ABCEFGHJKLNPRXZ"

const fontLarge = `
..##...#####...####..######.######..####..#....#....###.#....#.#......#....#.#####..#####..#....#.######
.#..#..#....#.#....#.#......#......#....#.#....#.....#..#...#..#......##...#.#....#.#....#.#....#......#
#....#.#....#.#......#......#......#......#....#.....#..#..#...#......##...#.#....#.#....#..#..#.......#
#....#.#....#.#......#......#......#......#....#.....#..#.#....#......#.#..#.#....#.#....#..#..#......#.
#....#.#####..#......#####..#####..#......######.....#..##.....#......#.#..#.#####..#####....##......#..
######.#....#.#......#......#......#..###.#....#.....#..##.....#......#..#.#.#......#..#.....##.....#...
#....#.#....#.#......#......#......#....#.#....#.....#..#.#....#......#..#.#.#......#...#...#..#...#....
#....#.#....#.#......#......#......#....#.#....#.#...#..#..#...#......#...##.#......#...#...#..#..#.....
#....#.#....#.#....#.#......#......#...##.#....#.#...#..#...#..#......#...##.#......#....#.#....#.#.....
#....#.#####...####..######.#.......###.#.#....#..###...#....#.######.#....#.#......#....#.#....#.######
`

func GetOCRKey(plot [][]rune, x, y, width, height int) string {
	var b strings.Builder
	for dy := y; dy < y+height; dy++ {
		fmt.Fprintf(&b, "%s", string(plot[dy][x:x+width]))
	}
	return b.String()
}

var lock = &sync.Mutex{}
var smallFontMap map[string]rune

func createFontMap(font, letters string, widths map[rune]int, height int) (result map[string]rune) {
	result = make(map[string]rune)
	lines := strings.Split(font, "\n")[1:]
	plot := make([][]rune, len(lines))
	for y := range plot {
		plot[y] = []rune(lines[y])
	}
	offset := 0
	for _, r := range letters {
		width := widths[r]
		result[GetOCRKey(plot, offset, 0, width, height)] = r
		offset += width + 1
	}
	return
}

func ensureMaps() {
	lock.Lock()
	defer lock.Unlock()
	if smallFontMap == nil {
		smallFontMap = createFontMap(fontSmall, fontSmallLetters, smallLetterWidths, 6)
	}
}

func GetLetterByOCRKey(ocrKey string) rune {
	ensureMaps()
	return smallFontMap[ocrKey]
}

func OCR2021Day13Part2(plot [][]rune) string {
	var b strings.Builder
	var width int
	for x := 0; x < len(plot[0]); x = x + width {
		letter := GetLetterByOCRKey(GetOCRKey(plot, x, 0, 4, 6))
		fmt.Fprintf(&b, "%s", string(letter))
		width = smallLetterWidths[letter] + 1
	}
	return b.String()
}
