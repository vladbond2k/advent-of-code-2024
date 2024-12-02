package main

func solve_a(reports [][]int) int {
	safeReportsCount := 0

	for _, report := range reports {
		if verifyLevelSafety(report) {
			safeReportsCount++
		}
	}

	return safeReportsCount
}
