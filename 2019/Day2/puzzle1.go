package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// 1 = add
// 2 = multiply
// 99 = finish program
// unknown = something went wrong

func calculateNewState(initialState []int) []int {
	return []int{}
}

func getListFromFile(inputFile string) []string {
	file, _ := os.Open(inputFile)
	reader := csv.NewReader(file)
	initialList, _ := reader.Read()
	return initialList
}

func toIntList(stringList []string) []int {
	var intList []int
	for value := range stringList {
		num, _ := strconv.Atoi(stringList[value])
		intList = append(intList, num)
	}
	return intList
}

// testing

func d2p1TestCase(initialState, expectedFinalState []int) {
	finalState := calculateNewState(initialState)
	fmt.Printf("initial: %v \n   final: %v \nexpected: %v", initialState, finalState, expectedFinalState)
}

func d2p1runTestCases() {
	d2p1TestCase([]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99})
	d2p1TestCase([]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99})
	d2p1TestCase([]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801})
	d2p1TestCase([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99})
}
