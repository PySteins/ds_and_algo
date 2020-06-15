package leetcode

// https://leetcode-cn.com/problems/remove-linked-list-elements/
// 203. 移除链表元素

func removeElements(head *ListNode, val int) *ListNode {
	newHead := new(ListNode)
	newTail := newHead
	for head != nil {
		if head.Val != val {
			newTail.Next = head
			newTail = head
		}
		head = head.Next
	}
	newTail.Next = nil
	return newHead.Next
}
