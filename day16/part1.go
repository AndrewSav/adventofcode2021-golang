package day16

import (
	"aoc2021/util"
	"encoding/hex"
	"fmt"
	"strings"
)

func sumVersions(p packet) int {
	if p.isLiteral() {
		return p.getVersion()
	}
	o := p.(*operator)
	result := o.version
	for _, c := range o.subPackets {
		result += sumVersions(c)
	}
	return result
}

func Part1(inputFile string) string {
	lines := util.ReadInput(inputFile)
	bytes, _ := hex.DecodeString(lines[0])
	var sb strings.Builder
	for _, b := range bytes {
		fmt.Fprintf(&sb, "%08b", b)
	}
	input := sb.String()
	v, _ := parse(input)
	//fmt.Println(reminder)
	//fmt.Println(v.string())
	return fmt.Sprint(sumVersions(v))
}
