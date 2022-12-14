# 141: Linked List Cycle
# https://leetcode.com/problems/linked-list-cycle/

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

    def createLoop(self, p):
        if self.head.next==None: return
        if p==-1: return

        tail, node = self.head, self.head
        while tail.next!=None: tail = tail.next
        for _ in range(p): node = node.next

        tail.next = node


class Solution:
    # SOLUTION
    def hasCycle(self, head: ListNode) -> bool:
        slow, fast = head, head

        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow == fast: return True
        
        return False


if __name__ == "__main__":
    o = Solution()
    ll = LinkedList()

    # INPUT
    li = [3,2,0,-4]
    ll.Input(li)
    ll.createLoop(1)

    # OUTPUT
    result = o.hasCycle(ll.head)
    print(result)