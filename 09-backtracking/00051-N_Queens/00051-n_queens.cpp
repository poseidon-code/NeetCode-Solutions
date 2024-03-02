// 51: N Queens
// https://leetcode.com/problems/n-queens


#include <iostream>
#include <vector>

using namespace std;

class Solution {
private:
    void backtrack(int n, int row, vector<string>& nQueens, vector<vector<string>>& result) {
        if (row == n) {
            result.push_back(nQueens);
            return;
        }

        for (int col=0; col<n; col++) {
            if (isValid(nQueens, row, col, n)) {
                nQueens[row][col] = 'Q';
                backtrack(n, row+1, nQueens, result);
                nQueens[row][col] = '.';
            }
        }
    }

    bool isValid(std::vector<std::string> &nQueens, int row, int col, int &n) {
        for (int i=0; i<row; i++)
            if (nQueens[i][col] == 'Q')
                return false;

        for (int i=row-1, j=col-1; i>=0 && j>=0; i--, j--)
            if (nQueens[i][j] == 'Q')
                return false;

        for (int i=row-1, j=col+1; i>=0 && j<n; i--, j++)
            if (nQueens[i][j] == 'Q')
                return false;

        return true;
    }

public:
    // SOLUTOIN
    vector<vector<string>> solveNQueens(int n) {
        vector<string> nQueens(n, string(n, '.'));
        vector<vector<string>> result;
        backtrack(n, 0, nQueens, result);
        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    int n = 4;

    // OUTPUT
    auto result = o.solveNQueens(n);

    cout << "["; for (auto x : result) {
        cout << "["; for (auto y : x) {
            cout << "\"" << y << "\", ";
        } cout << "\b\b], ";
    } cout << "\b\b]" << endl;

    return 0;
}
