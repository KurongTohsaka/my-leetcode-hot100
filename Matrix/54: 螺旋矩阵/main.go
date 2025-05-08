package main

import "math"

func main() {}

func spiralOrder(matrix [][]int) []int {
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上

	m, n := len(matrix), len(matrix[0])
	ans := make([]int, m*n)
	i, j, direction := 0, 0, 0

	for k := range ans {
		ans[k] = matrix[i][j]
		matrix[i][j] = math.MaxInt // 标记
		// 计算下一步位置
		nextI, nextJ := i+dirs[direction][0], j+dirs[direction][1]
		// 这一步是精髓
		if nextI < 0 || nextI >= m || nextJ < 0 || nextJ >= n || matrix[nextI][nextJ] == math.MaxInt {
			direction = (direction + 1) % 4 // 右转 90°
		}
		// 走一步
		i += dirs[direction][0]
		j += dirs[direction][1]
	}

	return ans
}
