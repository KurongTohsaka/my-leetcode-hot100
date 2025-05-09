package main

import "slices"

func main() {

}

func rotate(matrix [][]int) {
	n := len(matrix)

	// 直接转置得不到目标矩阵
	// 先转置
	for i := range n {
		for j := range i {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 再把每行翻转
	for i := range n {
		slices.Reverse(matrix[i])
	}
}
