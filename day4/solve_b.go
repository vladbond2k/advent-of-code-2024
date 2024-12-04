package main

func solveB(input []string) int {
	xmasCounter := 0
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input)-1; j++ {
			if input[i][j] == 'A' {
				topLeft, topRight := input[i-1][j-1], input[i-1][j+1]
				bottomLeft, bottomRight := input[i+1][j-1], input[i+1][j+1]
				if (topLeft == 'M' && bottomRight == 'S' || topLeft == 'S' && bottomRight == 'M') &&
					(topRight == 'M' && bottomLeft == 'S' || topRight == 'S' && bottomLeft == 'M') {
					xmasCounter++
				}
			}
		}
	}
	return xmasCounter
}
