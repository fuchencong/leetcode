package leetcode

import (
    "fmt"
    "testing"
)

type question93 struct {
    para93
    ans93
}

type para93 struct {
    para string
}

type ans93 struct {
    results []string
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question93 {
        question93 {
            para93 {
                "25525511135",
            },
            ans93 {
                []string{"255.255.11.135","255.255.111.35" },
            },
        },
        question93 {
            para93 {
                "0000",
            },
            ans93 {
                []string{"0.0.0.0"},
            },
        },
        question93 {
            para93 {
                "11111",
            },
            ans93 {
                []string{"1.1.1.11", "1.1.11.1", "1.11.1.1", "11.1.1.1"},
            },
        },
        question93 {
            para93 {
                "101111",
            },
            ans93 {
                []string{"1.0.1.111", "1.0.11.11", "1.0.111.1", "10.1.1.11", "10.1.11.1", "10.11.1.1", "101.1.1.1"},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para)
		fmt.Printf("Expect: %v\n", c.results)
		fmt.Printf("Result: %v\n\n", restoreIpAddresses(c.para))
	}
}
