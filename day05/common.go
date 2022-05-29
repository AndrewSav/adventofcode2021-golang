package day05

import (
	"aoc2021/util"
	"fmt"
	"image"
	"image/color"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func toNum(s string) int {
	if i, err := strconv.ParseInt(strings.Trim(s, " "), 10, 32); err != nil {
		log.Fatalf("cannot convert string '%s' to a number: %v", s, err)
	} else {
		return int(i)
	}
	panic("unexpected code path")
}

func parseLine(s string) image.Rectangle {
	r := regexp.MustCompile(`^(\d{1,3}),(\d{1,3}) -> (\d{1,3}),(\d{1,3})$`)
	rr := r.FindStringSubmatch(s)
	if rr == nil {
		log.Fatalf("cannot parse '%s'", s)
	}
	return image.Rectangle{image.Point{toNum(rr[1]), toNum(rr[2])}, image.Point{toNum(rr[3]), toNum(rr[4])}}
}

type myImage struct {
	image.Gray
}

func newImage(r image.Rectangle) *myImage {
	return &myImage{*image.NewGray(r)}
}

func (img *myImage) drawOrthogonal(r image.Rectangle) {
	if r := r.Canon(); r.Min.Y == r.Max.Y {
		for ; r.Min.X <= r.Max.X; r.Min.X++ {
			img.SetGray(r.Min.X, r.Min.Y, color.Gray{img.GrayAt(r.Min.X, r.Min.Y).Y + 1})
		}
	}
	if r := r.Canon(); r.Min.X == r.Max.X {
		for ; r.Min.Y <= r.Max.Y; r.Min.Y++ {
			img.SetGray(r.Min.X, r.Min.Y, color.Gray{img.GrayAt(r.Min.X, r.Min.Y).Y + 1})
		}
	}
}

func (img *myImage) drawDiagonal(r image.Rectangle) {
	if abs(r.Min.Y-r.Max.Y) == abs(r.Min.X-r.Max.X) {
		stepx, stepy := 1, 1
		if r.Min.Y > r.Max.Y {
			stepy = -1
		}
		if r.Min.X > r.Max.X {
			stepx = -1
		}
		for ; r.Min.X != r.Max.X; r.Min.X, r.Min.Y = r.Min.X+stepx, r.Min.Y+stepy {
			img.SetGray(r.Min.X, r.Min.Y, color.Gray{img.GrayAt(r.Min.X, r.Min.Y).Y + 1})
		}
		img.SetGray(r.Min.X, r.Min.Y, color.Gray{img.GrayAt(r.Min.X, r.Min.Y).Y + 1})
	}
}

func solve(inputFile string, includeDiagonal bool) string {
	lines := util.ReadInput(inputFile)
	data := []image.Rectangle{}
	bound := image.Rectangle{}
	for _, l := range lines {
		item := parseLine(l)
		data = append(data, item)
		bound = bound.Union(item.Canon())
	}
	bound = bound.Inset(-1)
	img := newImage(bound)
	for _, d := range data {
		img.drawOrthogonal(d)
		if includeDiagonal {
			img.drawDiagonal(d)
		}
	}
	count := 0
	for x := bound.Min.X; x <= bound.Max.X; x++ {
		for y := bound.Min.Y; y <= bound.Max.Y; y++ {
			if img.GrayAt(x, y).Y > 1 {
				count++
			}
		}
	}
	return fmt.Sprint(count)
}
