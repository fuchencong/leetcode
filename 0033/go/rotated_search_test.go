package leetcode

import (
    "testing"
    "fmt"
)

type question33 struct {
    para33
    ans33
}

type para33 struct {
    nums []int
    target int
}

type ans33 struct {
    result int
}

func TestFindMedianSortedArrays(t *testing.T) {
    cases := []question33 {
        question33 {
            para33 {
                []int{4, 5, 6, 7, 0, 1, 2},
                0,
            },
            ans33 {
                4,
            },
        },

        question33 {
            para33 {
                []int{4, 5, 6, 7, 0, 1, 2},
                3,
            },
            ans33 {
                -1,
            },
        },

        question33 {
            para33 {
                []int{1},
                0,
            },
            ans33 {
                -1,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v %v\n", c.nums, c.target)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", search(c.nums, c.target))
	}
}
