#!/usr/bin/env python

# Copyright (C) fuchencong.com

class Solution:
    def intToRoman(self, num):
        roman_map = [(1000, "M"),   (900, "CM"),
                     (500, "D"),    (400, "CD"),
                     (100, "C"),    (90, "XC"),
                     (50, "L"),     (40, "XL"),
                     (10, "X"),     (9, "IX"),
                     (5, "V"),      (4, "IV"),
                     (1, "I")]

        result = ""
        curr_index = 0
        while num > 0:
            div, reminder = divmod(num, roman_map[curr_index][0])
            result += div * roman_map[curr_index][1]
            num = reminder

            if div == 0:
                curr_index += 1
        return result


def test():
    solution = Solution()

    def output_case(case, expect, result):
        case_result = "ok" if expect == result else "fail"
        print("%s: %s" % (case, case_result))
        if case_result != "ok":
            print("  expect: %s, result: %s" % (expect, result))

    num = 3
    expect = "III"
    result = solution.intToRoman(num)
    output_case("case01:", expect, result)

    num = 4
    expect = "IV"
    result = solution.intToRoman(num)
    output_case("case02:", expect, result)

    num = 9
    expect = "IX"
    result = solution.intToRoman(num)
    output_case("case03:", expect, result)

    num = 58
    expect = "LVIII"
    result = solution.intToRoman(num)
    output_case("case04:", expect, result)

    num = 1994
    expect = "MCMXCIV"
    result = solution.intToRoman(num)
    output_case("case05:", expect, result)

    num = 6
    expect = "VI"
    result = solution.intToRoman(num)
    output_case("case06:", expect, result)

    num = 11
    expect = "XI"
    result = solution.intToRoman(num)
    output_case("case07:", expect, result)

    num = 21
    expect = "XXI"
    result = solution.intToRoman(num)
    output_case("case08:", expect, result)

    num = 91
    expect = "XCI"
    result = solution.intToRoman(num)
    output_case("case09:", expect, result)



if __name__ == "__main__":
    test()



