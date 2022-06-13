package leetcode

import (
    "fmt"
    "testing"
)

type question110 struct {
    para110
    ans110
}

type para110 struct {
    root *TreeNode
}

type ans110 struct {
    result bool
}

func createTree1() *TreeNode {
    root := &TreeNode{
        10,
        nil,
        nil,
    }

    left1 := &TreeNode{
        20,
        nil,
        nil,
    }

    right1 := &TreeNode{
        30,
        nil,
        nil,
    }

    root.Left = left1
    root.Right = right1
    return root
}

func createTree2() *TreeNode {
    root := &TreeNode{
        10,
        nil,
        nil,
    }

    right1 := &TreeNode{
        30,
        &TreeNode{40, nil, nil},
        &TreeNode{50, nil, nil},
    }

    root.Left = nil
    root.Right = right1
    return root
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question110 {
        question110 {
            para110 {
                createTree1(),
            },
            ans110 {
                true,
            },
        },
        question110 {
            para110 {
                createTree2(),
            },
            ans110 {
                false,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para110)
		fmt.Printf("Expect: %v\n", c.ans110)
		fmt.Printf("Result: %v\n\n", isBalanced(c.root))
	}
}
