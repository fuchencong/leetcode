package leetcode

import (
	"fmt"
	"testing"
)

type question70 struct {
	para70
	ans70
}

type para70 struct {
	n   int
}

type ans70 struct {
	result int
}

func TestTwoSum(t *testing.T) {
	qs := []question70{
		{
			para70{2},
			ans70{2},
		},
		{
			para70{3},
			ans70{3},
		},
		{
			para70{1},
			ans70{1},
		},
	}

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para70)
		fmt.Printf("Expect: %v\n", q.ans70)
		fmt.Printf("Result: %v\n\n", climbStairs(q.para70.n))
	}
}
