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

	numberList := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.Atoi(line)
		numberList = append(numberList, val)
	}
	return numberList
}

func checkNextNumber(arr []int, nextNum int) bool {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)
	minIdx := 0
	maxIdx := len(sortedArr) - 1
	solutionFound := true
	for {
		sum := sortedArr[minIdx] + sortedArr[maxIdx]
		if minIdx >= maxIdx {
			solutionFound = false
			break
		}
		if sum == nextNum {
			break
		} else if sum > nextNum {
			maxIdx--
		} else {
			minIdx++
		}
	}
	return solutionFound

}

func findContigArr(arr []int, targetNum int) []int {
	contigArr := []int{}
	for sz := 2; sz < len(arr); sz++ { // try different sizes of the contiguous arrays
		for idx := 0; idx <= len(arr)-sz; idx++ { // try different starting points
			sum := 0
			for i := idx; i < idx+sz; i++ { // sum the numbers in the current contigous array
				sum += arr[i]
			}
			if sum == targetNum {
				contigArr = arr[idx : idx+sz]
				break

			}
		}
	}

	return contigArr
}

func main() {

	const filename string = "input.txt"
	const buffer int = 25
	numberList := readInput(filename)
	nonValidNum := -1
	for i := buffer; i < len(numberList); i++ {
		arr := numberList[i-buffer : i]
		nextNum := numberList[i]
		if !checkNextNumber(arr, nextNum) {
			nonValidNum = nextNum
			break
		}
	}
	fmt.Println(">> First non valid number (part1) :", nonValidNum)

	contigArr := findContigArr(numberList, nonValidNum)
	sort.Ints(contigArr)
	encrWeak := contigArr[0] + contigArr[len(contigArr)-1]
	fmt.Println(">> Encryption weakness (part2) :", encrWeak)
}
