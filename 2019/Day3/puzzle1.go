package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	runPuzzle1()
	//runTestCasesPuzzle1()
}

func runPuzzle1() {
	fileWires := getWiresFromFile("2019/Day3/input_file.txt")
	wires := processWires(fileWires)
	intersections := findIntersections(wires[0], wires[1])
	dists := findDistances(intersections)
	min := findMin(dists)
	fmt.Printf("Min distance: %d", min)
}

func findMin(distances []int) int {
	min := distances[1]
	for _, dist := range distances[1:] {
		if dist < min && dist != 0 {
			min = dist
		}
	}
	return min
}

func findDistances(intersections [][]int) []int {
	var distances []int
	for _, coord := range intersections {
		distances = append(distances, manhattenDistance(coord))
	}
	return distances
}

func manhattenDistance(coord []int) int {
	x := coord[0]
	y := coord[1]
	if coord[0] < 0 {
		x = x * -1
	}
	if coord[1] < 0 {
		y = y * -1
	}
	return x + y
}

func findIntersections(wire1, wire2 *Wire) [][]int {
	var intersections [][]int
	for _, coord1 := range wire1.coords {
		for _, coord2 := range wire2.coords {
			if coord1[0] == coord2[0] && coord1[1] == coord2[1] {
				fmt.Printf("Found intersection at (%d, %d)\n", coord1[0], coord1[1])
				intersections = append(intersections, coord1)
			}
		}
	}
	return intersections
}

func processWires(wireStrs [][]string) []*Wire {
	var wires []*Wire
	for _, wireStr := range wireStrs {
		wires = append(wires, handleWire(wireStr))
	}
	return wires
}

func getWiresFromFile(inputFile string) [][]string {
	file, _ := os.Open(inputFile)
	reader := csv.NewReader(file)
	wires, _ := reader.ReadAll()
	return wires
}

func handleWire(moves []string) *Wire {
	wire := newWire()
	for _, move := range moves {
		wire = wire.addMovement(move)
	}
	return wire

}

// testing
func puzzle1TestCase(wiresList [][]string, expected int) {
	wires := processWires(wiresList)
	fmt.Printf("Wire1 coords: %v\n", wires[0].coords)
	intersections := findIntersections(wires[0], wires[1])
	fmt.Printf("Intersections: %v\n", intersections)
	distances := findDistances(intersections)
	fmt.Printf("Distances: %v\n", distances)
	minDist := findMin(distances)
	pass := minDist == expected
	fmt.Printf("Pass? %t    Min dist: %d\n---\n", pass, minDist)

}

func runTestCasesPuzzle1() {
	wire1 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	wire2 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	puzzle1TestCase([][]string{wire1, wire2}, 159)
	wire3 := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	wire4 := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
	puzzle1TestCase([][]string{wire3, wire4}, 135)
}
