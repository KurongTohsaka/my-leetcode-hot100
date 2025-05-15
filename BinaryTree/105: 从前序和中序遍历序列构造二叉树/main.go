package main

import "slices"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}

	// 返回的索引 leftSize 就是根节点在中序遍历中的位置
	// 由于在中序遍历中，根节点左边的所有元素都属于左子树，
	// 所以 leftSize 也代表了当前子树的左子树的节点个数
	leftSize := slices.Index(inorder, preorder[0])
	// preorder[1:1+leftSize]: 在前序遍历中，根节点之后紧跟着的是左子树的节点。
	// 左子树的节点个数是 leftSize，所以我们取从索引 1 开始的 leftSize 个元素。
	// inorder[:leftSize]: 在中序遍历中，根节点之前的 leftSize 个元素组成了左子树的中序遍历结果。
	left := buildTree(preorder[1:1+leftSize], inorder[:leftSize])
	// preorder[1+leftSize:]: 在前序遍历中，左子树的节点之后就是右子树的节点。
	// 我们取从索引 1+leftSize 开始到末尾的所有元素。
	// inorder[1+leftSize:]: 在中序遍历中，根节点之后的元素组成了右子树的中序遍历结果。
	// 我们取从索引 1+leftSize 开始到末尾的所有元素。
	right := buildTree(preorder[1+leftSize:], inorder[1+leftSize:])

	return &TreeNode{
		Val:   preorder[0],
		Left:  left,
		Right: right,
	}
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	// 存储根节点索引值
	indexCache := make(map[int]int, n)
	for i, x := range inorder {
		indexCache[x] = i
	}

	var dfs func(preLeft, preRight, inLeft, inRight int) *TreeNode
	dfs = func(preLeft, preRight, inLeft, inRight int) *TreeNode {
		// 空节点
		if preLeft == preRight {
			return nil
		}

		// 从这开始与上面的解法一样，只不过换成了区间上下限
		leftSize := indexCache[preorder[preLeft]] - inLeft
		left := dfs(preLeft+1, preLeft+leftSize+1, inLeft, inLeft+leftSize)
		right := dfs(preLeft+leftSize+1, preRight, inLeft+leftSize+1, inRight)
		return &TreeNode{
			Val:   preorder[preLeft],
			Left:  left,
			Right: right,
		}
	}

	// 左闭右开区间
	return dfs(0, n, 0, n)
}
