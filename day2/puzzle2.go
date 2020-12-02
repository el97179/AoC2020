package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkPassword(minRep int, maxRep int, letter string, password string) bool {
	rep := strings.Count(password, letter)
	if rep <= maxRep && rep >= minRep {
		return true
	}
	return false
}

func checkPassword2(pos1 int, pos2 int, letter string, password string) bool {
	if len(password) < pos1 || len(password) < pos2 {
		fmt.Println("Position exceeds password's length")
		return false
	}
	if string(password[pos1-1]) == letter && string(password[pos2-1]) != letter {
		return true
	} else if string(password[pos1-1]) != letter && string(password[pos2-1]) == letter {
		return true
	}
	return false
}
func main() {

	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	goodPasswords := 0
	goodPasswords2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splits1 := strings.Split(line, "-")
		if len(splits1) == 2 {
			minRep, _ := strconv.Atoi(splits1[0])
			splits2 := strings.Split(splits1[1], " ")
			if len(splits2) == 3 {
				maxRep, _ := strconv.Atoi(splits2[0])
				letter := strings.Split(splits2[1], ":")[0]
				password := splits2[2]
				fmt.Println(line)
				fmt.Println(minRep)
				fmt.Println(maxRep)
				fmt.Println(letter)
				fmt.Println(password)
				isOk := checkPassword(minRep, maxRep, letter, password)
				if isOk {
					fmt.Println("Password", password, "is good!")
					goodPasswords++
				} else {
					fmt.Println("Password", password, "is broken!")
				}
				isOk2 := checkPassword2(minRep, maxRep, letter, password)
				if isOk2 {
					fmt.Println("Password", password, "is good!")
					goodPasswords2++
				} else {
					fmt.Println("Password", password, "is broken!")
				}
			}
		} else {
			log.Fatal("wrong input line")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Good passwords:", goodPasswords)
	fmt.Println("Good passwords2:", goodPasswords2)
}
