package leetcode

import (
	"fmt"
	"testing"
)

type question38 struct {
	para38
	ans38
}

type para38 struct {
    n int
}

type ans38 struct {
	result string
}

func TestCountAndSay(t *testing.T) {
	cases := []question38{
		question38{
			para38{
                1,
			},
            ans38{
                "1",
			},
		},
		question38{
			para38{
                4,
			},
            ans38{
                "1211",
			},
		},
		question38{
			para38{
                5,
			},
            ans38{
                "111221",
			},
		},
		question38{
			para38{
                6,
			},
            ans38{
                "312211",
			},
		},
	}

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para38)
		fmt.Printf("Expect: %v\n", c.ans38)
		fmt.Printf("Result: %v\n\n", countAndSay(c.n))
	}
}
