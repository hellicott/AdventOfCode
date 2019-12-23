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

func main() {
	//rund2p1()
	//d2p1runTestCases()
	rund2p2()
}

func rund2p1() {
	list := toIntList(getListFromFile("2019/Day2/puzzle_input.txt"))
	final := calculateNewState(list)
	fmt.Printf("%v", final)
}

func calculateNewState(state []int) []int {
	fmt.Printf("new: %v", state)
	maxPointer := len(state)
	pointer := 0
	for pointer < maxPointer {
		fmt.Printf("POINTER: %d\n", pointer)
		opcode := state[pointer]
		if opcode == 99 {
			fmt.Printf("Found opcode %d\n", opcode)
			break
		}
		num1 := state[pointer+1]
		num2 := state[pointer+2]
		output := state[pointer+3]
		state = handleOpcode(opcode, num1, num2, output, state)
		pointer += 4
	}
	return state
}

func handleOpcode(opcode, num1, num2, output int, state []int) []int {
	if opcode == 1 {
		fmt.Printf("state[%d] = state[%d] (%d) + state[%d] (%d) \n", output, num1, state[num1], num2, state[num2])
		state[output] = state[num1] + state[num2]
		fmt.Printf("RESULT: %d \n---\n", state[output])

	} else if opcode == 2 {
		fmt.Printf("state[%d] = state[%d] (%d) * state[%d] (%d) \n", output, num1, state[num1], num2, state[num2])
		state[output] = state[num1] * state[num2]
		fmt.Printf("RESULT: %d \n---\n", state[output])
	} else {
		fmt.Printf("Something Went Wrong :( got opcode %d\n", opcode)
	}
	return state
}

func getListFromFile(inputFile string) []string {
	file, _ := os.Open(inputFile)
	reader := csv.NewReader(file)
	initialList, _ := reader.Read()
	fmt.Printf("str: %v\n", initialList)
	return initialList
}

func toIntList(stringList []string) []int {
	var intList []int
	for value := range stringList {
		num, _ := strconv.Atoi(stringList[value])
		intList = append(intList, num)
	}
	fmt.Printf("int: %v\n", intList)
	return intList
}

// testing

func d2p1TestCase(initialState, expectedFinalState []int) {
	finalState := calculateNewState(initialState)
	fmt.Printf("   final: %v \nexpected: %v\n\n", finalState, expectedFinalState)
}

func d2p1runTestCases() {
	d2p1TestCase([]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99})
	d2p1TestCase([]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99})
	d2p1TestCase([]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801})
	d2p1TestCase([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99})
}
