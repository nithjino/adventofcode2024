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

func checkRules(pages []string, rules [][]string) bool {
	for _, rule := range rules {
		x := rule[0]
		y := rule[1]

		xIndex := slices.Index(pages, x)
		yIndex := slices.Index(pages, y)

		if xIndex == -1 || yIndex == -1 {
			continue
		}

		if xIndex > yIndex {
			return false
		}
	}
	return true
}

func processInputText(content []string, delim string) [][]string {
	processed := [][]string{}
	for _, c := range content {
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

	rules := processInputText(strings.Split(content[0], "\n"), "|")
	pageNumberUpdates := processInputText(strings.Split(content[1], "\n"), ",")

	middlePageSum := 0
	correctedMiddlePageSum := 0

	for _, pages := range pageNumberUpdates {
		passesRules := checkRules(pages, rules)

		if passesRules {
			middlePageNum, err := strconv.Atoi(pages[len(pages)/2])

			if err != nil {
				log.Fatalln(err)
			}

			middlePageSum += middlePageNum
		}
	}
	fmt.Printf("\nDay 5 Part 1: %d\n", middlePageSum)
	fmt.Printf("Day 5 Part 2: %d\n", correctedMiddlePageSum)
}
