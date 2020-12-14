package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInput(filename string) (int, []int) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	mytime, _ := strconv.Atoi(lines[0])
	busesStr := strings.Split(lines[1], ",")
	buses := []int{}
	for _, d := range busesStr {
		val, err := strconv.Atoi(d)
		if err != nil {
			buses = append(buses, -1)
		} else {
			buses = append(buses, val)
		}
	}
	return mytime, buses
}

func findFirstBus(mytime int, buses []int) (int, int) {
	firstBus := 0
	waitTime := 0
	for _, busID := range buses {
		if busID == -1 {
			continue
		}
		if mytime%busID == 0 {
			return busID, 0
		}
		if mytime%busID > waitTime {
			waitTime = mytime % busID
			firstBus = busID
		}
	}
	return firstBus, firstBus - waitTime
}

func findStartTime(buses []int) int {
	startTime := 0
	breakFlag := true
	for true {
		increament := 1
		for idx, val := range buses {
			breakFlag = true
			if val == -1 {
				continue
			}
			if (startTime+idx)%val != 0 {
				for i := 0; i < idx; i++ {
					if buses[i] == -1 {
						continue
					}
					increament *= buses[i]
				}
				startTime += increament
				breakFlag = false
				break
			}
		}
		if breakFlag {
			break
		}
	}
	return startTime
}

func main() {

	const filename string = "input.txt"
	time, buses := readInput(filename)
	firstBus, waitTime := findFirstBus(time, buses)
	fmt.Println("(Part1) First bus ID:", firstBus, "wait time:", waitTime, "answer:", firstBus*waitTime)

	startTime := findStartTime(buses)
	fmt.Println("(Part2) Start time:", startTime)

}
