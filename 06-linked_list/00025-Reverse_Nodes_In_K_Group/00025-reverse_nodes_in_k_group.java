// 25: Reverse Nodes in K-Group
// https://leetcode.com/problems/reverse-nodes-in-k-group/


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
    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode dummy = new ListNode();
        dummy.next = head;

        ListNode previous = dummy;
        ListNode current = dummy.next;
        ListNode t = null;

        int count = k;

        while (current != null) {
            if (count > 1) {
                t = previous.next;
                previous.next = current.next;
                current.next = current.next.next;
                previous.next.next = t;
                count--;
            } else {
                previous = current;
                current = current.next;
                count = k;

                ListNode end = current;
                for (int i = 0; i < k; i++) {
                    if (end == null) return dummy.next;
                    end = end.next;
                }
            }
        }

        return dummy.next;
    }


    public static void main(String[] args) {
        Solution o = new Solution();
        LinkedList ll = new LinkedList();

        // INPUT
        int[] li = {1,2,3,4,5};
        int k = 3;
        llInput(ll, li);

        // OUTPUT
        var result = o.reverseKGroup(ll.head, k);
        llOutput(result);
    }
}