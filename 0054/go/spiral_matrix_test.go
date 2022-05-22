package leetcode

import (
    "fmt"
    "testing"
)

type question54 struct {
    para54
    ans54
}

type para54 struct {
    matrix [][]int
}

type ans54 struct {
    result []int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question54 {
        question54 {
            para54 {
                [][]int{{1,2,3},{4,5,6},{7,8,9}},
            },
            ans54 {
                []int{1,2,3,6,9,8,7,4,5},
            },
        },
        question54 {
            para54 {
                [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}},
            },
            ans54 {
                []int{1,2,3,4,8,12,11,10,9,5,6,7},
            },
        },
        question54 {
            para54 {
                [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12},{13,14,15,16}},
            },
            ans54 {
                []int{1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10},
            },
        },
        question54 {
            para54 {
                [][]int{{1}},
            },
            ans54 {
                []int{1},
            },
        },
        question54 {
            para54 {
                [][]int{{1, 2}},
            },
            ans54 {
                []int{1,2},
            },
        },
        question54 {
            para54 {
                [][]int{{1},{2}},
            },
            ans54 {
                []int{1,2},
            },
        },
        question54 {
            para54 {
                [][]int{{1, 2}, {3, 4}},
            },
            ans54 {
                []int{1,2,4,3},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para54)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", spiralOrder(c.matrix))
	}
}
