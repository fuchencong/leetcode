package leetcode

import (
    "fmt"
    "testing"
)

type question131 struct {
    para131
    ans131
}

type para131 struct {
    s string
}

type ans131 struct {
    results [][]string
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question131 {
        question131 {
            para131 {
                "aab",
            },
            ans131 {
                [][]string{{"a","a","b"}, {"aa","b"}},
            },
        },
        question131 {
            para131 {
                "a",
            },
            ans131 {
                [][]string{{"a"}},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para131)
		fmt.Printf("Expect: %v\n", c.results)
		fmt.Printf("Result: %v\n\n", partition(c.s))
	}
}

