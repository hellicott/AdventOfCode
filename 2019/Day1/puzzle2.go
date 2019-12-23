package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func totalFuel(inputFile string) {
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)
	totalFuelRequired := 0
	for scanner.Scan() {
		line := scanner.Text()
		mass, _ := strconv.Atoi(line)
		totalFuelRequired = totalFuelRequired + calculateRequiredFuel(mass) + calculateAdditionalFuel(calculateRequiredFuel(mass))
	}
	fmt.Printf("Fuel Required: %d", totalFuelRequired)
}

func calculateAdditionalFuel(fuel int) int {
	requiredFuel := calculateRequiredFuel(fuel)
	fmt.Printf("Required Additional Fuel for [%d] is [%d]\n", fuel, requiredFuel)
	if requiredFuel <= 0 {
		return 0
	} else {
		return requiredFuel + calculateAdditionalFuel(requiredFuel)
	}
}

func puzzle2TestCase(startingFuel, expectedAdditionalFuel int) {
	fuel := calculateAdditionalFuel(startingFuel)
	correct := fuel == expectedAdditionalFuel
	fmt.Printf("%t | starting = %d, additional = %d, expected = %d \n", correct, startingFuel, fuel, expectedAdditionalFuel)
}

func runTestCasesPuzzle2() {
	puzzle2TestCase(2, 0)
	puzzle2TestCase(654, 312)
	puzzle2TestCase(33583, 16763)
}
