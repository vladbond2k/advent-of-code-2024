package main

func solveB(input [][]rune) int {
	labMap, guard := parseBoard(input)
	guard.GoOutOfTheLab(labMap)

	numOfLoops := 0
	for _, visitedPos := range labMap.getVisitedPositions() {
		if visitedPos.x == guard.startPosition.x && visitedPos.y == guard.startPosition.y {
			continue
		}
		labMap.Reset()
		guard.position = guard.startPosition
		guard.direction = guard.startDirection

		labMap.updateCell(visitedPos, cellNewObstacle)

		if isLooped := guard.GoOutOfTheLab(labMap); isLooped {
			numOfLoops++
		}
	}
	return numOfLoops
}
