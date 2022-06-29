package leetcode

import (
    "fmt"
    "testing"
)

type question139 struct {
    para139
    ans139
}

type para139 struct {
    s string
    wordDict []string
}

type ans139 struct {
    results bool
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question139 {
        question139 {
            para139 {
                "leetcode",
                []string{"leet", "code"},
            },
            ans139 {
                true,
            },
        },
        question139 {
            para139 {
                "applepenapple",
                []string{"apple","pen"},
            },
            ans139 {
                true,
            },
        },
        question139 {
            para139 {
                "catsandog",
                []string{"cats","dog","sand","and","cat"},
            },
            ans139 {
                false,
            },
        },
        question139 {
            para139 {
                "abcd",
                []string{"abcd"},
            },
            ans139 {
                true,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para139)
		fmt.Printf("Expect: %v\n", c.results)
		fmt.Printf("Result: %v\n\n", wordBreak(c.s, c.wordDict))
	}
}

