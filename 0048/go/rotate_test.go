package leetcode

import (
    "fmt"
    "testing"
)

type question48 struct {
    para48
    ans48
}

type para48 struct {
    matrix [][]int
}

type ans48 struct {
    result [][]int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question48 {
        question48 {
            para48 {
                [][]int{{1,2,3},{4,5,6},{7,8,9}},
            },
            ans48 {
                [][]int{{7,4,1},{8,5,2},{9,6,3}},
            },
        },
        question48 {
            para48 {
                [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}},
            },
            ans48 {
                [][]int{{15,13,2,5},{14,3,4,1},{12,6,8,9},{16,7,10,11}},
            },
        },
        question48 {
            para48 {
                [][]int{{1,2},{3, 4}},
            },
            ans48 {
                [][]int{{3,1},{4, 2}},
            },
        },
        question48 {
            para48 {
                [][]int{{1}},
            },
            ans48 {
                [][]int{{1}},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para48)
		fmt.Printf("Expect: %v\n", c.result)
        rotate(c.matrix)
		fmt.Printf("Result: %v\n\n", c.matrix)
	}
}
