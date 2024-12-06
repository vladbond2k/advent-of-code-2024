package main

func solveA(input [][]rune) int {
	labMap, guard := parseBoard(input)
	guard.GoOutOfTheLab(labMap)
	return labMap.countVisits()
}
