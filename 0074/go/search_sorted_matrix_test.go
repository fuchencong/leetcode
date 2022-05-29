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
                [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},
                3,
            },
            ans240 {
                true,
            },
        },
        question240 {
            para240 {
                [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},
                13,
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
