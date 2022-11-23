// 206: Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/


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
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) return head;

        ListNode previous = null;
        ListNode current = head;
        ListNode next = head.next;

        while (current != null) {
            next = current.next;
            current.next = previous;
            previous = current;
            current = next;
        }

        return previous;
    }


    public static void main(String[] args) {
        Solution o = new Solution();
        LinkedList ll = new LinkedList();

        // INPUT
        int[] li = {1,2,3,4,5};
        llInput(ll, li);

        // OUTPUT
        var result = o.reverseList(ll.head);
        llOutput(result);
    }
}