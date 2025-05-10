package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	// 找中间节点，此时 slow 为后半链表的头节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 原地反转后半链表
	var newHead, cur *ListNode = nil, slow
	for cur != nil {
		next := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = next
	}

	// 对比两段链表
	for newHead != nil {
		if head.Val != newHead.Val {
			return false
		}
		head = head.Next
		newHead = newHead.Next
	}
	return true
}
