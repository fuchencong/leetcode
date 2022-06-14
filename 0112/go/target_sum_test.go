package leetcode

import (
    "fmt"
    "testing"
)

type question112 struct {
    para112
    ans112
}

type para112 struct {
    root *TreeNode
    target int
}

type ans112 struct {
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
    cases := []question112 {
        question112 {
            para112 {
                createTree1(),
                30,
            },
            ans112 {
                true,
            },
        },
        question112 {
            para112 {
                createTree1(),
                50,
            },
            ans112 {
                false,
            },
        },
        question112 {
            para112 {
                createTree2(),
                40,
            },
            ans112 {
                false,
            },
        },
        question112 {
            para112 {
                createTree2(),
                100,
            },
            ans112 {
                false,
            },
        },
        question112 {
            para112 {
                createTree2(),
                80,
            },
            ans112 {
                true,
            },
        },
        question112 {
            para112 {
                createTree2(),
                90,
            },
            ans112 {
                true,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para112)
		fmt.Printf("Expect: %v\n", c.ans112)
		fmt.Printf("Result: %v\n\n", hasPathSum(c.root, c.target))
	}
}
