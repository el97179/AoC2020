package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(filename string) map[int][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// var allAnswers [][]string
	allAnswers := make(map[int][]string, 0)
	groupAnswers := make([]string, 0)
	personAnswer := ""
	scanner := bufio.NewScanner(file)
	groupID := 0
	for scanner.Scan() {
		personAnswer = scanner.Text()
		if personAnswer == "" { // new group
			allAnswers[groupID] = groupAnswers
			groupAnswers = nil
			groupID++
			continue
		}
		groupAnswers = append(groupAnswers, personAnswer)
	}
	if personAnswer != "" { // to add the last group if the file doesn't include an empty line at the end
		allAnswers[groupID] = append(allAnswers[groupID], groupAnswers...)
		groupID++
	}

	return allAnswers
}

func convertGroupAnswersToMaps(groupAnswers []string) []map[rune]bool {
	groupAnswersMap := make([]map[rune]bool, 0)
	for _, personAnswer := range groupAnswers {
		personAnswerMap := make(map[rune]bool)
		for _, letter := range personAnswer {
			if 'a' <= letter && letter <= 'z' {
				personAnswerMap[letter] = true
			}
		}
		groupAnswersMap = append(groupAnswersMap, personAnswerMap)
	}
	return groupAnswersMap
}

func countGroupAnswersAny(groupAnswers []string) int {
	groupAnswersMap := convertGroupAnswersToMaps(groupAnswers)
	unionAnswers := groupAnswersMap[0]
	for _, personAnswerMap := range groupAnswersMap {
		for k := range personAnswerMap {
			unionAnswers[k] = true
		}
	}
	return countTrues(unionAnswers)
}

func countGroupAnswersAll(groupAnswers []string) int {
	groupAnswersMap := convertGroupAnswersToMaps(groupAnswers)
	intersectionAnswers := groupAnswersMap[0]
	for _, personAnswerMap := range groupAnswersMap {
		for k := range intersectionAnswers {
			if !personAnswerMap[k] {
				intersectionAnswers[k] = false
			}
		}
	}
	return countTrues(intersectionAnswers)
}

func countTrues(boolMap map[rune]bool) int {
	groupAnswers := 0
	for _, val := range boolMap {
		if val {
			groupAnswers++
		}
	}
	return groupAnswers
}

func main() {

	allAnswers := readInput("input6.txt")
	countAnswers1, countAnswers2 := 0, 0
	for _, groupAnswers := range allAnswers {
		countAnswers1 += countGroupAnswersAny(groupAnswers)
		countAnswers2 += countGroupAnswersAll(groupAnswers)
	}
	fmt.Println("Group answers any:", countAnswers1)
	fmt.Println("Group answers all:", countAnswers2)

}
