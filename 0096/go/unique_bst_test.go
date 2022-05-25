package leetcode

import (
    "fmt"
    "testing"
)

type question96 struct {
    para96
    ans96
}

type para96 struct {
    n int
}

type ans96 struct {
    result int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question96 {
        question96 {
            para96 {
                1,
            },
            ans96 {
                1,
            },
        },
        question96 {
            para96 {
                2,
            },
            ans96 {
                2,
            },
        },
        question96 {
            para96 {
                3,
            },
            ans96 {
                5,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para96)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", numTrees(c.n))
	}
}
