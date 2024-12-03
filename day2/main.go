package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafe(levels []int) bool {
	//scanIndex := 0
	prevLevel := 0
	prevDifference := 0
	incrementSafeReport := true
	//fmt.Printf("going to process %d\n", reports)
	for index, level := range levels {
		if index == 0 {
			prevLevel = level
		} else {
			//fmt.Printf("prevLevel: %d\n", prevLevel)
			//fmt.Printf("level: %d\n", level)

			difference := level - prevLevel
			//fmt.Printf("prevDifference: %d\n", prevDifference)
			//fmt.Printf("difference: %d\n\n", difference)

			if (index == 1 && difference == 0) || (index > 1 && !(difference != 0 && ((prevDifference < 0 && difference < 0) || (prevDifference > 0 && difference > 0)))) {
				//fmt.Printf("%d is unsafe because the prevDifference and difference are different signs\nprevDifference: %d\ndifference: %d\n", reports, prevDifference, difference)
				//fmt.Printf("current level: %d\nprevLevel: %d\n\n", level, prevLevel)
				incrementSafeReport = false
			} else {
				prevDifference = difference
				prevLevel = level
			}

			if difference < 0 {
				difference *= -1
			}
			if difference < 1 || difference > 3 {
				//fmt.Printf("%d is unsafe because of the difference is too big or 0.\ndifference: %d\n\n", reports, difference)
				//fmt.Printf("current level: %d\nprevLevel: %d\n\n", level, prevLevel)
				incrementSafeReport = false
			}
		}
	}
	return incrementSafeReport
}

func main() {
	f, err := os.Open("day2.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	safeReports := 0
	dampenedSafeReports := 0
	for scanner.Scan() {
		report := strings.Fields(scanner.Text())
		var levels []int
		for _, number := range report {
			n, _ := strconv.Atoi(number)
			levels = append(levels, n)
		}
		if isSafe(levels) {
			safeReports++
		} else {
			for i := range levels {
				//fmt.Printf("pre deletelevels: %d\n", levels)
				leftSide := make([]int, i)
				rightSide := make([]int, len(levels)-len(leftSide)-1)
				copy(leftSide, levels[:i])
				copy(rightSide, levels[i+1:])
				dampenedLevels := append(leftSide, rightSide...)
				//fmt.Printf("levels: %d\nremoving at index %d: %d\ndamepenedLevels: %d\n\n", levels, i, n, dampenedLevels)
				if isSafe(dampenedLevels) {
					dampenedSafeReports++
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil { // if with assignment
		log.Fatalln(err)
	}

	fmt.Printf("number of safe reports: %d\n", safeReports)
	fmt.Printf("number of safe reports including dampened: %d\n", safeReports+dampenedSafeReports)
}
