package day16

import (
	"aoc2021/util"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type header struct {
	version int
	typeId  int
}

func (h header) getVersion() int {
	return h.version
}

func (h header) isLiteral() bool {
	return h.typeId == 4
}

type literal struct {
	header
	value int64
}

type operator struct {
	header
	lengthTypeId int // 0 or 1
	length       int // different meaning depending on the value of lengthTypeId as per puzzle description
	subPackets   []packet
}

type packet interface {
	isLiteral() bool
	getVersion() int
	getValue() int64
}

func bitStringToNum(s string) int {
	i, _ := strconv.ParseInt(s, 2, 32)
	return int(i)
}

type bitStream string

func (b *bitStream) getBits(width int) int {
	result := bitStringToNum(string((*b)[:width]))
	*b = (*b)[width:]
	return result
}

func (b *bitStream) parseLiteral(h header) packet {
	value, stop := int64(0), false
	for !stop {
		stop = b.getBits(1) == 0
		value = value*16 + int64(b.getBits(4))
	}
	return &literal{header: h, value: value}
}

func (b *bitStream) parseOperator(h header) packet {
	result := operator{header: h, lengthTypeId: b.getBits(1)}
	if result.lengthTypeId == 0 {
		result.length = b.getBits(15)
		targetLength := len(*b) - result.length
		for len(*b) > targetLength {
			result.subPackets = append(result.subPackets, b.parse())
		}
	} else {
		result.length = b.getBits(11)
		for i := 0; i < result.length; i++ {
			result.subPackets = append(result.subPackets, b.parse())
		}
	}
	return &result
}

func (b *bitStream) parse() packet {
	h := header{version: b.getBits(3), typeId: b.getBits(3)}
	if h.isLiteral() {
		return b.parseLiteral(h)
	} else {
		return b.parseOperator(h)
	}
}

func loadAndParse(inputFile string) packet {
	lines := util.ReadInput(inputFile)
	bytes, _ := hex.DecodeString(lines[0])
	var sb strings.Builder
	// We are getting a string of 1's and 0's here
	for _, b := range bytes {
		fmt.Fprintf(&sb, "%08b", b)
	}
	s := bitStream(sb.String())
	return s.parse()
}
