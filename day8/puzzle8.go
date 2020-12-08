package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Instruction is a structure to hold the name and value of an instruction
type Instruction struct {
	id    int    // execution id
	name  string // name of the operation
	value int    // value for the operation
}

var glAccumulator int = 0
var glInstructionList = []Instruction{}
var glExecutedInstructionIds = make(map[int]bool)

func initGlobalVars(filename string) {
	glAccumulator = 0
	glInstructionList = readInput(filename)
	glExecutedInstructionIds = make(map[int]bool)
}

func readInput(filename string) []Instruction {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	instructionsList := []Instruction{}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		op := lineSplit[0]
		val, _ := strconv.Atoi(lineSplit[1])
		instr := Instruction{id: i, name: op, value: val}
		instructionsList = append(instructionsList, instr)
		i++
	}
	return instructionsList
}

func getInstruction(id int) Instruction {
	if id < 0 || id > len(glInstructionList) {
		log.Fatal("Operation id is beyond bounds, returning empty Instruction")
		return Instruction{}
	} else if id == len(glInstructionList) {
		return Instruction{id: id, name: "EOF", value: 0}
	}
	return glInstructionList[id]
}

func execInstruction(instr Instruction) int {
	// fmt.Println("Executing:", instr.name, instr.value)
	id := instr.id
	glExecutedInstructionIds[id] = true
	switch instr.name {
	case "acc":
		execAcc(instr.value)
		id++
	case "jmp":
		id = execJmp(instr.value, id)
	case "nop":
		id = execNop(id)
	default:
		log.Fatal("Unknown instruction!")
	}
	nextInstruction := getInstruction(id)
	if glExecutedInstructionIds[nextInstruction.id] {
		fmt.Println("Instruction", nextInstruction, "has been executed before, we entered a loop!")
		return -1
	}
	if nextInstruction.name == "EOF" {
		fmt.Println("Reached end of the code!")
		return 0
	}
	return execInstruction(nextInstruction)
}

func execAcc(val int) {
	glAccumulator += val
}

func execJmp(val int, id int) int {
	id += val
	return id
}

func execNop(id int) int {
	id++
	return id
}

func main() {

	filename := "input8.txt"
	initGlobalVars(filename)
	for idx, instr := range glInstructionList {
		fmt.Println("Run", idx+1, "/", len(glInstructionList), ":")
		if instr.name != "jmp" && instr.name != "nop" {
			continue
		}
		retCode := execInstruction(glInstructionList[0])
		if idx == 0 {
			fmt.Println(">> Accumulator value after hitting first loop (part1) :", glAccumulator)
		}
		if retCode == 0 {
			fmt.Println(">> Accumulator value after completing code (part2) :", glAccumulator)
			break
		} else if retCode == -1 { // in a loop, try swapping exactly one jmp or nop instructions
			initGlobalVars(filename)
			if instr.name == "jmp" {
				glInstructionList[idx].name = "nop"
			}
			if instr.name == "nop" {
				glInstructionList[idx].name = "jmp"
			}
			continue
		}
	}
}
