package day16

import (
	"fmt"
	"strings"
)

func (h header) string() string {
	return fmt.Sprintf("Version: %d, Type id: %d", h.version, h.typeId)
}

func (l *literal) string() string {
	return fmt.Sprintf("%s, Value: %d", l.header.string(), l.value)
}

func (l *operator) string() string {
	return strings.Join(l.strings(), "\n")
}

func (l *operator) strings() []string {
	result := []string{fmt.Sprintf("%s, Length type id: %d, Lenght: %d, Value: %d;", l.header.string(), l.lengthTypeId, l.length, l.getValue())}
	for i, p := range l.subPackets {
		if p.isLiteral() {
			result = append(result, fmt.Sprintf("  %s;<-%d", p.string(), i))
		} else {
			for _, s := range p.(*operator).strings() {
				result = append(result, fmt.Sprintf("  %s<-%d", s, i))
			}
		}
	}
	return result
}
