// 211: Design Add And Search Words Data Structure
// https://leetcode.com/problems/design-add-and-search-words-data-structure/



class WordDictionary {
    private WordDictionary[] children;
    private boolean isEndOfWord;

    public WordDictionary() {
        this.children = new WordDictionary[26];
        this.isEndOfWord = false;
    }
    

    public void addWord(String word) {
        WordDictionary current = this;

        for (char c: word.toCharArray()) {
            if (current.children[c - 'a'] == null)
                current.children[c - 'a'] = new WordDictionary();
            current = current.children[c - 'a'];
        }

        current.isEndOfWord = true;
    }
    

    public boolean search(String word) {
        WordDictionary current = this;

        for (int i=0; i<word.length(); i++){
            char c = word.charAt(i);

            if (c == '.') {
                for (WordDictionary ch: current.children)
                    if( ch != null && ch.search(word.substring(i+1))) return true;
                return false;
            }

            if (current.children[c - 'a'] == null) return false;
            current = current.children[c - 'a'];
        }

        return current != null && current.isEndOfWord;
    }

    public static void main(String[] args) {
        WordDictionary wd = new WordDictionary();

        // OPERATIONS
        wd.addWord("bad");
        wd.addWord("dad");
        wd.addWord("mad");
        System.out.println((wd.search("pad") ? "true" : "false"));
        System.out.println((wd.search("bad") ? "true" : "false"));
        System.out.println((wd.search(".ad") ? "true" : "false"));
        System.out.println((wd.search("b..") ? "true" : "false"));
    }
}
