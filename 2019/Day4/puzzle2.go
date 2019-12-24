package main

import "fmt"

func runPuzzle2() {
	possPasswords := findPossiblePasswordsPuzzle2(206938, 679128)
	fmt.Printf("%v\n p2Amount: %d\n", possPasswords, len(possPasswords))
}

func findPossiblePasswordsPuzzle2(min, max int) []int {
	possPasswords := findPossiblePasswords(min, max)
	fmt.Printf("%v\n", possPasswords)
	var newPossPasswords []int
	for i := 0; i < len(possPasswords); i++ {
		password := possPasswords[i]
		digits := toList(password)
		if pairNotPartOfLargerGrouping(getGroupings(digits)) {
			newPossPasswords = append(newPossPasswords, password)
		}
	}
	fmt.Printf("%v\n", newPossPasswords)
	return newPossPasswords
}

func pairNotPartOfLargerGrouping(groupings []int) bool {
	fmt.Printf("groupings: %v\n", groupings)
	for i := 0; i < len(groupings); i++ {
		group := groupings[i]
		if group == 2 {
			return true
		}
	}
	return false
}

func getGroupings(nums []int) []int {
	var groupings []int
	rep := 1
	for i := 0; i <= 4; i++ {
		if nums[i] == nums[i+1] {
			rep += 1
		} else {
			groupings = append(groupings, rep)
			rep = 1
		}
	}
	groupings = append(groupings, rep)
	return groupings
}
