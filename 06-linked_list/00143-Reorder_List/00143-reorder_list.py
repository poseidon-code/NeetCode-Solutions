# 143: Reorder List
# https://leetcode.com/problems/reorder-list/

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
    def reorderList(self, head: ListNode) -> ListNode:
        if head == None or head.next == None: return head
        
        prev, slow, fast, l1, l2 = None, head, head, head, None

        while fast!=None and fast.next!=None:
            prev = slow
            slow = slow.next
            fast = fast.next.next
        prev.next = None

        p, c, n = None, slow, None
        while c!=None:
            n=c.next
            c.next=p
            p=c
            c=n
        l2 = p

        while l1!=None:
            n1, n2 = l1.next, l2.next
            l1.next = l2
            if n1==None: break

            l2.next=n1
            l1=n1
            l2=n2

        return head


if __name__ == "__main__":
    o = Solution()
    ll = LinkedList()

    # INPUT
    li = [1,2,3,4]
    ll.Input(li)

    # OUTPUT
    result = o.reorderList(ll.head)
    ll.Output(result)
