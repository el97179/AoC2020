package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type instructionBlock struct {
	bitmask   string
	instrList []map[int]int
}

func readInput(filename string) []instructionBlock {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var instructionList []instructionBlock
	var bitmapBlock instructionBlock
	var instructionsBlock []map[int]int
	bitmask := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:4] == "mask" {
			instructionList = append(instructionList, bitmapBlock) // save data of the previous block
			bitmask = line[7:]
			bitmapBlock.bitmask = bitmask
			instructionsBlock = make([]map[int]int, 0)
		} else {
			instructions := strings.Split(line[4:], "=")
			addr, _ := strconv.Atoi(instructions[0][:len(instructions[0])-2])
			value, _ := strconv.Atoi(instructions[1][1:])
			instructionsBlock = append(instructionsBlock, map[int]int{addr: value})
			bitmapBlock.instrList = instructionsBlock
		}
	}
	instructionList = append(instructionList, bitmapBlock) // save data of the last block
	instructionList = instructionList[1:]                  // remove empty first block

	return instructionList
}

func processInstructions(instructionList []instructionBlock) map[int]int {
	mem := make(map[int]int)
	for _, block := range instructionList {
		bitmask := block.bitmask
		for _, instr := range block.instrList {
			addr := 0
			value := 0
			for key := range instr { // we always have one key in the map
				addr = key
			}
			value = instr[addr]
			for i, v := range block.bitmask {
				if v == '0' {
					value = value & int(math.Pow(2, float64(len(bitmask)))-1-math.Pow(2, float64(len(bitmask)-i-1)))
				} else if v == '1' {
					value = value | int(math.Pow(2, float64(len(bitmask)-i-1)))
				}
			}
			mem[addr] = value
		}
	}
	return mem
}

// 101010 (42) + X0000X ==> 001010, 001011, 101010, 101011 ==> 10, 11, 42, 43
func createAddresses(address int, floatingBits []int) map[int]bool {
	addressSize := 36
	addresses := make(map[int]bool)
	if len(floatingBits) == 0 {
		addresses[address] = true
	} else {
		for _, bitPos := range floatingBits {
			reducedFloatingBits := floatingBits[1:]

			addr0 := address & int(math.Pow(2, float64(addressSize))-1-math.Pow(2, float64(addressSize-bitPos-1)))
			extendedAddresses0 := createAddresses(addr0, reducedFloatingBits)
			for addr := range extendedAddresses0 {
				// fmt.Println("Appending0:", addr, "into", addresses)
				addresses[addr] = true
			}

			addr1 := address | int(math.Pow(2, float64(addressSize-bitPos-1)))
			extendedAddresses1 := createAddresses(addr1, reducedFloatingBits)
			for addr := range extendedAddresses1 {
				// fmt.Println("Appending1:", addr, "into", addresses)
				addresses[addr] = true
			}
		}
	}
	return addresses
}

func processInstructionsAddress(instructionList []instructionBlock) map[int]int {
	mem := make(map[int]int)
	for idx, block := range instructionList {
		fmt.Println("processing block:", idx, "/", len(instructionList))
		bitmask := block.bitmask
		for _, instr := range block.instrList {
			addr := 0
			value := 0
			floatingBits := []int{}
			addrList := []int{}
			for key := range instr { // we always have one key in the map
				addr = key
			}
			value = instr[addr]
			for i, v := range block.bitmask {
				if v == '1' {
					addr = addr | int(math.Pow(2, float64(len(bitmask)-i-1)))
				} else if v == 'X' {
					floatingBits = append(floatingBits, i)
				}
			}
			addressRange := createAddresses(addr, floatingBits)
			for addr := range addressRange {
				addrList = append(addrList, addr)
			}
			// fmt.Println(addrList)
			for _, addr := range addrList {
				mem[addr] = value
			}
		}
	}
	return mem
}

func getFloatMasks(s string) []int {
	a := []int{}
	for idx, char := range s {
		if char == 'X' {
			a = append(a, 1<<(len(s)-idx-1))
		}
	}

	b := []int{0}
	for _, i := range a {
		for _, j := range b {
			b = append(b, i|j)
		}
	}

	sort.IntSlice(b).Sort()
	return b
}

func main() {
	start := time.Now()

	const filename string = "input.txt"
	instructionList := readInput(filename)

	mem := processInstructions(instructionList)
	sum := 0
	for _, val := range mem {
		sum += val
	}
	fmt.Println("sum of all values left in memory after it completes (part 1):", sum)

	mem = processInstructionsAddress(instructionList) // runtime: 4h8'24"...
	sum = 0
	for _, val := range mem {
		sum += val
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("sum of all values left in memory after it completes (part 2):", sum)
	fmt.Println("runtime:", elapsed)
}
