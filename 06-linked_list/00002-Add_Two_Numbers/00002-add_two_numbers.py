# 2: Add Two Numbers
# https://leetcode.com/problems/add-two-numbers/


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
    # SOLUTION
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        head = ListNode(0)
        t = head
        c = 0

        while c!=0 or l1!=None or l2!=None:
            if l1!=None:
                c += l1.val
                l1 = l1.next
            if l2!=None:
                c += l2.val
                l2 = l2.next

            t.next = ListNode(c%10)
            t = t.next
            c //= 10
        
        return head.next


if __name__ == "__main__":
    o = Solution()
    ll1, ll2 = LinkedList(), LinkedList()

    # INPUT
    li1, li2 = [2,4,3], [5,6,4]
    ll1.Input(li1)
    ll2.Input(li2)

    # OUTPUT
    result = o.addTwoNumbers(ll1.head, ll2.head)
    ll1.Output(result)
