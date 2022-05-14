package leetcode

import (
    "fmt"
    "testing"
)

type question69 struct {
    para69
    ans69
}

type para69 struct {
    para int
}

type ans69 struct {
    result int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question69 {
        question69 {
            para69 {
                4,
            },
            ans69 {
                2,
            },
        },
        question69 {
            para69 {
                122,
            },
            ans69 {
                11,
            },
        },
        question69 {
            para69 {
                3,
            },
            ans69 {
                1,
            },
        },
        question69 {
            para69 {
                99,
            },
            ans69 {
                9,
            },
        },
        question69 {
            para69 {
                101,
            },
            ans69 {
                10,
            },
        },
        question69 {
            para69 {
                2147483647,
            },
            ans69 {
                46340,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", mySqrt(c.para))
	}
}



