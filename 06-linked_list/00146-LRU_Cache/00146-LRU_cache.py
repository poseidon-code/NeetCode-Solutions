# 146: LRU Cache
# https://leetcode.com/problems/lru-cache/


from typing import Dict

# SOLUTION
class Node:
    def __init__(self, key: int, value: int) -> None:
        self.k = key
        self.v = value
        self.previous = None
        self.next = None

class LRUCache:
    def __init__(self, capacity: int) -> None:
        self.capacity: int = capacity
        self.cache: Dict[int, Node] = {}
        self.left: Node = Node(0, 0)
        self.right: Node = Node(0, 0)

        self.left.next = self.right
        self.right.previous = self.left

    def remove(self, node: Node) -> None:
        previous = node.previous
        next = node.next

        previous.next = next
        next.previous = previous

    def insert(self, node: Node) -> None:
        previous = self.right.previous
        next = self.right

        previous.next = node
        next.previous = node
        node.previous = previous
        node.next = next

    def get(self, key: int) -> int:
        if key in self.cache:
            self.remove(self.cache[key])
            self.insert(self.cache[key])
            return self.cache[key].v
        return -1


    def put(self, key: int, value: int) -> None:
        if key in self.cache: self.remove(self.cache[key])

        self.cache[key] = Node(key, value)
        self.insert(self.cache[key])

        if len(self.cache) > self.capacity:
            lru = self.left.next
            self.remove(lru)
            del self.cache[lru.k]


if __name__ == "__main__":
    # OPERATIONS
    o = LRUCache(2)
    o.put(1, 1)
    o.put(2, 2)
    print(o.get(1))
    o.put(3, 3)
    print(o.get(2))
    o.put(4, 4)
    print(o.get(1))
    print(o.get(3))
    print(o.get(4))
