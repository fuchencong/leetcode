#!/usr/bin/env python3

# Copyright (C) fuchencong.com

class Solution:
    def maxArea(self, height):
        def generate_height_index(height):
            index_list = []

            curr_highest = -1
            curr_highest_index = -1
            for i, h in enumerate(height):
                if h > curr_highest:
                    curr_highest = h
                    curr_highest_index = i
                index_list.append(curr_highest_index)
            return index_list

        result = 0
        curr_highest_index = generate_height_index(height)
        for i, h in enumerate(curr_highest_index):
            # not the left
            if i != h:
                continue

            # iterate the right
            start = curr_highest_index[-1]
            for j in range(start, len(height)):
                area_hight = min(height[i], height[j])
                result = max(result, area_hight * (j - i))
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
