package leetcode

import (
    "fmt"
    "testing"
)

type question154 struct {
    para154
    ans154
}

type para154 struct {
    para []int
}

type ans154 struct {
    result int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question154 {
        question154 {
            para154 {
                []int{1, 3, 5},
            },
            ans154 {
                1,
            },
        },
        question154 {
            para154 {
                []int{2,2,2,0,1},
            },
            ans154 {
                0,
            },
        },
        question154 {
            para154 {
                []int{5,5,5,5,5},
            },
            ans154 {
                5,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", findMin(c.para))
	}
}
