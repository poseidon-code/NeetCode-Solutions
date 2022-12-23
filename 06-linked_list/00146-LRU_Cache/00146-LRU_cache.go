// 146: LRU Cache
// https://leetcode.com/problems/lru-cache/

package main

import "fmt"

// SOLUTION
type Node struct {
    k int
    v int
    previous *Node
    next *Node
}

type LRUCache struct {
    capacity int
    cache map[int]*Node
    left *Node
    right *Node
}

func Constructor(capacity int) LRUCache {
    left, right := &Node{}, &Node{}
    left.next, right.previous = right, left
    return LRUCache{
        capacity: capacity,
        cache: map[int]*Node{},
        left: left,
        right: right,
    }
}

func (this *LRUCache) remove(node *Node) {
    previous := node.previous;
    next := node.next;

    previous.next = next;
    next.previous = previous;
}

func (this *LRUCache) insert(node *Node) {
    previous := this.right.previous
    next := this.right

    previous.next = node
    next.previous = node
    node.previous = previous
    node.next = next
}

func (this *LRUCache) Get(key int) int {
    if _, ok := this.cache[key]; ok {
        this.remove(this.cache[key])
        this.insert(this.cache[key])
        return this.cache[key].v
    }
    return -1
}

func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.cache[key]; ok {this.remove(this.cache[key])}

    this.cache[key] = &Node{k: key, v: value}
    this.insert(this.cache[key])

    if (len(this.cache) > this.capacity) {
        lru := this.left.next
        this.remove(lru)
        delete(this.cache, lru.k)
    }
}

func main() {
    // OPERATIONS
    o := Constructor(2)
    o.Put(1, 1)
    o.Put(2, 2)
    fmt.Println(o.Get(1))
    o.Put(3, 3);
    fmt.Println(o.Get(2))
    o.Put(4, 4);
    fmt.Println(o.Get(1))
    fmt.Println(o.Get(3))
    fmt.Println(o.Get(4))
}
