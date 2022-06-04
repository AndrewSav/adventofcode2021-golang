package day12

import (
	"aoc2021/util"
	"fmt"
	"strings"
)

type node struct {
	name  string
	peers []*node
}

func (n node) IsSmall() bool {
	return n.name[0] > 'Z'
}

type nodePath struct {
	lastNode        *node
	hasOnPath       map[*node]struct{}
	visitedTwoSmall bool
}

func (p *nodePath) Add(n *node) {
	p.lastNode = n
	if !n.IsSmall() {
		return
	}
	if _, ok := p.hasOnPath[n]; ok {
		p.visitedTwoSmall = true
	} else {
		p.hasOnPath[n] = struct{}{}
	}
}

func (p nodePath) Clone() *nodePath {
	result := p
	result.hasOnPath = make(map[*node]struct{})
	for k := range p.hasOnPath {
		result.hasOnPath[k] = struct{}{}
	}
	return &result
}

func newNodePath(n *node) *nodePath {
	result := &nodePath{hasOnPath: make(map[*node]struct{})}
	result.Add(n)
	return result
}

func solve(inputFile string, canProceed func(p nodePath, n *node) bool) string {
	lines := util.ReadInput(inputFile)
	nodes := make(map[string]*node)
	addNode := func(name string) {
		_, ok := nodes[name]
		if !ok {
			nodes[name] = &node{name: name}
		}
	}
	for _, l := range lines {
		parts := strings.Split(l, "-")
		addNode(parts[0])
		addNode(parts[1])
		nodes[parts[0]].peers = append(nodes[parts[0]].peers, nodes[parts[1]])
		nodes[parts[1]].peers = append(nodes[parts[1]].peers, nodes[parts[0]])
	}
	count := 0
	for queue := []*nodePath{newNodePath(nodes["start"])}; len(queue) > 0; {
		p := queue[0]
		queue = queue[1:]
		from := p.lastNode
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
