package leetcode

import (
	"fmt"
	"testing"
)

type question6 struct {
	para6
	ans6
}

type para6 struct {
	s       string
	numRows int
}

type ans6 struct {
	result string
}

func TestLongestPalindrome(t *testing.T) {
	cases := []question6{
		question6{
			para6{
				"PAYPALISHIRING",
				3,
			},
			ans6{
				"PAHNAPLSIIGYIR",
			},
		},
		question6{
			para6{
				"PAYPALISHIRING",
				4,
			},
			ans6{
				"PINALSIGYAHRPI",
			},
		},
		question6{
			para6{
				"A",
				1,
			},
			ans6{
				"A",
			},
		},
		question6{
			para6{
				"AB",
				2,
			},
			ans6{
				"AB",
			},
		},
		question6{
			para6{
				"AB",
				1,
			},
			ans6{
				"AB",
			},
		},
	}

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para6)
		fmt.Printf("Expect: %v\n", c.ans6)
		fmt.Printf("Result: %v\n\n", convert(c.s, c.numRows))
	}
}
