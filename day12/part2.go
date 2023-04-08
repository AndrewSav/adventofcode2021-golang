package day12

func (p nodePath) CanProceed2(n *node) bool {
	// we can visit a neighbour if it is large, or if it has not been already visited
	// or if we have not visited any small cave twice yet (we cannot return to start_
	return p.CanProceed(n) || n.name != "start" && !p.visitedTwoSmall
}

func Part2(inputFile string) string {
	return solve(inputFile, nodePath.CanProceed2)
}
