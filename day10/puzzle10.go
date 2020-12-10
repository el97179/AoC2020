package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func gradient(voltages []int) []int {
	gradients := []int{}
	grad := 0
	sort.Ints(voltages)
	for idx, val := range voltages {
		if idx == 0 { // we start from 0 volts
			grad = val
		} else {
			grad = val - voltages[idx-1]
		}
		gradients = append(gradients, grad)
	}
	return gradients
}

func distribution(arr []int) map[int]int {
	histo := make(map[int]int)
	for _, val := range arr {
		histo[val]++
	}
	return histo
}

func countConsecutiveOnes(arr []int) []int {
	onesArr := []int{}
	countOnes := 0
	for idx, val := range arr {
		if idx == 0 {
			continue
		}
		if arr[idx-1] == 1 {
			countOnes++
		}
		if val != arr[idx-1] {
			if countOnes > 1 {
				onesArr = append(onesArr, countOnes)
			}
			countOnes = 0
			continue
		}
	}
	return onesArr
}

func getBaseOfExp(aMap map[int]int, sz int) int {
	if sz < 1 {
		fmt.Println("Cannot process zero or negative values")
		return -1
	}
	if sz < 4 { // easy solution; ommitting (or not) one instance of the excessive ones: 2^n, where n=#consecutive_ones-1
		return int(math.Pow(2, float64(sz-1)))
	}
	// when more than 4 consecutive ones, we cant ommit more than 3 consecutive, so we count the number of triplets (111) by shifting them:
	// 2*n-1, where n=#permutations for sequence with length key-1 (reursively computed)
	return 2*getBaseOfExp(aMap, sz-1) - 1
}

func countPermutations(onesSeriesDistr map[int]int) int {
	count := 1.0
	for k, val := range onesSeriesDistr {
		base := getBaseOfExp(onesSeriesDistr, k)
		count *= math.Pow(float64(base), float64(val))
	}
	return int(count)
}

func main() {

	const filename string = "input.txt"
	voltages := readInput(filename)
	gradients := gradient(voltages)
	gradients = append(gradients, 3) // add the diff from the max to my device's input
	hist := distribution(gradients)
	fmt.Println(">> Product of 1s and 3s (part1) :", hist[1]*hist[3])

	consOnes := countConsecutiveOnes(gradients)
	hist = distribution(consOnes)
	fmt.Println("Number of consecutive ones :", consOnes)
	fmt.Println("Distribution of consecutive ones :", hist)
	fmt.Println(">> Number of distinct arrangements (part2) :", countPermutations(hist))
}
