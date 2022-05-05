package leetcode

import (
    "fmt"
    "testing"
)

type question5 struct {
    para5
    ans5
}

type para5 struct {
    para string
}

type ans5 struct {
    result string
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question5 {
        question5 {
            para5 {
                "babad",
            },
            ans5 {
                "bab",
            },
        },
        question5 {
            para5 {
                "cbbd",
            },
            ans5 {
                "bb",
            },
        },
        question5 {
            para5 {
                "ddcbbcdd",
            },
            ans5 {
                "ddcbbcdd",
            },
        },
        question5 {
            para5 {
                "x",
            },
            ans5 {
                "x",
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", longestPalindrome(c.para))
	}
}
