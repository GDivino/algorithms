package main

import (
	"fmt"
	"strings"
)

func main() {
	test1 := "100001"
	test2 := "10001100"
	test3 := "1"
	test4 := "10"

	fmt.Println(checkOnesSegment(test1))
	fmt.Println(checkOnesSegment(test2))
	fmt.Println(checkOnesSegment(test3))
	fmt.Println(checkOnesSegment(test4))
}

func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}

func solutionTwo(s string) bool {
	var has_segment bool = false
	var no_segment bool = false
	var count int = -1

	if rune(s[0]) == '1' {
		count++
		has_segment = !has_segment
	}

	for i := 1; i < len(s); i++ {
		if rune(s[count]) == '0' && has_segment {
			no_segment = true
		} else if rune(s[count]) == '1' && !no_segment{
			has_segment = true
		} else if rune(s[count]) == '1' && no_segment {
            return false
        }

        count++
	}

    return has_segment
}
