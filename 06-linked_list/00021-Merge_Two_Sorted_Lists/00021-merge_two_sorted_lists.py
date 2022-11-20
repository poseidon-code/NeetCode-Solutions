# 21: Merge Two Sorted Lists
# https://leetcode.com/problems/merge-two-sorted-lists/


class ListNode:
    def __init__(self, x: int=None):
        self.val = x
        self.next = None

class LinkedList:
    def __init__(self) -> None:
        self.head = None

    def Input(self, inputs: list[int]):
        for i in inputs:
            node = ListNode(i)
            if self.head:
                tail = self.head
                while tail.next != None: tail = tail.next
                tail.next = node
            else:
                self.head = node

    def Output(self, head: ListNode):
        while (head != None):
            print(head.val, "-> ", end='')
            head = head.next
        print("NULL")


class Solution:
    def mergeTwoLists(self, list1: ListNode, list2: ListNode) -> ListNode:
        if list1 == None: return list2
        if list2 == None: return list1

        p = list1
        if list1.val > list2.val:
            p = list2
            list2 = list2.next
        else:
            list1 = list1.next

        c = p
        while list1!=None and list2!=None:
            if list1.val < list2.val:
                c.next = list1
                list1 = list1.next
            else:
                c.next = list2
                list2 = list2.next
            c = c.next

        if list1 == None:
            c.next = list2
        else:
            c.next = list1

        return p


if __name__ == "__main__":
    o = Solution()
    ll1, ll2 = LinkedList(), LinkedList()

    # INPUT
    li1, li2 = [1,2,4], [1,3,4]
    ll1.Input(li1)
    ll2.Input(li2)

    # OUTPUT
    result = o.mergeTwoLists(ll1.head, ll2.head)
    ll1.Output(result)
