// 146: LRU Cache
// https://leetcode.com/problems/lru-cache/

#include <iostream>
#include <unordered_map>

using namespace std;

class Node {
public:
    int k;
    int v;
    Node* previous;
    Node* next;

    Node(int key, int value) : k(key), v(value) {
        previous = nullptr;
        next = nullptr;
    }
};

// SOLUTION
class LRUCache {
private:
    int capacity;
    unordered_map<int, Node*> cache;
    Node* left;
    Node* right;

    void remove(Node* node) {
        Node* previous = node->previous;
        Node* next = node->next;

        previous->next = next;
        next->previous = previous;
    }
    
    void insert(Node* node) {
        Node* previous = right->previous;
        Node* next = right;

        previous->next = node;
        next->previous = node;
        node->previous = previous;
        node->next = next;
    }

public:
    LRUCache(int capacity) {
        this->capacity = capacity;
        left = new Node(0, 0);
        right = new Node(0, 0);
        left->next = right;
        right->previous = left;
    }

    int get(int key) {
        if (cache.find(key) != cache.end()) {
            remove(cache[key]);
            insert(cache[key]);
            return cache[key]->v;
        }
        return -1;
    }

    void put(int key, int value) {
        if (cache.find(key) != cache.end()) remove(cache[key]);

        cache[key] = new Node(key, value);
        insert(cache[key]);

        if (cache.size() > capacity) {
            Node* lru = left->next;
            remove(lru);
            cache.erase(lru->k);
        }
    }
};


int main() {
    // OPERATIONS
    LRUCache *o = new LRUCache(2);
    o->put(1, 1);
    o->put(2, 2);
    cout << o->get(1) << endl;
    o->put(3, 3);
    cout << o->get(2) << endl;
    o->put(4, 4);
    cout << o->get(1) << endl;
    cout << o->get(3) << endl;
    cout << o->get(4) << endl;

    return 0;
}