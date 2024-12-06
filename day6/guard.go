package main

type Guard struct {
	startPosition  Position
	position       Position
	startDirection DirectionVector
	direction      DirectionVector
}

func (g *Guard) MakeATurn() {
	g.direction.TurnRight()
}

func (g *Guard) GoOutOfTheLab(lab LabMap) bool {
	visitedInDirection := make(map[Position]DirectionVector)
	for {
		inDirection, wasVisited := visitedInDirection[g.position]
		if wasVisited && inDirection == g.direction {
			return true
		}
		nextPos := g.position.Move(g.direction)

		if !lab.isInbounds(nextPos) {
			lab.updateCell(g.position, cellVisited)
			break
		}
		if lab.IsObstacle(nextPos) {
			g.MakeATurn()
		} else {
			lab.updateCell(g.position, cellVisited)
			visitedInDirection[g.position] = g.direction
			g.position = nextPos
		}
	}
	return false
}

var guardBeaconToDirection = map[rune]DirectionVector{
	'^': {x: -1, y: 0},
	'>': {x: 0, y: 1},
	'v': {x: 1, y: 0},
	'<': {x: 0, y: -1},
}
