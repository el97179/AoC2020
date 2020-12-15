package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) map[int]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bitmask := ""
	addr := 0
	value := 0
	mem := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:4] == "mask" {
			bitmask = line[7:]
		} else {
			instructions := strings.Split(line[4:], "=")
			addr, _ = strconv.Atoi(instructions[0][:len(instructions[0])-2])
			value, _ = strconv.Atoi(instructions[1][1:])

			// maskOr, _ := strconv.ParseInt(strings.Replace(bitmask, "X", "0", -1), 2, 64)
			// maskAnd, _ := strconv.ParseInt(strings.Replace(bitmask, "X", "1", -1), 2, 64)
			// value = value&int(maskAnd) + value | int(maskOr)
			// mem[addr] = value

			// value = 45         // 101101
			// bitmask = "11X0X0" // 11X0X0
			// maskOr, _ = strconv.ParseInt(strings.Replace(bitmask, "X", "0", -1), 2, 64)  // 110000  111101
			// maskAnd, _ = strconv.ParseInt(strings.Replace(bitmask, "X", "1", -1), 2, 64) // 111010  101000
			// value = (value & int(maskAnd)) | (value | int(maskOr))                       //         111000  111101

			for i, v := range bitmask {
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

func main() {

	const filename string = "input.txt"
	mem := readInput(filename)
	sum := 0
	for _, val := range mem {
		sum += val
	}
	fmt.Println("sum of all values left in memory after it completes:", sum)

}
