// 208: Implement Trie Prefix Tree
// https://leetcode.com/problems/implement-trie-prefix-tree


#include <iostream>
#include <string>

using namespace std;


class TrieNode {
public:
    TrieNode* children[26];
    bool isWord;

    TrieNode() {
        for (int i=0; i<26; i++) children[i] = NULL;
        isWord = false;
    }
};

class Trie {
private:
    TrieNode* root;

public:
    Trie() {
        this->root = new TrieNode();
    }
    
    void insert(string word) {
        TrieNode* node = root;
        int current = 0;
        
        for (int i=0; i<word.size(); i++) {
            current = word[i] - 'a';
            if (node->children[current] == NULL)
                node->children[current] = new TrieNode();
            node = node->children[current];
        }

        node->isWord = true;
    }

    bool search(string word) {
        TrieNode* node = root;
        int current = 0;

        for (int i=0; i<word.size(); i++) {
            current = word[i] - 'a';
            if (node->children[current] == NULL)
                return false;
            node = node->children[current];
        }
        
        return node->isWord;
    }

    bool startsWith(string prefix) {
        TrieNode* node = root;
        int current = 0;

        for (int i=0; i<prefix.size(); i++) {
            current = prefix[i] - 'a';
            if (node->children[current] == NULL) {
                return false;
            }
            node = node->children[current];
        }

        return true;
    }
};


int main() {
    Trie* trie = new Trie();

    // OPERATION
    trie->insert("apple");
    cout << (trie->search("apple") ? "true" : "false") << endl;
    cout << (trie->search("app") ? "true" : "false") << endl;
    cout << (trie->startsWith("app") ? "true" : "false") << endl;
    trie->insert("app");
    cout << (trie->search("app") ? "true" : "false") << endl;

    return 0;
}
