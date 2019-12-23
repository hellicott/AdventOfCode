package main

import (
	"fmt"
	"strconv"
)

// wire class/struct
type Wire struct {
	coords [][]int
}

func (w Wire) addMovement(move string) *Wire {
	currentPosition := w.coords[len(w.coords)-1]
	for i := 1; i <= getDistance(move); i++ {
		w.coords = append(w.coords, shift(currentPosition, getDirection(move), i))
	}
	return &w
}

func getDirection(move string) string {
	return move[:1]
}

func getDistance(move string) int {
	dist, _ := strconv.Atoi(move[1:])
	return dist
}

func shift(currentPosition []int, direction string, distance int) []int {
	if direction == "R" {
		return []int{currentPosition[0] + distance, currentPosition[1]}
	}
	if direction == "L" {
		return []int{currentPosition[0] - distance, currentPosition[1]}
	}
	if direction == "U" {
		return []int{currentPosition[0], currentPosition[1] + distance}
	}
	if direction == "D" {
		return []int{currentPosition[0], currentPosition[1] - distance}
	} else {
		fmt.Printf("Invalid Direction: %s\n", direction)
		return []int{0}
	}
}

func newWire() *Wire {
	coord := []int{0, 0}
	coords := [][]int{coord}
	wire := Wire{coords: coords}
	return &wire
}
