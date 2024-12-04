package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func IgnoreError[T any](val T, err error) T {
	return val
}

func main() {
	corruptedMemoryBytes, err := os.ReadFile("day3.txt")

	if err != nil {
		log.Fatalln(err)
	}

	regexPattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	if err != nil {
		log.Fatalln(err)
	}

	matches := regexPattern.FindAllString(string(corruptedMemoryBytes), -1)

	sum := 0
	regexPatternDigits := regexp.MustCompile(`\d{1,3}`)

	for _, match := range matches {
		numberPair := regexPatternDigits.FindAllString(match, -1)
		sum += IgnoreError(strconv.Atoi(numberPair[0])) * IgnoreError(strconv.Atoi(numberPair[1]))
	}

	fmt.Printf("Part 1: %d\n", sum)

	sum = 0
	regexPattern = regexp.MustCompile(`(do\(\))|(don't\(\))|(mul\(\d{1,3},\d{1,3}\))`)

	matches = regexPattern.FindAllString(string(corruptedMemoryBytes), -1)

	addMul := true
	for _, match := range matches {

		switch match {
		case "don't()":
			addMul = false
		case "do()":
			addMul = true
		default:
			if addMul {
				numberPair := regexPatternDigits.FindAllString(match, -1)
				sum += IgnoreError(strconv.Atoi(numberPair[0])) * IgnoreError(strconv.Atoi(numberPair[1]))
			}
		}
	}

	fmt.Printf("Part 2: %d\n", sum)
}
