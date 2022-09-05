package day20

import (
	"aoc2021/util"
)

// TODO: split
type image struct {
	data          [][]int
	onCount       int
	infinityValue int
}

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

func (im *image) Enhance(alg [512]int, start, length int) {
	//fmt.Printf("start: %d, length: %d\n", start, length)
	im.onCount = 0
	newImage := make([][]int, len(im.data))
	index := 0
	if im.infinityValue == 1 {
		index = 512 - 1
	}
	im.infinityValue = alg[index]
	for y := 0; y < len(im.data); y++ {
		newImage[y] = make([]int, len(im.data))
		for x := 0; x < len(im.data); x++ {
			if x >= start && x < start+length && y >= start && y < start+length {
				mask := im.GetMask(x, y)
				bit := alg[mask]
				//fmt.Printf("x: %d, y: %d, mask: %d, bit: %d\n", x, y, mask, bit)
				newImage[y][x] = bit
				if bit == 1 {
					im.onCount++
				}
			}
		}
	}
	for y := start - 2; y <= start+length+1; y++ {
		newImage[y][start-2] = im.infinityValue
		newImage[y][start+length+1] = im.infinityValue
		newImage[y][start-1] = im.infinityValue
		newImage[y][start+length] = im.infinityValue
	}
	for x := start - 2; x <= start+length+1; x++ {
		newImage[start-2][x] = im.infinityValue
		newImage[start+length+1][x] = im.infinityValue
		newImage[start-1][x] = im.infinityValue
		newImage[start+length][x] = im.infinityValue
	}
	im.data = newImage
}

func cell(c byte) int {
	if c == '#' {
		return 1
	}
	return 0
}

func loadInput(inputFile string, cycles int) int {
	var imageEnhancementAlgorithm [512]int
	var im image
	data := util.ReadInput(inputFile)
	for i := 0; i < len(imageEnhancementAlgorithm); i++ {
		imageEnhancementAlgorithm[i] = cell(data[0][i])
	}
	const offset = 2
	padding := cycles + 2
	maxDimension := len(data) - offset + padding*2

	im = image{data: make([][]int, maxDimension)}
	for i := 0; i < maxDimension; i++ {
		im.data[i] = make([]int, maxDimension)
	}
	for i := offset; i < len(data); i++ {
		for j := range data[i] {
			//fmt.Printf("x: %d, y: %d, val: %d\n", j, i, cell(data[i][j]))
			im.data[i-offset+padding][j+padding] = cell(data[i][j])
		}
	}
	//im.print()
	//fmt.Println("---")
	for i := 0; i < cycles; i++ {
		dimension := len(data) - offset + (i+1)*2
		start := (maxDimension - dimension) / 2
		im.Enhance(imageEnhancementAlgorithm, start, dimension)
		//if i <= 2 {
		//im.print()
		//fmt.Println("---")
		//}
	}
	return im.onCount
}
