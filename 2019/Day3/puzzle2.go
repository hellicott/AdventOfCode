package main

import "fmt"

func runPuzzle2() {
	fileWires := getWiresFromFile("2019/Day3/input_file.txt")
	wires := processWires(fileWires)
	_, steps := findIntersectionsAndSteps(wires[0], wires[1])
	min := findMin(steps)
	fmt.Printf("Min distance: %d", min)
}

func findIntersectionsAndSteps(wire1, wire2 *Wire) ([][]int, []int) {
	var intersections [][]int
	var steps []int
	for i, coord1 := range wire1.coords {
		for j, coord2 := range wire2.coords {
			if coord1[0] == coord2[0] && coord1[1] == coord2[1] {
				fmt.Printf("Found intersection at (%d, %d)\n", coord1[0], coord1[1])
				steps = append(steps, i+j)
				intersections = append(intersections, coord1)
			}
		}
	}
	return intersections, steps
}

func puzzle2TestCase(wiresList [][]string, expected int) {
	wires := processWires(wiresList)
	fmt.Printf("Wire1 coords: %v\n", wires[0].coords)
	_, steps := findIntersectionsAndSteps(wires[0], wires[1])
	fmt.Printf("Steps: %v\n", steps)
	minDist := findMin(steps)
	pass := minDist == expected
	fmt.Printf("Pass? %t    Min dist: %d\n---\n", pass, minDist)
}

func runTestCasesPuzzle2() {
	wire1 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	wire2 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	puzzle2TestCase([][]string{wire1, wire2}, 610)
	wire3 := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	wire4 := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
	puzzle2TestCase([][]string{wire3, wire4}, 410)
}
