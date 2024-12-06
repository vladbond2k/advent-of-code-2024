package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadInput(filename string) [][]rune {
	var mapMatrix [][]rune

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return mapMatrix
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		mapMatrix = append(mapMatrix, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return mapMatrix
}

func main() {
	labMap := loadInput("input.txt")
	fmt.Println("Result A:", solveA(labMap))

	labMap = loadInput("input.txt")
	fmt.Println("Result B:", solveB(labMap))
}
