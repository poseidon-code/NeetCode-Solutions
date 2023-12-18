# 208: Implement Trie Prefix Tree
# https://leetcode.com/problems/implement-trie-prefix-tree


class TrieNode:
    def __init__(self) -> None:
        self.children: TrieNode[26] = []
        self.isWord = False

        for i in range(26):
            self.children.append(None)
            self.isWord = False

class Trie:
    def __init__(self):
        self.root: TrieNode = TrieNode()


    def insert(self, word: str) -> None:
        node: TrieNode = self.root
        current = 0;
        
        for i in range(len(word)):
            current = ord(word[i]) - ord('a')
            if node.children[current] == None:
                node.children[current] = TrieNode()
            node = node.children[current];

        node.isWord = True

    def search(self, word: str) -> bool:
        node: TrieNode = self.root
        current = 0;
        
        for i in range(len(word)):
            current = ord(word[i]) - ord('a')
            if node.children[current] == None:
                return False
            node = node.children[current];

        return node.isWord

    def startsWith(self, prefix: str) -> bool:
        node: TrieNode = self.root
        current = 0;
        
        for i in range(len(prefix)):
            current = ord(prefix[i]) - ord('a')
            if node.children[current] == None:
                return False
            node = node.children[current];

        return True
    

if __name__ == "__main__":
    trie: Trie = Trie()

    # OPERATIONS
    trie.insert("apple")
    print(trie.search("apple"))
    print(trie.search("app"))
    print(trie.startsWith("app"))
    trie.insert("app")
    print(trie.search("app"))
