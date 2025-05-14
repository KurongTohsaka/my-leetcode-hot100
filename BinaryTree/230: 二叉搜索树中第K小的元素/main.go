package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	ans := 0
	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 因为最小的元素都在左子树上，所以每次遍历完左子树就 k-- ，当 k = 0 就说明找到了
		dfs(node.Left)
		k--
		if k == 0 {
			ans = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)

	return ans
}
