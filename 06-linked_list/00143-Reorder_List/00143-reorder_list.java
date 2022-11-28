// 143: Reorder List
// https://leetcode.com/problems/reorder-list/


class Solution {
    public static class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }

    public static class LinkedList {
        ListNode head;
        LinkedList() { this.head = null; }
    }

    public static void llInput(LinkedList list, int[] inputs) {
        for (int i=inputs.length-1; i>=0; i--) {
            ListNode node = new ListNode(inputs[i], list.head);
            list.head = node;
        }
    }

    public static void llOutput(ListNode head) {
        while (head != null) {
            System.out.print(head.val + " -> ");
            head = head.next;
        }
        System.out.println("NULL");
    }

    // SOLUTION
    public ListNode reorderList(ListNode head) {
        if (head == null || head.next == null) return head;

        ListNode prev=null, slow=head, fast=head, l1=head, l2=null;

        while (fast!=null && fast.next!=null) {
            prev = slow;
            slow = slow.next;
            fast = fast.next.next;
        }
        prev.next = null;

        ListNode p=null, c=slow, n=null;
        while (c!=null) {
            n=c.next;
            c.next=p;
            p=c;
            c=n;
        }
        l2 = p;

        while (l1!=null) {
            ListNode n1=l1.next, n2=l2.next;
            l1.next = l2;
            if (n1==null) break;

            l2.next = n1;
            l1 = n1;
            l2 = n2;
        }

        return head;
    }


    public static void main(String[] args) {
        Solution o = new Solution();
        LinkedList ll = new LinkedList();

        // INPUT
        int[] li = {1,2,3,4};
        llInput(ll, li);

        // OUTPUT
        var result = o.reorderList(ll.head);
        llOutput(result);
    }
}