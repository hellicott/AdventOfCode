package main

import "fmt"

// range 206938 to 679128

func main() {
	//runPuzzle1()
	runPuzzle2()
}

func runPuzzle1() {
	possPasswords := findPossiblePasswords(206938, 679128)
	fmt.Printf("%v\n Amount: %d\n", possPasswords, len(possPasswords))
}

func findPossiblePasswords(min, max int) []int {
	var possPasswords []int
	for num := min; num <= max; num++ {
		digits := toList(num)
		if neverDecreases(digits) && containsDouble(digits) {
			possPasswords = append(possPasswords, num)
		}
	}
	return possPasswords
}

func neverDecreases(digits []int) bool {
	for i := 0; i <= 4; i++ {
		if digits[i] < digits[i+1] {
			return false
		}
	}
	return true
}

func containsDouble(digits []int) bool {
	for i := 0; i <= 4; i++ {
		if digits[i] == digits[i+1] {
			return true
		}
	}
	return false
}

func toList(num int) []int {
	var digits []int
	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}
	return digits
}
