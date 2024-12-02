package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day1.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var rightList []int
	var leftList []int
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		lNumber, _ := strconv.Atoi(numbers[0]) // not using err return value
		rNumber, _ := strconv.Atoi(numbers[1])

		leftList = append(leftList, lNumber)
		rightList = append(rightList, rNumber)
	}

	if err := scanner.Err(); err != nil { // if with assignment
		log.Fatalln(err)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)
	totalDistanceSum := 0
	similarityScore := 0
	for i := range leftList { // using index only
		distance := leftList[i] - rightList[i]
		if distance < 0 {
			distance = distance * -1
		}
		totalDistanceSum = totalDistanceSum + distance
	}
	fmt.Printf("total distance sum: %d\n", totalDistanceSum)

	for _, leftValue := range leftList { // not using index but only value
		if slices.Contains(rightList, leftValue) {
			counter := 0
			for _, rightValue := range rightList {
				if rightValue == leftValue {
					counter++
				}
			}
			similarityScore = similarityScore + (leftValue * counter)
		}
	}
	fmt.Printf("similarity score: %d\n", similarityScore)
}
