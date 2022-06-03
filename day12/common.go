package day12

import (
	"aoc2021/util"
	"fmt"
	"strings"
	"unicode"
)

type node struct {
	name  string
	peers []*node
}

func (n node) IsSmall() bool {
	return unicode.IsLower([]rune(n.name)[0])
}

type nodePath struct {
	array           []*node
	lookup          map[*node]int
	visitedTwoSmall bool
}

func (p *nodePath) Add(n *node) {
	if i, ok := p.lookup[n]; ok {
		p.lookup[n] = i + 1
		if n.IsSmall() {
			p.visitedTwoSmall = true
		}
	} else {
		p.lookup[n] = 1
	}
	p.array = append(p.array, n)
}

func (p nodePath) Clone() *nodePath {
	result := &nodePath{lookup: make(map[*node]int), array: make([]*node, len(p.array))}
	copy(result.array, p.array)
	for k, v := range p.lookup {
		result.lookup[k] = v
	}
	result.visitedTwoSmall = p.visitedTwoSmall
	return result
}

func newNodePath(n *node) *nodePath {
	result := &nodePath{lookup: make(map[*node]int)}
	result.Add(n)
	return result
}

func solve(inputFile string, canProceed func(p nodePath, n *node) bool) string {
	lines := util.ReadInput(inputFile)
	nodes := make(map[string]*node)
	for _, l := range lines {
		parts := strings.Split(l, "-")
		addNode := func(name string) {
			_, ok := nodes[name]
			if !ok {
				nodes[name] = &node{name: name}
			}
		}
		addNode(parts[0])
		addNode(parts[1])
		nodes[parts[0]].peers = append(nodes[parts[0]].peers, nodes[parts[1]])
		nodes[parts[1]].peers = append(nodes[parts[1]].peers, nodes[parts[0]])
	}
	count := 0
	for queue := []*nodePath{newNodePath(nodes["start"])}; len(queue) > 0; {
		p := queue[0]
		queue = queue[1:]
		from := p.array[len(p.array)-1]
		for _, to := range from.peers {
			if !canProceed(*p, to) {
				continue
			}
			if to != nodes["end"] {
				newPath := p.Clone()
				newPath.Add(to)
				queue = append(queue, newPath)
			} else {
				count++
			}
		}
	}
	return fmt.Sprint(count)
}
