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

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

func main() {
	f, err := os.Open("sample.txt")

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
					tempSolutionCoords := strconv.Itoa(row) + "," + strconv.Itoa(column) + " "
					tempSolutionCoords += strconv.Itoa(row) + "," + strconv.Itoa(column+1) + " "
					tempSolutionCoords += strconv.Itoa(row) + "," + strconv.Itoa(column+2) + " "
					tempSolutionCoords += strconv.Itoa(row) + "," + strconv.Itoa(column+3)
					fmt.Printf("found forward solution at: %s\n\n", tempSolutionCoords)
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
					tempSolutionCoords := strconv.Itoa(row) + "," + strconv.Itoa(column) + " "
					tempSolutionCoords += strconv.Itoa(row) + "," + strconv.Itoa(column-1) + " "
					tempSolutionCoords += strconv.Itoa(row) + "," + strconv.Itoa(column-2) + " "
					tempSolutionCoords += strconv.Itoa(row) + "," + strconv.Itoa(column-3)
					fmt.Printf("found backward solution at: %s\n\n", tempSolutionCoords)
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
					tempSolutionCoords := strconv.Itoa(row) + "," + strconv.Itoa(column) + " "
					tempSolutionCoords += strconv.Itoa(row-1) + "," + strconv.Itoa(column) + " "
					tempSolutionCoords += strconv.Itoa(row-2) + "," + strconv.Itoa(column) + " "
					tempSolutionCoords += strconv.Itoa(row-3) + "," + strconv.Itoa(column)
					fmt.Printf("found up solution at: %s\n\n", tempSolutionCoords)
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
					tempSolutionCoords := strconv.Itoa(row) + "," + strconv.Itoa(column) + " "
					tempSolutionCoords += strconv.Itoa(row+1) + "," + strconv.Itoa(column) + " "
					tempSolutionCoords += strconv.Itoa(row+2) + "," + strconv.Itoa(column) + " "
					tempSolutionCoords += strconv.Itoa(row+3) + "," + strconv.Itoa(column)
					fmt.Printf("found down solution at: %s\n\n", tempSolutionCoords)
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
					tempSolutionCoords := strconv.Itoa(row) + "," + strconv.Itoa(column) + " "
					tempSolutionCoords += strconv.Itoa(row+1) + "," + strconv.Itoa(column+1) + " "
					tempSolutionCoords += strconv.Itoa(row+2) + "," + strconv.Itoa(column+2) + " "
					tempSolutionCoords += strconv.Itoa(row+3) + "," + strconv.Itoa(column+3)
					fmt.Printf("found downward slope diagonal solution at: %s\n\n", tempSolutionCoords)
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
					tempSolutionCoords := strconv.Itoa(row) + "," + strconv.Itoa(column) + " "
					tempSolutionCoords += strconv.Itoa(row-1) + "," + strconv.Itoa(column-1) + " "
					tempSolutionCoords += strconv.Itoa(row-2) + "," + strconv.Itoa(column-2) + " "
					tempSolutionCoords += strconv.Itoa(row-3) + "," + strconv.Itoa(column-3)
					fmt.Printf("found foward slope diagonal solution at: %s\n\n", tempSolutionCoords)
				}
			}
		}

	}
	fmt.Printf("Day 4 Part 1: %d\n", matches)
}
