package day16

import (
	"strconv"
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
	value int
}

type operator struct {
	header
	lengthTypeId int
	length       int
	subPackets   []packet
}

type packet interface {
	isLiteral() bool
	getVersion() int
	string() string
	getValue() int
}

func parseLiteral(input string, h header) (packet, string) {
	value := 0
	for previousPrefix := "1"; previousPrefix == "1"; previousPrefix, input = input[0:1], input[5:] {
		value = value*16 + parseBinary(input[1:5])
	}
	return &literal{header: h, value: value}, input
}

func parseOperator(input string, h header) (packet, string) {
	result := operator{header: h, lengthTypeId: parseBinary(input[0:1])}
	if result.lengthTypeId == 0 {
		result.length = parseBinary(input[1:16])
		input = input[16:]
		targetLength := len(input) - result.length
		var p packet
		for len(input) > targetLength {
			p, input = parse(input)
			result.subPackets = append(result.subPackets, p)
		}
		return &result, input
	} else {
		result.length = parseBinary(input[1:12])
		input = input[12:]
		var p packet
		for i := 0; i < result.length; i++ {
			p, input = parse(input)
			result.subPackets = append(result.subPackets, p)
		}
		return &result, input
	}
}

func parse(input string) (packet, string) {
	h := header{version: parseBinary(input[:3]), typeId: parseBinary(input[3:6])}
	if h.isLiteral() {
		return parseLiteral(input[6:], h)
	} else {
		return parseOperator(input[6:], h)
	}
}

func parseBinary(s string) int {
	i, _ := strconv.ParseInt(s, 2, 32)
	return int(i)
}
