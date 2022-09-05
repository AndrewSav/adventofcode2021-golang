package day20

import (
	"aoc2021/util"
	"fmt"
)

var imageEnhancementAlgorithm [512]bool

type image struct {
	data                   map[int]map[int]bool
	minX, minY, maxX, maxY int
	infinityValue          bool
	onCount                int
}

func (im *image) GetValue(x, y int) bool {
	value, _ := im.GetValueOrOk(x, y)
	return value
}

func (im *image) GetValueInt(x, y int) int {
	if im.GetValue(x, y) {
		return 1
	} else {
		return 0
	}
}

func (im *image) GetValueOrOk(x, y int) (bool, bool) {
	inner, ok := im.data[y]
	if !ok {
		return im.infinityValue, ok
	}
	result, ok := inner[x]
	if !ok {
		return im.infinityValue, ok
	}
	return result, ok
}

func (im *image) SetValue(x, y int, value bool) {
	if im.data[y] == nil {
		im.data[y] = map[int]bool{}
	}
	im.data[y][x] = value
	if value {
		im.onCount++
	}
}

func (im *image) GetMask(x, y int) int {
	result := im.GetValueInt(x-1, y-1)
	result <<= 1
	result |= im.GetValueInt(x, y-1)
	result <<= 1
	result |= im.GetValueInt(x+1, y-1)
	result <<= 1
	result |= im.GetValueInt(x-1, y)
	result <<= 1
	result |= im.GetValueInt(x, y)
	result <<= 1
	result |= im.GetValueInt(x+1, y)
	result <<= 1
	result |= im.GetValueInt(x-1, y+1)
	result <<= 1
	result |= im.GetValueInt(x, y+1)
	result <<= 1
	result |= im.GetValueInt(x+1, y+1)
	return result
}

func (im *image) Enhance(alg [512]bool) {
	newImage := image{data: map[int]map[int]bool{}}
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
	if im.infinityValue {
		index = 512 - 1
	}
	im.infinityValue = alg[index]
}

func Part1(inputFile string) string {
	data := util.ReadInput(inputFile)
	for i := 0; i < len(imageEnhancementAlgorithm); i++ {
		imageEnhancementAlgorithm[i] = (data[0][i]) == '#'
	}
	const offset = 2
	image := image{maxX: len(data) - offset, maxY: len(data[offset]), data: map[int]map[int]bool{}}
	for i := offset; i < len(data); i++ {
		for j, c := range data[i] {
			image.SetValue(j, i-offset, c == '#')
		}
	}

	image.Enhance(imageEnhancementAlgorithm)
	image.Enhance(imageEnhancementAlgorithm)

	return fmt.Sprintf("%d", image.onCount)
}
