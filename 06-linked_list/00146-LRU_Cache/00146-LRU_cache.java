// 146: LRU Cache
// https://leetcode.com/problems/lru-cache/

import java.util.HashMap;
import java.util.Map;


// SOLUTION
class LRUCache {
    private class Node {
        private int k;
        private int v;
        Node previous;
        Node next;
    
        public Node(int key, int value) {
            this.k = key;
            this.v = value;
        }
    };

    private int capacity;
    private Map<Integer, Node> cache;
    private Node left;
    private Node right;

    private void remove(Node node) {
        Node previous = node.previous;
        Node next = node.next;

        previous.next = next;
        next.previous = previous;
    }
    
    private void insert(Node node) {
        Node previous = this.right.previous;
        Node next = this.right;

        previous.next = node;
        next.previous = node;
        node.previous = previous;
        node.next = next;
    }

    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.cache = new HashMap<>();
        this.left = new Node(0, 0);
        this.right = new Node(0, 0);
        this.left.next = right;
        this.right.previous = left;
    }

    public int get(int key) {
        if (cache.containsKey(key)) {
            remove(cache.get(key));
            insert(cache.get(key));
            return cache.get(key).v;
        }
        return -1;
    }

    public void put(int key, int value) {
        if (cache.containsKey(key)) remove(cache.get(key));

        cache.put(key, new Node(key, value));
        insert(cache.get(key));

        if (cache.size() > capacity) {
            Node lru = this.left.next;
            remove(lru);
            cache.remove(lru.k);
        }
    }


    public static void main(String[] args) {
        // OPERATIONS
        LRUCache o = new LRUCache(2);
        o.put(1, 1);
        o.put(2, 2);
        System.out.println(o.get(1));
        o.put(3, 3);
        System.out.println(o.get(2));
        o.put(4, 4);
        System.out.println(o.get(1));
        System.out.println(o.get(3));
        System.out.println(o.get(4));
    }
}