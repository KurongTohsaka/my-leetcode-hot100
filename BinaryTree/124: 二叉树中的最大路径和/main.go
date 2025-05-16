package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMax := dfs(node.Left)   // 左子树最大链和
		rightMax := dfs(node.Right) // 右子树最大链和

		curSum := node.Val + leftMax + rightMax // 两条链拼成路径（如果只有其中一条链构成路径就不符合题意了）
		ans = max(ans, curSum)

		// 每条路径由两个子树链构成，这里返回本子树的最大链和
		return max(max(leftMax, rightMax)+node.Val, 0) // 当前子树最大链和（注意这里和 0 取最大值了）
	}

	dfs(root)

	return ans
}
