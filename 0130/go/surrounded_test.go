package leetcode

import (
    "fmt"
    "testing"
)

type question130 struct {
    para130
}

type para130 struct {
    matrix [][]byte
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question130 {
        question130 {
            para130 {
                [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'},{'X','O','X','X'}},
            },
        },
        question130 {
            para130 {
                [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'}},
            },
        },
        question130 {
            para130 {
                [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','O','O','X'}},
            },
        },
        question130 {
            para130 {
                [][]byte{{'O','X','X','O','X'},{'X','O','O','X','O'},{'X','O','X','O','X'}, {'O','X','O','O','O'},{'X','X','O','X','O'}},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %c\n", c.matrix)
		solve(c.matrix)
		fmt.Printf("Input: %c\n", c.matrix)
	}
}
