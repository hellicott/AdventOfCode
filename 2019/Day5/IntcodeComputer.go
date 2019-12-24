package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type IntcodeComputer struct {
	pointer      int
	MAXPOINTER   int
	program      []int
	parameterMap map[int]int
}

func newIntcodeComputer(inputFile string) *IntcodeComputer {
	pointer := 0
	prog := retrieveProgram(inputFile)
	paramMap := make(map[int]int)
	paramMap[01] = 3
	paramMap[02] = 3
	paramMap[03] = 1
	paramMap[04] = 1
	computer := IntcodeComputer{
		pointer:      pointer,
		MAXPOINTER:   len(prog),
		program:      prog,
		parameterMap: paramMap,
	}
	return &computer
}

func (comp IntcodeComputer) calculateNewState() []int {
	for comp.pointer < comp.MAXPOINTER {
		instruction := comp.program[comp.pointer]
		if instruction == 99 {
			break
		}
		comp.handleInstruction(instruction)
	}
	return comp.program
}

func (comp IntcodeComputer) handleInstruction(instruction int) {
	opcode := getOpcode(instruction)
	parameters := comp.getParams(opcode)
	parameterModes := fillEmptyParamModes(len(parameters), getParamModes(instruction))
	if opcode == 01 {
		comp.add(parameters, parameterModes)
	} else if opcode == 02 {
		comp.multiply(parameters, parameterModes)
	} else if opcode == 03 {
		comp.save(parameters, parameterModes)
	}
}

func (comp IntcodeComputer) add(params, paramModes []int) {
	num1 := comp.getValue(params[0], paramModes[0])
	num2 := comp.getValue(params[1], paramModes[1])
	output := params[2]
	comp.program[output] = num1 + num2
}

func (comp IntcodeComputer) multiply(params, paramModes []int) {
	num1 := comp.getValue(params[0], paramModes[0])
	num2 := comp.getValue(params[1], paramModes[1])
	output := params[2]
	comp.program[output] = num1 * num2
}

func (comp IntcodeComputer) save(params, paramModes []int) {
	input := comp.getValue(params[0], paramModes[0])
}

func (comp IntcodeComputer) get(location int) int {
	return comp.program[location]
}

func (comp IntcodeComputer) getValue(param, paramMode int) int {
	if paramMode == 0 {
		return comp.program[param]
	}
	return param
}

func (comp IntcodeComputer) getParams(opcode int) []int {
	numOfParams := comp.parameterMap[opcode]
	start := comp.pointer + 1
	end := start + numOfParams
	return comp.program[start:end]
}

// static methods

func fillEmptyParamModes(expectedLen int, paramModes []int) []int {
	diff := expectedLen - len(paramModes)
	if diff > 1 {
		for i := 0; i < diff; i++ {
			paramModes = append(paramModes, 0)
		}
	}
	return paramModes
}

func getParamModes(instruction int) []int {
	digits := toReverseDigits(instruction)
	return digits[2:]
}

func getOpcode(instruction int) int {
	digits := toReverseDigits(instruction)
	return digits[0] + digits[1]*10
}

func toReverseDigits(num int) []int {
	var digits []int
	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}
	return digits
}

func toDigits(num int) []int {
	return reverse(toReverseDigits(num))
}

func reverse(oldList []int) []int {
	var newList []int
	for i := len(oldList) - 1; i >= 0; i-- {
		newList = append(newList, oldList[i])
	}
	return newList
}

func retrieveProgram(inputFile string) []int {
	return toIntList(getListFromFile(inputFile))
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
