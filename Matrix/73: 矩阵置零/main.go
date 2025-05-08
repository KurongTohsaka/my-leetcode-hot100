package main

func main() {

}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	stack := make([][2]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				stack = append(stack, [2]int{i, j})
			}
		}
	}

	// 空间复杂度：O(m+n)
	for _, c := range stack {
		x, y := c[0], c[1]
		for i := range n {
			matrix[x][i] = 0
		}
		for j := range m {
			matrix[j][y] = 0
		}
	}
}
