package day12

func (p nodePath) CanProceed2(n *node) bool {
	i, alreadyVisited := p.lookup[n]
	return !alreadyVisited || !n.IsSmall() || n.name != "start" && i == 1 && !p.visitedTwoSmall
}

func Part2(inputFile string) string {
	return solve(inputFile, nodePath.CanProceed2)
}
