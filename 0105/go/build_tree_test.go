package leetcode

import (
    "fmt"
    "testing"
)

type question105 struct {
    para105
    ans105
}

type para105 struct {
    preorder []int
    inorder []int
}

type ans105 struct {
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
    cases := []question105 {
        question105 {
            para105 {
                []int{3,9,20,15,7},
                []int{9,3,15,20,7},
            },
            ans105 {
                []int{3,9,20,15,7},
            },
        },
        question105 {
            para105 {
                []int{-1},
                []int{-1},
            },
            ans105 {
                []int{-1},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para105)
		fmt.Printf("Expect: %v\n", c.ans105)
		fmt.Printf("Result: %v\n\n", buildTree(c.preorder, c.inorder).preorder())
	}
}
