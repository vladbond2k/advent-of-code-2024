package main

func solveB(rules map[int][]int, updates [][]int) int {
	result := 0
	for _, update := range updates {
		wrongNumbers, correctUpdate := breakUpdate(update, rules)
		if len(wrongNumbers) == 0 {
			continue
		}
		correctUpdate = fixUpdate(wrongNumbers, correctUpdate, rules)
		result += correctUpdate[len(correctUpdate)/2]
	}
	return result
}
