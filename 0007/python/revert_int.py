#!/usr/bin/env python3

# Copyright (C) fuchencong.com

class Solution:
    def reverse(self, x):
        def check_value_ok(str_max_value, str_value):
            if ((len(str_value) > len(str_max_value)) or
               (len(str_value) == len(str_max_value) and
               (str_value > str_max_value))):
                return False
            else:
                return True

        sx = str(x)
        if x < 0:
            smax = str(-2**31)
            sv = "-" + sx[:0:-1].lstrip("0")
            sr = sv if check_value_ok(smax, sv) else "0"
        elif x > 0:
            smax = str(2**31 - 1)
            sv = sx[::-1].lstrip("0")
            sr = sv if check_value_ok(smax, sv) else "0"
        else:
            sr = "0"

        return int(sr)


def test():
    solution = Solution()

    def output_case(case, expect, result):
        case_result = "ok" if expect == result else "fail"
        print("%s: %s" % (case, case_result))
        if case_result != "ok":
            print("  expect: %s, result: %s" % (expect, result))

    x = 123
    expect = 321
    result = solution.reverse(x)
    output_case("case01:", expect, result)

    x = -123
    expect = -321
    result = solution.reverse(x)
    output_case("case02:", expect, result)

    x = 120
    expect = 21
    result = solution.reverse(x)
    output_case("case03:", expect, result)

    x = 0
    expect = 0
    result = solution.reverse(x)
    output_case("case04:", expect, result)

    x = 2**31 - 1
    expect = 0
    result = solution.reverse(x)
    output_case("case05:", expect, result)

    x = -2**31
    expect = 0
    result = solution.reverse(x)
    output_case("case06:", expect, result)

    x = 54321
    expect = 12345
    result = solution.reverse(x)
    output_case("case07:", expect, result)

    x = -220022
    expect = -220022
    result = solution.reverse(x)
    output_case("case08:", expect, result)

    x = -1234500
    expect = -54321
    result = solution.reverse(x)
    output_case("case09:", expect, result)

    x = -10000000
    expect = -1
    result = solution.reverse(x)
    output_case("case10:", expect, result)

    x = 10000000
    expect = 1
    result = solution.reverse(x)
    output_case("case11:", expect, result)

    x = 299990
    expect = 99992
    result = solution.reverse(x)
    output_case("case12:", expect, result)


if __name__ == "__main__":
    test()
