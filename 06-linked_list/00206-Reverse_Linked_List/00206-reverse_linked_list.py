# 206: Reverse Linked List
# https://leetcode.com/problems/reverse-linked-list/


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
    def reverseList(self, head: ListNode) -> ListNode:
        if head == None or head.next == None: return head
        
        previous = None
        current = head
        next = current.next

        while (current != None):
            next = current.next
            current.next = previous
            previous = current
            current = next

        return previous


if __name__ == "__main__":
    o = Solution()
    ll = LinkedList()

    # INPUT
    li = [1,2,3,4,5]
    ll.Input(li)

    # OUTPUT
    result = o.reverseList(ll.head)
    ll.Output(result)
