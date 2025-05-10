package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 注意是移动后再比较
		if fast == slow { // 相遇说明有环
			return true
		}
	}
	return false
}
