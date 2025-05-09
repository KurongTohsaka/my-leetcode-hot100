package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	// 对每一行进行二分查找
	for i := 0; i < m; i++ {
		left, right := 0, n-1
		for left <= right {
			mid := (left + right) / 2
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	i, j := 0, n-1
	// 利用排序特性，从右上角进行比较，一旦 target 比它大就可以直接排除一整行元素
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i++ // 这一行剩余元素全部小于 target，排除
		} else {
			j-- // 这一列剩余元素全部大于 target，排除
		}
	}

	return false
}
