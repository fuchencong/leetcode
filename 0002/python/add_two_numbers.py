#!/usr/bin/env python3

# Copyright(C) fuchencong.com

from itertools import zip_longest


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def iter_list(head):
    while head:
        yield head.val
        head = head.next


class Solution:
    def addTwoNumbers(self, l1, l2):
        iter1 = iter_list(l1)
        iter2 = iter_list(l2)

        curr = 0
        result = None
        prev = None
        for i, j in zip_longest(iter1, iter2, fillvalue=0):
            value = (i + j + curr) % 10
            curr = (i + j + curr) // 10
            node = ListNode(value)

            if prev is None:
                result = node
            else:
                prev.next = node
            prev = node

        if curr:
            node = ListNode(curr)
            prev.next = node

        return result


def test():
    def create_list(values):
        head = None
        prev = None 
        for v in values:
            node = ListNode(v)

            if prev:
                prev.next = node
            else:
                head = node
            prev = node

        return head

    def gen_values(head):
        return [v for v in iter_list(head)]

    s = Solution()
    l1 = create_list([2, 4, 3])
    l2 = create_list([5, 6, 4])
    print(gen_values(s.addTwoNumbers(l1, l2)))

    l1 = create_list([0])
    l2 = create_list([0])
    print(gen_values(s.addTwoNumbers(l1, l2)))

    l1 = create_list([0, 1, 1])
    l2 = create_list([9, 9, 9])
    print(gen_values(s.addTwoNumbers(l1, l2)))

    l1 = create_list([9,9,9,9,9,9,9])
    l2 = create_list([9,9,9,9])
    print(gen_values(s.addTwoNumbers(l1, l2)))

if __name__ == "__main__":
    test()
