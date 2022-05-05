package leetcode

import (
    "testing"
    "fmt"
)

type question4 struct {
    para4
    ans4
}

type para4 struct {
    nums1 []int
    nums2 []int
}

type ans4 struct {
    result float64
}

func TestFindMedianSortedArrays(t *testing.T) {
    cases := []question4 {
        question4 {
            para4 {
                nums1: []int{1, 3},
                nums2: []int{2},
            },
            ans4 {
                result:2.0,
            },
        },

        question4 {
            para4 {
                nums1: []int{1, 2},
                nums2: []int{3, 4},
            },
            ans4 {
                result:2.5,
            },
        },

        question4 {
            para4 {
                nums1: []int{1},
                nums2: []int{1},
            },
            ans4 {
                result:1.0,
            },
        },

        question4 {
            para4 {
                nums1: []int{2},
                nums2: []int{1,3,4},
            },
            ans4 {
                result:2.5,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v %v\n", c.nums1, c.nums2)
		fmt.Printf("Expect: %v\n", c.result)
		fmt.Printf("Result: %v\n\n", findMedianSortedArrays(c.nums1, c.nums2))
	}
}
