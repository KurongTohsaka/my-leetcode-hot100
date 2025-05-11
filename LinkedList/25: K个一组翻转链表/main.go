package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}

	dummy := &ListNode{Next: head}
	p0 := dummy       // p0 始终指向当前待翻转组的前一个节点
	var pre *ListNode // 用于翻转的辅助指针，初始为nil
	cur := p0.Next    // 当前要处理的节点

	// k 个一组处理
	for ; n >= k; n -= k {
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		// 此时：
		// pre → 翻转后的新头
		// cur → 下一组的第一个节点
		// p0.Next → 原组的第一个节点（现在是新尾）

		// 连接翻转后的子链表
		nxt := p0.Next     // 保存原组的第一个节点（新尾）
		p0.Next.Next = cur // 新尾连接剩余链表
		p0.Next = pre      // 前一组尾连接新头
		p0 = nxt           // 移动p0到新尾（即下一组的前一个节点）
	}
	return dummy.Next
}
