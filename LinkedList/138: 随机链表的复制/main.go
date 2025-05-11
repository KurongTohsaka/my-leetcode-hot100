package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cachedNode := make(map[*Node]*Node)
	// 复制所有节点，并建立映射关系
	current := head
	for current != nil {
		newNode := &Node{Val: current.Val}
		cachedNode[current] = newNode
		current = current.Next
	}

	current = head
	for current != nil {
		// 获取当前节点对应的新节点
		newNode := cachedNode[current]

		// 连接 Next 指针
		if current.Next != nil {
			newNode.Next = cachedNode[current.Next]
		}

		// 连接 Random 指针
		if current.Random != nil {
			newNode.Random = cachedNode[current.Random]
		}

		current = current.Next
	}

	return cachedNode[head]
}
