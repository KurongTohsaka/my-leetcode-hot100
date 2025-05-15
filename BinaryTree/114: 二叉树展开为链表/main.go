package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var head *TreeNode // 代表链表头部
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 后续遍历（相当于从后往前构建链表）
		dfs(node.Right)
		dfs(node.Left)

		// 头插法的核心步骤
		// Left 设置为空是因为展平后不会再有左子节点
		// 将当前节点 node 的 Right 指针指向当前的 head，
		// 这意味着我们将当前节点插入到链表的头部
		node.Left, node.Right = nil, head
		head = node
	}

	dfs(root)
}
