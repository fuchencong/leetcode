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
	result []int
}

func TestTwoSum(t *testing.T) {
	qs := []question1{
		{
			para1{[]int{2, 7, 11, 5}, 9},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{3, 2, 4}, 6},
			ans1{[]int{1, 2}},
		},

		{
			para1{[]int{3, 3}, 6},
			ans1{[]int{0, 1}},
		},
	}

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para1)
		fmt.Printf("Expect: %v\n", q.ans1)
		fmt.Printf("Result: %v\n\n", twoSum(q.para1.nums, q.para1.target))
	}
}
