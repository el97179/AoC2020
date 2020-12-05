package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var seats []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seat := scanner.Text()
		seats = append(seats, seat)
	}

	return seats
}

func getSeatID(seat string) int {
	rowCode := seat[:7]
	row := convertToBinary(rowCode)
	colCode := seat[7:]
	col := convertToBinary(colCode)
	return row*8 + col
}

func convertToBinary(strCode string) int {
	strCode = strings.ReplaceAll(strCode, "F", "0")
	strCode = strings.ReplaceAll(strCode, "L", "0")
	strCode = strings.ReplaceAll(strCode, "B", "1")
	strCode = strings.ReplaceAll(strCode, "R", "1")
	intCode, err := strconv.ParseInt(strCode, 2, 64)
	if err != nil {
		fmt.Println("Error in converting the string code to int:", err)
		return -1
	}
	return int(intCode)
}

func findMySeat(seatIds []int) int {
	sort.Ints(seatIds)
	for idx, id := range seatIds {
		if idx == 0 {
			continue
		}
		if id-seatIds[idx-1] > 1 {
			return id - 1
		}
	}
	return -1
}

func main() {

	seats := readInput("input5.txt")
	seatIDs := []int{}
	maxSeatID := 0
	for _, seat := range seats {
		seatID := getSeatID(seat)
		seatIDs = append(seatIDs, seatID)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	mySeatID := findMySeat(seatIDs)

	fmt.Println("Maximum seatID:", maxSeatID)
	fmt.Println("My seatID:", mySeatID)
}
