package leetcode

import (
    "fmt"
    "testing"
)

type question125 struct {
    para125
    ans125
}

type para125 struct {
    s string
}

type ans125 struct {
    result bool
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question125 {
        question125 {
            para125 {
                "A man, a plan, a canal: Panama",
            },
            ans125 {
                true,
            },
        },
        question125 {
            para125 {
                "race a car",
            },
            ans125 {
                false,
            },
        },
        question125 {
            para125 {
                " ",
            },
            ans125 {
                true,
            },
        },
        question125 {
            para125 {
                "a b a",
            },
            ans125 {
                true,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para125)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", isPalindrome(c.s))
	}
}
