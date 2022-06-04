package day12

func (p nodePath) CanProceed(n *node) bool {
	_, alreadyVisited := p.hasOnPath[n]
	return !n.IsSmall() || !alreadyVisited
}

func Part1(inputFile string) string {
	return solve(inputFile, nodePath.CanProceed)
}
