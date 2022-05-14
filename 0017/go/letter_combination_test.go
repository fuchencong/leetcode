package leetcode

import (
	"fmt"
	"testing"
)

type question17 struct {
	para17
	ans17
}

type para17 struct {
	digits string
}

type ans17 struct {
	result []string
}

func TestTwoSum(t *testing.T) {
	qs := []question17{
		{
            para17{
                "23",
            },
			ans17{
                []string{
                    "ad","ae","af","bd","be","bf","cd","ce","cf",
                },
            },
		},

		{
            para17{
                "",
            },
			ans17{
                []string{},
            },
		},

		{
            para17{
                "2",
            },
			ans17{
                []string{
                    "a","b","c",
                },
            },
		},

		{
            para17{
                "9",
            },
			ans17{
                []string{
                    "w","x","y","z",
                },
            },
		},

	}

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para17)
		fmt.Printf("Expect: %v\n", q.ans17)
		fmt.Printf("Result: %v\n\n", letterCombinations(q.para17.digits))
	}
}
