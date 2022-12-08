// 2: Add Two Numbers
// https://leetcode.com/problems/add-two-numbers/


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
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode head = new ListNode();
        ListNode t = head;
        int c = 0;

        while (c!=0 || l1!=null || l2!=null) {
            c += (l1!=null ? l1.val : 0) + (l2!=null ? l2.val : 0);

            ListNode temp = new ListNode();
            temp.val = c%10;
            temp.next = null;

            t.next = temp;
            t = t.next;
            c /= 10;

            if (l1!=null) l1 = l1.next;
            if (l2!=null) l2 = l2.next;
        }

        return head.next;
    }


    public static void main(String[] args) {
        Solution o = new Solution();
        LinkedList ll1 = new LinkedList();
        LinkedList ll2 = new LinkedList();

        // INPUT
        int[] li1 = {2,4,3};
        int[] li2 = {5,6,4};
        llInput(ll1, li1);
        llInput(ll2, li2);

        // OUTPUT
        var result = o.addTwoNumbers(ll1.head, ll2.head);
        llOutput(result);
    }
}