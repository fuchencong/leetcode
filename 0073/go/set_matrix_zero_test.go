package leetcode

import (
    "fmt"
    "testing"
)

type question73 struct {
    para73
    ans73
}

type para73 struct {
    matrix [][]int
}

type ans73 struct {
    result [][]int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question73 {
        question73 {
            para73 {
                [][]int{{1,1,1},{1,0,1},{1,1,1}},
            },
            ans73 {
                [][]int{{1,0,1},{0,0,0},{1,0,1}},
            },
        },
        question73 {
            para73 {
                [][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}},
            },
            ans73 {
                [][]int{{0,0,0,0},{0,4,5,0},{0,3,1,0}},
            },
        },
        question73 {
            para73 {
                [][]int{{1,2},{3,4}},
            },
            ans73 {
                [][]int{{1,2},{3,4}},
            },
        },
        question73 {
            para73 {
                [][]int{{1,2,3,4},{5,0,7,8},{0,10,11,12},{13,14,15,0}},
            },
            ans73 {
                [][]int{{0,0,3,0},{0,0,0,0},{0,0,0,0},{0,0,0,0}},
            },
        },
        question73 {
            para73 {
                [][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}},
            },
            ans73 {
                [][]int{{0,0,0,0},{0,4,5,0},{0,3,1,0}},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para73)
		fmt.Printf("Expect: %v\n", c.result)
        setZeroes(c.matrix)
		fmt.Printf("Result: %v\n\n", c.matrix)
	}
}
