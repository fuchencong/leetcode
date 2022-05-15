package leetcode

import (
    "fmt"
    "testing"
)

type question50 struct {
    para50
    ans50
}

type para50 struct {
    x float64
    n int
}

type ans50 struct {
    result float64
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question50 {
        question50 {
            para50 {
                2.0,
                10,
            },
            ans50 {
                1024.0,
            },
        },
        question50 {
            para50 {
                2.1,
                3,
            },
            ans50 {
                9.261,
            },
        },
        question50 {
            para50 {
                2.0,
                -2,
            },
            ans50 {
                0.25,
            },
        },
        question50 {
            para50 {
                1,
                10,
            },
            ans50 {
                1,
            },
        },
        question50 {
            para50 {
                -1,
                9,
            },
            ans50 {
                -1,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para50)
		fmt.Printf("Expect: %v\n", c.ans50)
		fmt.Printf("Result: %v\n\n", myPow(c.x, c.n))
	}
}



