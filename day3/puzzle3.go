package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var treeMap []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		treeMap = append(treeMap, line)
	}

	return treeMap
}

func countTrees(down int, right int, treeMap []string) int {
	countTrees := 0
	rows := len(treeMap)
	cols := len(treeMap[0])
	c := 0
	for r := down; r < rows; r = r + down {
		c = (c + right) % cols
		if string(treeMap[r][c]) == "#" {
			countTrees++
		}
	}
	return countTrees
}

func main() {

	treeMap := readInput("input3.txt")

	numTreesPuzzle1a := countTrees(1, 3, treeMap)
	fmt.Println("Number of trees (puzzle1a):", numTreesPuzzle1a)

	numTreesPuzzle1b := 1
	slopes := []complex64{complex(1, 1), complex(1, 3), complex(1, 5), complex(1, 7), complex(2, 1)}
	for _, coord := range slopes {
		down := int(real(coord))
		right := int(imag(coord))
		numTreesPuzzle1b *= countTrees(down, right, treeMap)
	}
	fmt.Println("Number of trees (puzzle1b):", numTreesPuzzle1b)
}
