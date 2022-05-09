package leetcode

import (
    "fmt"
    "testing"
)

type question79 struct {
    para79
    ans79
}

type para79 struct {
    board [][]byte
    word string
}

type ans79 struct {
    result bool
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question79 {
        question79 {
            para79 {
                [][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")},
                "ABCCED",
            },
            ans79 {
                true,
            },
        },
        question79 {
            para79 {
                [][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")},
                "SEE",
            },
            ans79 {
                true,
            },
        },
        question79 {
            para79 {
                [][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")},
                "ABCB",
            },
            ans79 {
                false,
            },
        },
        question79 {
            para79 {
                [][]byte{[]byte("A"), []byte("S"), []byte("A")},
                "SA",
            },
            ans79 {
                true,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para79)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", exist(c.board, c.word))
	}
}
