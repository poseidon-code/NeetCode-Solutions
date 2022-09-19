// 74: Search a 2D Matrix
// https://leetcode.com/problems/search-a-2d-matrix/


class Solution {
    // SOLUTION
    public boolean searchMatrix(int[][] matrix, int target) {
        int rows = matrix.length;
		int cols = matrix[0].length;
        int row = 0;
        int col = cols-1;
			
        while (row<rows && col>-1) {
            int cur = matrix[row][col];
            if (cur == target) return true;
            if (target > cur) row++;
            else col--;
        }
        
        return false;
    }


    public static void main(String[] args) {
        Solution o = new Solution();
        
        // INPUT
        int[][] matrix = {{1,3,5,7},{10,11,16,20},{23,30,34,60}};
        int target = 3;

        // OUTPUT
        var result = o.searchMatrix(matrix, target);
        System.out.println(result ? "true" : "false");
    }
}