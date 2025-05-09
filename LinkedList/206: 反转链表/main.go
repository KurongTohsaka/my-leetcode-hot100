package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead, cur *ListNode = nil, head

	for cur != nil {
		next := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = next
	}
	return newHead
}
