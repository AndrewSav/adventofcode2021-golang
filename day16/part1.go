package day16

import (
	"fmt"
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
	return fmt.Sprint(sumVersions(loadAndParse(inputFile)))
}
