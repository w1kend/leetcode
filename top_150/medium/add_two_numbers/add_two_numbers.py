from typing import List, Optional

# https://leetcode.com/problems/add-two-numbers


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def list(self) -> List:
        r = list()
        i = self
        while i is not None:
            r.append(i.val)
            i = i.next

    def fromList(self, l: List):
        if len(l) == 0:
            return

        i = self
        i.val = l[0]
        for j in range(1, len(l)-1):
            i.next = ListNode(l[j])
            i = i.next
        return self


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        root = ListNode()

        overflow = 0
        currNode = root
        while l1 is not None or l2 is not None:
            sum = 0
            # calculate sum
            if l1 is not None:
                sum += l1.val
                l1 = l1.next
            if l2 is not None:
                sum += l2.val
                l2 = l2.next

            sum += overflow
            overflow = 0

            if sum >= 10:
                sum = sum - 10
                overflow = 1

            currNode.next = ListNode(sum)
            currNode = currNode.next

        if overflow > 0:
            currNode.next = ListNode(overflow)

        return root.next
