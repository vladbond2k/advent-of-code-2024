package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func parseReports(input []string) [][]int {
	parsedLevels := make([][]int, len(input))
	for i, report := range input {
		levels := strings.Fields(report)
		for _, level := range levels {
			levelInt, _ := strconv.Atoi(level)
			parsedLevels[i] = append(parsedLevels[i], levelInt)
		}
	}
	return parsedLevels
}

func main() {
	input := loadInput("input.txt")
	reports := parseReports(input)

	fmt.Println("Result A:", solve_a(reports))
	fmt.Println("Result B:", solve_b(reports))
}
