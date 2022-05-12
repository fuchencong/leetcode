package leetcode

import (
	"fmt"
	"testing"
)

type question15 struct {
	para15
	ans15
}

type para15 struct {
	nums   []int
}

type ans15 struct {
	result [][]int
}

func TestThreeSum(t *testing.T) {
	qs := []question15{
		{
			para15{[]int{-1, 0, 1, 2, -1, -4}},
			ans15{[][]int{
                {-1, -1, 2},
                {-1, 0, 1},
            }},
        },

		{
			para15{[]int{}},
			ans15{[][]int{
                {},
            }},
        },

		{
			para15{[]int{0}},
			ans15{[][]int{
                {},
            }},
        },
    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para15)
		fmt.Printf("Expect: %v\n", q.ans15)
		fmt.Printf("Result: %v\n\n", threeSum(q.nums))
	}
}
