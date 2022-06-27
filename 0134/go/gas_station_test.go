package leetcode

import (
    "fmt"
    "testing"
)

type question134 struct {
    para134
    ans134
}

type para134 struct {
    gas []int
    cost []int
}

type ans134 struct {
    result int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question134 {
        question134 {
            para134 {
                []int{1, 2, 3, 4, 5},
                []int{3, 4, 5, 1, 2},
            },
            ans134 {
                3,
            },
        },
        question134 {
            para134 {
                []int{2, 3, 4},
                []int{3, 4, 3},
            },
            ans134 {
                -1,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para134)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", canCompleteCircuit(c.gas, c.cost))
	}
}
