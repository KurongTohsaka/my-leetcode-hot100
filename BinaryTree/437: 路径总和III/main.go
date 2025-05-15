package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	// 这是第二次做了，又忘了💦
	// 存储不同路径和的计数
	cache := map[int]int{0: 1}
	ans := 0
	// prevSum 是当前节点的前缀路径和
	var dfs func(node *TreeNode, prevSum int)
	dfs = func(node *TreeNode, prevSum int) {
		if node == nil {
			return
		}
		cur := node.Val + prevSum
		ans += cache[cur-targetSum]
		cache[cur]++

		dfs(node.Left, cur)
		dfs(node.Right, cur)
		cache[cur]-- // 回溯，当前路径已经访问过，去掉
	}
	dfs(root, 0)
	return ans
}
