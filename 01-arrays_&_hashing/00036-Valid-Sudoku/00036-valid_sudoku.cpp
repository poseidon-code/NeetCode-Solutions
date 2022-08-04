// 36: Valid Sudoku
// https://leetcode.com/problems/valid-sudoku/

#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    // SOLUTION
    bool isValidSudoku(vector<vector<char>>& board) {
        int used1[9][9] = {0}, used2[9][9] = {0}, used3[9][9] = {0};
        
        for (int i=0; i<board.size(); ++i)
            for (int j=0; j<board[i].size(); ++j)
                if (board[i][j]!='.') {
                    int num = board[i][j]-'0' - 1;
                    int k = i/3 * 3 + j/3;
                    
                    if (used1[i][num] || used2[j][num] || used3[k][num])
                        return false;
                    
                    used1[i][num] = used2[j][num] = used3[k][num] = 1;
                }

        return true;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<vector<char>> board = 
    {
        {'5','3','.','.','7','.','.','.','.'},
        {'6','.','.','1','9','5','.','.','.'},
        {'.','9','8','.','.','.','.','6','.'},
        {'8','.','.','.','6','.','.','.','3'},
        {'4','.','.','8','.','3','.','.','1'},
        {'7','.','.','.','2','.','.','.','6'},
        {'.','6','.','.','.','.','2','8','.'},
        {'.','.','.','4','1','9','.','.','5'},
        {'.','.','.','.','8','.','.','7','9'}
    };

    // OUTPUT
    auto result = o.isValidSudoku(board);
    cout<< (result ? "true" : "false") << endl;

    return 0;
}