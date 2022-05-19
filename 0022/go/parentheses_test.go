package leetcode

import (
    "fmt"
    "testing"
)

type question22 struct {
    para22
    ans22
}

type para22 struct {
    n int
}

type ans22 struct {
    results []string
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question22 {
        question22 {
            para22 {
                3,
            },
            ans22 {
                []string{
                    "((()))","(()())","(())()","()(())","()()()",
                },
            },
        },
        question22 {
            para22 {
                1,
            },
            ans22 {
                []string{
                    "()",
                },
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para22)
		fmt.Printf("Expect: %v\n", c.results)
		fmt.Printf("Result: %v\n\n", generateParenthesis(c.n))
	}
}
