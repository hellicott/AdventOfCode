package main

import "fmt"

func rund2p2() {
Loop:
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			fmt.Printf("Trying combination [%d, %d]\n", noun, verb)
			state := getStartState("2019/Day2/puzzle_input.txt")
			endVal := tryCombination(noun, verb, state)
			fmt.Printf("Value at position 0 after processing: %d\n", endVal)
			if endVal == 19690720 {
				fmt.Printf("Value found with combination [%d, %d]\n", noun, verb)
				result := (noun * 100) + verb
				fmt.Printf("Advent of Code output: %d", result)
				break Loop
			}
		}

	}

}

func tryCombination(noun, verb int, startState []int) int {
	startState[1] = noun
	startState[2] = verb
	value := calculateNewState(startState)[0]
	return value
}

func getStartState(inputFile string) []int {
	return toIntList(getListFromFile(inputFile))
}
