package leetcode

import (
    "fmt"
    "testing"
)

type question103 struct {
    para103
    ans103
}

type para103 struct {
    root *TreeNode
}

type ans103 struct {
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

func TestZigzagLevelTree(t *testing.T) {
    cases := []question103 {
        question103 {
            para103 {
                createTree1(),
            },
            ans103 {
                [][]int{{10}, {30, 20}},
            },
        },
        question103 {
            para103 {
                createTree2(),
            },
            ans103 {
                [][]int{{10}, {30,20}, {40, 50}},
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para103)
		fmt.Printf("Expect: %v\n", c.ans103)
		fmt.Printf("Result: %v\n\n", zigzagLevelOrder(c.root))
	}
}
