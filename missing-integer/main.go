package main

import (
	"fmt"
)

func main() {
    var test1 = []int{1, 2, 3, 4, 5}
	var test2 = []int{5, 3, 15, 51, 123, 2, 4, 1, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
    var test3 = []int{-10, -3, -4, -5, -7, 23}
    var test4 = []int{1, 2, 5, 4, 5}

    fmt.Println(solution(test1))
	fmt.Println(solution(test2))
    fmt.Println(solution(test3))
    fmt.Println(solution(test4))
}

func solution(nums []int) int {
	var start int = 1
    var cont = true
    var num_map = make(map[int]int)

    for _, j := range nums {
        num_map[j] = 1
    }

    for cont {
        if num_map[start] == 0 {
            return start
        }
        start += 1
    }
    return 0
}

/*
Problem: find the smallest positive integer not in the list

Cases:
- number is within the list
- number is smaller than the min in the list
- number is bigger than the min in the list
- negative numbers

Possible approaches:
1. Start from number 1 and then go up
2. Sort the slice and start from the smallest number
*/
