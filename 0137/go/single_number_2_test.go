package leetcode

import (
    "testing"
    "fmt"
)

type question136 struct {
    para136
    ans136
}

type para136 struct {
    nums []int
}

type ans136 struct {
    result int
}

func TestFindMedianSortedArrays(t *testing.T) {
    cases := []question136 {
        question136 {
            para136 {
                []int{1, 1, 1, 3},
            },
            ans136 {
                result:3,
            },
        },

        question136 {
            para136 {
                nums: []int{2, 2, 2, 1},
            },
            ans136 {
                result:1,
            },
        },
        question136 {
            para136 {
                nums: []int{-2,-2,1,1,4,1,4,4,-4,-2},
            },
            ans136 {
                result:-4,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.nums)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", singleNumber(c.nums))
	}
}
