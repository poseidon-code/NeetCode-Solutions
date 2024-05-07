// 200: Number Of Islands
// https://leetcode.com/problems/number-of-islands/

package main

import "fmt"

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

// SOLUTION
func numIslands(grid [][]byte) int {
	result := 0

	for i, row := range grid {
		for j, point := range row {
			if point == '1' {
				result++
				dfs(grid, i, j)
			}
		}
	}

	return result
}

func main() {
	// INPUT
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	// OUTPUT
	result := numIslands(grid)
	fmt.Println(result)
}
