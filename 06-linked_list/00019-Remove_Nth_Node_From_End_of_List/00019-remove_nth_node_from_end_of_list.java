// 19: Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/


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
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode slow=head, fast=head;
        for (int i=0; i<=n; i++) fast = fast.next;

        while (fast != null) {
            slow = slow.next;
            fast = fast.next;
        }

        slow.next = slow.next.next;
        return head;
    }


    public static void main(String[] args) {
        Solution o = new Solution();
        LinkedList ll = new LinkedList();

        // INPUT
        int[] li = {1,2,3,4,5};
        int n = 2;
        llInput(ll, li);

        // OUTPUT
        var result = o.removeNthFromEnd(ll.head, n);
        llOutput(result);
    }
}