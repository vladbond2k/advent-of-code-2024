package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadInput(filename string) []string {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return lines
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return lines
}

func main() {
	input := loadInput("input.txt")

	fmt.Println("Result A:", solveA(input))
	fmt.Println("Result B:", solveB(input))
}
