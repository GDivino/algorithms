package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var nums_map = make(map[int]int)

	for idx1, num := range nums {
		if idx2, check := nums_map[target-num]; check {
			return []int{idx1, idx2}
		}

		nums_map[num]++
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 10}, 14))
}
