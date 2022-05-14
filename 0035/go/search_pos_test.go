package leetcode

import (
	"fmt"
	"testing"
)

type question35 struct {
	para35
	ans35
}

type para35 struct {
	nums   []int
	target int
}

type ans35 struct {
	result int
}

func TestTwoSum(t *testing.T) {
	qs := []question35{
		{
			para35{[]int{1, 3, 5, 6}, 5},
			ans35{2},
		},

		{
			para35{[]int{1, 3, 5, 6}, 2},
			ans35{1},
		},

		{
			para35{[]int{1, 3, 5, 6}, 7},
			ans35{4},
		},

		{
			para35{[]int{1}, 0},
			ans35{0},
		},

		{
			para35{[]int{1}, 2},
			ans35{1},
		},
	}

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para35)
		fmt.Printf("Expect: %v\n", q.ans35)
		fmt.Printf("Result: %v\n\n", searchInsert(q.para35.nums, q.para35.target))
	}
}
