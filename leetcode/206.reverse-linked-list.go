package leetcode

// https://leetcode-cn.com/problems/reverse-linked-list/
// 206. 反转链表

func reverseList(head *ListNode) *ListNode {
	newHead := (*ListNode)(nil)
	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}
	return newHead
}
