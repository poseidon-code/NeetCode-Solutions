// 695: Max Area Of Island
// https://leetcode.com/problems/max-area-of-island


class Solution {
    int result = 0;

    public int maxAreaOfIsland(int[][] grid, int i, int j) {
        if (i < 0 || j < 0 || i == grid.length || j == grid[0].length || grid[i][j] == 0) {
            return 0;  
        }
    
        grid[i][j] = 0;
        
        return (
            1 + maxAreaOfIsland(grid, i + 1, j)
            + maxAreaOfIsland(grid, i - 1, j)
            + maxAreaOfIsland(grid, i, j + 1)
            + maxAreaOfIsland(grid, i, j - 1)
        );
    }
    
    // SOLUTION
    public int maxAreaOfIsland(int[][] grid) {
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                result = Math.max(result, maxAreaOfIsland(grid, i, j));
            }
        }
    
        return result;
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[][] grid = {
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
        var result = o.maxAreaOfIsland(grid);
        System.out.println(result);
    }
}
