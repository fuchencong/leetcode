#!/usr/bin/env python3

# Copyright (C) fuchencong.com

class Solution:
    def myAtoi(self, s):
        negative = 0
        space_finish = sign_occur = False
        start = end = None
        for i, c in enumerate(s):
            if c == " " and not space_finish:
                continue

            space_finish = True

            if (c == "+" or c == "-") and not sign_occur:
                sign_occur = True
                negative = 1 if c == "-" else 0
                continue

            sign_occur = True

            if c.isdigit():
                end = i
                if start is None and c != "0":
                    start = i
            else:
                break

        if start is None:
            return 0

        def check_value_ok(value, max_value):
            if ((len(value) > len(max_value)) or
                ((len(value) == len(max_value)) and
                 (value > max_value))):
                return False
            return True

        def atoi(value):
            if value == str(-2**31):
                return -2**31

            result = 0
            sign = 1
            for i, c in enumerate(value):
                if i == 0 and c == '-':
                    sign = -1
                else:
                    result = (result * 10 + ord(c) - ord("0"))
            return result * sign

        sv = '-' * negative + s[start:end + 1]
        max_value = str(-2**31) if negative else str(2**31-1)
        if check_value_ok(sv, max_value):
            return atoi(sv)
        else:
            return -2**31 if negative else 2**31 - 1


def test():
    solution = Solution()

    def output_case(case, expect, result):
        case_result = "ok" if expect == result else "fail"
        print("%s: %s" % (case, case_result))
        if case_result != "ok":
            print("  expect: %s, result: %s" % (expect, result))

    s = "42"
    expect = 42
    result = solution.myAtoi(s)
    output_case("case01:", expect, result)

    s = "   -42"
    expect = -42
    result = solution.myAtoi(s)
    output_case("case02:", expect, result)

    s = "4193 with words"
    expect = 4193
    result = solution.myAtoi(s)
    output_case("case03:", expect, result)

    s = "words and 987"
    expect = 0
    result = solution.myAtoi(s)
    output_case("case04:", expect, result)

    s = "-91283472332"
    expect = -2147483648
    result = solution.myAtoi(s)
    output_case("case05:", expect, result)

    s = "-0"
    expect = 0
    result = solution.myAtoi(s)
    output_case("case06:", expect, result)

    s = "-00005000"
    expect = -5000
    result = solution.myAtoi(s)
    output_case("case07:", expect, result)

    s = "5000"
    expect = 5000
    result = solution.myAtoi(s)
    output_case("case08:", expect, result)

    s = "2147483648"
    expect = 2147483647
    result = solution.myAtoi(s)
    output_case("case09:", expect, result)

    s = "-2147483648"
    expect = -2147483648
    result = solution.myAtoi(s)
    output_case("case10:", expect, result)

    s = "-2147483647"
    expect = -2147483647
    result = solution.myAtoi(s)
    output_case("case11:", expect, result)

    s = "0"
    expect = 0
    result = solution.myAtoi(s)
    output_case("case12:", expect, result)

    s = "   +0 123"
    expect = 0
    result = solution.myAtoi(s)
    output_case("case13:", expect, result)


if __name__ == "__main__":
    test()
