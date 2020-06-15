package leetcode

// https://leetcode-cn.com/problems/partition-list/
// 86. 分隔链表

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	biggerHead := new(ListNode)
	biggerCur := biggerHead
	smallerHead := new(ListNode)
	smallerCur := smallerHead

	for head != nil {
		if head.Val < x {
			biggerCur.Next = head
			biggerCur = head
		} else {
			smallerCur.Next = head
			smallerCur = head
		}
		head = head.Next
	}
	smallerCur.Next = nil
	biggerCur.Next = smallerHead.Next
	return biggerHead.Next
}
