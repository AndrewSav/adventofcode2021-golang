package day12

func (p nodePath) CanProceed2(n *node) bool {
	return p.CanProceed(n) || n.name != "start" && !p.visitedTwoSmall
}

func Part2(inputFile string) string {
	return solve(inputFile, nodePath.CanProceed2)
}
