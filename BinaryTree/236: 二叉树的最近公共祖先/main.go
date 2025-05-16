package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//做第二遍了，还是不会😢

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 函数返回值代表的是 “最近公共祖先” 的候选节点
	// p、q 本身也可以作为公共祖先
	if root == nil || root == p || root == q {
		return root
	}

	// left、right 为空就代表公共祖先没在这半边子树上，就需要返回另一边不为空的作为祖先候选
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	// 当left、right 都存在，就代表候选祖先为root，或是在其两侧子树上，所以这里返回 root
	return root
}
