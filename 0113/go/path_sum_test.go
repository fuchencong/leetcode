package leetcode

import (
    "fmt"
    "testing"
)

type question113 struct {
    para113
    ans113
}

type para113 struct {
    root *TreeNode
    target int
}

type ans113 struct {
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
        &TreeNode{40, nil, nil},
    }

    root.Left = left1
    root.Right = right1
    return root
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question113 {
        question113 {
            para113 {
                createTree1(),
                30,
            },
            ans113 {
                [][]int{{10, 20}},
            },
        },
        question113 {
            para113 {
                createTree1(),
                50,
            },
            ans113 {
                [][]int{{}},
            },
        },
        question113 {
            para113 {
                createTree2(),
                80,
            },
            ans113 {
                [][]int{{10, 30, 40}, {10, 30, 40}},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para113)
		fmt.Printf("Expect: %v\n", c.ans113)
		fmt.Printf("Result: %v\n\n", pathSum(c.root, c.target))
	}
}
