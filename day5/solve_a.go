package main

func solveA(rules map[int][]int, updates [][]int) int {
	result := 0
	for _, update := range updates {
		wrongNumbers, _ := breakUpdate(update, rules)
		if len(wrongNumbers) == 0 {
			result += update[len(update)/2]
		}
	}
	return result
}

func fixUpdate(wrongNumbers, correctNumbers []int, rules map[int][]int) []int {
	for _, num := range wrongNumbers {
		earliestIndex := len(wrongNumbers) + len(correctNumbers) - 1
		for _, rule := range rules[num] {
			for i, v := range correctNumbers {
				if v == rule && i < earliestIndex {
					earliestIndex = i
				}
			}
		}
		correctNumbers = append(correctNumbers[:earliestIndex], append([]int{num}, correctNumbers[earliestIndex:]...)...)
	}
	return correctNumbers
}

func breakUpdate(update []int, rules map[int][]int) ([]int, []int) {
	wrongNumbers, newUpdate := make([]int, 0), make([]int, 0)
	for i, num := range update {
		numRules, isRuled := rules[num]
		if isRuled && doesBrakeRules(update[:i], numRules) {
			wrongNumbers = append(wrongNumbers, num)
		} else {
			newUpdate = append(newUpdate, num)
		}
	}
	return wrongNumbers, newUpdate
}

func doesBrakeRules(listBefore []int, rules []int) bool {
	for _, v := range listBefore {
		for _, rule := range rules {
			if v == rule {
				return true
			}
		}
	}
	return false
}
