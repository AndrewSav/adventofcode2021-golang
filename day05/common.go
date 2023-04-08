package day05

import (
	"aoc2021/util"
	"fmt"
	"image"
	"image/color"
	"regexp"
)

func parseLine(s string) image.Rectangle {
	r := regexp.MustCompile(`^(\d{1,3}),(\d{1,3}) -> (\d{1,3}),(\d{1,3})$`)
	rr := r.FindStringSubmatch(s)
	return image.Rectangle{image.Pt(util.MustAtoi(rr[1]), util.MustAtoi(rr[2])), image.Pt(util.MustAtoi(rr[3]), util.MustAtoi(rr[4]))}
}

type myImage struct {
	// we are using this type to represent how many lines go through a given point, using the gray value for that number
	// this saves us defining our own structure
	image.Gray
}

func newImage(r image.Rectangle) *myImage {
	return &myImage{*image.NewGray(r)}
}

func (img *myImage) drawOrthogonal(r image.Rectangle) {
	if r := r.Canon(); r.Min.Y == r.Max.Y { // horizontal line
		for ; r.Min.X <= r.Max.X; r.Min.X++ {
			img.SetGray(r.Min.X, r.Min.Y, color.Gray{img.GrayAt(r.Min.X, r.Min.Y).Y + 1}) // increment each point
		}
	}
	if r := r.Canon(); r.Min.X == r.Max.X { // vertical line
		for ; r.Min.Y <= r.Max.Y; r.Min.Y++ {
			img.SetGray(r.Min.X, r.Min.Y, color.Gray{img.GrayAt(r.Min.X, r.Min.Y).Y + 1}) // increment each point
		}
	}
}

func (img *myImage) drawDiagonal(r image.Rectangle) {
	if util.Abs(r.Min.Y-r.Max.Y) == util.Abs(r.Min.X-r.Max.X) { // diagonal line
		stepx, stepy := 1, 1
		if r.Min.Y > r.Max.Y {
			stepy = -1
		}
		if r.Min.X > r.Max.X {
			stepx = -1
		}
		for ; r.Min.X != r.Max.X; r.Min.X, r.Min.Y = r.Min.X+stepx, r.Min.Y+stepy {
			img.SetGray(r.Min.X, r.Min.Y, color.Gray{img.GrayAt(r.Min.X, r.Min.Y).Y + 1}) // increment each point
		}
		img.SetGray(r.Min.X, r.Min.Y, color.Gray{img.GrayAt(r.Min.X, r.Min.Y).Y + 1})
	}
}

func solve(inputFile string, includeDiagonal bool) string {
	lines := util.ReadInput(inputFile)
	data := []image.Rectangle{}
	bounds := image.Rectangle{}
	for _, l := range lines {
		item := parseLine(l)
		data = append(data, item) // collects all input lines in `data` array
		// Union cannot combine "empty" rectangles (horizontal and vertical lines), so we are using Inset to make them non empty
		bounds = bounds.Union(item.Canon().Inset(-1)) // and also determines the combined boundss of all lines
	}
	img := newImage(bounds)
	for _, d := range data {
		img.drawOrthogonal(d)
		if includeDiagonal {
			img.drawDiagonal(d)
		}
	}
	count := 0
	for x := bounds.Min.X; x <= bounds.Max.X; x++ {
		for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
			if img.GrayAt(x, y).Y > 1 {
				count++
			}
		}
	}
	return fmt.Sprint(count)
}
