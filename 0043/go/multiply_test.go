package leetcode

import (
	"fmt"
	"testing"
)

type question49 struct {
	para49
	ans49
}

type para49 struct {
    num1 string
    num2 string
}

type ans49 struct {
    result string
}

func TestMultiply(t *testing.T) {
	cases := []question49{
		question49{
			para49{
                "2",
                "3",
			},
            ans49{
                "6",
			},
		},
		question49{
			para49{
                "111",
                "0",
			},
            ans49{
                "0",
			},
		},

		question49{
			para49{
                "12",
                "12",
			},
            ans49{
                "144",
			},
		},

		question49{
			para49{
                "100",
                "10000",
			},
            ans49{
                "1000000",
			},
		},

		question49{
			para49{
                "0",
                "0",
			},
            ans49{
                "0",
			},
		},

		question49{
			para49{
                "123",
                "456",
			},
            ans49{
                "56088",
			},
		},
	}

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para49)
		fmt.Printf("Expect: %v\n", c.ans49)
		fmt.Printf("Result: %v\n\n", multiply(c.num1, c.num2))
	}
}
