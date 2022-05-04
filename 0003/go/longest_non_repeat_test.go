package leetcode

import (
	"fmt"
	"testing"
)

type question3 struct {
	para3 string
	ans3  int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []question3{
		question3{
			"abcabcbb",
			3,
		},
		question3{
			"bbbbb",
			1,
		},
		question3{
			"pwwkew",
			3,
		},
	}
	for _, c := range cases {
		fmt.Printf("Input: %v", c.para3)
		fmt.Printf("Expect: %v\n", c.ans3)
		fmt.Printf("Result: %v\n\n", lengthOfLongestSubstring(c.para3))
	}

}
