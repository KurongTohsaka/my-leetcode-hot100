package main

import (
	"slices"
	"strings"
)

func main() {

}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	path := make([]string, 2*n) // 预分配空间确实有些用
	var dfs func(i, left int)   // i 为括号位置，left 为左括号数量
	dfs = func(i, left int) {
		if i == 2*n {
			ans = append(ans, strings.Join(slices.Clone(path), ""))
			return
		}

		// 左括号最大数量为 n ，先填左括号
		if left < n {
			path[i] = "("
			dfs(i+1, left+1)
		}

		// 如果右括号比左括号数量少。再填右括号
		if i-left < left {
			path[i] = ")"
			dfs(i+1, left)
		}
	}

	dfs(0, 0)

	return ans
}
