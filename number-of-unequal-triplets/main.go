package main

import "fmt"

func unequalTriplets(nums []int) int {
	nums_freq_map := make(map[int]int)
	left, count, right := 0, 0, len(nums)

	for _, num := range nums {
		nums_freq_map[num]++
	}

	for _, freq := range nums_freq_map {
		right -= freq
		count += left * freq * right
		left += freq
	}

	return count
}

func main() {
	fmt.Println(unequalTriplets([]int{1, 2, 3, 4, 5, 6}))
}
