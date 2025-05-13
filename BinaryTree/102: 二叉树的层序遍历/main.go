package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	// 使用队列存储节点
	queue := make([]*TreeNode, 0)
	ans := make([][]int, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		// 需要统计当前层节点个数，确定循环次数
		levelSize := len(queue)
		temp := make([]int, 0)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:] // 出队
			temp = append(temp, node.Val)

			// 添加节点操作在完成当前节点后就进行
			if node.Left != nil {
				queue = append(queue, node.Left) // 左子节点入队
			}
			if node.Right != nil {
				queue = append(queue, node.Right) // 右子节点入队
			}
		}

		ans = append(ans, temp)
	}

	return ans
}
