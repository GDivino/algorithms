package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	var num int = 1
	var i int = 0
	var arr_idx int = 0

	if arr[0]-1 >= k {
		return k
	}

	for {
		if arr_idx == len(arr) {
			return arr[len(arr)-1] + k - i
		}

		if k == i {
			return num - 1
		}

		if num == arr[arr_idx] {
			arr_idx++
		} else {
			i++
		}

		num++
	}
}

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
}
