package leetcode

import (
    "fmt"
    "testing"
)

type question102 struct {
    para102
    ans102
}

type para102 struct {
    root *TreeNode
}

type ans102 struct {
    result [][]int
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

    left1 := &TreeNode{
        20,
        nil,
        nil,
    }

    right1 := &TreeNode{
        30,
        &TreeNode{40, nil, nil},
        &TreeNode{50, nil, nil},
    }

    root.Left = left1
    root.Right = right1
    return root
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question102 {
        question102 {
            para102 {
                createTree1(),
            },
            ans102 {
                [][]int{{10}, {20, 30}},
            },
        },
        question102 {
            para102 {
                createTree2(),
            },
            ans102 {
                [][]int{{10}, {20, 30}, {40, 50}},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para102)
		fmt.Printf("Expect: %v\n", c.ans102)
		fmt.Printf("Result: %v\n\n", levelOrder(c.root))
	}
}
