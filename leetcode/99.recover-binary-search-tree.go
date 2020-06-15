package leetcode

// https://leetcode-cn.com/problems/recover-binary-search-tree/
// 99. 恢复二叉搜索树

var first, second, prev *TreeNode

func recoverTree(root *TreeNode) {
	first, second, prev = (*TreeNode)(nil), (*TreeNode)(nil), (*TreeNode)(nil)
	findWrongNode(root)
	first.Val, second.Val = second.Val, first.Val
}

func findWrongNode(root *TreeNode) {
	if root == nil {
		return
	}
	findWrongNode(root.Left)
	if prev != nil && prev.Val > root.Val {
		if first != nil {
			second = root
			return
		}
		second, first = root, prev
	}
	prev = root
	findWrongNode(root.Right)
	return
}
