package main

import (
	"fmt"
	"time"
)

func findNumber(numbers []int, iterations int) int {
	var numMap = map[int]int{}
	for idx, val := range numbers {
		numMap[val] = idx
	}
	delete(numMap, numbers[len(numbers)-1])
	sz := len(numbers)
	lastNum := numbers[sz-1]
	for i := sz; i < iterations; i++ {
		nextNum := 0
		if idx, seenBefore := numMap[lastNum]; seenBefore {
			nextNum = sz - 1 - idx
		} else {
			nextNum = 0
		}
		numMap[lastNum] = sz - 1
		lastNum = nextNum
		sz++
	}
	return lastNum
}

func main() {
	numbers := []int{7, 12, 1, 0, 16, 2}

	num1 := findNumber(numbers, 2020)
	fmt.Println("the 2020th number spoken:", num1)

	start := time.Now()
	num2 := findNumber(numbers, 30000000)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("the 30000000th number spoken (runtime", elapsed, "):", num2)
}
