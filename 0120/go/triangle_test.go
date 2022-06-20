package leetcode

import (
	"fmt"
	"testing"
)

type question118 struct {
	para118
	ans118
}

type para118 struct {
	triangle [][]int
}

type ans118 struct {
	result   int
}

func TestThreeSum(t *testing.T) {
	qs := []question118{
		{
			para118{[][]int{
                {10},
            }},
			ans118{10},
        },
		{
			para118{[][]int{
                {2}, 
                {3,4}, 
                {6,5,7}, 
                {4,1,8,3}, 
            }},
			ans118{11},
        },
		{
			para118{[][]int{
                {1}, 
                {1,4}, 
                {1,5,7}, 
                {1,0,8,3}, 
            }},
			ans118{3},
        },
    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para118)
		fmt.Printf("Expect: %v\n", q.ans118)
		fmt.Printf("Result: %v\n\n", minimumTotal(q.triangle))
	}
}
