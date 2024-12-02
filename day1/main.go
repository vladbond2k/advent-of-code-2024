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

func loadLists(input []string) ([]int, []int) {
	left, right := make([]int, len(input)), make([]int, len(input))
	for i, line := range input {
		nums := strings.Fields(line)
		left[i], _ = strconv.Atoi(nums[0])
		right[i], _ = strconv.Atoi(nums[1])
	}
	return left, right
}

func main() {
	leftList, rightList := loadLists(loadInput("input.txt"))

	fmt.Println("Result A:", solve_a(leftList, rightList))
	fmt.Println("Result B:", solve_b(leftList, rightList))
}
