package main

import "fmt"

func makeSmallestPalindrome(s string) string {
	var half int = len(s) / 2
	var byte_str = make([]byte, len(s))

	for i := 0; i <= half; i++ {
		var opposite byte = s[len(s)-1-i]

		if s[i] > opposite {
			byte_str[i] = opposite
			byte_str[len(s)-1-i] = opposite
		} else if i == half {
			byte_str[i] = s[i]
		} else {
			byte_str[i] = s[i]
			byte_str[len(s)-1-i] = s[i]
		}
	}

	return string(byte_str)
}

func main() {
	fmt.Println(makeSmallestPalindrome("test"))
}
