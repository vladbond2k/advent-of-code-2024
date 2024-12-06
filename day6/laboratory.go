package main

type Cell int

const (
	cellObstacle Cell = iota
	cellNewObstacle
	cellFree
	cellVisited
)

var cellMapping = map[rune]Cell{
	'#': cellObstacle,
	'O': cellNewObstacle,
	'X': cellVisited,
	'.': cellFree,
	'^': cellFree, 'v': cellFree, '<': cellFree, '>': cellFree,
}

type LabMap struct {
	cells [][]Cell
}

func (m *LabMap) updateCell(pos Position, newCell Cell) {
	m.cells[pos.x][pos.y] = newCell
}

func (m *LabMap) countVisits() int {
	visits := 0
	for _, row := range m.cells {
		for _, cell := range row {
			if cell == cellVisited {
				visits++
			}
		}
	}
	return visits
}

func (m *LabMap) Reset() {
	for x, row := range m.cells {
		for y, cell := range row {
			if cell != cellObstacle && cell != cellFree {
				m.cells[x][y] = cellFree
			}
		}
	}
}

func (m *LabMap) getVisitedPositions() []Position {
	visited := make([]Position, 0)
	for x, row := range m.cells {
		for y, cell := range row {
			if cell == cellVisited {
				visited = append(visited, Position{x, y})
			}
		}
	}
	return visited
}

func (m *LabMap) isInbounds(pos Position) bool {
	return pos.x >= 0 && pos.x < len(m.cells) && pos.y >= 0 && pos.y < len(m.cells[0])
}

func (m *LabMap) IsObstacle(pos Position) bool {
	return m.cells[pos.x][pos.y] == cellObstacle || m.cells[pos.x][pos.y] == cellNewObstacle
}

func parseBoard(labMap [][]rune) (LabMap, Guard) {
	var guard Guard
	lab := LabMap{}
	lab.cells = make([][]Cell, len(labMap))

	for i, row := range labMap {
		lab.cells[i] = make([]Cell, len(row))
		for j, char := range row {
			if direction, isGuard := guardBeaconToDirection[char]; isGuard {
				guard = Guard{
					startDirection: direction, direction: direction,
					startPosition: Position{i, j}, position: Position{i, j},
				}
			}
			lab.cells[i][j] = cellMapping[char]
		}
	}
	return lab, guard
}
