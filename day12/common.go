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
	nodes           []*node
	hasOnPath       map[*node]struct{}
	visitedTwoSmall bool
}

func (p *nodePath) Add(n *node) {
	if _, ok := p.hasOnPath[n]; ok {
		if n.IsSmall() {
			p.visitedTwoSmall = true
		}
	} else {
		p.hasOnPath[n] = struct{}{}
	}
	p.nodes = append(p.nodes, n)
}

func (p nodePath) Clone() *nodePath {
	result := &nodePath{hasOnPath: make(map[*node]struct{}), nodes: make([]*node, len(p.nodes))}
	copy(result.nodes, p.nodes)
	for k, v := range p.hasOnPath {
		result.hasOnPath[k] = v
	}
	result.visitedTwoSmall = p.visitedTwoSmall
	return result
}

func newNodePath(n *node) *nodePath {
	result := &nodePath{hasOnPath: make(map[*node]struct{})}
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
		from := p.nodes[len(p.nodes)-1]
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
