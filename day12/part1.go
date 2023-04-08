package day12

func (p nodePath) CanProceed(n *node) bool {
	// we can visit a neighbour if it is large, or if it has not been already visited
	_, alreadyVisited := p.hasOnPath[n]
	return !n.IsSmall() || !alreadyVisited
}

func Part1(inputFile string) string {
	return solve(inputFile, nodePath.CanProceed)
}
