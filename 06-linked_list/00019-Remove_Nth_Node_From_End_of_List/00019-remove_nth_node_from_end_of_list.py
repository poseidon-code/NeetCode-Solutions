# 19: Remove Nth Node From End of List
# https://leetcode.com/problems/remove-nth-node-from-end-of-list/


class ListNode:
    def __init__(self, x: int):
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
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        slow, fast = head, head

        for _ in range(n+1): fast = fast.next
        
        while fast != None:
            slow = slow.next
            fast = fast.next
        
        slow.next = slow.next.next
        return head


if __name__ == "__main__":
    o = Solution()
    ll = LinkedList()

    # INPUT
    li = [1,2,3,4,5]
    n = 2
    ll.Input(li)

    # OUTPUT
    result = o.removeNthFromEnd(ll.head, n)
    ll.Output(result)
