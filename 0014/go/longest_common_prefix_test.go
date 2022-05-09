package leetcode

import (
    "fmt"
    "testing"
)

type question14 struct {
    para14
    ans14
}

type para14 struct {
    para []string
}

type ans14 struct {
    result string
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question14 {
        question14 {
            para14 {
                []string{"flower","flow","flight",},
            },
            ans14 {
                "fl",
            },
        },
        question14 {
            para14 {
                []string{"flower","flower","flower",},
            },
            ans14 {
                "flower",
            },
        },
        question14 {
            para14 {
                []string{"flower","lower","ower",},
            },
            ans14 {
                "",
            },
        },
        question14 {
            para14 {
                []string{"","","",},
            },
            ans14 {
                "",
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", longestCommonPrefix(c.para))
	}
}
