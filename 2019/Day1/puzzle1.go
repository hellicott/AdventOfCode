package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	//d1p1RunTestCases()
	//processFile("2019/Day1/puzzle1_input.txt")
	//d1p2RunTestCases()
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

func d1p1testCase(mass, expectedFuel int) {
	fuel := calculateRequiredFuel(mass)
	correct := fuel == expectedFuel
	fmt.Printf("%t | mass = %d, fuel = %d, expected = %d \n", correct, mass, fuel, expectedFuel)
}

func d1p1RunTestCases() {
	d1p1testCase(12, 2)
	d1p1testCase(14, 2)
	d1p1testCase(1969, 654)
	d1p1testCase(100756, 33583)
}
