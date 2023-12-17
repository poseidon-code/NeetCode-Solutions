// 208: Implement Trie Prefix Tree
// https://leetcode.com/problems/implement-trie-prefix-tree


class Trie {
    class TrieNode {
        public TrieNode[] children = new TrieNode[26];
        public boolean isWord; 

        public TrieNode() {
            for (int i=0; i<26; i++) children[i] = null;
            isWord = false;
        }
    }

    public TrieNode root;

    public Trie() {
        root = new TrieNode();
    }

    public void insert(String word) {
        TrieNode node = root;
        int current = 0;
        
        for (int i=0; i<word.length(); i++) {
            current = word.charAt(i) - 'a';
            if (node.children[current] == null)
                node.children[current] = new TrieNode();
            node = node.children[current];
        }

        node.isWord = true;
    }

    public boolean search(String word) {
        TrieNode node = root;
        int current = 0;

        for (int i=0; i<word.length(); i++) {
            current = word.charAt(i) - 'a';
            if (node.children[current] == null)
                return false;
            node = node.children[current];
        }
        
        return node.isWord;
    }

    public boolean startsWith(String prefix) {
        TrieNode node = root;
        int current = 0;

        for (int i=0; i<prefix.length(); i++) {
            current = prefix.charAt(i) - 'a';
            if (node.children[current] == null) {
                return false;
            }
            node = node.children[current];
        }

        return true;
    }

    public static void main(String[] args) {
        Trie trie = new Trie();

        // OPERATION
        trie.insert("apple");
        System.out.println(trie.search("apple") ? "true" : "false");
        System.out.println(trie.search("app") ? "true" : "false");
        System.out.println(trie.startsWith("app") ? "true" : "false");
        trie.insert("app");
        System.out.println(trie.search("app") ? "true" : "false");
    }
}
