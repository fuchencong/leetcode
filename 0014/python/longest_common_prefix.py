#!/usr/bin/env python3

# Copyright (C) fuchencong.com

class Solution(object):
    def longestCommonPrefix(self, strs):
        result = ""
        for s in zip(*strs):
            if len(set(s)) == 1:
                result += s[0]
            else:
                break
        return result


def test():
    solution = Solution()

    def output_case(case, expect, result):
        case_result = "ok" if expect == result else "fail"
        print("%s: %s" % (case, case_result))
        if case_result != "ok":
            print("  expect: %s, result: %s" % (expect, result))

    strs = ["flower", "flow", "flight"]
    expect = "fl"
    result = solution.longestCommonPrefix(strs)
    output_case("case01:", expect, result)

    strs = ["dog", "racecar", "car"]
    expect = ""
    result = solution.longestCommonPrefix(strs)
    output_case("case02:", expect, result)

    strs = ["dog", "dog1", "dog2"]
    expect = "dog"
    result = solution.longestCommonPrefix(strs)
    output_case("case03:", expect, result)

    strs = ["d", "xy", "dxy"]
    expect = ""
    result = solution.longestCommonPrefix(strs)
    output_case("case04:", expect, result)

    strs = ["x", "y", "z"]
    expect = ""
    result = solution.longestCommonPrefix(strs)
    output_case("case04:", expect, result)

if __name__ == "__main__":
    test()
