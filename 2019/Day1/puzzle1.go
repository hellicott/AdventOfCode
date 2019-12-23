package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	//runTestCasesPuzzle1()
	//processFile("2019/Day1/puzzle1_input.txt")
	//runTestCasesPuzzle2()
	totalFuel("2019/Day1/puzzle1_input.txt")
}

func processFile(inputFile string) int {
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)
	totalFuelRequired := 0
	for scanner.Scan() {
		line := scanner.Text()
		mass, _ := strconv.Atoi(line)
		totalFuelRequired = totalFuelRequired + calculateRequiredFuel(mass)
	}
	fmt.Printf("Fuel Required: %d", totalFuelRequired)
	return totalFuelRequired
}

func calculateRequiredFuel(mass int) int {
	return int(math.Floor(float64(mass/3))) - 2
}

// Testing

func puzzle1TestCase(mass, expectedFuel int) {
	fuel := calculateRequiredFuel(mass)
	correct := fuel == expectedFuel
	fmt.Printf("%t | mass = %d, fuel = %d, expected = %d \n", correct, mass, fuel, expectedFuel)
}

func runTestCasesPuzzle1() {
	puzzle1TestCase(12, 2)
	puzzle1TestCase(14, 2)
	puzzle1TestCase(1969, 654)
	puzzle1TestCase(100756, 33583)
}
