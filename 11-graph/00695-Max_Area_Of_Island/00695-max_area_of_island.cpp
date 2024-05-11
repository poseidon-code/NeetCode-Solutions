// 695: Max Area Of Island
// https://leetcode.com/problems/max-area-of-island

#include <vector>
#include <iostream>

using namespace std;

class Solution {
private:
    int dfs(vector<vector<int>>& grid, int i, int j, int m, int n) {
        if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0) {
            return 0;
        }
        grid[i][j] = 0;

        return
            1 + dfs(grid, i - 1, j, m, n)
            + dfs(grid, i + 1, j, m, n)
            + dfs(grid, i, j - 1, m, n)
            + dfs(grid, i, j + 1, m, n);
    }

public:
    // SOLUTION
    int maxAreaOfIsland(vector<vector<int>>& grid) {
        int m = grid.size();
        int n = grid[0].size();

        int result = 0;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 1) {
                    result = max(result, dfs(grid, i, j, m, n));
                }
            }
        }

        return result;
    }

};


int main() {
    Solution o;

    // INPUT
    vector<vector<int>> grid = {
        {0,0,1,0,0,0,0,1,0,0,0,0,0},
        {0,0,0,0,0,0,0,1,1,1,0,0,0},
        {0,1,1,0,1,0,0,0,0,0,0,0,0},
        {0,1,0,0,1,1,0,0,1,0,1,0,0},
        {0,1,0,0,1,1,0,0,1,1,1,0,0},
        {0,0,0,0,0,0,0,0,0,0,1,0,0},
        {0,0,0,0,0,0,0,1,1,1,0,0,0},
        {0,0,0,0,0,0,0,1,1,0,0,0,0}
    };

    // OUTPUT
    auto result = o.maxAreaOfIsland(grid);
    cout << result << endl;

    return 0;
}
