#!/usr/bin/env python3

# Copyright (C) fuchencong.com

class Solution:
    def twoSum(self, nums, target):
        pair = ((nums[i], nums[j], i, j)
                for i in range(len(nums)) 
                for j in range(i + 1, len(nums)))
        for x, y, i, j in pair:
            if x + y == target:
                return [i, j]


def test():
    s = Solution()

    nums = [2, 7, 11, 15]
    target = 9
    print(s.twoSum(nums, target))

    nums = [3, 2, 4]
    target = 6
    print(s.twoSum(nums, target))

    nums = [3, 3]
    target = 6
    print(s.twoSum(nums, target))


if __name__ == "__main__":
    test()
