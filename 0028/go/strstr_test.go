package leetcode

import (
    "fmt"
    "testing"
)

type question28 struct {
    para28
    ans28
}

type para28 struct {
    haystack string
    needle string
}

type ans28 struct {
    result int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question28 {
        question28 {
            para28 {
                "hello",
                "ll",
            },
            ans28 {
                2,
            },
        },
        question28 {
            para28 {
                "aaaaa",
                "bba",
            },
            ans28 {
                -1,
            },
        },
        question28 {
            para28 {
                "aaaaa",
                "aaa",
            },
            ans28 {
                0,
            },
        },
        question28 {
            para28 {
                "aabbb",
                "aabbb",
            },
            ans28 {
                0,
            },
        },
        question28 {
            para28 {
                "a",
                "a",
            },
            ans28 {
                0,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para28)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", strStr(c.haystack, c.needle))
	}
}
