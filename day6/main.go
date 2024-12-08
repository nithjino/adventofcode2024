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

func printMap(mapGrid [][]string) {
	for _, row := range mapGrid {
		fmt.Println(row)
	}
}

func findGuardPosition(mapGrid [][]string) ([]int, []int, string) { //position, movementDelta, character
	for y := range mapGrid {
		for x := range mapGrid[y] {

			switch mapGrid[y][x] {
			case "<":
				return []int{x, y}, []int{-1, 0}, mapGrid[y][x]
			case ">":
				return []int{x, y}, []int{1, 0}, mapGrid[y][x]
			case "^":
				return []int{x, y}, []int{0, -1}, mapGrid[y][x]
			case "v":
				return []int{x, y}, []int{0, 1}, mapGrid[y][x]
			}
		}
	}
	return nil, nil, ""
}

func rotateGuardPosition(position []int) ([]int, string) { //movementDelta, character
	if slices.Equal(position, []int{0, -1}) { //pointing up
		return []int{1, 0}, ">" //point right
	}
	if slices.Equal(position, []int{1, 0}) { //pointing right
		return []int{0, 1}, "v" //point down
	}
	if slices.Equal(position, []int{0, 1}) { //pointing down
		return []int{-1, 0}, "<" //point left
	}
	if slices.Equal(position, []int{-1, 0}) { //pointing left
		return []int{0, -1}, "^" //point up
	}
	return nil, ""
}

func findAllUnique(mapGrid [][]string) int {
	count := 0
	for y := range mapGrid {
		for x := range mapGrid[y] {
			if mapGrid[y][x] == "X" {
				count++
			}
		}
	}
	return count
}

func main() {
	f, err := os.Open("day6.txt")

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

	lastY := len(mapGrid)
	lastX := len(mapGrid[0])

	guardPosition, movementDelta, guardCursor := findGuardPosition(mapGrid)
	if guardPosition == nil || movementDelta == nil {
		log.Fatalln("couldn't find guard in map grid")
	}

	for {
		nextPlaceX := guardPosition[0] + movementDelta[0]
		nextPlaceY := guardPosition[1] + movementDelta[1]
		if nextPlaceX >= lastX || nextPlaceY >= lastY || nextPlaceX < 0 || nextPlaceY < 0 {
			mapGrid[guardPosition[1]][guardPosition[0]] = "X"
			break
		}

		if mapGrid[nextPlaceY][nextPlaceX] != "#" {
			mapGrid[guardPosition[1]][guardPosition[0]] = "X"
			guardPosition[0], guardPosition[1] = nextPlaceX, nextPlaceY
			mapGrid[guardPosition[1]][guardPosition[0]] = guardCursor
		} else {
			movementDelta, guardCursor = rotateGuardPosition(movementDelta)
			mapGrid[guardPosition[1]][guardPosition[0]] = guardCursor
		}

	}

	fmt.Printf("Day 6 Part 1: %d\n", findAllUnique(mapGrid))

}
