package leetcode

import (
	"fmt"
	"testing"
)

type question122 struct {
	para122
	ans122
}

type para122 struct {
	prices   []int
}

type ans122 struct {
	result int
}

func TestTwoSum(t *testing.T) {
	qs := []question122{
		{
			para122{[]int{7,1,5,3,6,4}},
			ans122{7},
		},

		{
			para122{[]int{7,6,4,3,1}},
			ans122{0},
		},

		{
			para122{[]int{1,2,3,4,5}},
			ans122{4},
        },
		{
			para122{[]int{1}},
			ans122{0},
        },
    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para122)
		fmt.Printf("Expect: %v\n", q.ans122)
		fmt.Printf("Result: %v\n\n", maxProfit(q.para122.prices))
	}
}
