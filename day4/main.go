package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func IgnoreError[T any](val T, err error) T {
	return val
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

					tempSolutionCoords := strconv.Itoa(row) + "-" + strconv.Itoa(column) + ","
					tempSolutionCoords += strconv.Itoa(row) + "-" + strconv.Itoa(column+1) + ","
					tempSolutionCoords += strconv.Itoa(row) + "-" + strconv.Itoa(column+2) + ","
					tempSolutionCoords += strconv.Itoa(row) + "-" + strconv.Itoa(column+3)
					fmt.Printf("foward match temp solution: %s\n", tempSolution)
					fmt.Printf("match found at %s\n\n", tempSolutionCoords)
					if !matchesDict[tempSolutionCoords] {
						matchesDict[tempSolutionCoords] = true
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

					tempSolutionCoords := strconv.Itoa(row) + "-" + strconv.Itoa(column) + ","
					tempSolutionCoords += strconv.Itoa(row) + "-" + strconv.Itoa(column-1) + ","
					tempSolutionCoords += strconv.Itoa(row) + "-" + strconv.Itoa(column-2) + ","
					tempSolutionCoords += strconv.Itoa(row) + "-" + strconv.Itoa(column-3)
					fmt.Printf("backward match temp solution: %s\n", tempSolution)
					fmt.Printf("match found at %s\n\n", tempSolutionCoords)
					if !matchesDict[tempSolutionCoords] {
						matchesDict[tempSolutionCoords] = true
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

					tempSolutionCoords := strconv.Itoa(row) + "-" + strconv.Itoa(column) + ","
					tempSolutionCoords += strconv.Itoa(row-1) + "-" + strconv.Itoa(column) + ","
					tempSolutionCoords += strconv.Itoa(row-2) + "-" + strconv.Itoa(column) + ","
					tempSolutionCoords += strconv.Itoa(row-3) + "-" + strconv.Itoa(column)
					fmt.Printf("up match temp solution: %s\n", tempSolution)
					fmt.Printf("match found at %s\n\n", tempSolutionCoords)
					if !matchesDict[tempSolutionCoords] {
						matchesDict[tempSolutionCoords] = true
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

					tempSolutionCoords := strconv.Itoa(row) + "-" + strconv.Itoa(column) + ","
					tempSolutionCoords += strconv.Itoa(row+1) + "-" + strconv.Itoa(column) + ","
					tempSolutionCoords += strconv.Itoa(row+2) + "-" + strconv.Itoa(column) + ","
					tempSolutionCoords += strconv.Itoa(row+3) + "-" + strconv.Itoa(column)
					fmt.Printf("down match temp solution: %s\n", tempSolution)
					fmt.Printf("match found at %s\n\n", tempSolutionCoords)
					if !matchesDict[tempSolutionCoords] {
						matchesDict[tempSolutionCoords] = true
					}
				}
			}
			//check forward diagonal
			if column+3 < lastColumn && row+3 < lastRow {
				tempSolution := wordSeachPuzzle[row][column]
				tempSolution += wordSeachPuzzle[row+1][column+1]
				tempSolution += wordSeachPuzzle[row+2][column+2]
				tempSolution += wordSeachPuzzle[row+3][column+3]
				if tempSolution == "XMAS" || tempSolution == "SAMX" {
					matches++

					tempSolutionCoords := strconv.Itoa(row) + "-" + strconv.Itoa(column) + ","
					tempSolutionCoords += strconv.Itoa(row+1) + "-" + strconv.Itoa(column+1) + ","
					tempSolutionCoords += strconv.Itoa(row+2) + "-" + strconv.Itoa(column+2) + ","
					tempSolutionCoords += strconv.Itoa(row+3) + "-" + strconv.Itoa(column+3)
					fmt.Printf("foward diagonal match temp solution: %s\n", tempSolution)
					fmt.Printf("match found at %s\n\n", tempSolutionCoords)
					if !matchesDict[tempSolutionCoords] {
						matchesDict[tempSolutionCoords] = true
					}
				}
			}
			//check back diagonal
			if column-3 > -1 && row-3 > -1 {
				tempSolution := wordSeachPuzzle[row][column]
				tempSolution += wordSeachPuzzle[row-1][column-1]
				tempSolution += wordSeachPuzzle[row-2][column-2]
				tempSolution += wordSeachPuzzle[row-3][column-3]
				if tempSolution == "XMAS" || tempSolution == "SAMX" {
					matches++

					tempSolutionCoords := strconv.Itoa(row) + "-" + strconv.Itoa(column) + ","
					tempSolutionCoords += strconv.Itoa(row-1) + "-" + strconv.Itoa(column-1) + ","
					tempSolutionCoords += strconv.Itoa(row-2) + "-" + strconv.Itoa(column-2) + ","
					tempSolutionCoords += strconv.Itoa(row-3) + "-" + strconv.Itoa(column-3)
					fmt.Printf("back diagonal match temp solution: %s\n", tempSolution)
					fmt.Printf("match found at %s\n\n", tempSolutionCoords)
					if !matchesDict[tempSolutionCoords] {
						matchesDict[tempSolutionCoords] = true
					}
				}
			}
		}

	}
	fmt.Printf("Day 4 Part 1: %d\n", matches)
	fmt.Printf("Day 4 Part 1 dict: %d\n", len(matchesDict))

}
