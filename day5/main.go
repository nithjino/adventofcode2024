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

func main() {
	contentBytes, err := os.ReadFile("day5.txt")

	if err != nil {
		log.Fatalln(err)
	}

	content := strings.Split(string(contentBytes), "\n\n")
	rules := strings.Split(content[0], "\n")
	pageNumberUpdates := strings.Split(content[1], "\n")
	middlePageSum := 0

	for _, pageNumberUpdate := range pageNumberUpdates {
		passesRules := true
		pages := strings.Split(pageNumberUpdate, ",")

		for _, rule := range rules {
			rulesSplt := strings.Split(rule, "|")
			x := rulesSplt[0]
			y := rulesSplt[1]

			xIndex := slices.Index(pages, x)
			yIndex := slices.Index(pages, y)

			if xIndex == -1 || yIndex == -1 {
				continue
			}

			if xIndex > yIndex {
				passesRules = false
				break
			}
		}

		if passesRules {
			middlePageNum, err := strconv.Atoi(pages[len(pages)/2])

			if err != nil {
				log.Fatalln(err)
			}

			middlePageSum += middlePageNum
		}
	}
	fmt.Printf("\nDay 5 Part 1: %d\n", middlePageSum)
}
