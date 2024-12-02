package main

func solve_b(reports [][]int) int {
	safeReportsCount := 0

	for _, report := range reports {
		isSafe := verifyLevelSafety(report)
		if !isSafe {
			alteredReports := genAlteredReports(report)
			for _, altReport := range alteredReports {
				if verifyLevelSafety(altReport) {
					isSafe = true
					break
				}
			}
		}
		if isSafe {
			safeReportsCount++
		}
	}

	return safeReportsCount
}

func genAlteredReports(report []int) [][]int {
	var result [][]int

	for i := range report {
		variant := append([]int{}, report[:i]...)
		variant = append(variant, report[i+1:]...)
		result = append(result, variant)
	}
	return result
}

func verifyLevelSafety(report []int) bool {
	isIncreasing := report[0] < report[1]
	previous := report[0]
	for _, current := range report[1:] {
		if diff := current - previous; isIncreasing && !(diff >= 1 && diff <= 3) || !isIncreasing && !(diff >= -3 && diff <= -1) {
			return false
		}
		previous = current
	}
	return true
}
