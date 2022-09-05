package day20

import (
	"aoc2021/util"
)

type image struct {
	data                   map[int]map[int]int
	minX, minY, maxX, maxY int
	infinityValue          int
	onCount                int
}

func (im *image) GetValue(x, y int) int {
	inner, ok := im.data[y]
	if !ok {
		return im.infinityValue
	}
	result, ok := inner[x]
	if !ok {
		return im.infinityValue
	}
	return result
}

func (im *image) SetValue(x, y, value int) {
	if im.data[y] == nil {
		im.data[y] = map[int]int{}
	}
	im.data[y][x] = value
	if value == 1 {
		im.onCount++
	}
}

func (im *image) GetMask(x, y int) int {
	result := im.GetValue(x-1, y-1)
	result <<= 1
	result |= im.GetValue(x, y-1)
	result <<= 1
	result |= im.GetValue(x+1, y-1)
	result <<= 1
	result |= im.GetValue(x-1, y)
	result <<= 1
	result |= im.GetValue(x, y)
	result <<= 1
	result |= im.GetValue(x+1, y)
	result <<= 1
	result |= im.GetValue(x-1, y+1)
	result <<= 1
	result |= im.GetValue(x, y+1)
	result <<= 1
	result |= im.GetValue(x+1, y+1)
	return result
}

func (im *image) Enhance(alg [512]int) {
	newImage := image{data: map[int]map[int]int{}}
	for y := im.minY - 1; y < im.maxY+1; y++ {
		for x := im.minX - 1; x < im.maxX+1; x++ {
			newImage.SetValue(x, y, alg[im.GetMask(x, y)])
		}
	}
	im.data = newImage.data
	im.onCount = newImage.onCount
	im.minY -= 1
	im.minX -= 1
	im.maxX += 1
	im.maxY += 1
	index := 0
	if im.infinityValue == 1 {
		index = 512 - 1
	}
	im.infinityValue = alg[index]
}

func cell(c byte) int {
	if c == '#' {
		return 1
	}
	return 0
}

func loadInput(inputFile string) (imageEnhancementAlgorithm [512]int, im image) {
	data := util.ReadInput(inputFile)
	for i := 0; i < len(imageEnhancementAlgorithm); i++ {
		imageEnhancementAlgorithm[i] = cell(data[0][i])
	}
	const offset = 2
	im = image{maxX: len(data) - offset, maxY: len(data[offset]), data: map[int]map[int]int{}}
	for i := offset; i < len(data); i++ {
		for j := range data[i] {
			im.SetValue(j, i-offset, cell(data[i][j]))
		}
	}
	return
}
