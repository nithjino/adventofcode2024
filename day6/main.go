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

func findGuardPosition(mapGrid [][]string) ([]int, []int) {
	for y := range mapGrid {
		for x := range mapGrid[y] {

			switch mapGrid[x][y] {
			case "<":
				return []int{y, x}, []int{-1, 0}
			case ">":
				return []int{y, x}, []int{1, 0}
			case "^":
				return []int{y, x}, []int{0, 1}
			case "v":
				return []int{y, x}, []int{0, -1}
			}
		}
	}
	return nil, nil
}

func printMap(mapGrid [][]string) {
	for _, row := range mapGrid {
		fmt.Println(row)
	}
}

func rotateGuardPosition(position []int) []int {
	fmt.Printf("entering rotateGuardPosition with %d\n", position)
	if slices.Equal(position, []int{0, 1}) { //pointing up
		fmt.Println("match pointing up")
		return []int{1, 0} //point right
	}
	if slices.Equal(position, []int{1, 0}) { //pointing right
		fmt.Println("match pointing right")
		return []int{0, -1} //point down
	}
	if slices.Equal(position, []int{0, -1}) { //pointing down
		fmt.Println("match pointing down")
		return []int{-1, 0} //point left
	}
	if slices.Equal(position, []int{-1, 0}) { //pointing left
		fmt.Println("match pointing left")
		return []int{0, 1} //point up
	}
	return nil
}

func main() {
	f, err := os.Open("sample.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	mapGrid := [][]string{}
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Fields(text)
		mapGrid = append(mapGrid, strings.Split(line[0], ""))
	}

	lastRow := len(mapGrid)
	lastColumn := len(mapGrid[0])
	count := 0

	guardPosition, movementDelta := findGuardPosition(mapGrid)
	if guardPosition == nil || movementDelta == nil {
		log.Fatalln("couldn't find guard in map grid")
	}

	printMap(mapGrid)
	fmt.Printf("movementDelta: %d\n", movementDelta)
	fmt.Printf("guardPosition: %d\n", guardPosition)
	fmt.Printf("lastRow: %d\n", lastRow)
	fmt.Printf("lastColumn: %d\n", lastColumn)
	fmt.Println("")

	for {
		nextPlaceX := guardPosition[0] + movementDelta[0]
		nextPlaceY := guardPosition[1] + movementDelta[1]

		fmt.Printf("nextPlaceX: %d\n", nextPlaceX)
		fmt.Printf("nextPlaceY: %d\n", nextPlaceY)

		if nextPlaceX >= lastColumn || nextPlaceY >= lastRow || nextPlaceX < 0 || nextPlaceY < 0 {
			break
		}

		if mapGrid[nextPlaceX][nextPlaceY] != "#" {
			fmt.Println("setting old guard position as marked")
			mapGrid[guardPosition[0]][guardPosition[1]] = "X"
			fmt.Println("advancing guard")
			guardPosition[0], guardPosition[1] = nextPlaceX, nextPlaceY
			fmt.Printf("new guardPosition: %d\n", guardPosition)
			count++
		} else {
			fmt.Printf("ran into obstacle at %d, %d. rotating guard position\n", nextPlaceX, nextPlaceY)
			movementDelta = rotateGuardPosition(movementDelta)
			fmt.Printf("new movementDelta: %d\n", movementDelta)
		}
		printMap(mapGrid)
	}

	fmt.Printf("Day 6 Part 1: %d\n", count)
}
