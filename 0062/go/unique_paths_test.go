package leetcode

import (
    "fmt"
    "testing"
)

type question62 struct {
    para62
    ans62
}

type para62 struct {
    m int
    n int
}

type ans62 struct {
    result int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question62 {
        question62 {
            para62 {
                3,
                7,
            },
            ans62 {
                28,
            },
        },
        question62 {
            para62 {
                3,
                2,
            },
            ans62 {
                3,
            },
        },
        question62 {
            para62 {
                2,
                2,
            },
            ans62 {
                2,
            },
        },
        question62 {
            para62 {
                1,
                2,
            },
            ans62 {
                1,
            },
        },
        question62 {
            para62 {
                5,
                1,
            },
            ans62 {
                1,
            },
        },
        question62 {
            para62 {
                1,
                1,
            },
            ans62 {
                1,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para62)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", uniquePaths(c.m, c.n))
	}
}



