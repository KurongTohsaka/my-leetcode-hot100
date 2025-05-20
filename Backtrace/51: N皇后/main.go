package main

import "strings"

func main() {

}

func solveNQueens(n int) (ans [][]string) {
	queens := make([]int, n)     // 皇后放在 (r,queens[r])
	col := make([]bool, n)       // col 数组记录哪些列已经被占用
	diag1 := make([]bool, n*2-1) // diag1 数组记录哪些主对角线被占用
	diag2 := make([]bool, n*2-1) // diag2 数组记录哪些副对角线被占用
	var dfs func(int)            // r 为当前行
	dfs = func(r int) {
		if r == n {
			board := make([]string, n) // 根据 queens 数组中记录的皇后位置，构建棋盘的字符串表示
			for i, c := range queens {
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)
			}
			ans = append(ans, board)
			return
		}
		// 在 (r,c) 放皇后
		for c, ok := range col {
			// 计算副对角线的索引
			rc := r - c + n - 1
			// 检查当前位置 (r, c) 是否合法：
			// 1. col[c] 为 false：表示第 c 列没有被占用
			// 2. diag1[r+c] 为 false：表示主对角线 (r+c) 没有被占用
			// 3. diag2[rc] 为 false：表示副对角线 (r-c) 没有被占用
			if !ok && !diag1[r+c] && !diag2[rc] {
				queens[r] = c                                    // 直接覆盖，无需恢复现场
				col[c], diag1[r+c], diag2[rc] = true, true, true // 皇后占用了 c 列和两条斜线
				dfs(r + 1)
				col[c], diag1[r+c], diag2[rc] = false, false, false // 恢复现场
			}
		}
	}
	dfs(0)
	return
}
