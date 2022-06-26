package leetcode

import (
    "fmt"
    "testing"
)

type question129 struct {
    para129
    ans129
}

type para129 struct {
    root *TreeNode
}

type ans129 struct {
    result int
}

func createTree1() *TreeNode {
    root := &TreeNode{
        1,
        nil,
        nil,
    }

    left1 := &TreeNode{
        2,
        nil,
        nil,
    }

    right1 := &TreeNode{
        3,
        nil,
        nil,
    }

    root.Left = left1
    root.Right = right1
    return root
}

func createTree2() *TreeNode {
    root := &TreeNode{
        4,
        &TreeNode{0, nil, nil},
        nil,
    }

    right1 := &TreeNode{
        9,
        &TreeNode{5, nil, nil},
        &TreeNode{1, nil, nil},
    }

    root.Right = right1
    return root
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question129 {
        question129 {
            para129 {
                createTree1(),
            },
            ans129 {
                25,
            },
        },
        question129 {
            para129 {
                createTree2(),
            },
            ans129 {
                1026,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para129)
		fmt.Printf("Expect: %v\n", c.ans129)
		fmt.Printf("Result: %v\n\n", sumNumbers(c.root))
	}
}
