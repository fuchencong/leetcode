package leetcode

import (
    "fmt"
    "testing"
)

type question63 struct {
    para63
    ans63
}

type para63 struct {
    matrix [][]int
}

type ans63 struct {
    result int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question63 {
        question63 {
            para63 {
                [][]int{{0,0,0},{0,1,0},{0,0,0}},
            },
            ans63 {
                2,
            },
        },
        question63 {
            para63 {
                [][]int{{0,1},{0,0}},
            },
            ans63 {
                1,
            },
        },
        question63 {
            para63 {
                [][]int{{0,1},{1, 0}},
            },
            ans63 {
                0,
            },
        },
        question63 {
            para63 {
                [][]int{{0,0}},
            },
            ans63 {
                1,
            },
        },
        question63 {
            para63 {
                [][]int{{0},{0}},
            },
            ans63 {
                1,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para63)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", uniquePathsWithObstacles(c.matrix))
	}
}
