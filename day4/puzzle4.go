package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var passports []string
	passport := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && passport != "" { // new passport
			passports = append(passports, passport)
			passport = ""
			continue
		}
		passport = passport + " " + line
	}
	if passport != "" {
		passports = append(passports, passport)
		passport = ""
	}

	return passports
}

func checkPassport(passport string, required []string, optional []string) (bool, map[string]string) {
	passportMap := make(map[string]string)
	fieldsAndValues := strings.Split(passport, " ")
	for _, fvs := range fieldsAndValues {
		if fvs == "" {
			continue
		}
		fv := strings.Split(fvs, ":")
		field := fv[0]
		value := fv[1]
		passportMap[field] = value
	}

	for _, reqKey := range required {
		_, keyExists := passportMap[reqKey]
		if !keyExists {
			return false, passportMap
		}
	}
	return true, passportMap
}

func checkValidPassport(passportMap map[string]string, required []string, optional []string) bool {

	byr, err := strconv.Atoi(passportMap["byr"])
	if err != nil || byr < 1920 || byr > 2002 {
		fmt.Println("Invalid birthday year: ", byr)
		return false
	}

	iyr, err := strconv.Atoi(passportMap["iyr"])
	if err != nil || iyr < 2010 || iyr > 2020 {
		fmt.Println("Invalid issue year: ", iyr)
		return false
	}

	eyr, err := strconv.Atoi(passportMap["eyr"])
	if err != nil || eyr < 2020 || eyr > 2030 {
		fmt.Println("Invalid expiration year: ", eyr)
		return false
	}

	hgtUnit := passportMap["hgt"][len(passportMap["hgt"])-2:]
	hgt, err := strconv.Atoi(passportMap["hgt"][:len(passportMap["hgt"])-2])
	if err != nil || (hgtUnit == "cm" && (hgt < 150 || hgt > 193)) || (hgtUnit == "in" && (hgt < 59 || hgt > 76)) {
		fmt.Println("Invalid height: ", hgt, hgtUnit)
		return false
	}

	hcl := passportMap["hcl"]
	if len(hcl) != 7 || hcl[0] != '#' {
		fmt.Println("Invalid hair color: ", hcl)
		return false
	}
	for _, letter := range hcl[1:] {
		isValidChar := false
		isValidNum := false
		if letter >= 'a' && letter <= 'f' {
			isValidChar = true
		}
		if letter >= '0' && letter <= '9' {
			isValidNum = true
		}
		if !isValidChar && !isValidNum {
			fmt.Println("Invalid hair color: ", hcl)
			return false
		}
	}

	validEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	isValidEye := false
	for _, eyeColor := range validEyeColors {
		if eyeColor == passportMap["ecl"] {
			isValidEye = true
			break
		}
	}
	if !isValidEye {
		fmt.Println("Invalid eye color: ", hcl)
		return false
	}

	pid, err := strconv.Atoi(passportMap["pid"])
	if ((err != nil) || len(passportMap["pid"]) != 9) || (pid < 0 || pid > 999999999) {
		fmt.Println("Invalid passport id: ", passportMap["pid"])
		return false
	}

	return true
}

func main() {

	passports := readInput("input4.txt")
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	optional := []string{"cid"}
	countPassports := 0
	countValidPasports := 0
	for idx, passport := range passports {
		fmt.Println(idx, ":", passport)
		isPassport, passportMap := checkPassport(passport, required, optional)
		if isPassport {
			countPassports++
			isValidPassport := checkValidPassport(passportMap, required, optional)
			if isValidPassport {
				countValidPasports++
			}
		}
	}

	fmt.Println("Number of passports:", countPassports)
	fmt.Println("Number of valid passports:", countValidPasports)
}
