#!/usr/bin/env python3

# Copyright (C) fuchencong.com

class Solution:
    def convert(self, s, numRows):
        def gen_line(s, total_num, curr_num):
            result = ""
            length = len(s)
            extra_offset = 2 * (total_num - curr_num - 1)
            while (curr_num < length):
                result += s[curr_num]
                next_num = curr_num + 2 * (total_num - 1)

                extra = curr_num + extra_offset
                if (extra != curr_num and
                    extra != next_num and
                    extra < length):
                    result += s[extra]
                curr_num = next_num
            return result

        if numRows == 1:
            return s

        result = ""
        for i in range(numRows):
            result += gen_line(s, numRows, i)
        return result


def test():
    solution = Solution()

    def output_case(case, expect, result):
        case_result = "ok" if expect == result else "fail"
        print("%s: %s" % (case, case_result))
        if case_result != "ok":
            print("  expect: %s, result: %s" % (expect, result))

    s = "PAYPALISHIRING"
    num_rows = 3
    expect = "PAHNAPLSIIGYIR"
    result = solution.convert(s, num_rows)
    output_case("case01:", expect, result)

    s = "PAYPALISHIRING"
    num_rows = 4
    expect = "PINALSIGYAHRPI"
    result = solution.convert(s, num_rows)
    output_case("case02:", expect, result)

    s = "A"
    num_rows = 1
    expect = "A"
    result = solution.convert(s, num_rows)
    output_case("case03:", expect, result)

    s = "PAYPALISHIRING"
    num_rows = 1
    expect = "PAYPALISHIRING"
    result = solution.convert(s, num_rows)
    output_case("case04:", expect, result)

    s = "ABCDE"
    num_rows = 2
    expect = "ACEBD"
    result = solution.convert(s, num_rows)
    output_case("case05:", expect, result)


if __name__ == "__main__":
    test()
