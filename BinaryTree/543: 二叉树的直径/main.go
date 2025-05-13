package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftMaxDepth := dfs(root.Left)
		rightMaxDepth := dfs(root.Right)
		ans = max(leftMaxDepth+rightMaxDepth, ans)

		return max(leftMaxDepth, rightMaxDepth) + 1
	}

	dfs(root)

	return ans
}
