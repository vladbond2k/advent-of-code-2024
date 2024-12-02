package main

func solve_b(leftList, rightList []int) int {
	occurrences := make(map[int]int)
	for _, rightItem := range rightList {
		value, exists := occurrences[rightItem]
		if exists {
			occurrences[rightItem] = value + 1
		} else {
			occurrences[rightItem] = 1
		}
	}
	result := 0
	for _, leftItem := range leftList {
		result += leftItem * occurrences[leftItem]
	}
	return result
}
