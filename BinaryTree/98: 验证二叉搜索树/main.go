package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	pre := math.MinInt // 当前节点前一个节点的值

	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		// 左
		if !dfs(node.Left) {
			return false
		}

		// 中
		if node.Val <= pre {
			return false
		}
		pre = node.Val

		// 右
		if !dfs(node.Right) {
			return false
		}

		return true
	}

	return dfs(root)
}
