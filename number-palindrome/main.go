package main

import "fmt"

func isPalindrome(x int) bool {
	var reversed_x int = 0
	var num = x

	for num > 0 {
		reversed_x = reversed_x*10 + num%10
		num /= 10
	}

	return reversed_x == x
}

func main() {
	fmt.Println(isPalindrome(121121))
}
