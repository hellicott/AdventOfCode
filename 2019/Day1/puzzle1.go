package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	//runTestCases()
	processFile("2019/Day1/puzzle1_input.txt")
}

func processFile(inputFile string) {
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)
	totalFuelRequired := 0
	for scanner.Scan() {
		line := scanner.Text()
		mass, _ := strconv.Atoi(line)
		totalFuelRequired = totalFuelRequired + calculateRequiredFuel(mass)
	}
	fmt.Printf("Fuel Required: %d", totalFuelRequired)
}

func calculateRequiredFuel(mass int) int {
	return int(math.Floor(float64(mass/3))) - 2
}

// Testing

func testCase(mass, expectedFuel int) {
	fuel := calculateRequiredFuel(mass)
	correct := fuel == expectedFuel
	fmt.Printf("%t | mass = %d, fuel = %d, expected = %d \n", correct, mass, fuel, expectedFuel)
}

func runTestCases() {
	testCase(12, 2)
	testCase(14, 2)
	testCase(1969, 654)
	testCase(100756, 33583)
}
