package leetcode

import (
    "fmt"
    "testing"
)

type question111 struct {
    para111
    ans111
}

type para111 struct {
    root *TreeNode
}

type ans111 struct {
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

    right2 := &TreeNode{
        30,
        nil,
        &TreeNode{40, nil, nil},
    }

    right1 := &TreeNode{
        20,
        nil,
        right2,
    }


    root.Left = nil
    root.Right = right1
    return root
}

func TestLongestPalindrome(t *testing.T) {
    cases := []question111 {
        question111 {
            para111 {
                createTree1(),
            },
            ans111 {
                2,
            },
        },
        question111 {
            para111 {
                createTree2(),
            },
            ans111 {
                4,
            },
        },
    }

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para111)
		fmt.Printf("Expect: %v\n", c.ans111)
		fmt.Printf("Result: %v\n\n", minDepth(c.root))
	}
}
