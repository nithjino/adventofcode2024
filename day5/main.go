package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func IgnoreError[T any](val T, err error) T {
	return val
}

func checkRules(pages []string, rules [][]string) (bool, []string) {
	for _, rule := range rules {
		x := rule[0]
		y := rule[1]

		xIndex := slices.Index(pages, x)
		yIndex := slices.Index(pages, y)

		if xIndex == -1 || yIndex == -1 {
			continue
		}

		if xIndex > yIndex {
			return false, rule
		}
	}
	return true, []string{"-1", "-1"}
}

func processInputText(content string, delim string) [][]string {
	processed := [][]string{}
	for _, c := range strings.Split(content, "\n") {
		processed = append(processed, strings.Split(c, delim))
	}
	return processed
}

func main() {
	contentBytes, err := os.ReadFile("day5.txt")

	if err != nil {
		log.Fatalln(err)
	}

	content := strings.Split(string(contentBytes), "\n\n")

	rules := processInputText(content[0], "|")
	pageNumberUpdates := processInputText(content[1], ",")

	middlePageSum := 0
	correctedMiddlePageSum := 0
	failedCheckPages := [][][]string{}

	for _, pages := range pageNumberUpdates {
		passesRules, failedRule := checkRules(pages, rules)

		if passesRules {
			middlePageNum, err := strconv.Atoi(pages[len(pages)/2])

			if err != nil {
				log.Fatalln(err)
			}

			middlePageSum += middlePageNum
		} else {
			failedCheckPages = append(failedCheckPages, [][]string{pages, failedRule})
		}
	}

	for _, failedPageSet := range failedCheckPages {
		fmt.Println(failedPageSet)
		failedPages := failedPageSet[0]
		failedRule := failedPageSet[1]
		failedX := failedRule[0]
		failedY := failedRule[1]

		failedXIndex := slices.Index(failedPages, failedX)
		failedYIndex := slices.Index(failedPages, failedY)
		failedPages[failedXIndex], failedPages[failedYIndex] = failedPages[failedYIndex], failedPages[failedXIndex]
		passTest := false
		passTest, failedRule = checkRules(failedPages, rules)

		for !passTest {
			failedX = failedRule[0]
			failedY = failedRule[1]
			failedXIndex = slices.Index(failedPages, failedX)
			failedYIndex = slices.Index(failedPages, failedY)
			fmt.Printf("Before swap: %s\n", failedPages)
			fmt.Printf("failed rule: %s\n", failedRule)
			fmt.Printf("failedXIndex: %d\n", failedXIndex)
			fmt.Printf("failedYIndex: %d\n", failedYIndex)
			failedPages[failedXIndex], failedPages[failedYIndex] = failedPages[failedYIndex], failedPages[failedXIndex]
			fmt.Printf("After swap: %s\n", failedPages)
			passTest, failedRule = checkRules(failedPages, rules)
		}
		fmt.Printf("\n%s now passes all rule checks\n", failedPages)
		middlePageNum, err := strconv.Atoi(failedPages[len(failedPages)/2])

		if err != nil {
			log.Fatalln(err)
		}

		correctedMiddlePageSum += middlePageNum

	}

	fmt.Printf("\nDay 5 Part 1: %d\n", middlePageSum)
	fmt.Printf("Day 5 Part 2: %d\n", correctedMiddlePageSum)
}
