package leetcode

// https://leetcode-cn.com/problems/add-two-numbers/
// 2. 两数相加

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummyNode := new(ListNode)
	last := dummyNode
	// 进位
	carry := 0

	var v1, v2 int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		last.Next = &ListNode{Val: sum % 10}
		last = last.Next
	}

	if carry > 0 {
		last.Next = &ListNode{Val: carry}
	}
	return dummyNode.Next
}
