package leetcode

import (
    "fmt"
    "testing"
)

type question67 struct {
    para67
    ans67
}

type para67 struct {
    a string
    b string
}

type ans67 struct {
    result string
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question67 {
        question67 {
            para67 {
                "11",
                "1",
            },
            ans67 {
                "100",
            },
        },
        question67 {
            para67 {
                "1010",
                "1011",
            },
            ans67 {
                "10101",
            },
        },
        question67 {
            para67 {
                "1",
                "1",
            },
            ans67 {
                "10",
            },
        },
        question67 {
            para67 {
                "0",
                "0",
            },
            ans67 {
                "0",
            },
        },
        question67 {
            para67 {
                "10",
                "110",
            },
            ans67 {
                "1000",
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para67)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", addBinary(c.a, c.b))
	}
}
