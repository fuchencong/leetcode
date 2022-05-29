package leetcode

import (
    "fmt"
    "testing"
)

type question75 struct {
    para75
    ans75
}

type para75 struct {
    para []int
}

type ans75 struct {
    result []int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question75 {
        question75 {
            para75 {
                []int{2,0,2,1,1,0},
            },
            ans75 {
                []int{0,0,1,1,2,2},
            },
        },
        question75 {
            para75 {
                []int{2,0,1},
            },
            ans75 {
                []int{0,1,2},
            },
        },
        question75 {
            para75 {
                []int{0,0,0,0,0},
            },
            ans75 {
                []int{0,0,0,0,0},
            },
        },
        question75 {
            para75 {
                []int{0,0,1,1,2,2},
            },
            ans75 {
                []int{0,0,1,1,2,2},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para)
		fmt.Printf("Expect: %v\n", c.result)
        sortColors(c.para)
		fmt.Printf("Result: %v\n\n", c.para)
	}
}
