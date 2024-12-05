package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadInput(filename string) (map[int][]int, [][]int) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := make(map[int][]int)
	updates := make([][]int, 0)

	part1 := true
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			part1 = false
			continue
		}
		if part1 {
			firstNum, secondNum := parseRule(text)
			if _, exists := rules[firstNum]; !exists {
				rules[firstNum] = []int{secondNum}
			} else {
				rules[firstNum] = append(rules[firstNum], secondNum)
			}
		} else {
			updates = append(updates, parseUpdate(text))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return rules, updates
}

func parseRule(line string) (int, int) {
	numbers := strings.Split(line, "|")
	if len(numbers) == 2 {
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		return num1, num2
	}
	return 0, 0
}

func parseUpdate(line string) []int {
	numbers := make([]int, 0)
	unparsedNumbers := strings.Split(line, ",")
	for _, num := range unparsedNumbers {
		parsedNum, _ := strconv.Atoi(num)
		numbers = append(numbers, parsedNum)
	}
	return numbers
}

func main() {
	rules, updates := loadInput("input.txt")

	fmt.Println("Result A:", solveA(rules, updates))
	fmt.Println("Result B:", solveB(rules, updates))
}
