package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	arr_map := make(map[int]int)
	count := 1

	for _, num := range arr {
		arr_map[num]++
	}

	for k > 0 {
		if _, check := arr_map[count]; !check {
			k--
		}

		count++
	}

	return count - 1
}

func main() {
	fmt.Println(findKthPositive([]int{1, 3, 4, 6}, 5))
}
