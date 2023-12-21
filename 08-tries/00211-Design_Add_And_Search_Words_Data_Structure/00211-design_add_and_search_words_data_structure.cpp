// 211: Design Add And Search Words Data Structure
// https://leetcode.com/problems/design-add-and-search-words-data-structure/


#include <iostream>
#include <vector>
#include <string>

using namespace std;


class WordDictionary {
private:
    vector<WordDictionary*> children;
    bool isEndOfWord;

public:
    WordDictionary(): isEndOfWord(false){
        this->children = vector<WordDictionary*>(26, nullptr);
        this->isEndOfWord = false;
    }
    

    void addWord(string word) {
        WordDictionary *current = this;

        for (char c: word) {
            if (current->children[c - 'a'] == nullptr)
                current->children[c - 'a'] = new WordDictionary();
            current = current->children[c - 'a'];
        }

        current->isEndOfWord = true;
    }


    bool search(string word) {
        WordDictionary *current = this;

        for (int i=0; i<word.length(); i++){
            char c = word[i];

            if (c == '.') {
                for (auto ch: current->children)
                    if (ch && ch->search(word.substr(i+1))) return true;
                return false;
            }

            if (current->children[c - 'a'] == nullptr) return false;
            current = current->children[c - 'a'];
        }

        return current && current->isEndOfWord;
    }
};


int main() {
    WordDictionary* wd = new WordDictionary();

    // OPERATIONS
    wd->addWord("bad");
    wd->addWord("dad");
    wd->addWord("mad");
    cout << (wd->search("pad") ? "true" : "false") << endl;
    cout << (wd->search("bad") ? "true" : "false") << endl;
    cout << (wd->search(".ad") ? "true" : "false") << endl;
    cout << (wd->search("b..") ? "true" : "false") << endl;

    return 0;
}
