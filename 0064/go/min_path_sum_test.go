package leetcode

import (
    "fmt"
    "testing"
)

type question64 struct {
    para64
    ans64
}

type para64 struct {
    matrix [][]int
}

type ans64 struct {
    result int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question64 {
        question64 {
            para64 {
                [][]int{{1,3,1},{1,5,1},{4,2,1}},
            },
            ans64 {
                7,
            },
        },
        question64 {
            para64 {
                [][]int{{1,2,3},{4,5,6}},
            },
            ans64 {
                12,
            },
        },
        question64 {
            para64 {
                [][]int{{1,1},{1, 1}},
            },
            ans64 {
                3,
            },
        },
        question64 {
            para64 {
                [][]int{{1}},
            },
            ans64 {
                1,
            },
        },
        question64 {
            para64 {
                [][]int{{1,2}},
            },
            ans64 {
                3,
            },
        },
        question64 {
            para64 {
                [][]int{{1},{2}},
            },
            ans64 {
                3,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para64)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", minPathSum(c.matrix))
	}
}
