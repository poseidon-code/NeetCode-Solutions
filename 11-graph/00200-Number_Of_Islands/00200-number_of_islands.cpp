// 200: Number Of Islands
// https://leetcode.com/problems/number-of-islands/

#include <vector>
#include <iostream>

using namespace std;


class Solution {
public:
    // SOLUTION
    int numIslands(vector<vector<char>>& grid) {
        int m = grid.size();
        int n = grid[0].size();

        int result = 0;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '1') {
                    dfs(grid, i, j, m, n);
                    result++;
                }
            }
        }

        return result;
    }
private:
    void dfs(vector<vector<char>>& grid, int i, int j, int m, int n) {
        if (
            i < 0 ||
            i >= m ||
            j < 0 ||
            j >= n ||
            grid[i][j] == '0'
        ) {
            return;
        }

        grid[i][j] = '0';

        dfs(grid, i - 1, j, m, n);
        dfs(grid, i + 1, j, m, n);
        dfs(grid, i, j - 1, m, n);
        dfs(grid, i, j + 1, m, n);
    }
};


int main() {
    Solution o;

    // INPUT
    vector<vector<char>> grid = {
        {'1','1','1','1','0'},
        {'1','1','0','1','0'},
        {'1','1','0','0','0'},
        {'0','0','0','0','0'}
    };

    // OUTPUT
    auto result = o.numIslands(grid);
    cout << result << endl;

    return 0;
}
