// 21: Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/


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
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        if (list1 == null) return list2;
        if (list2 == null) return list1;

        ListNode p = list1;
        if (list1.val > list2.val) {
            p = list2;
            list2 = list2.next;
        } else {
            list1 = list1.next;
        }

        ListNode c = p;
        while (list1!=null && list2!=null) {
            if (list1.val < list2.val) {
                c.next = list1;
                list1 = list1.next;
            } else {
                c.next = list2;
                list2 = list2.next;
            }
            c = c.next;
        }

        if (list1 == null) c . next = list2;
        else c.next = list1;

        return p;
    }


    public static void main(String[] args) {
        Solution o = new Solution();
        LinkedList head1 = new LinkedList();
        LinkedList head2 = new LinkedList();

        // INPUT
        int[] li1 = {1,2,4};
        int[] li2 = {1,3,4};
        llInput(head1, li1);
        llInput(head2, li2);

        // OUTPUT
        var result = o.mergeTwoLists(head1.head, head2.head);
        llOutput(result);
    }
}