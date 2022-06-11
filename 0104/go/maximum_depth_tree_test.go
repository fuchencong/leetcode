package leetcode

import (
    "fmt"
    "testing"
)

type question104 struct {
    para104
    ans104
}

type para104 struct {
    root *TreeNode
}

type ans104 struct {
    result int
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
    cases := []question104 {
        question104 {
            para104 {
                createTree1(),
            },
            ans104 {
                2,
            },
        },
        question104 {
            para104 {
                createTree2(),
            },
            ans104 {
                3,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para104)
		fmt.Printf("Expect: %v\n", c.ans104)
		fmt.Printf("Result: %v\n\n", maxDepth(c.root))
	}
}
