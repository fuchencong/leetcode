package leetcode

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

type para1 struct {
	nums   []int
	target int
}

type ans1 struct {
	result [][]int
}

func TestTwoSum(t *testing.T) {
	qs := []question1{
		{
			para1{[]int{2,3,6,7}, 7},
			ans1{[][]int{{2,2,3},{7}}},
		},

		{
			para1{[]int{2,3,5}, 8},
			ans1{[][]int{{2,2,2,2},{2,3,3},{3,5}}},
		},
		{
			para1{[]int{1}, 5},
            ans1{[][]int{{1,1,1,1,1}}},
		},
	}

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para1)
		fmt.Printf("Expect: %v\n", q.ans1)
		fmt.Printf("Result: %v\n\n", combinationSum(q.para1.nums, q.para1.target))
	}
}
