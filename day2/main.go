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
	prevLevel := 0
	prevDifference := 0
	incrementSafeReport := true
	for index, level := range levels {
		if index == 0 {
			prevLevel = level
		} else {
			difference := level - prevLevel

			if (index == 1 && difference == 0) || (index > 1 && !(difference != 0 && ((prevDifference < 0 && difference < 0) || (prevDifference > 0 && difference > 0)))) {
				incrementSafeReport = false
			} else {
				prevDifference = difference
				prevLevel = level
			}

			if difference < 0 {
				difference *= -1
			}
			if difference < 1 || difference > 3 {
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
				leftSide := make([]int, i)
				rightSide := make([]int, len(levels)-len(leftSide)-1)
				copy(leftSide, levels[:i])
				copy(rightSide, levels[i+1:])
				dampenedLevels := append(leftSide, rightSide...)
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
