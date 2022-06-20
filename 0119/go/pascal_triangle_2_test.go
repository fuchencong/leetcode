package leetcode

import (
	"fmt"
	"testing"
)

type question119 struct {
	para119
	ans119
}

type para119 struct {
	nums   int
}

type ans119 struct {
	result []int
}

func TestThreeSum(t *testing.T) {
	qs := []question119{
		{
			para119{0},
			ans119{[]int{
                1,
            }},
        },
		{
			para119{1},
			ans119{[]int{
                1,1,
            }},
        },
		{
			para119{4},
			ans119{[]int{
                1,4,6,4,1,
            }},
        },
    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para119)
		fmt.Printf("Expect: %v\n", q.ans119)
		fmt.Printf("Result: %v\n\n", getRow(q.nums))
	}
}
