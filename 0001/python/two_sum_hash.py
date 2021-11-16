#!/usr/bin/env python3

# Copyright (C) fuchencong.com

class Solution:
    def twoSum(self, nums, target):
        hashtable = {}
        for index, value in enumerate(nums):
            reminder = target - value
            reminder_index = hashtable.get(reminder, None)
            if reminder_index is not None:
                return [index, reminder_index]
            else:
                hashtable[value] = index



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
