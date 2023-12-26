// 212: Word Search Ii
// https://leetcode.com/problems/word-search-ii/


#include<iostream>
#include<vector>
#include<string>

using namespace std;


class Solution {
private:
    vector<string> result;

    struct TrieNode {
        TrieNode* children[26] = {};
        string* word;
        
        void addWord(string& word) {
            TrieNode* current = this;

            for (char c : word) {
                c -= 'a';
                if (current->children[c] == nullptr)
                    current->children[c] = new TrieNode();
                current = current->children[c];
            }
            current->word = &word;
        }
    };

    void dfs(vector<vector<char>>& board, int r, int c, TrieNode* current) {
        if (r<0 || r==board.size() || c<0 || c==board[0].size() || board[r][c]=='#' || current->children[board[r][c]-'a']==nullptr) return;
        vector<int> dirs = {0, 1, 0, -1, 0};

        char oc = board[r][c];
        current = current->children[oc-'a'];

        if (current->word != nullptr) {
            result.push_back(*current->word);
            current->word = nullptr;
        }

        board[r][c] = '#';
        for (int i=0; i<4; ++i) dfs(board, r+dirs[i], c+dirs[i+1], current);
        board[r][c] = oc;
    }

public:
    // SDLUTION
    vector<string> findWords(vector<vector<char>>& board, vector<string>& words) {
        TrieNode node;
        for (string& word : words) node.addWord(word);

        for (int r=0; r<board.size(); ++r)
            for (int c=0; c<board[0].size(); ++c)
                dfs(board, r, c, &node);

        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<vector<char>> board = {
        {'o','a','a','n'},
        {'e','t','a','e'},
        {'i','h','k','r'},
        {'i','f','l','v'}
    };
    vector<string> words = {"oath","pea","eat","rain"};

    // OUTPUT
    auto result = o.findWords(board, words);
    cout << "["; for (auto x : result) cout<< x << ", "; cout<<"\b\b]" << endl;

    return 0;
}
