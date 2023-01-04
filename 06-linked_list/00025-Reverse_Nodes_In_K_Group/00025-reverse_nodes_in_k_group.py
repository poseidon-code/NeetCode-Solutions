# 25: Reverse Nodes in K-Group
# https://leetcode.com/problems/reverse-nodes-in-k-group/


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
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        dummy = ListNode(None)
        dummy.next = head

        previous = dummy
        current = dummy.next
        t = ListNode(None)

        count = k

        while (current != None):
            if (count > 1):
                t = previous.next
                previous.next = current.next
                current.next = current.next.next
                previous.next.next = t
                count-=1
            else:
                previous = current
                current = current.next
                count = k

                end = current
                for i in range(k):
                    if (end == None): return dummy.next
                    end = end.next

        return dummy.next


if __name__ == "__main__":
    o = Solution()
    ll = LinkedList()

    # INPUT
    li = [1,2,3,4,5]
    k = 3
    ll.Input(li)

    # OUTPUT
    result = o.reverseKGroup(ll.head, k)
    ll.Output(result)
