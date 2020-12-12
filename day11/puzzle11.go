package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	array2d := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		intLine := []int{}
		for _, val := range line {
			if val == '.' { // floor
				intLine = append(intLine, 0)
			} else if val == 'L' { // empty
				intLine = append(intLine, 1)
			} else if val == '#' { // occupied
				intLine = append(intLine, -1)
			}
		}
		array2d = append(array2d, intLine)
	}
	return array2d
}

type coordinate struct {
	row int
	col int
}

func countOccupiedNeighbors(seatArray2d [][]int, coord coordinate) int {
	if len(seatArray2d) == 0 {
		fmt.Println("Empty array of seat arrangement")
		return -1
	}
	countOccupied := 0
	rows := len(seatArray2d)
	cols := len(seatArray2d[0])
	i := coord.row
	j := coord.col
	for r := -1; r < 2; r++ {
		if i+r < 0 || i+r >= rows {
			continue
		}
		for c := -1; c < 2; c++ {
			if j+c < 0 || j+c >= cols || (r == 0 && c == 0) {
				continue
			}
			if seatArray2d[i+r][j+c] == -1 { // occupied
				countOccupied++
			}
		}
	}
	return countOccupied
}

func countOccupiedLines(seatArray2d [][]int, coord coordinate) int {
	if len(seatArray2d) == 0 {
		fmt.Println("Empty array of seat arrangement")
		return -1
	}
	countOccupied := 0
	countOccupied += countN(seatArray2d, coord)
	countOccupied += countNE(seatArray2d, coord)
	countOccupied += countE(seatArray2d, coord)
	countOccupied += countSE(seatArray2d, coord)
	countOccupied += countS(seatArray2d, coord)
	countOccupied += countSW(seatArray2d, coord)
	countOccupied += countW(seatArray2d, coord)
	countOccupied += countNW(seatArray2d, coord)
	return countOccupied
}

func countN(seatArray2d [][]int, coord coordinate) int {
	if coord.row == 0 {
		return 0
	}
	for r := coord.row - 1; r >= 0; r-- {
		if seatArray2d[r][coord.col] == -1 {
			return 1
		} else if seatArray2d[r][coord.col] == 1 {
			return 0
		}
	}
	return 0
}

func countNE(seatArray2d [][]int, coord coordinate) int {
	if coord.row == 0 || coord.col >= len(seatArray2d[0])-1 {
		return 0
	}
	for r, c := coord.row-1, coord.col+1; r >= 0 && c < len(seatArray2d[0]); r, c = r-1, c+1 {
		if seatArray2d[r][c] == -1 {
			return 1
		} else if seatArray2d[r][c] == 1 {
			return 0
		}
	}
	return 0
}

func countE(seatArray2d [][]int, coord coordinate) int {
	if coord.col >= len(seatArray2d[0])-1 {
		return 0
	}
	for c := coord.col + 1; c < len(seatArray2d[0]); c++ {
		if seatArray2d[coord.row][c] == -1 {
			return 1
		} else if seatArray2d[coord.row][c] == 1 {
			return 0
		}
	}
	return 0
}

func countSE(seatArray2d [][]int, coord coordinate) int {
	if coord.row >= len(seatArray2d)-1 || coord.col >= len(seatArray2d[0])-1 {
		return 0
	}
	for r, c := coord.row+1, coord.col+1; r < len(seatArray2d) && c < len(seatArray2d[0]); r, c = r+1, c+1 {
		if seatArray2d[r][c] == -1 {
			return 1
		} else if seatArray2d[r][c] == 1 {
			return 0
		}
	}
	return 0
}

func countS(seatArray2d [][]int, coord coordinate) int {
	if coord.row >= len(seatArray2d)-1 {
		return 0
	}
	for r := coord.row + 1; r < len(seatArray2d); r++ {
		if seatArray2d[r][coord.col] == -1 {
			return 1
		} else if seatArray2d[r][coord.col] == 1 {
			return 0
		}
	}
	return 0
}

func countSW(seatArray2d [][]int, coord coordinate) int {
	if coord.row >= len(seatArray2d)-1 || coord.col == 0 {
		return 0
	}
	for r, c := coord.row+1, coord.col-1; r < len(seatArray2d) && c >= 0; r, c = r+1, c-1 {
		if seatArray2d[r][c] == -1 {
			return 1
		} else if seatArray2d[r][c] == 1 {
			return 0
		}
	}
	return 0
}

func countW(seatArray2d [][]int, coord coordinate) int {
	if coord.col == 0 {
		return 0
	}
	for c := coord.col - 1; c >= 0; c-- {
		if seatArray2d[coord.row][c] == -1 {
			return 1
		} else if seatArray2d[coord.row][c] == 1 {
			return 0
		}
	}
	return 0
}

func countNW(seatArray2d [][]int, coord coordinate) int {
	if coord.row == 0 || coord.col == 0 {
		return 0
	}
	for r, c := coord.row-1, coord.col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if seatArray2d[r][c] == -1 {
			return 1
		} else if seatArray2d[r][c] == 1 {
			return 0
		}
	}
	return 0
}

func update(seatArray2d [][]int, method string, minOccupiedSeats int) ([][]int, bool) {
	updateFlag := false
	updatedSeatArray2d := make([][]int, len(seatArray2d))
	for i := range seatArray2d {
		updatedSeatArray2d[i] = make([]int, len(seatArray2d[i]))
		copy(updatedSeatArray2d[i], seatArray2d[i])
	}

	for r := range seatArray2d {
		for c := range seatArray2d[r] {
			if seatArray2d[r][c] == 0 { // floor
				continue
			}
			var counting func([][]int, coordinate) int
			if method == "neighbors" {
				counting = countOccupiedNeighbors
			} else if method == "lines" {
				counting = countOccupiedLines
			} else {
				fmt.Println("Uknown update policy", method)
				return nil, false
			}
			countOccupied := counting(seatArray2d, coordinate{r, c})
			if (seatArray2d[r][c] == 1 && countOccupied == 0) || (seatArray2d[r][c] == -1 && countOccupied >= minOccupiedSeats) {
				updatedSeatArray2d[r][c] *= -1
				updateFlag = true
			}
		}
	}
	return updatedSeatArray2d, updateFlag
}

func countOccupied(seatArray2d [][]int) int {
	countOccupied := 0
	for r := range seatArray2d {
		for c := range seatArray2d[r] {
			if seatArray2d[r][c] == -1 { // occupied
				countOccupied++
			}
		}
	}
	return countOccupied
}

func printArray(arr [][]int) {
	array2d := [][]string{}
	for _, line := range arr {
		strLine := []string{}
		for _, val := range line {
			if val == 0 { // floor
				strLine = append(strLine, ".")
			} else if val == 1 { // empty
				strLine = append(strLine, "L")
			} else if val == -1 { // occupied
				strLine = append(strLine, "#")
			}
		}
		array2d = append(array2d, strLine)
	}

	for _, line := range array2d {
		fmt.Println(line)
	}
	fmt.Println()
}

func main() {

	const filename string = "input.txt"
	methods := []string{"neighbors", "lines"}
	minOccupiedSeats := []int{4, 5}
	for idx, method := range methods {
		seatArray2d := readInput(filename)
		updateFlag := true
		for {
			if !updateFlag {
				break
			}
			seatArray2d, updateFlag = update(seatArray2d, method, minOccupiedSeats[idx])
			// printArray(seatArray2d)
		}
		occupied := countOccupied(seatArray2d)
		fmt.Println("Occupied seats after chaos (method:", method, ") :", occupied)
	}
}
