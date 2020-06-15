package leetcode

// https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
// 538. 把二叉搜索树转换为累加树

func convertBST(root *TreeNode) *TreeNode {
	num := 0
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		f(node.Right)
		num += node.Val
		node.Val = num
		f(node.Left)
	}
	f(root)
	return root
}
