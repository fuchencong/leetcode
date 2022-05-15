package leetcode

import (
	"fmt"
	"testing"
)

type question49 struct {
	para49
	ans49
}

type para49 struct {
	s       []string
}

type ans49 struct {
	result [][]string
}

func TestLongestPalindrome(t *testing.T) {
	cases := []question49{
		question49{
			para49{
                []string{"eat","tea","tan","ate","nat","bat"},
			},
            ans49{
                [][]string {
                    []string{"bat"},
                    []string{"nat","tan"},
                    []string{"ate","eat","tea"},
                },
			},
		},
		question49{
			para49{
                []string{""},
			},
            ans49{
                [][]string {
                    []string{""},
                },
			},
		},
		question49{
			para49{
                []string{"a"},
			},
            ans49{
                [][]string {
                    []string{"a"},
                },
			},
		},
	}

	for _, c := range cases {
		fmt.Printf("Input: %v\n", c.para49)
		fmt.Printf("Expect: %v\n", c.ans49)
		fmt.Printf("Result: %v\n\n", groupAnagrams(c.s))
	}
}
