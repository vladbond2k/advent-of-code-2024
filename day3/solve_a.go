package main

import (
	"regexp"
	"strconv"
	"sync"
)

var reA = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func solveA(input []string) int {
	results := make(chan int, len(input))
	var wg sync.WaitGroup

	for _, line := range input {
		wg.Add(1)
		go processLineA(line, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for res := range results {
		sum += res
	}
	return sum
}

func processLineA(line string, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	lineSum := 0
	matches := reA.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		lineSum += num1 * num2
	}
	results <- lineSum
}
