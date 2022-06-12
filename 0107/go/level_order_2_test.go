package leetcode

import (
    "fmt"
    "testing"
)

type question107 struct {
    para107
    ans107
}

type para107 struct {
    root *TreeNode
}

type ans107 struct {
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
    cases := []question107 {
        question107 {
            para107 {
                createTree1(),
            },
            ans107 {
                [][]int{{20, 30}, {10}},
            },
        },
        question107 {
            para107 {
                createTree2(),
            },
            ans107 {
                [][]int{{40, 50}, {20, 30}, {10}},
            },
        },
        question107 {
            para107 {
                nil,
            },
            ans107 {
                [][]int{},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para107)
		fmt.Printf("Expect: %v\n", c.ans107)
		fmt.Printf("Result: %v\n\n", levelOrderBottom(c.root))
	}
}
