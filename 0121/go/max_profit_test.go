package leetcode

import (
	"fmt"
	"testing"
)

type question121 struct {
	para121
	ans121
}

type para121 struct {
	prices   []int
}

type ans121 struct {
	result int
}

func TestTwoSum(t *testing.T) {
	qs := []question121{
		{
			para121{[]int{7,1,5,3,6,4}},
			ans121{5},
		},

		{
			para121{[]int{7,6,4,3,1}},
			ans121{0},
		},

		{
			para121{[]int{1,2,3,4,5}},
			ans121{4},
        },
		{
			para121{[]int{1}},
			ans121{0},
        },
    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para121)
		fmt.Printf("Expect: %v\n", q.ans121)
		fmt.Printf("Result: %v\n\n", maxProfit(q.para121.prices))
	}
}
