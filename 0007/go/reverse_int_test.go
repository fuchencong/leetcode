package leetcode

import (
    "fmt"
    "testing"
)

type question7 struct {
    para7
    ans7
}

type para7 struct {
    para int
}

type ans7 struct {
    result int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question7 {
        question7 {
            para7 {
                123,
            },
            ans7 {
                321,
            },
        },
        question7 {
            para7 {
                -123,
            },
            ans7 {
                -321,
            },
        },
        question7 {
            para7 {
                120,
            },
            ans7 {
                21,
            },
        },
        question7 {
            para7 {
                0,
            },
            ans7 {
                0,
            },
        },
        question7 {
            para7 {
                2147483647,
            },
            ans7 {
                0,
            },
        },
        question7 {
            para7 {
                -2147483648,
            },
            ans7 {
                0,
            },
        },
        question7 {
            para7 {
                74847,
            },
            ans7 {
                74847,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", reverse(c.para))
	}
}



