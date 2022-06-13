package leetcode

import (
    "fmt"
    "testing"
)

type question108 struct {
    para108
    ans108
}

type para108 struct {
    nums []int
}

type ans108 struct {
    result []int
}

func (tree *TreeNode) preorder() []int {
    values := []int{}
    values = append(values, tree.Val)
    if tree.Left != nil {
        values = append(values, tree.Left.preorder()...)
    }
    if tree.Right != nil {
        values = append(values, tree.Right.preorder()...)
    }
    return values
}


func TestLongestPalindrome(t *testing.T) {
    cases := []question108 {
        question108 {
            para108 {
                []int{-10,-3,0,5,9},
            },
            ans108 {
                []int{-10,-3,0,5,9},
            },
        },
        question108 {
            para108 {
                []int{1,3},
            },
            ans108 {
                []int{1,3},
            },
        },
        question108 {
            para108 {
                []int{-1},
            },
            ans108 {
                []int{-1},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para108)
		fmt.Printf("Expect: %v\n", c.ans108)
		fmt.Printf("Result: %v\n\n", sortedArrayToBST(c.nums).preorder())
	}
}
