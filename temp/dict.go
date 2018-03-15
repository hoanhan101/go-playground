package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	tempDict := make(map[int]int)

	for index, value := range nums {
		if _, exist := tempDict[target-value]; exist == false {
			tempDict[value] = index
		} else {
			return []int{tempDict[target-value], index}
		}
	}

	// Return [0, 0] if cannot find any
	return []int{0, 0}
}
