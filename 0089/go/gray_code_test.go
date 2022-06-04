package leetcode

import (
	"fmt"
	"testing"
)

type question89 struct {
	para89
	ans89
}

type para89 struct {
	para int
}

type ans89 struct {
	result []int
}

func TestLongestPalindrome(t *testing.T) {
	cases := []question89{
		question89{
			para89{
				2,
			},
			ans89{
				[]int{0,1,3,2},
			},
		},
		question89{
			para89{
				1,
			},
			ans89{
				[]int{0,1},
			},
		},
	}

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", grayCode(c.para))
	}
}
