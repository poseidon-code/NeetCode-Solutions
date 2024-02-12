// 79: Word Search
// https://leetcode.com/problems/word-search


#include <iostream>
#include <vector>

using namespace std;

class Solution {
private:
    bool backtrack(vector<vector<char>>& board, int i, int j, string& word) {
        if (!word.size()) return true;
        if (i<0 || i>=board.size() || j<0 || j>=board[0].size() || board[i][j]!=word[0]) return false;

        char c = board[i][j];
        board[i][j] = '*';
        string s = word.substr(1);
        
        bool result =  backtrack(board, i-1, j, s)
                    || backtrack(board, i+1, j, s)
                    || backtrack(board, i, j-1, s)
                    || backtrack(board, i, j+1, s);
        board[i][j] = c;
        
        return result;
    }

public:
    // SOLUTOIN
    bool exist(vector<vector<char>>& board, string word) {
        for (int i=0; i<board.size(); i++)
            for (int j=0; j<board[0].size(); j++)
                if (backtrack(board, i, j, word))
                    return true;
        
        return false;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<vector<char>> board = {
        {'A','B','C','E'},
        {'S','F','C','S'},
        {'A','D','E','E'}
    };
    string word = "ABCCED";

    // OUTPUT
    auto result = o.exist(board, word);

    cout << boolalpha << result << endl;

    return 0;
}
