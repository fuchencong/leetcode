#!/usr/bin/env python3

# Copyright (C) fuchencong.com

class Solution:
    def lengthOfLongestSubstring(self, s):
        result = 0

        for i in range(len(s)):
            occur_map = {}
            occur_map[s[i]] = True
            j = i + 1
            for j in range(i + 1, len(s)):
                if s[j] in occur_map:
                    break
                occur_map[s[j]] = True
            else:
                j = len(s)

            if j - i > result:
                result = j - i

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


                    
