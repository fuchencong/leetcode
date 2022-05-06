package leetcode

import (
    "fmt"
    "testing"
)

type question240 struct {
    para240
    ans240
}

type para240 struct {
    matrix [][]int
    target int
}

type ans240 struct {
    result bool
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question240 {
        question240 {
            para240 {
                [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}},
                5,
            },
            ans240 {
                true,
            },
        },
        question240 {
            para240 {
                [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}},
                20,
            },
            ans240 {
                false,
            },
        },
        question240 {
            para240 {
                [][]int{{1,2},{3, 4}},
                2,
            },
            ans240 {
                true,
            },
        },
        question240 {
            para240 {
                [][]int{{1}},
                2,
            },
            ans240 {
                false,
            },
        },
        question240 {
            para240 {
                [][]int{{1, 2, 3, 4, 5}},
                5,
            },
            ans240 {
                true,
            },
        },
        question240 {
            para240 {
                [][]int{{1}, {2}, {3}, {4}, {5}},
                6,
            },
            ans240 {
                false,
            },
        },
        question240 {
            para240 {
                [][]int{{-5}},
                -5,
            },
            ans240 {
                true,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para240)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", searchMatrix(c.matrix, c.target))
	}
}
