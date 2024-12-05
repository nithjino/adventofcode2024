package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func IgnoreError[T any](val T, err error) T {
	return val
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

func main() {
	f, err := os.Open("day4.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	wordSeachPuzzle := [][]string{}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		wordSeachPuzzle = append(wordSeachPuzzle, strings.Split(line[0], ""))
	}

	lastRow := len(wordSeachPuzzle) - 1
	lastColumn := len(wordSeachPuzzle[0]) - 1
	matchesDict := make(map[string]bool)
	matches := 0

	for row := range wordSeachPuzzle {
		for column := range wordSeachPuzzle[row] {
			letter := wordSeachPuzzle[row][column]
			if !(letter == "X" || letter == "S") {
				continue
			}
			//fmt.Printf("letter: %s\nrow: %d\ncolumn: %d\n", letter, row, column)
			//check forward
			if column+3 < lastColumn {
				tempSolution := wordSeachPuzzle[row][column]
				tempSolution += wordSeachPuzzle[row][column+1]
				tempSolution += wordSeachPuzzle[row][column+2]
				tempSolution += wordSeachPuzzle[row][column+3]
				if tempSolution == "XMAS" || tempSolution == "SAMX" {
					matches++

					tempSolutionCoords := []int{row, column}
					tempSolutionCoords = append(tempSolutionCoords, row, column+1)
					tempSolutionCoords = append(tempSolutionCoords, row, column+2)
					tempSolutionCoords = append(tempSolutionCoords, row, column+3)
					fmt.Printf("forward match temp solution: %s\n", tempSolution)
					fmt.Printf("match found at %d\n", tempSolutionCoords)
					slices.Sort(tempSolutionCoords)
					fmt.Printf("sorted tempCoords: %d\n", tempSolutionCoords)
					fmt.Printf("string sorted tempCoords: %s\n", arrayToString(tempSolutionCoords, ""))
					if !matchesDict[arrayToString(tempSolutionCoords, "")] {
						matchesDict[arrayToString(tempSolutionCoords, "")] = true
						fmt.Println("")
					} else {
						fmt.Printf("match with those coords already found: %s\n\n", arrayToString(tempSolutionCoords, ""))
					}
				}
			}
			//check backward
			if column-3 > -1 {
				tempSolution := wordSeachPuzzle[row][column]
				tempSolution += wordSeachPuzzle[row][column-1]
				tempSolution += wordSeachPuzzle[row][column-2]
				tempSolution += wordSeachPuzzle[row][column-3]
				if tempSolution == "XMAS" || tempSolution == "SAMX" {
					matches++

					tempSolutionCoords := []int{row, column}
					tempSolutionCoords = append(tempSolutionCoords, row, column-1)
					tempSolutionCoords = append(tempSolutionCoords, row, column-2)
					tempSolutionCoords = append(tempSolutionCoords, row, column-3)
					fmt.Printf("backward match temp solution: %s\n", tempSolution)
					fmt.Printf("match found at %d\n", tempSolutionCoords)
					slices.Sort(tempSolutionCoords)
					fmt.Printf("sorted tempCoords: %d\n", tempSolutionCoords)
					fmt.Printf("string sorted tempCoords: %s\n", arrayToString(tempSolutionCoords, ""))
					if !matchesDict[arrayToString(tempSolutionCoords, "")] {
						matchesDict[arrayToString(tempSolutionCoords, "")] = true
						fmt.Println("")
					} else {
						fmt.Printf("match with those coords already found: %s\n\n", arrayToString(tempSolutionCoords, ""))
					}
				}
			}
			//check up
			if row-3 > -1 {
				tempSolution := wordSeachPuzzle[row][column]
				tempSolution += wordSeachPuzzle[row-1][column]
				tempSolution += wordSeachPuzzle[row-2][column]
				tempSolution += wordSeachPuzzle[row-3][column]
				if tempSolution == "XMAS" || tempSolution == "SAMX" {
					matches++

					tempSolutionCoords := []int{row, column}
					tempSolutionCoords = append(tempSolutionCoords, row-1, column)
					tempSolutionCoords = append(tempSolutionCoords, row-2, column)
					tempSolutionCoords = append(tempSolutionCoords, row-3, column)
					fmt.Printf("up match temp solution: %s\n", tempSolution)
					fmt.Printf("match found at %d\n", tempSolutionCoords)
					slices.Sort(tempSolutionCoords)
					fmt.Printf("sorted tempCoords: %d\n", tempSolutionCoords)
					fmt.Printf("string sorted tempCoords: %s\n", arrayToString(tempSolutionCoords, ""))
					if !matchesDict[arrayToString(tempSolutionCoords, "")] {
						matchesDict[arrayToString(tempSolutionCoords, "")] = true
						fmt.Println("")
					} else {
						fmt.Printf("match with those coords already found: %s\n\n", arrayToString(tempSolutionCoords, ""))
					}
				}
			}
			//check down
			if row+3 < lastRow {
				tempSolution := wordSeachPuzzle[row][column]
				tempSolution += wordSeachPuzzle[row+1][column]
				tempSolution += wordSeachPuzzle[row+2][column]
				tempSolution += wordSeachPuzzle[row+3][column]
				if tempSolution == "XMAS" || tempSolution == "SAMX" {
					matches++

					tempSolutionCoords := []int{row, column}
					tempSolutionCoords = append(tempSolutionCoords, row+1, column)
					tempSolutionCoords = append(tempSolutionCoords, row+2, column)
					tempSolutionCoords = append(tempSolutionCoords, row+3, column)
					fmt.Printf("down match temp solution: %s\n", tempSolution)
					fmt.Printf("match found at %d\n", tempSolutionCoords)
					slices.Sort(tempSolutionCoords)
					fmt.Printf("sorted tempCoords: %d\n", tempSolutionCoords)
					fmt.Printf("string sorted tempCoords: %s\n", arrayToString(tempSolutionCoords, ""))
					if !matchesDict[arrayToString(tempSolutionCoords, "")] {
						matchesDict[arrayToString(tempSolutionCoords, "")] = true
						fmt.Println("")
					} else {
						fmt.Printf("match with those coords already found: %s\n\n", arrayToString(tempSolutionCoords, ""))
					}
				}
			}
			//check downward slope diagonal
			if column+3 < lastColumn && row+3 < lastRow {
				tempSolution := wordSeachPuzzle[row][column]
				tempSolution += wordSeachPuzzle[row+1][column+1]
				tempSolution += wordSeachPuzzle[row+2][column+2]
				tempSolution += wordSeachPuzzle[row+3][column+3]
				if tempSolution == "XMAS" || tempSolution == "SAMX" {
					matches++

					tempSolutionCoords := []int{row, column}
					tempSolutionCoords = append(tempSolutionCoords, row+1, column+1)
					tempSolutionCoords = append(tempSolutionCoords, row+2, column+2)
					tempSolutionCoords = append(tempSolutionCoords, row+3, column+3)
					fmt.Printf("downward slope match temp solution: %s\n", tempSolution)
					fmt.Printf("match found at %d\n", tempSolutionCoords)
					slices.Sort(tempSolutionCoords)
					fmt.Printf("sorted tempCoords: %d\n", tempSolutionCoords)
					fmt.Printf("string sorted tempCoords: %s\n", arrayToString(tempSolutionCoords, ""))
					if !matchesDict[arrayToString(tempSolutionCoords, "")] {
						matchesDict[arrayToString(tempSolutionCoords, "")] = true
						fmt.Println("")
					} else {
						fmt.Printf("match with those coords already found: %s\n\n", arrayToString(tempSolutionCoords, ""))
					}
				}
			}
			//check upward slope diagonal
			if column-3 > -1 && row-3 > -1 {
				tempSolution := wordSeachPuzzle[row][column]
				tempSolution += wordSeachPuzzle[row-1][column-1]
				tempSolution += wordSeachPuzzle[row-2][column-2]
				tempSolution += wordSeachPuzzle[row-3][column-3]
				if tempSolution == "XMAS" || tempSolution == "SAMX" {
					matches++

					tempSolutionCoords := []int{row, column}
					tempSolutionCoords = append(tempSolutionCoords, row-1, column-1)
					tempSolutionCoords = append(tempSolutionCoords, row-2, column-2)
					tempSolutionCoords = append(tempSolutionCoords, row-3, column-3)
					fmt.Printf("upward slope match temp solution: %s\n", tempSolution)
					fmt.Printf("match found at %d\n", tempSolutionCoords)
					slices.Sort(tempSolutionCoords)
					fmt.Printf("sorted tempCoords: %d\n", tempSolutionCoords)
					fmt.Printf("string sorted tempCoords: %s\n", arrayToString(tempSolutionCoords, ""))
					if !matchesDict[arrayToString(tempSolutionCoords, "")] {
						matchesDict[arrayToString(tempSolutionCoords, "")] = true
						fmt.Println("")
					} else {
						fmt.Printf("match with those coords already found: %s\n\n", arrayToString(tempSolutionCoords, ""))
					}
				}
			}
		}

	}
	fmt.Printf("Day 4 Part 1: %d\n", matches)
	fmt.Printf("Day 4 Part 1 dict: %d\n", len(matchesDict))

}
