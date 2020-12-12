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

type position struct {
	x           int
	y           int
	orientation string
}

type movement struct {
	dir  string
	unit int
}

func readInput(filename string) []movement {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	instructions := []movement{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		d := string(line[0])
		u, _ := strconv.Atoi(line[1:])
		instructions = append(instructions, movement{dir: d, unit: u})
	}
	return instructions
}

func navigate(pos position, instr movement, withWaypoint bool, posShip position) (position, position) {
	switch instr.dir {
	case "N":
		pos = moveN(pos, instr.unit)
		break
	case "E":
		pos = moveE(pos, instr.unit)
		break
	case "S":
		pos = moveS(pos, instr.unit)
		break
	case "W":
		pos = moveW(pos, instr.unit)
		break
	case "R":
		if withWaypoint {
			pos = rotateWaypoint(pos, instr.unit)
		} else {
			pos = rotate(pos, instr.unit)
		}
		break
	case "L":
		instr.unit = 360 - instr.unit
		if withWaypoint {
			pos = rotateWaypoint(pos, instr.unit)
		} else {
			pos = rotate(pos, instr.unit)
		}
		break
	case "F":
		if withWaypoint {
			posShip = forwardShip(posShip, pos, instr.unit)
		} else {
			pos = forward(pos, instr.unit)
		}
		break
	default:
		log.Fatal("Unknown instruction:", instr.dir)
	}
	return pos, posShip
}

func moveN(pos position, unit int) position {
	pos.y += unit
	return pos
}

func moveE(pos position, unit int) position {
	pos.x += unit
	return pos
}

func moveS(pos position, unit int) position {
	pos.y -= unit
	return pos
}

func moveW(pos position, unit int) position {
	pos.x -= unit
	return pos
}

func forward(pos position, unit int) position {
	switch pos.orientation {
	case "N":
		return moveN(pos, unit)
	case "E":
		return moveE(pos, unit)
	case "S":
		return moveS(pos, unit)
	case "W":
		return moveW(pos, unit)
	default:
		return pos
	}
}

func forwardShip(pos position, wptPos position, unit int) position {
	pos.x += unit * wptPos.x
	pos.y += unit * wptPos.y
	return pos
}

func rotate(pos position, degrees int) position {
	orientations := "NESW"
	idx := strings.Index(orientations, pos.orientation)
	rotationShift := degrees / 90
	idx = (idx + rotationShift) % len(orientations)
	pos.orientation = string(orientations[idx])
	return pos
}

func rotateWaypoint(wptPos position, degrees int) position {
	x := float64(wptPos.x)
	y := float64(wptPos.y)
	r11 := math.Cos(float64(degrees) * 2 * math.Pi / 360)
	r12 := math.Sin(float64(degrees) * 2 * math.Pi / 360)
	r21 := math.Sin(-float64(degrees) * 2 * math.Pi / 360)
	r22 := math.Cos(float64(degrees) * 2 * math.Pi / 360)
	wptPos.x = int(math.Round(r11*x + r12*y))
	wptPos.y = int(math.Round(r21*x + r22*y))
	return wptPos
}

func main() {

	const filename string = "input.txt"
	instructions := readInput(filename)

	pos := position{0, 0, "E"}
	posShip := pos
	withWaypoint := false
	for _, instr := range instructions {
		pos, _ = navigate(pos, instr, withWaypoint, posShip)
	}
	manDist := int(math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
	fmt.Println("(Part1) Final position:", pos, "\n\tManhattan distance:", manDist)

	posShip = position{0, 0, ""}
	withWaypoint = true
	wptPos := position{10, 1, ""}
	for _, instr := range instructions {
		wptPos, posShip = navigate(wptPos, instr, withWaypoint, posShip)
	}
	manDist = int(math.Abs(float64(posShip.x)) + math.Abs(float64(posShip.y)))
	fmt.Println("(Part2) Final position:", posShip, "\n\tManhattan distance:", manDist)
}
