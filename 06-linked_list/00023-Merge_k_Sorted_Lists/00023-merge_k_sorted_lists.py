# 23: Merge k Sorted Lists
# https://leetcode.com/problems/merge-k-sorted-lists


from queue import PriorityQueue


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
    def mergeKLists(self, lists: list[ListNode]) -> ListNode:
        heap = PriorityQueue()
        head = ListNode(None)
        c = head

        for i, node in enumerate(lists):
            if node:
                heap.put((node.val, i, node))

        while heap.qsize()>0:
            t = heap.get()
            c.next, i = t[2], t[1]
            c = c.next
            if c.next:
                heap.put((c.next.val, i, c.next))
        
        return head.next


if __name__ == "__main__":
    o = Solution()
    ll1, ll2, ll3 = LinkedList(), LinkedList(), LinkedList()

    # INPUT
    li1, li2, li3 = [1,4,5], [1,3,4], [2,6]
    ll1.Input(li1)
    ll2.Input(li2)
    ll3.Input(li3)
    lists: list[ListNode] = [ll1.head, ll2.head, ll3.head]

    # OUTPUT
    result = o.mergeKLists(lists)
    ll1.Output(result)
