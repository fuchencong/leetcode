
package leetcode

import (
	"fmt"
	"testing"
)

type question16 struct {
	para16
	ans16
}

type para16 struct {
	nums   []int
    target int
}

type ans16 struct {
    result int
}

func TestThreeSum(t *testing.T) {
	qs := []question16{
		{
			para16{
                []int{-1, 2, 1, -4},
                1,

            },
			ans16{
                2,
            },
        },

		{
			para16{
                []int{0, 0, 0},
                1,

            },
			ans16{
                0,
            },
        },

		{
			para16{
                []int{-1, 2, 1, -4},
                -4,

            },
			ans16{
                -4,
            },
        },
    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para16)
		fmt.Printf("Expect: %v\n", q.ans16)
		fmt.Printf("Result: %v\n\n", threeSumClosest(q.nums, q.target))
	}
}
