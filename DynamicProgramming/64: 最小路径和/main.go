package main

import "math"

func main() {

}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cache := make([][]int, m+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}
	for j := 2; j <= n; j++ {
		cache[0][j] = math.MaxInt
	}

	for i := 0; i < m; i++ {
		cache[i+1][0] = math.MaxInt
		for j := 0; j < n; j++ {
			cache[i+1][j+1] = min(cache[i+1][j], cache[i][j+1]) + grid[i][j]
		}
	}

	return cache[m][n]
}
