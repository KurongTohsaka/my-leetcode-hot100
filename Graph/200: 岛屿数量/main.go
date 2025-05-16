package main

func main() {

}

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	var dfs func(int, int)

	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != '1' {
			return
		}
		// 标记，防止重复访问
		grid[x][y] = '2'

		// 四个方向进行遍历，寻找岛屿
		dfs(x, y+1)
		dfs(x, y-1)
		dfs(x+1, y)
		dfs(x-1, y)
	}

	for i, row := range grid {
		for j, c := range row {
			// 只对岛屿内进行 dfs 遍历，遇到海就结束
			if c == '1' {
				dfs(i, j)
				ans++
			}
		}
	}

	return ans
}
