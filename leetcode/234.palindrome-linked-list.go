package leetcode

// https://leetcode-cn.com/problems/palindrome-linked-list/
// 234. 回文链表

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	if head.Next.Next == nil {
		return head.Val == head.Next.Val
	}

	mid := findMid(head)
	newHead := reverseLinkList(mid.Next)

	for newHead != nil {
		if head.Val != newHead.Val {
			return false
		}
		head = head.Next
		newHead = newHead.Next
	}

	return true
}

// 找到中间节点
func findMid(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 翻转链表
func reverseLinkList(head *ListNode) (newHead *ListNode) {
	for head != nil {
		tmp := head.Next
		head.Next = newHead
		newHead = head
		head = tmp
	}
	return newHead
}
