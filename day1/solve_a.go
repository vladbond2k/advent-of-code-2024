package main

import (
	"math"
	"sort"
)

func solve_a(leftList, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)
	distance := 0
	for i := range len(leftList) {
		distance += int(math.Abs(float64(leftList[i] - rightList[i])))
	}
	return distance
}
