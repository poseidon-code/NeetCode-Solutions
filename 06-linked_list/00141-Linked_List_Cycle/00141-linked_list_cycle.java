// 141: Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/


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

    public static void createLoop(LinkedList ll, int p) {
        if (ll.head.next==null) return;
        if (p==-1) return;

        ListNode tail = ll.head, node = ll.head;
        while (tail.next!=null) tail = tail.next;
        for (int i=0; i<p; i++) node = node.next;

        tail.next = node;
    }

    // SOLUTION
    public boolean hasCycle(ListNode head) {
        if (head==null) return false;
        ListNode fast=head, slow=head;
        
        while (fast != null && fast.next != null) {
            fast = fast.next.next;
            slow = slow.next;
            if (fast == slow) return true;
        }

        return false;
    }


    public static void main(String[] args) {
        Solution o = new Solution();
        LinkedList ll = new LinkedList();

        // INPUT
        int[] li = {3,2,0,-4};
        llInput(ll, li);
        createLoop(ll, 1);

        // OUTPUT
        var result = o.hasCycle(ll.head);
        System.out.println(result ? "true" : "false");
    }
}