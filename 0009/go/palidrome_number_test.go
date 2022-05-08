package leetcode

import (
    "fmt"
    "testing"
)

type question9 struct {
    para9
    ans9
}

type para9 struct {
    para int
}

type ans9 struct {
    result bool
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question9 {
        question9 {
            para9 {
                121,
            },
            ans9 {
                true,
            },
        },
        question9 {
            para9 {
                -121,
            },
            ans9 {
                false,
            },
        },
        question9 {
            para9 {
                10,
            },
            ans9 {
                false,
            },
        },
        question9 {
            para9 {
                123321,
            },
            ans9 {
                true,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", isPalindrome(c.para))
	}
}
