package leetcode

import (
	"fmt"
	"testing"
)

type question509 struct {
	para509
	ans509
}

type para509 struct {
	para int
}

type ans509 struct {
	result int
}

func TestLongestPalindrome(t *testing.T) {
	cases := []question509{
		question509{
			para509{
				2,
			},
			ans509{
				1,
			},
		},
		question509{
			para509{
				3,
			},
			ans509{
				2,
			},
		},
		question509{
			para509{
				4,
			},
			ans509{
				3,
			},
		},
	}

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", fib(c.para))
	}
}
