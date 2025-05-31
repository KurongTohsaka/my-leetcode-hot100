package main

func main() {

}

func uniquePaths(m int, n int) int {
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}
	cache[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// i = 0, 只能往右走
			if i == 0 && j > 0 {
				cache[i][j] = cache[i][j-1]
			}
			// j = 0, 只能向下走
			if j == 0 && i > 0 {
				cache[i][j] = cache[i-1][j]
			}
			if j > 0 && i > 0 {
				cache[i][j] = cache[i-1][j] + cache[i][j-1]
			}
		}
	}

	return cache[m-1][n-1]
}
