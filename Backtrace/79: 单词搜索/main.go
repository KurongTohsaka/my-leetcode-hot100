package main

func main() {

}

func exist(board [][]byte, word string) bool {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m, n := len(board), len(board[0])
	var dfs func(x, y, cur int) bool // 返回结果为 (x,y) 的搜索结果是否正确

	dfs = func(x, y, cur int) bool {
		if cur == len(word) {
			return true
		}

		// 越界或不相等就返回
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[cur] {
			return false
		}

		// 标记
		temp := board[x][y]
		board[x][y] = 0

		for _, direction := range directions {
			nx, ny := x+direction[0], y+direction[1]
			// 如果下一个搜索方向正确，就返回 true
			if dfs(nx, ny, cur+1) {
				board[nx][ny] = temp
				return true
			}
		}

		// 回溯
		board[x][y] = temp
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}
