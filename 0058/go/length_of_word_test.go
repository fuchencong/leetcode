package leetcode

import (
    "fmt"
    "testing"
)

type question58 struct {
    para58
    ans58
}

type para58 struct {
    s string
}

type ans58 struct {
    result int
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question58 {
        question58 {
            para58 {
                "Hello World",
            },
            ans58 {
                5,
            },
        },
        question58 {
            para58 {
                "   fly me   to   the moon  ",
            },
            ans58 {
                4,
            },
        },
        question58 {
            para58 {
                "      a    ",
            },
            ans58 {
                1,
            },
        },
        question58 {
            para58 {
                "aabbb",
            },
            ans58 {
                5,
            },
        },
        question58 {
            para58 {
                "a bc       ",
            },
            ans58 {
                2,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para58)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", lengthOfLastWord(c.s))
	}
}
