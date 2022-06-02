package leetcode

import (
    "fmt"
    "testing"
)

type question91 struct {
    para91
    ans91
}

type para91 struct {
    s string
}

type ans91 struct {
    result int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question91 {
        question91 {
            para91 {
                "12",
            },
            ans91 {
                2,
            },
        },
        question91 {
            para91 {
                "226",
            },
            ans91 {
                3,
            },
        },
        question91 {
            para91 {
                "06",
            },
            ans91 {
                0,
            },
        },
        question91 {
            para91 {
                "101",
            },
            ans91 {
                1,
            },
        },
        question91 {
            para91 {
                "26",
            },
            ans91 {
                2,
            },
        },
        question91 {
            para91 {
                "111111111111111111111111111111111111111111111",
            },
            ans91 {
                1836311903,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para91)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", numDecodingsDP(c.s))
	}
}
