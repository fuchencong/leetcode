package leetcode

import (
	"fmt"
	"testing"
)

type question27 struct {
	para27
	ans27
}

type para27 struct {
	nums   []int
    val    int
}

type ans27 struct {
	result int
}

func TestThreeSum(t *testing.T) {
	qs := []question27{
		{
			para27{
                []int{3,2,2,3},
                3,
            },
			ans27{2},
        },

		{
			para27{
                []int{0,1,2,2,3,0,4,2},
                2,
            },
			ans27{5},
        },

		{
			para27{
                []int{0,0,0,0,0},
                0,
            },
			ans27{0},
        },

		{
			para27{
                []int{0,0,0,0,0},
                5,
            },
			ans27{5},
        },


    }

	for _, q := range qs {
		fmt.Printf("Input: %v\n", q.para27)
		fmt.Printf("Expect: %v\n", q.ans27)
		fmt.Printf("Result: %v\n", removeElement(q.nums, q.val))
		fmt.Printf("Result: %v\n", q.nums)
	}
}
