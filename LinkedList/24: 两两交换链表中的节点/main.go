package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	prev := dummyHead

	for prev.Next != nil && prev.Next.Next != nil {
		slow := prev.Next
		fast := prev.Next.Next

		slow.Next = fast.Next
		fast.Next = slow
		prev.Next = fast

		// 更新哨兵
		prev = slow
	}

	return dummyHead.Next
}
