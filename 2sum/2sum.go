package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var nums_map = make(map[int]int)

	for _, num := range nums {
		if _, check := nums_map[target-num]; check {
			return []int{num, target - num}
		}

		nums_map[num]++
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 10}, 14))
}
