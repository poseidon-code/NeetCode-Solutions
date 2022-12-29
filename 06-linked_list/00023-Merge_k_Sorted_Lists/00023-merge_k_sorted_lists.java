// 23: Merge k Sorted Lists
// https://leetcode.com/problems/merge-k-sorted-lists


import java.util.PriorityQueue;

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
    public ListNode mergeKLists(ListNode[] lists) {
        PriorityQueue<ListNode> heap = new PriorityQueue<>(lists.length, (l,r)-> l.val - r.val);
        ListNode head = new ListNode();
        ListNode c = head;

        for (ListNode node : lists)
            if (node!=null)
                heap.add(node);

        while (!heap.isEmpty()) {
            c.next = heap.poll();
            c = c.next;
            if (c.next!=null) heap.add(c.next);
        }

        return head.next;
    }


    public static void main(String[] args) {
        Solution o = new Solution();
        LinkedList ll1 = new LinkedList();
        LinkedList ll2 = new LinkedList();
        LinkedList ll3 = new LinkedList();

        // INPUT
        int[] li1 = {1,4,5};
        int[] li2 = {1,3,4};
        int[] li3 = {2,6};
        llInput(ll1, li1);
        llInput(ll2, li2);
        llInput(ll3, li3);
        ListNode[] lists = {ll1.head, ll2.head, ll3.head};

        // OUTPUT
        var result = o.mergeKLists(lists);
        llOutput(result);
    }
}