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
	lastNode        *node              // which cave are we in right now?
	hasOnPath       map[*node]struct{} // small caves we have visited so far (we do not care about large ones)
	visitedTwoSmall bool               // have we visited any small cave twice yet?
}

func (p *nodePath) Add(n *node) { // visit a cave
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
	nodes := make(map[string]*node) // cave name to cave object mapping
	addNode := func(name string) {
		_, ok := nodes[name]
		if !ok {
			nodes[name] = &node{name: name}
		}
	}
	for _, l := range lines { // load input data to our object model
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
			if !canProceed(*p, to) { // if we cannot go to the peer, skip over it
				continue
			}
			if to == nodes["end"] { // we found a path! count it
				count++
			} else {
				newPath := p.Clone() // we have to clone a current path, because if we do not, on a fork both continuations will be added to the same path, which does not make sense
				newPath.Add(to)
				queue = append(queue, newPath)
			}
		}
	}
	return fmt.Sprint(count)
}
