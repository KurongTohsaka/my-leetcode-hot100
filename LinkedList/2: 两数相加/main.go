package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		// 加上上一次的进位
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next
	}

	// 处理最后的进位
	if carry > 0 {
		cur.Next = &ListNode{carry, nil}
	}

	return dummy.Next
}
