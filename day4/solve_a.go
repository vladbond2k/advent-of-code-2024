package main

var word = "XMAS"

type Direction struct {
	x, y int
}

func solveA(input []string) int {
	directionsToCheck := getDirections()
	wordCounter := 0
	for i, line := range input {
		for j := range len(line) {
		outer:
			for _, direction := range directionsToCheck {
				for charIndex := range len(word) {
					x := i + charIndex*direction.x
					y := j + charIndex*direction.y
					if x < 0 || x >= len(input) || y < 0 || y >= len(line) {
						continue outer
					}
					if input[x][y] != word[charIndex] {
						continue outer
					}
				}
				wordCounter++
			}
		}
	}
	return wordCounter
}

func getDirections() []Direction {
	directions := make([]Direction, 0)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			directions = append(directions, Direction{i, j})
		}
	}
	return directions
}
