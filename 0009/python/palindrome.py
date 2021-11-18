#!/usr/bin/env python3

# Copyright (C) fuchencong.com

class Solution:
    def isPalindrome(self, x):
        if x < 0:
            return False

        def num_list(x):
            result = []
            while True:
                if x < 10:
                    result.append(x)
                    break
                x, reminder = divmod(x, 10)
                result.append(reminder)
            return result

        result = True
        l = num_list(x)
        i = 0
        j = len(l) - 1
        while i < j:
            if l[i] != l[j]:
                result = False
                break
            i += 1
            j -= 1
        return result


def test():
    solution = Solution()

    def output_case(case, expect, result):
        case_result = "ok" if expect == result else "fail"
        print("%s: %s" % (case, case_result))
        if case_result != "ok":
            print("  expect: %s, result: %s" % (expect, result))


    x = 121
    expect = True
    result = solution.isPalindrome(x)
    output_case("case01:", expect, result)

    x = -121
    expect = False
    result = solution.isPalindrome(x)
    output_case("case02:", expect, result)

    x = 10
    expect = False
    result = solution.isPalindrome(x)
    output_case("case03:", expect, result)

    x = -101
    expect = False
    result = solution.isPalindrome(x)
    output_case("case04:", expect, result)

    x = 1000
    expect = False
    result = solution.isPalindrome(x)
    output_case("case05:", expect, result)

    x = 1111
    expect = True
    result = solution.isPalindrome(x)
    output_case("case06:", expect, result)

    x = 1221
    expect = True
    result = solution.isPalindrome(x)
    output_case("case07:", expect, result)

    x = 3003
    expect = True
    result = solution.isPalindrome(x)
    output_case("case08:", expect, result)


if __name__ == "__main__":
    test()



                 
