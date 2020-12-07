package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Bag is a structure to hold the name and count of the bags
type Bag struct {
	color string
	count int
}

func readInput(filename string) (map[string][]Bag, map[string][]string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	containsBagsMap := make(map[string][]Bag)
	containedInBagsMap := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		key := lineSplit[0] + " " + lineSplit[1]
		bagsInside := []Bag{}
		for idx := 4; idx < len(lineSplit); idx += 4 {
			count, _ := strconv.Atoi(lineSplit[idx])
			color := lineSplit[idx+1] + " " + lineSplit[idx+2]
			currentBag := Bag{color, count}

			bagsInside = append(bagsInside, currentBag)
			bagOutside := containedInBagsMap[color]
			bagOutside = append(bagOutside, key)
			containedInBagsMap[color] = bagOutside
		}
		containsBagsMap[key] = bagsInside
	}
	return containsBagsMap, containedInBagsMap
}

func getOuterBags(containedInBagsMap map[string][]string, targetBag string, validBags map[string]bool) map[string]bool {
	for _, bag := range containedInBagsMap[targetBag] {
		validBags[bag] = true
		validBags = getOuterBags(containedInBagsMap, bag, validBags)
	}
	return validBags
}

func getInnerBags(containsBagsMap map[string][]Bag, targetBag string, validBags map[string]int, numBags int) map[string]int {
	for _, bag := range containsBagsMap[targetBag] {
		validBags[bag.color] += numBags * bag.count
		validBags = getInnerBags(containsBagsMap, bag.color, validBags, numBags*bag.count)
	}
	return validBags
}

func main() {

	containsBagsMap, containedInBagsMap := readInput("input7.txt")
	targetBag := "shiny gold"

	outerBags := make(map[string]bool)
	outerBags = getOuterBags(containedInBagsMap, targetBag, outerBags)
	fmt.Println("Number of outermost colored bags:", len(outerBags))

	innerBags := make(map[string]int)
	countInnerBags := 0
	innerBags = getInnerBags(containsBagsMap, targetBag, innerBags, 1)
	for _, bagCount := range innerBags {
		countInnerBags += bagCount
	}
	fmt.Println("Number inner contained bags:", countInnerBags)
}
