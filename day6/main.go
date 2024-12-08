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
	fmt.Printf("entering rotateGuardPosition with %d\n", position)
	if slices.Equal(position, []int{0, -1}) { //pointing up
		fmt.Println("match pointing up")
		return []int{1, 0}, ">" //point right
	}
	if slices.Equal(position, []int{1, 0}) { //pointing right
		fmt.Println("match pointing right")
		return []int{0, 1}, "v" //point down
	}
	if slices.Equal(position, []int{0, 1}) { //pointing down
		fmt.Println("match pointing down")
		return []int{-1, 0}, "<" //point left
	}
	if slices.Equal(position, []int{-1, 0}) { //pointing left
		fmt.Println("match pointing left")
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
	count := 0

	guardPosition, movementDelta, guardCursor := findGuardPosition(mapGrid)
	if guardPosition == nil || movementDelta == nil {
		log.Fatalln("couldn't find guard in map grid")
	}

	printMap(mapGrid)
	fmt.Printf("movementDelta: %d\n", movementDelta)
	fmt.Printf("guardPosition: %d\n", guardPosition)
	fmt.Printf("lastY: %d\n", lastY)
	fmt.Printf("lastX: %d\n", lastX)
	fmt.Println("")

	for {
		nextPlaceX := guardPosition[0] + movementDelta[0]
		nextPlaceY := guardPosition[1] + movementDelta[1]

		fmt.Printf("nextPlaceX: %d\n", nextPlaceX)
		fmt.Printf("nextPlaceY: %d\n", nextPlaceY)

		if nextPlaceX >= lastX || nextPlaceY >= lastY || nextPlaceX < 0 || nextPlaceY < 0 {
			mapGrid[guardPosition[1]][guardPosition[0]] = "X"
			count++
			break
		}

		if mapGrid[nextPlaceY][nextPlaceX] != "#" {
			fmt.Println("setting old guard position as marked")
			if mapGrid[guardPosition[1]][guardPosition[0]] != "X" {
				mapGrid[guardPosition[1]][guardPosition[0]] = "X"
				count++
			}

			fmt.Println("advancing guard")
			guardPosition[0], guardPosition[1] = nextPlaceX, nextPlaceY
			fmt.Printf("new guardPosition: %d\n", guardPosition)
			mapGrid[guardPosition[1]][guardPosition[0]] = guardCursor
		} else {
			fmt.Printf("ran into obstacle at %d, %d. rotating guard position\n", nextPlaceX, nextPlaceY)
			movementDelta, guardCursor = rotateGuardPosition(movementDelta)
			mapGrid[guardPosition[1]][guardPosition[0]] = guardCursor
			fmt.Printf("new movementDelta: %d\n", movementDelta)
		}
		printMap(mapGrid)
		//os.Exit(0)
	}

	printMap(mapGrid)
	fmt.Printf("Day 6 Part 1: %d\n", findAllUnique(mapGrid))

}
