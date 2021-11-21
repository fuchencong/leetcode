#!/usr/bin/env python3

# Copyright (C) fuchencong.com

class Solution:
    def romanToInt(self, s):
        roman_map = {
            "I": 1,
            "V": 5,
            "X": 10,
            "L": 50,
            "C": 100,
            "D": 500,
            "M": 1000
        }

        result = i = 0
        while i < len(s):
            if (i != len(s) - 1 and
                    roman_map[s[i]] < roman_map[s[i + 1]]):
                result += roman_map[s[i + 1]] - roman_map[s[i]]
                i += 2
            else:
                result += roman_map[s[i]]
                i += 1

        return result


def test():
    solution = Solution()

    def output_case(case, expect, result):
        case_result = "ok" if expect == result else "fail"
        print("%s: %s" % (case, case_result))
        if case_result != "ok":
            print("  expect: %s, result: %s" % (expect, result))

    s = "III"
    expect = 3
    result = solution.romanToInt(s)
    output_case("case01:", expect, result)

    s = "IV"
    expect = 4
    result = solution.romanToInt(s)
    output_case("case02:", expect, result)

    s = "IX"
    expect = 9
    result = solution.romanToInt(s)
    output_case("case03:", expect, result)

    s = "LVIII"
    expect = 58
    result = solution.romanToInt(s)
    output_case("case04:", expect, result)

    s = "MCMXCIV"
    expect = 1994
    result = solution.romanToInt(s)
    output_case("case05:", expect, result)


if __name__ == "__main__":
    test()
