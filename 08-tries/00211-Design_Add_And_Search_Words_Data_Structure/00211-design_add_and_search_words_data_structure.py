# 211: Design Add And Search Words Data Structure
# https://leetcode.com/problems/design-add-and-search-words-data-structure/



class WordDictionary:
    def __init__(self):
        self.children = [None]*26
        self.isEndOfWord = False

    def addWord(self, word: str) -> None:
        current = self

        for c in word:
            if current.children[ord(c) - ord('a')] == None:
                current.children[ord(c) - ord('a')] = WordDictionary()
            current = current.children[ord(c) - ord('a')]
        
        current.isEndOfWord = True
        

    def search(self, word: str) -> bool:
        current = self

        for i in range(len(word)):
            c = word[i]
            if c == '.':
                for ch in current.children:
                    if ch != None and ch.search(word[i+1:]): return True
                return False
            
            if current.children[ord(c) - ord('a')] == None: return False
            current = current.children[ord(c) - ord('a')]
        
        return current != None and current.isEndOfWord
    


if __name__ == "__main__":
    wd = WordDictionary()

    # OPERATIONS
    wd.addWord("bad");
    wd.addWord("dad");
    wd.addWord("mad");
    print(wd.search("pad"))
    print(wd.search("bad"))
    print(wd.search(".ad"))
    print(wd.search("b.."))
