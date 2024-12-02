package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day2.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	safeReports := 0
	//scanIndex := 0
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		var reports []int
		for _, number := range numbers {
			n, _ := strconv.Atoi(number)
			reports = append(reports, n)
		}

		prevReport := 0
		prevDifference := 0
		incrementSafeReport := true
		//fmt.Printf("going to process %d\n", reports)
		for index, report := range reports {
			if index == 0 {
				prevReport = report
			} else {
				//fmt.Printf("prevReport: %d\n", prevReport)
				//fmt.Printf("report: %d\n", report)

				difference := report - prevReport
				//fmt.Printf("prevDifference: %d\n", prevDifference)
				//fmt.Printf("difference: %d\n\n", difference)

				if (index == 1 && difference == 0) || (index > 1 && !(difference != 0 && ((prevDifference < 0 && difference < 0) || (prevDifference > 0 && difference > 0)))) {
					//fmt.Printf("%d is unsafe because the prevDifference and difference are different signs\nprevDifference: %d\ndifference: %d\n", reports, prevDifference, difference)
					//fmt.Printf("current report: %d\nprevReport: %d\n\n", report, prevReport)
					incrementSafeReport = false
				} else {
					prevDifference = difference
					prevReport = report
				}

				if difference < 0 {
					difference *= -1
				}
				if difference < 1 || difference > 3 {
					//fmt.Printf("%d is unsafe because of the difference is too big or 0.\ndifference: %d\n\n", reports, difference)
					//fmt.Printf("current report: %d\nprevReport: %d\n\n", report, prevReport)
					incrementSafeReport = false
				}
			}
		}
		if incrementSafeReport {
			safeReports++
		}

	}

	if err := scanner.Err(); err != nil { // if with assignment
		log.Fatalln(err)
	}
	fmt.Printf("number of safe reports: %d\n", safeReports)
}
