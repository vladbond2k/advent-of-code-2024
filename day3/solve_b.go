package main

import (
	"regexp"
	"strconv"
)

var reB = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

func solveB(input []string) int {
	enabled := true
	sum := 0

	for _, line := range input {
		matches := reB.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[1] != "" && match[2] != "" && enabled {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				sum += num1 * num2
			} else {
				enabled = match[0] == "do()"
			}
		}
	}
	return sum
}
