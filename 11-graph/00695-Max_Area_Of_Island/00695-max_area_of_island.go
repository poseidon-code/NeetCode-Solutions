// 695: Max Area Of Island
// https://leetcode.com/problems/max-area-of-island

package main

import "fmt"

func dfs(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}

	if grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0

	return 1 + dfs(grid, i+1, j) + dfs(grid, i-1, j) + dfs(grid, i, j+1) + dfs(grid, i, j-1)
}

// SOLUTION
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	result := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				curr := dfs(grid, i, j)
				if curr > result {
					result = curr
				}
			}
		}
	}

	return result
}

func main() {
	// INPUT
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	// OUTPUT
	result := maxAreaOfIsland(grid)
	fmt.Println(result)

}
