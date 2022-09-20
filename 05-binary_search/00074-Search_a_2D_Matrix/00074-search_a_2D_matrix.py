# 74: Search a 2D Matrix
# https://leetcode.com/problems/search-a-2d-matrix/


class Solution:
    # SOLUTION
    def searchMatrix(self, matrix: list[list[int]], target: int) -> bool:
        rows = len(matrix)
        cols = len(matrix[0])
        row = 0
        col = cols-1
			
        while row<rows and col>-1:
            cur = matrix[row][col]
            if cur == target: return True
            if target > cur: row+=1
            else: col-=1
        
        return False


if __name__ == "__main__":
    o = Solution()

    # INPUT
    matrix: list[list[int]] = [[1,3,5,7],[10,11,16,20],[23,30,34,60]]
    target: int = 3

    # OUTPUT
    result = o.searchMatrix(matrix, target)
    print(result)