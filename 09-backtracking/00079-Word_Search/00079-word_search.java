// 79: Word Search
// https://leetcode.com/problems/word-search


class Solution {
    public boolean backtrack(char[][] board, int i, int j, String word) {
        if (word.length() == 0) return true;
        if (i<0 || i>=board.length || j<0 || j>=board[0].length || board[i][j]!=word.charAt(0)) return false;
        
        char c = board[i][j];
        board[i][j] = '*';
        String s = word.substring(1);
        
        boolean result =  backtrack(board, i-1, j, s)
        || backtrack(board, i+1, j, s)
        || backtrack(board, i, j-1, s)
        || backtrack(board, i, j+1, s);
        board[i][j] = c;
        
        return result;
    }


    // SOLUTION
    public boolean exist(char[][] board, String word) {
        for (int i=0; i<board.length; i++)
            for (int j=0; j<board[0].length; j++)
                if (backtrack(board, i, j, word))
                    return true;
        
        return false;
    }


    public static void main(String[] args) {
        Solution s = new Solution();

        // INPUT :
        char[][] board = {
            {'A','B','C','E'},
            {'S','F','C','S'},
            {'A','D','E','E'}
        };
        String word = "ABCCED";

        // OUTPUT :
        System.out.println(s.exist(board, word));
    }
}
