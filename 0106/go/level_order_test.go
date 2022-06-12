package leetcode

import (
    "fmt"
    "testing"
)

type question106 struct {
    para106
    ans106
}

type para106 struct {
    inorder []int
    postorder []int
}

type ans106 struct {
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
    cases := []question106 {
        question106 {
            para106 {
                []int{9,3,15,20,7},
                []int{9,15,7,20,3},
            },
            ans106 {
                []int{3,9,20,15,7},
            },
        },
        question106 {
            para106 {
                []int{-1},
                []int{-1},
            },
            ans106 {
                []int{-1},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para106)
		fmt.Printf("Expect: %v\n", c.ans106)
		fmt.Printf("Result: %v\n\n", buildTree(c.inorder, c.postorder).preorder())
	}
}
