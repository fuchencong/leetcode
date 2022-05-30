package leetcode

import (
    "testing"
    "fmt"
)

type question88 struct {
    para88
    ans88
}

type para88 struct {
    nums1 []int
    m int
    nums2 []int
    n int
}

type ans88 struct {
    result []int
}

func TestFindMedianSortedArrays(t *testing.T) {
    cases := []question88 {
        question88 {
            para88 {
                nums1: []int{1,2,3,0,0,0},
                m:3,
                nums2: []int{2,5,6},
                n:3,
            },
            ans88 {
               []int{1,2,2,3,5,6},
            },
        },

        question88 {
            para88 {
                nums1: []int{1},
                m:1,
                nums2: []int{},
                n:0,
            },
            ans88 {
               []int{1},
            },
        },

        question88 {
            para88 {
                nums1: []int{1,0,0,0,0,0},
                m:1,
                nums2: []int{-1,0,1,2,3},
                n:5,
            },
            ans88 {
               []int{-1,0,1,1,2,3},
            },
        },

    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para88)
		fmt.Printf("Expect: %v\n", c.result)
        merge(c.nums1, c.m, c.nums2, c.n)
		fmt.Printf("Result: %v\n\n", c.nums1)
	}
}
