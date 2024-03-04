// 51: N Queens
// https://leetcode.com/problems/n-queens


import java.util.List;
import java.util.ArrayList;

class Solution {
    private void backtrack(int n, int row, List<String> nQueens, List<List<String>> result) {
        if (row == n) {
            result.add(new ArrayList<>(nQueens));
            return;
        }

        for (int col=0; col<n; col++) {
            if (isValid(nQueens, row, col, n)) {
                StringBuilder r = new StringBuilder(nQueens.get(row));
                r.setCharAt(col, 'Q');
                nQueens.set(row, r.toString());
                backtrack(n, row + 1, nQueens, result);
                r.setCharAt(col, '.');
                nQueens.set(row, r.toString());
            }
        }
    }

    private boolean isValid(List<String> nQueens, int row, int col, int n) {
        for (int i=0; i<row; i++)
            if (nQueens.get(i).charAt(col) == 'Q')
                return false;

        for (int i=row-1, j=col-1; i>=0 && j>=0; i--, j--)
            if (nQueens.get(i).charAt(j) == 'Q')
                return false;

        for (int i=row-1, j=col+1; i>=0 && j<n; i--, j++)
            if (nQueens.get(i).charAt(j) == 'Q')
                return false;

        return true;
    }

    
    // SOLUTION
    public List<List<String>> solveNQueens(int n) {
        List<List<String>> result = new ArrayList<>();
        List<String> nQueens = new ArrayList<String>();
        for (int i=0; i<n; i++) {
            nQueens.add(".".repeat(n));
        }
        backtrack(n, 0, nQueens, result);
        return result;
    }

    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int n = 4;

        // OUTPUT
        var result = o.solveNQueens(n);
        System.out.println(result);
    }
}
