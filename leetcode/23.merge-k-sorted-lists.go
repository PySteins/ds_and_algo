package leetcode

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
// 23. 合并K个排序链表

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge2Node(lists, 0, len(lists)-1)
}

func merge2Node(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	mid := (l + r) / 2
	return mergeNode(merge2Node(lists, l, mid), merge2Node(lists, mid+1, r))
}

func mergeNode(n1, n2 *ListNode) *ListNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.Val < n2.Val {
		n1.Next = mergeNode(n1.Next, n2)
		return n1
	} else {
		n2.Next = mergeNode(n1, n2.Next)
		return n2
	}
}
