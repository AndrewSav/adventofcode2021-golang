package day20

import (
	"aoc2021/util"
)

type image struct {
	data          [][]int
	infinityValue int // This is what's contained in all cells that we are not keeping track of in .data
}

// This is getting the number corrsepending to the 3x3 square
// centered at (x,y) as described in the puzzle
func (im *image) GetMask(x, y int) int {
	result := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			result <<= 1
			result |= im.data[y+dy][x+dx]
		}
	}
	return result
}

// The algorythm runs on ever expanding square, that expand once cell each side
// every time Enhance is called.
// start is where in image that sqare starts (both x and y)
// and dimension is how many cells long is that square side
func (im *image) Enhance(alg [512]int, start, length int) int {
	var (
		onCount  = 0
		newImage = make([][]int, len(im.data))
		index    = 0 // The bit mask is either all zeroes
	)
	if im.infinityValue == 1 {
		index = 512 - 1 // Or the bit mask is all ones
	}
	im.infinityValue = alg[index]
	for y := 0; y < len(im.data); y++ {
		newImage[y] = make([]int, len(im.data[y]))
		if y >= start && y < start+length {
			for x := start; x < start+length; x++ {
				bit := alg[im.GetMask(x, y)]
				newImage[y][x] = bit
				if bit == 1 {
					onCount++
				}
			}
		}
	}
	// We do not really need to fill in all empty cells in .data
	// but we need a 2-wide border for .GetMask to work correctly
	for i := start - 2; i < start+length+2; i++ {
		newImage[i][start-2] = im.infinityValue
		newImage[i][start+length+1] = im.infinityValue
		newImage[i][start-1] = im.infinityValue
		newImage[i][start+length] = im.infinityValue
		newImage[start-2][i] = im.infinityValue
		newImage[start+length+1][i] = im.infinityValue
		newImage[start-1][i] = im.infinityValue
		newImage[start+length][i] = im.infinityValue
	}
	im.data = newImage
	return onCount
}

func cell(c byte) int {
	if c == '#' {
		return 1
	}
	return 0
}

func loadInput(inputFile string, cycles int) int {
	var imageEnhancementAlgorithm [512]int
	data := util.ReadInput(inputFile)
	for i := 0; i < len(imageEnhancementAlgorithm); i++ {
		imageEnhancementAlgorithm[i] = cell(data[0][i])
	}
	const offset = 2 // First two lines in the file are the Image Enhancment Algorythm and a new line
	var (
		im           image
		padding      = cycles + 2                     // add padding one cell for each enhancement cycle plus two-wide to account for infinite space on each side
		maxDimension = len(data) - offset + padding*2 // the final square side
	)
	im = image{data: make([][]int, maxDimension)}
	for i := 0; i < maxDimension; i++ {
		im.data[i] = make([]int, maxDimension)
	}
	for i := offset; i < len(data); i++ {
		for j := range data[i] {
			im.data[i-offset+padding][j+padding] = cell(data[i][j])
		}
	}
	onCount := 0
	for i := 0; i < cycles; i++ {
		dimension := len(data) - offset + (i+1)*2
		start := (maxDimension - dimension) / 2
		onCount = im.Enhance(imageEnhancementAlgorithm, start, dimension)
	}
	return onCount
}
