package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		// 当本行最大元素大于等于 target 时，就确定了 target 所在行
		if matrix[i][n-1] >= target {
			left, right := 0, n-1
			for left <= right {
				mid := (left + right) / 2
				if matrix[i][mid] > target {
					right = mid - 1
				} else if matrix[i][mid] < target {
					left = mid + 1
				} else {
					return true
				}
			}
		}
	}

	return false
}
