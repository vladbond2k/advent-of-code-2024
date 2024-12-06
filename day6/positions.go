package main

type Position struct {
	x, y int
}

func (pos Position) Move(direction DirectionVector) Position {
	return Position{pos.x + direction.x, pos.y + direction.y}
}

type DirectionVector struct {
	x, y int
}

func (dir *DirectionVector) TurnRight() {
	if dir.x == -1 {
		dir.x = 0
		dir.y = 1
	} else if dir.y == 1 {
		dir.x = 1
		dir.y = 0
	} else if dir.x == 1 {
		dir.x = 0
		dir.y = -1
	} else if dir.y == -1 {
		dir.x = -1
		dir.y = 0
	}
}

func (dir *DirectionVector) TurnLeft() {
	for i := 0; i < 3; i++ {
		dir.TurnRight()
	}
}
