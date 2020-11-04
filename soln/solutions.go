package soln

func collectSolutions() map[int]func() {
	solutions := make(map[int]func())
	solutions[1] = solve1
	solutions[2] = solve2
	solutions[3] = solve3

	return solutions
}

// GetSolution : returns the solution function for a day given its day number n
func GetSolution(n int) func() {
	solutions := collectSolutions()
	return solutions[n]
}