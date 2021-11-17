#!/usr/bin/env python3

# Copyright (C) fuchencong.com

class Solution:
    def longestPalindrome(self, s):
        def get_palindrome_str(s, left, right):
            length = len(s)
            while left >= 0 and right < length:
                if s[left] != s[right]:
                    break
                left -= 1
                right += 1
            return s[left + 1:right]

        result = ""
        for i in range(len(s)):
            j = i
            substr = get_palindrome_str(s, i, j)
            result = substr if len(substr) > len(result) else result

            j = i + 1
            substr = get_palindrome_str(s, i, j)
            result = substr if len(substr) > len(result) else result

        return result


def test():
    solution = Solution()
    s = "babad"
    print(solution.longestPalindrome(s))

    s = "cbbd"
    print(solution.longestPalindrome(s))

    s = "a"
    print(solution.longestPalindrome(s))

    s = "ac"
    print(solution.longestPalindrome(s))

    s = "xabbaz"
    print(solution.longestPalindrome(s))

    s = ""
    print(solution.longestPalindrome(s))

    s = "zzabbazz"
    print(solution.longestPalindrome(s))


if __name__ == "__main__":
    test()
