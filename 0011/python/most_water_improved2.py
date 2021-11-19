#!/usr/bin/env python3

# Copyright (C) fuchencong.com

class Solution:
    def maxArea(self, height):
        i = 0
        j = len(height) - 1
        result = 0
        while i < j:
            if height[i] < height[j]:
                result = max(result, (j - i) * height[i])
                i += 1
            else:
                result = max(result, (j - i) * height[j])
                j -= 1
        return result


def test():
    solution = Solution()

    def output_case(case, expect, result):
        case_result = "ok" if expect == result else "fail"
        print("%s: %s" % (case, case_result))
        if case_result != "ok":
            print("  expect: %s, result: %s" % (expect, result))

    height = [1, 8, 6, 2, 5, 4, 8, 3, 7]
    expect = 49
    result = solution.maxArea(height)
    output_case("case01:", expect, result)

    height = [1, 1]
    expect = 1
    result = solution.maxArea(height)
    output_case("case02:", expect, result)

    height = [4, 3, 2, 1, 4]
    expect = 16
    result = solution.maxArea(height)
    output_case("case03:", expect, result)

    height = [1, 2, 1]
    expect = 2
    result = solution.maxArea(height)
    output_case("case04:", expect, result)


if __name__ == "__main__":
    test()
