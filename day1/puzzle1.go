package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readInput(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var arr []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, num)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return arr
}

func puzzle1a(arr []int, targetSum int) (int, int, int, error) {
	sort.Ints(arr)
	minIdx := 0
	maxIdx := len(arr) - 1
	solutionFound := true
	for {
		var sum int = arr[minIdx] + arr[maxIdx]
		if minIdx >= maxIdx {
			solutionFound = false
			break
		}
		if sum == targetSum {
			break
		} else if sum > targetSum {
			maxIdx--
		} else {
			minIdx++
		}
	}
	if !solutionFound {
		return 0, 0, 0, fmt.Errorf("No two values found that sum to %d", targetSum)
	}
	prod := arr[minIdx] * arr[maxIdx]
	return arr[minIdx], arr[maxIdx], prod, nil

}

func puzzle1b(arr []int, targetSum int) (int, int, int, int, error) {
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		var arrMinusOne = make([]int, len(arr))
		copy(arrMinusOne, arr)
		arrLow := arrMinusOne[:i]
		arrBig := arrMinusOne[i+1:]
		arrMinusOne = append(arrLow, arrBig...)
		val1, val2, prod2, err := puzzle1a(arrMinusOne, targetSum-arr[i])
		if err != nil {
			log.Println(err)
		} else {
			prod := prod2 * arr[i]
			return val1, val2, arr[i], prod, nil
		}

	}
	return 0, 0, 0, 0, fmt.Errorf("No three values found that sum to %d", targetSum)
}

func main() {

	var arr []int = readInput("input1.txt")
	targetSum := 2020
	val1, val2, prod, err := puzzle1a(arr, targetSum)
	if err == nil {
		fmt.Println("Product of the two values", val1, "and", val2, "is", prod)
	} else {
		log.Println(err)
	}

	val1, val2, val3, prod, err := puzzle1b(arr, targetSum)
	if err == nil {
		fmt.Println("Product of the three values", val1, ",", val2, "and", val3, "is", prod)
	} else {
		log.Println(err)
	}
}
