#!/usr/bin/env python3

# Copyright (C) fuchencong.com

class Solution:
    def lengthOfLongestSubstring(self, s):
        result = 0

        start = 0
        occur_map = {}
        for i, c in enumerate(s):
            index = occur_map.get(c, None)
            if index is not None and index >= start:
                start = index + 1
            else:
                result = max(result, i - start + 1)
            occur_map[c] = i

        return result


def test():
    sl = Solution()

    s = "abcabcbb"
    print(sl.lengthOfLongestSubstring(s))

    s = "bbbbb"
    print(sl.lengthOfLongestSubstring(s))

    s = "pwwkew"
    print(sl.lengthOfLongestSubstring(s))

    s = ""
    print(sl.lengthOfLongestSubstring(s))

    s = " "
    print(sl.lengthOfLongestSubstring(s))

    s = "au"
    print(sl.lengthOfLongestSubstring(s))


if __name__ == "__main__":
    test()


                    
