package leetcode

import (
    "testing"
    "fmt"
)

type question53 struct {
    para53
    ans53
}

type para53 struct {
    nums []int
}

type ans53 struct {
    result int
}

func TestFindMedianSortedArrays(t *testing.T) {
    cases := []question53 {
        question53 {
            para53 {
                nums: []int{-2,1,-3,4,-1,2,1,-5,4},
            },
            ans53 {
                result:6,
            },
        },

        question53 {
            para53 {
                nums: []int{1},
            },
            ans53 {
                result:1,
            },
        },

        question53 {
            para53 {
                nums: []int{5,4,-1,7,8},
            },
            ans53 {
                23,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.nums)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", maxSubArray(c.nums))
	}
}
